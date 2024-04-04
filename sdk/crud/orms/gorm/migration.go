package gorm

import (
	"errors"
	"fmt"

	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var _ database.Migrator = (*GormMigrator)(nil)

type GormMigrator struct {
	DB         *gorm.DB
	Migrations []*gormigrate.Migration
	Models     []interface{}
}

func NewGormMigrator(db *database.Database) *GormMigrator {
	migrations := &GormMigrator{
		DB: db.Adapter.Gorm(),
	}

	migrations.Models = append(migrations.Models, &DBSeeder{})

	return migrations
}

func (migrator *GormMigrator) AddMigrations(migrations database.Migrations) {
	migrator.Migrations = append(migrator.Migrations, migrations.GormMigrations()...)
	migrator.Models = append(migrator.Models, migrations.Models()...)
}

func (migrator *GormMigrator) AddModels(models []interface{}) {
	migrator.Models = append(migrator.Models, models)
}

func (migrator *GormMigrator) Run() error {
	if err := migrator.migrateModels(); err != nil {
		fmt.Printf("Failed to migrate models: %s\n", err)
		return err
	}

	if migrator.Migrations == nil {
		return errors.New("No migrations to run!")
	}

	migration := gormigrate.New(migrator.DB, gormigrate.DefaultOptions, migrator.Migrations)
	migration.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(migrator.Models...)
		if err != nil {
			fmt.Printf("Init Schema failed: %s\n", err)
			return err
		}
		return nil
	})

	if err := migration.Migrate(); err != nil && err != gorm.ErrInvalidField {
		fmt.Printf("Failed to run migrations: %s\n", err)
		return err
	}
	return nil
}

func (migrator *GormMigrator) migrateModels() error {
	stmt := &gorm.Statement{DB: migrator.DB}

	if len(migrator.Models) == 0 {
		return errors.New("No models to migrate!")
	}

	for _, model := range migrator.Models {
		if err := stmt.Parse(&model); err != nil {
			return err
		}

		id := fmt.Sprintf("create_%v", stmt.Schema.Table)
		migration := &gormigrate.Migration{
			ID: id,
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model)
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(stmt.Schema.Table)
			},
		}
		migrator.Migrations = append(migrator.Migrations, migration)
	}

	return nil
}

