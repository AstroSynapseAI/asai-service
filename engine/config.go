package engine

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type Config struct {
	DB     *database.Database
	Avatar *models.Avatar
}

var _ AvatarConfig = (*Config)(nil)

func NewConfig(db *database.Database) *Config {
	return &Config{
		DB: db,
	}
}

func (cnf *Config) LoadConfig(avatarID uint) {
	avatars := repositories.NewAvatarsRepository(cnf.DB)

	avatar, err := avatars.Fetch(avatarID)
	if err != nil {
		fmt.Println("Error loading avatar:", err)
		return
	}

	cnf.Avatar = &avatar
}

func (cnf *Config) GetDB() *database.Database {
	return cnf.DB
}

func (cnf *Config) GetAvatarLLM() llms.LanguageModel {
	avatarLLM := cnf.Avatar.LLM
	activeLLMs := cnf.Avatar.ActiveLLMs
	if len(activeLLMs) == 0 {
		fmt.Println("No active LLMs")
		return nil
	}

	switch avatarLLM.Slug {
	case "gpt-4":
		var activeLLM models.ActiveLLM
		// extract active llm where activeLLM.llmID == avatarLLM.ID
		for _, active := range activeLLMs {
			if active.LLM.ID == avatarLLM.ID {
				activeLLM = active
			}
		}

		LLM, err := openai.NewChat(
			openai.WithToken(activeLLM.Token),
			openai.WithModel("gpt-4"))
		if err != nil {
			fmt.Println("Error setting gpt-4:", err)
			return nil
		}
		return LLM
	default:
		fmt.Println("Unknown LLM:", avatarLLM.Slug)
		return nil
	}
}

func (cnf *Config) GetAvatarName() string {
	return cnf.Avatar.Name
}

func (cnf *Config) GetAvatarPrimer() string {
	return cnf.Avatar.Primer
}

func (cnf *Config) GetAvatarMemorySize() int {
	return 4048
}

func (cnf *Config) AvatarIsPublic() bool {
	return cnf.Avatar.IsPublic
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

// Active LLM Config
type ActiveLLM struct {
	DB  *database.Database
	LLM llms.LanguageModel
}

var _ LLMConfig = (*ActiveLLM)(nil)

func NewActiveLLM(db *database.Database) *ActiveLLM {
	return &ActiveLLM{
		DB: db,
	}
}

func (cnf *ActiveLLM) GetAPI() string {
	return ""
}

func (cnf *ActiveLLM) GetToken() string {
	return ""
}

// Active Agent Config
type ActiveAgent struct {
	DB          *database.Database
	ActiveAgent *models.ActiveAgent
}

var _ AgentConfig = (*ActiveAgent)(nil)

func NewActiveAgent(db *database.Database) *ActiveAgent {
	return &ActiveAgent{
		DB: db,
	}
}

func (cnf *ActiveAgent) GetAgentName(agentID string) string {
	return cnf.ActiveAgent.Agent.Name
}

func (cnf *ActiveAgent) GetAgentModel(agentID string) *llms.LLM {
	return nil
}

func (cnf *ActiveAgent) GetAgentPrimer(agentID string) string {
	return cnf.ActiveAgent.Primer
}

func (cnf *ActiveAgent) IsAgentPublic(agentID string) bool {
	return cnf.ActiveAgent.IsPublic
}

func (cnf *ActiveAgent) IsAgentActive(agentID string) bool {
	return cnf.ActiveAgent.IsActive
}

func (cnf *ActiveAgent) GetAgentTools(agentID string) []any {
	return nil
}

// Active Tool Config
type ActiveTool struct {
}

var _ ToolConfig = (*ActiveTool)(nil)

func NewActiveTool() *ActiveTool {
	return &ActiveTool{}
}

func (cnf *ActiveTool) GetName() string {
	return ""
}

func (cnf *ActiveTool) GetToken() string {
	return ""
}

// Active Plugin Config
type ActivePlugin struct {
}

var _ PluginConfig = (*ActivePlugin)(nil)

func NewActivePlugin() *ActivePlugin {
	return &ActivePlugin{}
}

func (cnf *ActivePlugin) GetName() string {
	return ""
}

func (cnf *ActivePlugin) GetToken() string {
	return ""
}
