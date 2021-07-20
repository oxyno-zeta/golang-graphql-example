package database

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return errors.WithStack(err)
	}

	// Save new id
	base.ID = uuid.String()

	return nil
}
