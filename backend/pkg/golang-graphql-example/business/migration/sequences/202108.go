package sequences

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"gorm.io/gorm"
)

var Seq202108List = []*gormigrate.Migration{
	// Add todos
	{
		ID: "202108201848",
		Migrate: func(tx *gorm.DB) error {
			type Todo struct {
				database.Base
				Text string `gorm:"type:varchar(200)"`
				Done bool
			}

			return tx.AutoMigrate(&Todo{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("todos")
		},
	},
	// Increase text field size
	{
		ID: "202108201849",
		Migrate: func(tx *gorm.DB) error {
			type Todo struct {
				Text string `gorm:"type:varchar(2000)"`
			}

			return tx.AutoMigrate(&Todo{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Raw(`ALTER TABLE "todos" ALTER COLUMN "text" TYPE varchar(200)`).Error
		},
	},
}
