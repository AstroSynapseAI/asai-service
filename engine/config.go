package engine

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type Config struct {
	DB *database.Database
}

var _ AvatarConfig = (*Config)(nil)

func NewConfig(db *database.Database) *Config {
	return &Config{
		DB: db,
	}
}

func (cnf *Config) GetDB() *database.Database {
	return cnf.DB
}

func (cnf *Config) GetAvatarLLM() llms.LanguageModel {
	LLM, err := openai.NewChat(openai.WithModel("gpt-4"))
	if err != nil {
		fmt.Println("Error creating default LLM:", err)
		return nil
	}
	return LLM
}

func (cnf *Config) GetAvatarName() string {
	return "Asai"
}

func (cnf *Config) GetAvatarPrimer() string {
	return ""
}

func (cnf *Config) GetAvatarMemorySize() int {
	return 4048
}

func (cnf *Config) AvatarIsPublic() bool {
	return true
}

func (cnf *Config) GetAgents() []AgentConfig {
	return nil
}

func (cnf *Config) GetTools() []ToolConfig {
	return nil
}

func (cnf *Config) GetPlugins() []PluginConfig {
	return nil
}
