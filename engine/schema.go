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
	GetAgents() []AgentConfig
	GetTools() []ToolConfig
	GetPlugins() []PluginConfig
}

type AgentConfig interface {
	GetAgentName() string
	GetAgentSlug() string
	GetAgentLLM() llms.LanguageModel
	GetAgentPrimer() string
	IsAgentPublic() bool
	IsAgentActive() bool
	GetAgentTools() []ToolConfig
}

type ToolConfig interface {
	GetName() string
	GetToken() string
	IsPublic() bool
	IsActive() bool
}

type PluginConfig interface {
	GetName() string
	GetToken() string
	IsActive() bool
	IsPublic() bool
}
