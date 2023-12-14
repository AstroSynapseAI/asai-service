package app

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/go-gormigrate/gormigrate/v2"
)

type Migrations struct{}

var _ database.Migrations = (*Migrations)(nil)

func (*Migrations) Models() []interface{} {
	return []interface{}{
		&models.Account{},
		&models.Agent{},
		&models.Avatar{},
		&models.Document{},
		&models.LLM{},
		&models.Plugin{},
		&models.Role{},
		&models.Tool{},
		&models.User{},
	}
}

func (*Migrations) GormMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		// Example for adding a new field Test to users
		// {
		// 	ID: "add_test_to_user",
		// 	Migrate: func(tx *gorm.DB) error {
		// 		type User struct {
		// 			Test int
		// 		}
		// 		return tx.AutoMigrate(&User{})
		// 	},
		// 	Rollback: func(tx *gorm.DB) error {
		// 		return tx.Migrator().DropColumn(&model.User{}, "test")
		// 	},
		// },
	}
}
