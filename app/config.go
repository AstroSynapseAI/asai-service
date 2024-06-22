package app

import (
	"fmt"
	"os"

	"github.com/AstroSynapseAI/asai-service/models"
	"github.com/GoLangWebSDK/crud/database"
	"github.com/GoLangWebSDK/crud/database/adapters"
	"github.com/GoLangWebSDK/crud/orms/gorm"
)

const (
	DefaultDevDSN = "postgresql://asai-admin:asai-password@asai-db:5432/asai-db"
)

type ServerAdapter interface {
	Run(*database.Database) error
}

type Config struct {
	DB  *database.Database
	ENV string
	DSN string
}

func NewConfig() *Config {
	config := &Config{
		ENV: os.Getenv("ENVIRONMENT"),
	}

	return config
}

func (cnf *Config) InitDB() {
	if cnf.DSN == "" {
		fmt.Println("Empty DSN, setting default.")
		cnf.DSN = DefaultDevDSN
	}

	adapter := adapters.NewPostgres(
		database.WithDSN(cnf.DSN),
	)

	cnf.DB = database.New(adapter)

	migration := gorm.NewGormMigrator(cnf.DB)
	migration.AddMigrations(&Migrations{})
	migration.Run()

	seeder := gorm.NewGormSeeder(cnf.DB)
	seeders := seeder.AddSeeder(
		&models.Agent{},
		&models.LLM{},
		&models.Plugin{},
		&models.Role{},
		&models.Tool{},
		&models.User{},
	)

	err := seeders.Run()
	if err != nil {
		fmt.Println("Error running seeders:", err)
		return
	}
}

func (cnf *Config) RunServer(server ServerAdapter) error {
	return server.Run(cnf.DB)
}

func (cnf *Config) LoadEnvironment() {
	fmt.Println("Current Environment:", cnf.ENV)

	if cnf.ENV == "LOCAL DEV" {
		cnf.setupLocalDev()
		return
	}
	if cnf.ENV == "HEROKU DEV" {
		cnf.setupHeroku()
		return
	}

	if cnf.ENV == "AWS DEV" {
		cnf.setupAWS()
		return
	}

	if cnf.ENV == "AWS PROD" {
		cnf.setupAWSProd()
		return
	}

	fmt.Println("Unknown Environment")
}

func (cnf *Config) setupAWSProd() {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	cnf.DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
}

func (cnf *Config) setupAWS() {
	username := os.Getenv("RDS_USERNAME")
	password := os.Getenv("RDS_PASSWORD")
	database := os.Getenv("RDS_DBNAME")
	host := os.Getenv("RDS_HOST")
	port := os.Getenv("RDS_PORT")

	cnf.DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
}

func (cnf *Config) setupHeroku() {
	cnf.DSN = os.Getenv("DATABASE_URL")
}

func (cnf *Config) setupLocalDev() {
	cnf.DSN = DefaultDevDSN
}
