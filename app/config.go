package app

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadEnvironment() {
	env := os.Getenv("ENVIRONMENT")

	fmt.Println("Current Environment:", env)

	if env == "HEROKU DEV" {
		setupHerokuDev()
		return
	}

	if env == "LOCAL DEV" {
		setupLocalDev()
		return
	}
}

func setupHerokuDev() {

}

// setupLocalDev initializes the local development environment by setting the necessary environment variables.
//
// No parameters.
// No return values.
func setupLocalDev() {
	var Config struct {
		OpenAPIKey    string `yaml:"open_api_key"`
		SerpAPIKey    string `yaml:"serpapi_api_key"`
		DiscordApiKey string `yaml:"discord_api_key"`
	}

	keys, err := os.ReadFile("./config/keys.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(keys, &Config)
	if err != nil {
		panic(err)
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

func SetupPostgreDSN() string {
	if os.Getenv("ENVIRONMENT") == "HEROKU DEV" {
		return os.Getenv("DATABASE_URL")
	}

	dsn := "postgresql://asai-admin:asai-password@asai-db:5432/asai-db"
	return dsn
}
