package email

import (
	"crypto/tls"
	"time"

	"emperror.dev/errors"

	spmail "github.com/xhit/go-simple-mail/v2"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

const (
	CfgEncryptionNone        = "NONE"
	CfgEncryptionTLS         = "TLS"
	CfgEncryptionSSL         = "SSL"
	CfgAuthenticationPlain   = "PLAIN"
	CfgAuthenticationLogin   = "LOGIN"
	CfgAuthenticationCRAMMD5 = "CRAM-MD5"
	CfgDefaultTimeout        = 10 * time.Second
)

type service struct {
	logger     log.Logger
	cfgManager config.Manager
	server     *spmail.SMTPServer
}

func (*service) NewEmail() Email {
	return &email{
		firstBodySet: false,
		spemail:      spmail.NewMSG(),
	}
}

func (s *service) InitializeAndReload() error {
	// Get configuration
	cfg := s.cfgManager.GetConfig().SMTP

	// Check if configuration exists
	// If it is equal to nil, skip connect
	if cfg == nil {
		s.logger.Info("SMTP configuration not present, server creation skipped")

		// Force flush created service
		// In case of reload done with previously created service
		s.server = nil

		return nil
	}

	s.logger.Debug("Trying to create SMTP server")

	// Create SMTP client server
	server := spmail.NewSMTPClient()

	// SMTP Server configuration
	server.Host = cfg.Host
	server.Port = cfg.Port

	// Configure username
	if cfg.Username != nil {
		server.Username = cfg.Username.Value
	}
	// Configure password
	if cfg.Password != nil {
		server.Password = cfg.Password.Value
	}

	// Encryption (default value is TLS)
	switch cfg.Encryption {
	case CfgEncryptionTLS:
		server.Encryption = spmail.EncryptionTLS
	case CfgEncryptionNone:
		server.Encryption = spmail.EncryptionNone
	case CfgEncryptionSSL:
		server.Encryption = spmail.EncryptionSSL
	default:
		server.Encryption = spmail.EncryptionTLS
	}

	// Put authentication only if username and password exists
	if cfg.Username != nil && cfg.Password != nil {
		// Authentication type
		switch cfg.AuthenticationType {
		case CfgAuthenticationPlain:
			server.Authentication = spmail.AuthPlain
		case CfgAuthenticationLogin:
			server.Authentication = spmail.AuthLogin
		case CfgAuthenticationCRAMMD5:
			server.Authentication = spmail.AuthCRAMMD5
		default:
			server.Authentication = spmail.AuthPlain
		}
	}

	// Variable to keep alive connection
	server.KeepAlive = cfg.KeepAlive

	// Check if connect timeout isn't set
	if cfg.ConnectTimeout == "" {
		// Timeout for connect to SMTP Server
		server.ConnectTimeout = CfgDefaultTimeout
	} else {
		// Parse connect timeout duration
		connectTimeoutDur, err := time.ParseDuration(cfg.ConnectTimeout)
		// Check error
		if err != nil {
			return errors.WithStack(err)
		}
		// Timeout for connect to SMTP Server
		server.ConnectTimeout = connectTimeoutDur
	}

	// Check if send timeout isn't set
	if cfg.SendTimeout == "" {
		// Timeout for send the data and wait respond
		server.SendTimeout = CfgDefaultTimeout
	} else {
		// Parse send timeout duration
		sendTimeoutDur, err := time.ParseDuration(cfg.SendTimeout)
		// Check error
		if err != nil {
			return errors.WithStack(err)
		}
		// Timeout for send the data and wait respond
		server.SendTimeout = sendTimeoutDur
	}

	// Check if skip tls verify exists and is set
	if cfg.TLSSkipVerify {
		// Set TLSConfig to provide custom TLS configuration. For example,
		// to skip TLS verification (useful for testing):
		server.TLSConfig = &tls.Config{InsecureSkipVerify: true} //nolint: gosec // TLS Skip wanted
	}

	// Save SMTP server
	s.server = server

	s.logger.Info("SMTP server created")

	return nil
}

func (s *service) Check() error {
	// Check if server exists, if not, skip send
	if s.server == nil {
		return nil
	}

	// Connect server
	client, err := s.server.Connect()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}
	// Defer close client
	defer client.Close()

	return nil
}

func (s *service) Send(em Email) error {
	// Check if server exists, if not, skip send
	if s.server == nil {
		s.logger.Debug(
			"SMTP server not present (because configuration wasn't present probably), send skipped",
		)

		return nil
	}

	// Connect server
	client, err := s.server.Connect()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}
	// Defer close client
	defer client.Close()

	// Get email object
	e := em.GetEmail()
	// Send email
	err = e.Send(client)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}
