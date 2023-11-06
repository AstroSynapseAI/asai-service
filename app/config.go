package app

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var CONFIG *Config

type ServerAdapter interface {
	Run() error
}

type Config struct {
	ENV string
	DSN string
}

func NewConfig() *Config {
	config := &Config{
		ENV: os.Getenv("ENVIRONMENT"),
	}

	CONFIG = config
	return config
}

func (cnf *Config) RunServer(server ServerAdapter) error {
	return server.Run()
}

func (cnf *Config) InitDB() {
	if cnf.ENV == "HEROKU DEV" {
		cnf.DSN = os.Getenv("DATABASE_URL")
		return
	}
	cnf.DSN = "postgresql://asai-admin:asai-password@asai-db:5432/asai-db"
}

func (cnf *Config) LoadEnvironment() {
	fmt.Println("Current Environment:", cnf.ENV)
	if cnf.ENV == "LOCAL DEV" {
		setupLocalDev()
		return
	}

}

func setupLocalDev() {
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
