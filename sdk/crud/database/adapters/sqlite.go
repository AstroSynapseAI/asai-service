package adapters

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ database.Adapter = (*SQLite)(nil)

type SQLite struct {
	config *database.DBConfig
	gorm *gorm.DB
}

func NewSQLite(options ...database.DatabaseOptions) *SQLite {
	adapter := &SQLite{
		config: &database.DBConfig{},
	}

	for _, option := range options {
		option(adapter.config)
	}

	if adapter.config.DSN == "" && adapter.config.DBName == "" {
		fmt.Println("Missing DSN or database configuration for SQLite adapter.") 
		return nil
	}

	gorm, err := gorm.Open(adapter.open(), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open SQLite database:", err)
		return nil
	}

	adapter.gorm = gorm
	return adapter
}

func (adapter *SQLite) Gorm() *gorm.DB {
	return adapter.gorm
}

func (adapter *SQLite) open() gorm.Dialector {
	var dsn string
	if adapter.config.DSN == "" {
		dsn = adapter.config.DBName
	}	
	return sqlite.Open(dsn)
}