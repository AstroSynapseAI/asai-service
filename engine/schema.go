package engine

import "github.com/tmc/langchaingo/llms"

type EngineConfig interface {
	GetAvatarName() string
	GetAvatarModel() *llms.ChatLLM
	GetAvatarPrimer() string
	AvatarIsPublic() bool
	GetAgents() []any
	GetTools() []any
	GetPlugins() []any
}

type AgentConfig interface {
	GetName() string
	GetModel() *llms.LLM
	GetPrimer() string
	IsPublic() bool
	IsActive() bool
	GetTools() []any
}

type ToolConfig interface {
	GetName() string
	GetToken() string
}

type PluginConfig interface {
	GetName() string
	GetToken() string
}