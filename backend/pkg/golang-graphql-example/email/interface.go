package email

import (
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	spmail "github.com/xhit/go-simple-mail/v2"
)

//go:generate mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/email Service
type Service interface {
	// InitializeAndReload service.
	// If configuration isn't set, the setup will be skipped.
	InitializeAndReload() error
	// Check service health.
	// If configuration isn't set, the check will be skipped.
	Check() error
	// Send will send the email.
	// If configuration isn't set, the send action will be skipped.
	// NOTE: Connect SMTP server can take 1 second.
	Send(em Email) error
	// Create a new Email object.
	NewEmail() Email
}

type Priority string

const (
	HighPriority Priority = "HIGH"
	LowPriority  Priority = "LOW"
)

//go:generate mockgen -destination=./mocks/mock_Email.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/email Email
type Email interface {
	// SetFrom sets the From address.
	SetFrom(address string)
	// SetSender sets the Sender address.
	SetSender(address string)
	// SetReplyTo sets the Reply-To address.
	SetReplyTo(address string)
	// AddTo adds a To address. You can provide multiple
	// addresses at the same time.
	AddTo(addresses ...string)
	// AddCc adds a Cc address. You can provide multiple
	// addresses at the same time.
	AddCc(addresses ...string)
	// AddBcc adds a Bcc address. You can provide multiple
	// addresses at the same time.
	AddBcc(addresses ...string)
	// SetDate sets the Date header with the provided date/time (in RFC1123Z format).
	SetDate(dateTime time.Time)
	// SetSubject sets the subject of the email message.
	SetSubject(subject string)
	// SetTextBody sets text body.
	SetTextBody(body string)
	// SetHTMLBody sets html body.
	SetHTMLBody(body string)
	// SetPriority will add headers for priority.
	SetPriority(priority Priority)
	// AddAttachment is used to attach content from a byte array to the email.
	// Required parameters include a byte array and the desired filename for the attachment. The MIME-Type is optional.
	AddAttachment(data []byte, filename string, mimeType string) error
	// AddAttachmentFile is used to attach content to the email.
	// It attempts to open the file referenced by filename and, if successful, creates an Attachment.
	// This Attachment is then appended to the email.
	AddAttachmentFile(filePath string) error
	// AddInlineAttachmentFile is used to attach content to the email as HTML inline attachment.
	// It attempts to open the file referenced by filename and, if successful, creates an Attachment.
	// This Attachment is then appended to the email.
	AddInlineAttachmentFile(filePath string) error
	// AddInlineAttachment is used to attach content from a byte array to the email as HTML inline attachment.
	// Required parameters include a byte array and the desired filename for the attachment. The MIME-Type is optional.
	AddInlineAttachment(data []byte, filename string, mimeType string) error
	// GetEmail will get email object (used internally for sending).
	GetEmail() *spmail.Email
}

func NewService(cfgManager config.Manager, logger log.Logger) Service {
	return &service{cfgManager: cfgManager, logger: logger}
}
