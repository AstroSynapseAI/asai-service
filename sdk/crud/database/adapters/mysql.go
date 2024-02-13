package adapters

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySql struct {
	config *database.DBConfig
	gorm *gorm.DB
}

func NewMySQL(options ...database.DatabaseOptions) *MySql {
	adapter := &MySql{
		config: &database.DBConfig{},
	}

	for _, option := range options {
		option(adapter.config)
	}

	if adapter.config.DSN == "" && adapter.config.DBName == "" {
		fmt.Println("Missing DSN or database configuration for MySQL adapter.") 
		return nil
	}

	gorm, err := gorm.Open(adapter.open(), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open MySQL database:", err)
		return nil
	}

	adapter.gorm = gorm
	return adapter
}

func (adapter *MySql) Gorm() *gorm.DB {
	return adapter.gorm
}

func (adapter *MySql) open() gorm.Dialector {
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	if adapter.config.DSN == "" {
		adapter.config.DSN = fmt.Sprintf(dsn,
			adapter.config.DBUser,
			adapter.config.DBPass,
			adapter.config.DBHost,
			adapter.config.DBPort,
			adapter.config.DBName,
		)
	}	
	return mysql.Open(dsn)
}