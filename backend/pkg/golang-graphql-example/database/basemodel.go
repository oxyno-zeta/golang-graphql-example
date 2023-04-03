package database

import (
	"time"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	ID        string     `gorm:"primary_key"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	// Check if ID is set to avoid erasing it.
	// This is useful when it is asked to save object for the first
	// time with a fixed id.
	if base.ID == "" {
		// Generate new id
		uuid, err := uuid.NewV4()
		if err != nil {
			return errors.WithStack(err)
		}

		// Save new id
		base.ID = uuid.String()
	}

	return nil
}
