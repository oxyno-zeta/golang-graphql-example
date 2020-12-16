package email

import (
	"crypto/tls"
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	spmail "github.com/xhit/go-simple-mail/v2"
)

const CfgEncryptionNone = "NONE"
const CfgEncryptionTLS = "TLS"
const CfgEncryptionSSL = "SSL"
const CfgAuthenticationPlain = "PLAIN"
const CfgAuthenticationLogin = "LOGIN"
const CfgAuthenticationCRAMMD5 = "CRAM-MD5"
const CfgDefaultTimeout = 10 * time.Second

type service struct {
	logger     log.Logger
	cfgManager config.Manager
	server     *spmail.SMTPServer
}

func (s *service) Initialize() error {
	// Get configuration
	cfg := s.cfgManager.GetConfig().SMTP

	// Check if configuration exists
	// If it is equal to nil, skip connect
	if cfg == nil {
		s.logger.Info("SMTP configuration not present, server creation skipped")

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
	switch {
	case cfg.Encryption == CfgEncryptionTLS:
		server.Encryption = spmail.EncryptionTLS
	case cfg.Encryption == CfgEncryptionNone:
		server.Encryption = spmail.EncryptionNone
	case cfg.Encryption == CfgEncryptionSSL:
		server.Encryption = spmail.EncryptionSSL
	default:
		server.Encryption = spmail.EncryptionTLS
	}

	// Put authentication only if username and password exists
	if cfg.Username != nil && cfg.Password != nil {
		// Authentication type
		switch {
		case cfg.AuthenticationType == CfgAuthenticationPlain:
			server.Authentication = spmail.AuthPlain
		case cfg.AuthenticationType == CfgAuthenticationLogin:
			server.Authentication = spmail.AuthLogin
		case cfg.AuthenticationType == CfgAuthenticationCRAMMD5:
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
			return err
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
			return err
		}
		// Timeout for send the data and wait respond
		server.SendTimeout = sendTimeoutDur
	}

	// Check if skip tls verify exists and is set
	if cfg.TLSSkipVerify {
		// Set TLSConfig to provide custom TLS configuration. For example,
		// to skip TLS verification (useful for testing):
		server.TLSConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec // TLS Skip wanted
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
		return err
	}
	// Defer close client
	defer client.Close()

	return nil
}

func (s *service) Send(em Email) error {
	// Check if server exists, if not, skip send
	if s.server == nil {
		s.logger.Debug("SMTP server not present (because configration wasn't present probably), send skipped")

		return nil
	}

	// Connect server
	client, err := s.server.Connect()
	// Check error
	if err != nil {
		return err
	}
	// Defer close client
	defer client.Close()

	// Get email object
	e := em.GetEmail()
	// Send email
	return e.Send(client)
}
