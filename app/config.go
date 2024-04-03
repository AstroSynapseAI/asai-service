package app

import (
	"fmt"
	"os"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database/adapters"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
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

	fmt.Println("Unknown Environment")
}

func (cnf *Config) setupAWS() {
	username := os.Getenv("RDS_USERNAME")
	password := os.Getenv("RDS_PASSWORD")
	database := os.Getenv("RDS_DB_NAME")
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
