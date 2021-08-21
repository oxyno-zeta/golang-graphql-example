package sequences

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"gorm.io/gorm"
)

var Seq201608List = []*gormigrate.Migration{
	// create persons table
	{
		ID: "201608301400",
		Migrate: func(tx *gorm.DB) error {
			// it's a good practice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type Person struct {
				database.Base
				Name string
			}

			return tx.AutoMigrate(&Person{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("people")
		},
	},
	// add age column to persons
	{
		ID: "201608301415",
		Migrate: func(tx *gorm.DB) error {
			// when table already exists, it just adds fields as columns
			type Person struct {
				Age int
			}

			return tx.AutoMigrate(&Person{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropColumn("people", "age")
		},
	},
	// add pets table
	{
		ID: "201608301430",
		Migrate: func(tx *gorm.DB) error {
			type Pet struct {
				database.Base
				Name     string
				PersonID int
			}

			return tx.AutoMigrate(&Pet{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("pets")
		},
	},
}
