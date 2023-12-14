package models

// type Avatar struct {
// 	gorm.Model
// 	Name          string
// 	Slug          string
// 	LLMID         sql.NullInt64
// 	LLM           LLM
// 	DefaultPrimer string
// 	Primer        string
// 	IsPublic      bool
// 	Roles         []Role         `gorm:"foreignKey:AvatarID;"`
// 	ActiveAgents  []ActiveAgents `gorm:"foreignKey:AvatarID;"`
// 	ActiveTools   []ActiveTool   `gorm:"foreignKey:AvatarID;"`
// 	ActivePlugins []ActivePlugin `gorm:"foreignKey:AvatarID;"`
// }

// type ActiveAgents struct {
// 	gorm.Model
// 	IsActive bool
// 	IsPublic bool
// 	Primer   string
// 	AvatarID sql.NullInt64
// 	AgentID  sql.NullInt64
// 	Avatar   *Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// 	Agent    *Agent  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }

// type ActiveTool struct {
// 	gorm.Model
// 	IsActive bool
// 	IsPublic bool
// 	Token    string
// 	AvatarID sql.NullInt64
// 	ToolID   sql.NullInt64
// 	Avatar   Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// 	Tool     Tool   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }

// type ActivePlugin struct {
// 	gorm.Model
// 	IsActive bool
// 	IsPublic bool
// 	Token    string
// 	AvatarID sql.NullInt64
// 	PluginID sql.NullInt64
// 	Avatar   Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// 	Plugin   Plugin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }
