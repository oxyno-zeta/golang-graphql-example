package database

import (
	"time"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// BaseWithoutAllIndexes contains common columns for all tables.
type BaseWithoutAllIndexes struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
	ID        string     `json:"id"                  gorm:"primary_key"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseWithoutAllIndexes) BeforeCreate(_ *gorm.DB) error {
	// Check if ID is set to avoid erasing it.
	// This is useful when it is asked to save object for the first
	// time with a fixed id.
	if base.ID == "" {
		// Generate new id
		uuidGenerated, err := uuid.NewV7()
		if err != nil {
			return errors.WithStack(err)
		}

		// Save new id
		base.ID = uuidGenerated.String()
	}

	return nil
}
