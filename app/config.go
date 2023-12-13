package app

// Ollama example for dockerized ollama, keep alive for now

// cnf.LLM, err = ollama.New(
// 	ollama.WithModel("mistral"),
// 	ollama.WithServerURL("http://host.docker.internal:11434/"),
// )

import (
	"fmt"
	"os"

	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database/adapters"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	"gopkg.in/yaml.v2"
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

	fmt.Println("Unknown Environment")
}

func (cnf *Config) setupHeroku() {
	cnf.DSN = os.Getenv("DATABASE_URL")
}

func (cnf *Config) setupLocalDev() {
	cnf.DSN = DefaultDevDSN

	var Config struct {
		OpenAPIKey    string `yaml:"open_api_key"`
		SerpAPIKey    string `yaml:"serpapi_api_key"`
		DiscordApiKey string `yaml:"discord_api_key"`
	}

	keys, err := os.ReadFile("./app/keys.yaml")
	if err != nil {
		fmt.Println("Error reading keys.yaml:", err)
		return
	}

	err = yaml.Unmarshal(keys, &Config)
	if err != nil {
		fmt.Println("Error unmarshalling keys.yaml:", err)
		return
	}

	// Set the Openai API key as env variable
	err = os.Setenv("OPENAI_API_KEY", Config.OpenAPIKey)
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}

	// Set the SerpAPI API key as env variable
	err = os.Setenv("SERPAPI_API_KEY", Config.SerpAPIKey)
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}

	//Set the Discord API key as env variable
	err = os.Setenv("DISCORD_API_KEY", Config.DiscordApiKey)
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}
}
