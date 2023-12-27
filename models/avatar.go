package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Avatar struct {
	gorm.Model
	Name          string         `json:"name,omitempty"`
	Slug          string         `json:"slug,omitempty"`
	LLMID         sql.NullInt64  `json:"llm_id"`
	LLM           LLM            `json:"llm"`
	DefaultPrimer string         `json:"default_primer,omitempty"`
	Primer        string         `json:"primer,omitempty"`
	IsPublic      bool           `json:"is_public,omitempty"`
	Roles         []AvatarRole   `gorm:"foreignKey:AvatarID;" json:"roles"`
	ActiveAgents  []ActiveAgent  `gorm:"foreignKey:AvatarID;" json:"active_agents"`
	ActiveTools   []ActiveTool   `gorm:"foreignKey:AvatarID;" json:"active_tools"`
	ActivePlugins []ActivePlugin `gorm:"foreignKey:AvatarID;" json:"active_plugins"`
	Documents     []Document     `gorm:"foreignKey:AvatarID;" json:"documents"`
}

type ActiveAgent struct {
	gorm.Model
	IsActive bool          `json:"is_active,omitempty"`
	IsPublic bool          `json:"is_public,omitempty"`
	Primer   string        `json:"primer,omitempty"`
	LLMID    sql.NullInt64 `json:"llm_id"`
	AvatarID sql.NullInt64 `json:"avatar_id"`
	AgentID  sql.NullInt64 `json:"agent_id"`
	LLM      LLM           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"llm"`
	Avatar   Avatar        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar"`
	Agent    Agent         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"agent"`
}

type ActiveTool struct {
	gorm.Model
	IsActive bool          `json:"is_active,omitempty"`
	IsPublic bool          `json:"is_public,omitempty"`
	Token    string        `json:"token,omitempty"`
	AvatarID sql.NullInt64 `json:"avatar_id"`
	ToolID   sql.NullInt64 `json:"tool_id"`
	Avatar   Avatar        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar"`
	Tool     Tool          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tool"`
}

type ActivePlugin struct {
	gorm.Model
	IsActive bool          `json:"is_active,omitempty"`
	IsPublic bool          `json:"is_public,omitempty"`
	Token    string        `json:"token,omitempty"`
	AvatarID sql.NullInt64 `json:"avatar_id"`
	PluginID sql.NullInt64 `json:"plugin_id"`
	Avatar   Avatar        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar"`
	Plugin   Plugin        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"plugin"`
}
