package models

import (
	"database/sql"

	"gorm.io/gorm"
)

// I keep my models in a single file when I need to work on table relationships,
// makes it easier to manage when you have overview of all models in a single file.

// USER
type User struct {
	gorm.Model
	Username    string       `json:"username,omitempty"`
	Password    string       `json:"password"`
	ApiToken    string       `json:"api_token"`
	InviteToken string       `json:"invite_token"`
	Accounts    []Account    `json:"accounts"`
	Roles       []AvatarRole `gorm:"foreignKey:UserID;" json:"roles"`
}

type Account struct {
	gorm.Model
	UserID    int    `json:"user_id"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	DOB       string `json:"dob,omitempty"`
	Type      string `json:"type,omitempty"`
}

type AvatarRole struct {
	gorm.Model
	RoleID   sql.NullInt64 `json:"role_id"`
	UserID   sql.NullInt64 `json:"user_id"`
	AvatarID sql.NullInt64 `json:"avatar_id"`
	Role     Role          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	User     User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Avatar   Avatar        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar"`
}

// AVATAR
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

// AGENT
type Agent struct {
	gorm.Model
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	Slug         string        `json:"slug,omitempty"`
	Primer       string        `json:"primer,omitempty"`
	ActiveAgents []ActiveAgent `gorm:"foreignKey:AgentID;" json:"active_agents"`
	AgentTool    []AgentTool   `gorm:"foreignKey:AgentID;" json:"agent_tool"`
}

type AgentTool struct {
	gorm.Model
	IsActive bool          `json:"is_active,omitempty"`
	IsPublic bool          `json:"is_public,omitempty"`
	Token    string        `json:"token,omitempty"`
	AgentID  sql.NullInt64 `json:"agent_id"`
	ToolID   sql.NullInt64 `json:"tool_id"`
	Agent    Agent         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"agent"`
	Tool     Tool          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tool"`
}

// DISCRETE MODELS

// type Tool struct {
// 	gorm.Model
// 	Name        string `json:"name,omitempty"`
// 	Description string `json:"description,omitempty"`
// 	Slug        string `json:"slug,omitempty"`
// 	Token       string `json:"token,omitempty"`
// 	ActiveTools []ActiveTool `gorm:"foreignKey:ToolID;" json:"active_tools"`
// 	AgentTools  []AgentTool  `gorm:"foreignKey:ToolID;" json:"agent_tools"`
// }

type Plugin struct {
	gorm.Model
	Name          string         `json:"name,omitempty"`
	Slug          string         `json:"slug,omitempty"`
	ActivePlugins []ActivePlugin `gorm:"foreignKey:PluginID;" json:"active_plugins"`
}

type LLM struct {
	gorm.Model
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	Slug         string        `json:"slug,omitempty"`
	Avatars      []Avatar      `gorm:"foreignKey:LLMID;" json:"avatars"`
	ActiveAgents []ActiveAgent `gorm:"foreignKey:LLMID;" json:"active_agents"`
}

type Role struct {
	gorm.Model
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Permission  string       `json:"permission,omitempty"`
	Slug        string       `json:"slug,omitempty"`
	AvatarRoles []AvatarRole `gorm:"foreignKey:RoleID;" json:"avatar_roles"`
}

type Document struct {
	gorm.Model
	Path        string        `json:"path,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	AvatarID    sql.NullInt64 `json:"avatar_id"`
	Avatar      Avatar        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar"`
}
