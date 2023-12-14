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
	Username string
	Password string
	Accounts []Account
	Roles    []AvatarRole `gorm:"foreignKey:UserID;"`
}

type Account struct {
	gorm.Model
	UserID    int
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	DOB       string
	Type      string
}

type AvatarRole struct {
	gorm.Model
	RoleID   sql.NullInt64
	UserID   sql.NullInt64
	AvatarID sql.NullInt64
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Avatar   Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// AVATAR
type Avatar struct {
	gorm.Model
	Name          string
	Slug          string
	LLMID         sql.NullInt64
	LLM           LLM
	DefaultPrimer string
	Primer        string
	IsPublic      bool
	Roles         []AvatarRole   `gorm:"foreignKey:AvatarID;"`
	ActiveAgents  []ActiveAgents `gorm:"foreignKey:AvatarID;"`
	ActiveTools   []ActiveTool   `gorm:"foreignKey:AvatarID;"`
	ActivePlugins []ActivePlugin `gorm:"foreignKey:AvatarID;"`
	Documents     []Document     `gorm:"foreignKey:AvatarID;"`
}

type ActiveAgents struct {
	gorm.Model
	IsActive bool
	IsPublic bool
	Primer   string
	AvatarID sql.NullInt64
	AgentID  sql.NullInt64
	Avatar   *Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Agent    *Agent  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ActiveTool struct {
	gorm.Model
	IsActive bool
	IsPublic bool
	Token    string
	AvatarID sql.NullInt64
	ToolID   sql.NullInt64
	Avatar   Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tool     Tool   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ActivePlugin struct {
	gorm.Model
	IsActive bool
	IsPublic bool
	Token    string
	AvatarID sql.NullInt64
	PluginID sql.NullInt64
	Avatar   Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Plugin   Plugin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// AGENT
type Agent struct {
	gorm.Model
	Name         string
	Description  string
	Slug         string
	Primer       string
	ActiveAgents []ActiveAgents `gorm:"foreignKey:AgentID;"`
	AgentTool    []AgentTool    `gorm:"foreignKey:AgentID;"`
}

type AgentTool struct {
	gorm.Model
	IsActive bool
	IsPublic bool
	Token    string
	AgentID  sql.NullInt64
	ToolID   sql.NullInt64
	Agent    Agent `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tool     Tool  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// DISCRETE MODELS

type Tool struct {
	gorm.Model
	Name        string
	Description string
	Slug        string
	Token       string
	ActiveTools []ActiveTool `gorm:"foreignKey:ToolID;"`
	AgentTools  []AgentTool  `gorm:"foreignKey:ToolID;"`
}

type Plugin struct {
	gorm.Model
	Name          string
	Slug          string
	ActivePlugins []ActivePlugin `gorm:"foreignKey:PluginID;"`
}

type LLM struct {
	gorm.Model
	Name        string
	Description string
	Slug        string
	Avatars     []Avatar
	Agents      []Agent
}

type Role struct {
	gorm.Model
	Name        string
	Description string
	Permission  string
	Slug        string
	AvatarRoles []AvatarRole `gorm:"foreignKey:RoleID;"`
}

type Document struct {
	gorm.Model
	Path        string
	Name        string
	Description string
	AvatarID    sql.NullInt64
	Avatar      Avatar
}
