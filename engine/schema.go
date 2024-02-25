package engine

import (
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/tmc/langchaingo/llms"
)

type AvatarConfig interface {
	LoadConfig(userID uint)
	GetDB() *database.Database
	GetAvatarName() string
	GetAvatarLLM() llms.LanguageModel
	GetAvatarPrimer() string
	GetAvatarMemorySize() int
	AvatarIsPublic() bool
	GetLLM() LLMConfig
	GetAgents() []AgentConfig
	GetTools() []ToolConfig
	GetPlugins() []PluginConfig
}

type LLMConfig interface {
	GetAPI() string
	GetToken() string
}

type AgentConfig interface {
	GetAgentName(agentID string) string
	GetAgentModel(agentID string) *llms.LLM
	GetAgentPrimer(agentID string) string
	IsAgentPublic(agentID string) bool
	IsAgentActive(agentID string) bool
	GetAgentTools(agentID string) []any
}

type ToolConfig interface {
	GetName() string
	GetToken() string
}

type PluginConfig interface {
	GetName() string
	GetToken() string
}
