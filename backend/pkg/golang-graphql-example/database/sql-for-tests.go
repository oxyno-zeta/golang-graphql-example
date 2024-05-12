//go:build unit

package database

import "gorm.io/gorm"

// This method is added only for unit tests where injecting a gorm db mocked is important
// That's why it is built only on this target
func (sdb *sqldb) SetGormDB(db *gorm.DB) {
	sdb.db = db
}
