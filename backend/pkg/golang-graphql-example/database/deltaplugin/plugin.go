package deltaplugin

import (
	"gorm.io/gorm"
)

type deltaPlugin struct {
	notificationChan chan *Delta
}

func New(notificationChan chan *Delta) gorm.Plugin {
	return &deltaPlugin{notificationChan: notificationChan}
}

func (*deltaPlugin) Name() string {
	return "gorm-delta-plugin"
}

func (dtp *deltaPlugin) Initialize(db *gorm.DB) error {
	// Get callback object
	cb := db.Callback()

	// Register after create
	err := cb.Create().After("gorm:create").Register("delta:create:after", func(db *gorm.DB) {
		// Check if error is present to ignore message and that there are rows affected
		if db.Error == nil && db.RowsAffected != 0 {
			// Build delta object
			delt := &Delta{
				Action: CREATE,
				Table:  db.Statement.Table,
				Result: db.Statement.Model,
			}

			// Send to channel
			dtp.notificationChan <- delt
		}
	})
	// Check error
	if err != nil {
		return err
	}

	// Register after delete
	err = cb.Delete().After("gorm:delete").Register("delta:delete:after", func(db *gorm.DB) {
		// Check if error is present to ignore message and that there are rows affected
		if db.Error == nil && db.RowsAffected != 0 {
			// Build delta object
			delt := &Delta{
				Action: DELETE,
				Table:  db.Statement.Table,
				Result: db.Statement.Model,
			}

			// Send to channel
			dtp.notificationChan <- delt
		}
	})
	// Check error
	if err != nil {
		return err
	}

	// Register after update
	err = cb.Update().After("gorm:update").Register("delta:update:after", func(db *gorm.DB) {
		// Check if error is present to ignore message and that there are rows affected
		if db.Error == nil && db.RowsAffected != 0 {
			// Init patch object
			var patch map[string]interface{}
			// Init action
			var action = UPDATE
			// Check type
			v, ok := db.Statement.Dest.(map[string]interface{})
			// Check if it is ok
			if ok {
				patch = v
				// Update action
				action = PATCH
			}

			// Build delta object
			delt := &Delta{
				Action: action,
				Table:  db.Statement.Table,
				Result: db.Statement.Model,
				Patch:  patch,
			}

			// Send to channel
			dtp.notificationChan <- delt
		}
	})
	// Check error
	if err != nil {
		return err
	}

	return nil
}
