package app

import (
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"gopkg.in/yaml.v2"
)

var CONFIG *Config

type ServerAdapter interface {
	Run() error
}

type Config struct {
	LLM        llms.LanguageModel
	ENV        string
	DSN        string
	MemorySize int
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
		cnf.setupLocalDev()
		return
	}

	if cnf.ENV == "HEROKU DEV" {
		cnf.setupHeroku()
		return
	}
}

func (cnf *Config) setupHeroku() {
	var err error
	cnf.LLM, err = openai.NewChat(openai.WithModel("gpt-4"))
	cnf.MemorySize = 4048

	if err != nil {
		fmt.Println("Error creating default LLM:", err)
		return
	}
}

func (cnf *Config) setupLocalDev() {

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

	cnf.LLM, err = openai.NewChat(openai.WithModel("gpt-4"))
	cnf.MemorySize = 20048

	// cnf.MemorySize = 4024
	// cnf.LLM, err = ollama.New(
	// 	ollama.WithModel("mistral"),
	// 	ollama.WithServerURL("http://host.docker.internal:11434/"),
	// )
	if err != nil {
		fmt.Println("Error creating default LLM:", err)
		return
	}
}
