package models

import (
	"gorm.io/gorm"
)

type Avatar struct {
	gorm.Model
	Name          string         `json:"name,omitempty"`
	Slug          string         `json:"slug,omitempty"`
	LLMID         uint           `json:"llm_id,omitempty"`
	LLM           LLM            `json:"llm,omitempty"`
	DefaultPrimer string         `json:"default_primer,omitempty"`
	Primer        string         `json:"primer,omitempty"`
	IsPublic      bool           `json:"is_public"`
	Roles         []AvatarRole   `gorm:"foreignKey:AvatarID;" json:"roles,omitempty"`
	ActiveAgents  []ActiveAgent  `gorm:"foreignKey:AvatarID;" json:"active_agents"`
	ActiveTools   []ActiveTool   `gorm:"foreignKey:AvatarID;" json:"active_tools"`
	ActivePlugins []ActivePlugin `gorm:"foreignKey:AvatarID;" json:"active_plugins"`
	ActiveLLMs    []ActiveLLM    `gorm:"foreignKey:AvatarID;" json:"active_llms"`
	Documents     []Document     `gorm:"foreignKey:AvatarID;" json:"documents"`
}

type ActiveAgent struct {
	gorm.Model
	IsActive         bool              `json:"is_active"`
	IsPublic         bool              `json:"is_public"`
	Primer           string            `json:"primer,omitempty"`
	LLMID            uint              `json:"llm_id,omitempty"`
	AvatarID         uint              `json:"avatar_id,omitempty"`
	AgentID          uint              `json:"agent_id,omitempty"`
	LLM              LLM               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"llm,omitempty"`
	Avatar           Avatar            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar,omitempty"`
	Agent            Agent             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"agent,omitempty"`
	ActiveAgentTools []ActiveAgentTool `gorm:"foreignKey:ActiveAgentID;" json:"active_agent_tools"`
}

type ActiveAgentTool struct {
	gorm.Model
	IsActive      bool        `json:"is_active"`
	IsPublic      bool        `json:"is_public"`
	Token         string      `json:"token,omitempty"`
	ToolID        uint        `json:"tool_id"`
	ActiveAgentID uint        `json:"active_agent_id"`
	Tool          Tool        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tool"`
	ActiveAgent   ActiveAgent `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"active_agent"`
}

type ActiveTool struct {
	gorm.Model
	IsActive bool   `json:"is_active"`
	IsPublic bool   `json:"is_public"`
	Token    string `json:"token,omitempty"`
	AvatarID uint   `json:"avatar_id,omitempty"`
	ToolID   uint   `json:"tool_id,omitempty"`
	Avatar   Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar,omitempty"`
	Tool     Tool   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tool,omitempty"`
}

type ActivePlugin struct {
	gorm.Model
	IsActive bool   `json:"is_active"`
	IsPublic bool   `json:"is_public"`
	Token    string `json:"token,omitempty"`
	AvatarID uint   `json:"avatar_id,omitempty"`
	PluginID uint   `json:"plugin_id,omitempty"`
	Avatar   Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar,omitempty"`
	Plugin   Plugin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"plugin,omitempty"`
}

type ActiveLLM struct {
	gorm.Model
	IsActive bool   `json:"is_active"`
	Token    string `json:"token,omitempty"`
	AvatarID uint   `json:"avatar_id,omitempty"`
	Avatar   Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar,omitempty"`
	LLM      LLM    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"llm,omitempty"`
	LLMID    uint   `json:"llm_id,omitempty"`
}
