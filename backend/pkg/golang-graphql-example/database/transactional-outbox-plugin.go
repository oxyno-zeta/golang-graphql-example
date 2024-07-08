package database

import (
	"encoding/json"
	"fmt"

	"emperror.dev/errors"
	"gorm.io/gorm"
)

type transactionalOutbox struct {
}

func NewTransactionalOutboxPlugin() gorm.Plugin {
	return &transactionalOutbox{}
}

func (*transactionalOutbox) Name() string {
	return "gorm-transactional-outbox-plugin"
}

func (*transactionalOutbox) Initialize(db *gorm.DB) error {
	// Get callback object
	cb := db.Callback()

	cb.Create().Before("gorm:create").Register("toto", func(db *gorm.DB) {
		if db.Statement.Table == "transactional_outbox" {
			db.AddError(errors.New("fail !"))

			return
		}
	})

	// Register after create
	err := cb.Create().After("gorm:create").Register("transactional-outbox:create:after", func(db *gorm.DB) {
		fmt.Println(db.Statement.SQL.String())
		// Check if error is present to ignore message or that there aren't rows affected
		if db.Error != nil || db.RowsAffected == 0 {
			return
		}

		// // Try to cast to have Object schema version
		// casted, ok := db.Statement.Model.(ObjectSchemaVersionTransactionalOutbox)
		// // Check if transactional outbox isn't supported
		// if !ok {
		// 	// Ignore it
		// 	return
		// }

		// JSON Marshal model
		result, err := json.Marshal(db.Statement.Model)
		// Check error
		if err != nil {
			// Save error and return
			db.AddError(errors.WithStack(err))

			return
		}

		// Build transactional outbox object
		to := &TransactionalOutbox{
			Action:    TransactionalOutboxCreateAction,
			Table:     db.Statement.Table,
			EventDate: db.NowFunc(),
			// ObjectSchemaVersion: casted.ObjectSchemaVersion(),
			Result: result,
		}
		fmt.Println(to)

		// Save event
		ddb := db.Session(&gorm.Session{NewDB: true})
		// ddb.Statement = &gorm.Statement{}
		err = ddb.Create(to).Error
		// Check error
		if err != nil {
			// Save error and return
			db.AddError(errors.WithStack(err))

			return
		}
	})
	// Check error
	if err != nil {
		return err
	}

	// // Register after delete
	// err = cb.Delete().After("gorm:delete").Register("delta:delete:after", func(db *gorm.DB) {
	// 	// Check if error is present to ignore message and that there are rows affected
	// 	if db.Error == nil && db.RowsAffected != 0 {
	// 		// Build delta object
	// 		delt := &Delta{
	// 			Action:    TransactionalOutboxDeleteAction,
	// 			Table:     db.Statement.Table,
	// 			Result:    db.Statement.Model,
	// 			EventDate: db.NowFunc(),
	// 		}

	// 		// Send to channel
	// 		dtp.notificationChan <- delt
	// 	}
	// })
	// // Check error
	// if err != nil {
	// 	return err
	// }

	// err = cb.Update().Before("gorm:update").Register("delta:update:before", func(db *gorm.DB) { //nolint: wsl
	// 	// Get reflect value of models to check if it is an array.
	// 	// If it is, then this should be a bulk operation and then we should add or complete a RETURNING clause
	// 	// To add "id" and "updated_at" columns in the result

	// 	// Reflect models
	// 	v := reflect.ValueOf(db.Statement.Model)
	// 	// Indirect it
	// 	v = reflect.Indirect(v)
	// 	// Check type
	// 	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
	// 		// Check if there is a returning clause
	// 		db.Clauses(&clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "updated_at"}}})
	// 	}
	// })
	// // Check error
	// if err != nil {
	// 	return err
	// }

	// // Register after update
	// err = cb.Update().After("gorm:update").Register("delta:update:after", func(db *gorm.DB) {
	// 	// Check if error is present to ignore message and that there are rows affected
	// 	if db.Error == nil && db.RowsAffected != 0 {
	// 		eventDate := db.NowFunc()
	// 		// Init patch object
	// 		var patch map[string]interface{}
	// 		// Init action
	// 		var action = TransactionalOutboxUpdateAction
	// 		// Check type
	// 		v, ok := db.Statement.Dest.(map[string]interface{})
	// 		// Check if it is ok
	// 		if ok {
	// 			patch = v
	// 			// Update action
	// 			action = TransactionalOutboxPatchAction
	// 		}

	// 		// Reflect models
	// 		vr := reflect.ValueOf(db.Statement.Model)
	// 		// Indirect it
	// 		vr = reflect.Indirect(vr)
	// 		// Check type
	// 		if vr.Kind() == reflect.Array || vr.Kind() == reflect.Slice {
	// 			for i := 0; i < vr.Len(); i++ {
	// 				item := vr.Index(i)

	// 				// Build delta object
	// 				delt := &Delta{
	// 					Action:    action,
	// 					Table:     db.Statement.Table,
	// 					Result:    item.Interface(),
	// 					Patch:     patch,
	// 					EventDate: eventDate,
	// 				}

	// 				// Send to channel
	// 				dtp.notificationChan <- delt
	// 			}
	// 		} else {
	// 			// Build delta object
	// 			delt := &Delta{
	// 				Action:    action,
	// 				Table:     db.Statement.Table,
	// 				Result:    db.Statement.Model,
	// 				Patch:     patch,
	// 				EventDate: eventDate,
	// 			}

	// 			// Send to channel
	// 			dtp.notificationChan <- delt
	// 		}
	// 	}
	// })
	// // Check error
	// if err != nil {
	// 	return err
	// }

	return nil
}
