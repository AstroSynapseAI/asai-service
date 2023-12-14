package models

// type Agent struct {
// 	gorm.Model
// 	ID           int
// 	Name         string
// 	Description  string
// 	Slug         string
// 	Primer       string
// 	ActiveAgents []ActiveAgents `gorm:"foreignKey:AgentID;"`
// 	AgentTools   []AgentTool    `gorm:"foreignKey:AgentID;"`
// }

// type AgentTool struct {
// 	gorm.Model
// 	IsActive bool
// 	IsPublic bool
// 	Token    string
// 	AgentID  int
// 	ToolID   int
// 	Agent    *Agent `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// 	Tool     *Tool  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }
