package database

import (
	"time"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type ActionType string

const (
	TransactionalOutboxUpdateAction     ActionType = "UPDATE"
	TransactionalOutboxCreateAction     ActionType = "CREATE"
	TransactionalOutboxBulkCreateAction ActionType = "BULK_CREATE"
	TransactionalOutboxPatchAction      ActionType = "PATCH"
	TransactionalOutboxBulkPatchAction  ActionType = "BULK_PATCH"
	TransactionalOutboxDeleteAction     ActionType = "DELETE"
	TransactionalOutboxBulkDeleteAction ActionType = "BULK_DELETE"
)

type TransactionalOutbox struct {
	ID           string     `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `gorm:"index"       json:"deletedAt,omitempty"`
	EventDate    time.Time
	ConsumedDate *time.Time
	Table        string
	Action       ActionType `gorm:"type:varchar(50)"`
	Patch        []byte
	Result       []byte
}

// TableName.
func (TransactionalOutbox) TableName() string {
	return "transactional_outbox"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (t *TransactionalOutbox) BeforeCreate(_ *gorm.DB) error {
	// Check if ID is set to avoid erasing it.
	// This is useful when it is asked to save object for the first
	// time with a fixed id.
	if t.ID == "" {
		// Generate new id
		uuidGenerated, err := uuid.NewV7()
		if err != nil {
			return errors.WithStack(err)
		}

		// Save new id
		t.ID = uuidGenerated.String()
	}

	return nil
}

type ObjectSchemaVersionTransactionalOutbox interface {
	ObjectSchemaVersion() int
}

type MapPatchToJSONTransactionalOutbox interface {
	MapPatchToJSON(patch map[string]interface{}) (map[string]interface{}, error)
}

type Delta struct {
	EventDate time.Time              `json:"eventDate"`
	Patch     map[string]interface{} `json:"patch"`
	Result    interface{}            `json:"result"`
	Table     string                 `json:"table"`
	Action    ActionType             `json:"action"`
}
