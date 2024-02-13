package adapters

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	config *database.DBConfig
	gorm *gorm.DB
}

func NewPostgres(options ...database.DatabaseOptions) *Postgres {
	adapter := &Postgres{
		config: &database.DBConfig{},
	}

	for _, option := range options {
		option(adapter.config)
	}

	if adapter.config.DSN == "" && adapter.config.DBName == "" {
		fmt.Println("Missing DSN or database configuration for Postgres adapter.") 
		return nil
	}

	gorm, err := gorm.Open(adapter.open(), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open Postgres database:", err)
		return nil
	}

	adapter.gorm = gorm

	return adapter	
}

func (adapter *Postgres) Gorm() *gorm.DB {
	return adapter.gorm
}

func (adapter *Postgres) open() gorm.Dialector {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	if adapter.config.DSN == "" {
		adapter.config.DSN = fmt.Sprintf(dsn,
			adapter.config.DBHost,
			adapter.config.DBUser,
			adapter.config.DBPass,
			adapter.config.DBName,
			adapter.config.DBPort,
		)
	}	
	return postgres.Open(adapter.config.DSN)
}