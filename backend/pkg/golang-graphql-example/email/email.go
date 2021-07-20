package email

import (
	"time"

	"github.com/pkg/errors"
	spmail "github.com/xhit/go-simple-mail/v2"
)

const dateFormat = "2006-01-02 15:04:05 MST"

type email struct {
	spemail      *spmail.Email
	firstBodySet bool
}

// SetFrom sets the From address.
func (e *email) SetFrom(address string) {
	e.spemail.SetFrom(address)
}

// SetSender sets the Sender address.
func (e *email) SetSender(address string) {
	e.spemail.SetSender(address)
}

// SetReplyTo sets the Reply-To address.
func (e *email) SetReplyTo(address string) {
	e.spemail.SetReplyTo(address)
}

// AddTo adds a To address. You can provide multiple
// addresses at the same time.
func (e *email) AddTo(addresses ...string) {
	e.spemail.AddTo(addresses...)
}

// AddCc adds a Cc address. You can provide multiple
// addresses at the same time.
func (e *email) AddCc(addresses ...string) {
	e.spemail.AddCc(addresses...)
}

// AddBcc adds a Bcc address. You can provide multiple
// addresses at the same time.
func (e *email) AddBcc(addresses ...string) {
	e.spemail.AddBcc(addresses...)
}

// SetDate sets the Date header with the provided date/time (in RFC1123Z format).
func (e *email) SetDate(dateTime time.Time) {
	e.spemail.SetDate(dateTime.Format(dateFormat))
}

// SetSubject sets the subject of the email message.
func (e *email) SetSubject(subject string) {
	e.spemail.SetSubject(subject)
}

// SetTextBody sets text body.
func (e *email) SetTextBody(body string) {
	if !e.firstBodySet {
		// Set first body
		e.spemail.SetBody(spmail.TextPlain, body)
		// Turn on flag
		e.firstBodySet = true

		return
	}

	e.spemail.AddAlternative(spmail.TextPlain, body)
}

// SetHTMLBody sets html body.
func (e *email) SetHTMLBody(body string) {
	if !e.firstBodySet {
		// Set first body
		e.spemail.SetBody(spmail.TextHTML, body)
		// Turn on flag
		e.firstBodySet = true

		return
	}

	e.spemail.AddAlternative(spmail.TextHTML, body)
}

// SetPriority will add headers for priority.
func (e *email) SetPriority(priority Priority) {
	if priority == HighPriority {
		e.spemail.SetPriority(spmail.PriorityHigh)

		return
	}

	e.spemail.SetPriority(spmail.PriorityLow)
}

// AddAttachment is used to attach content from a byte array to the email.
// Required parameters include a byte array and the desired filename for the attachment. The MIME-Type is optional.
func (e *email) AddAttachment(data []byte, filename string, mimeType string) error {
	// Add attachment
	e.spemail.AddAttachmentData(data, filename, mimeType)

	// Get error
	err := e.spemail.GetError()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}

// AddAttachmentFile is used to attach content to the email.
// It attempts to open the file referenced by filename and, if successful, creates an Attachment.
// This Attachment is then appended to the email.
func (e *email) AddAttachmentFile(filePath string) error {
	// Add attachment
	e.spemail.AddAttachment(filePath)

	// Get error
	err := e.spemail.GetError()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}

// AddInlineAttachmentFile is used to attach content to the email as HTML inline attachment.
// It attempts to open the file referenced by filename and, if successful, creates an Attachment.
// This Attachment is then appended to the email.
func (e *email) AddInlineAttachmentFile(filePath string) error {
	// Add attachment
	e.spemail.AddInline(filePath)

	// Get error
	err := e.spemail.GetError()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}

// AddInlineAttachment is used to attach content from a byte array to the email as HTML inline attachment.
// Required parameters include a byte array and the desired filename for the attachment. The MIME-Type is optional.
func (e *email) AddInlineAttachment(data []byte, filename string, mimeType string) error {
	// Add attachment
	e.spemail.AddInlineData(data, filename, mimeType)

	// Get error
	err := e.spemail.GetError()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}

// GetEmail will get email object (used internally for sending).
func (e *email) GetEmail() *spmail.Email {
	return e.spemail
}
