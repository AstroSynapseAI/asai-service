package repositories

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	db "gorm.io/gorm"
)

type ToolsRepository struct {
	Repo   *gorm.Repository[models.Tool]
	Active *gorm.Repository[models.ActiveTool]
	Agent  *gorm.Repository[models.AgentTool]
}

func NewToolsRepository(db *database.Database) *ToolsRepository {
	return &ToolsRepository{
		Repo:   gorm.NewRepository[models.Tool](db, models.Tool{}),
		Active: gorm.NewRepository[models.ActiveTool](db, models.ActiveTool{}),
		Agent:  gorm.NewRepository[models.AgentTool](db, models.AgentTool{}),
	}
}

func (tool *ToolsRepository) SaveActiveTool(avatarData models.ActiveTool) error {
	result := tool.Active.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tool *ToolsRepository) SaveAgentTool(avatarData models.AgentTool) error {
	result := tool.Agent.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tool *ToolsRepository) ToggleAvatarTool(avatarID uint, toolID uint, active bool) error {
	var activeTool models.ActiveTool

	result := tool.Active.DB.Where("avatar_id = ? AND tool_id = ?", avatarID, toolID).First(&activeTool)
	if result.Error == db.ErrRecordNotFound {
		activeTool.ToolID = toolID
		activeTool.AvatarID = avatarID
	}

	activeTool.IsActive = active

	result = tool.Active.DB.Save(&activeTool)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (tool *ToolsRepository) ToggleAgentTool(agentID uint, toolID uint, active bool) error {
	var agentTool models.AgentTool

	result := tool.Agent.DB.Where("agent_id = ? AND tool_id = ?", agentID, toolID).First(&agentTool)
	if result.Error == db.ErrRecordNotFound {
		agentTool.ToolID = toolID
		agentTool.AgentID = agentID
	}

	agentTool.IsActive = active

	result = tool.Agent.DB.Save(&agentTool)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (tool *ToolsRepository) GetAvatarTools(avatarID uint) []models.ActiveTool {
	query := tool.Active.DB
	query = query.Preload("Tool")
	query = query.Where("avatar_id = ?", avatarID)

	var activeTools []models.ActiveTool

	result := query.Find(&activeTools)
	if result.Error != nil {
		return []models.ActiveTool{}
	}

	return activeTools
}

func (tool *ToolsRepository) GetAvatarTool(ID uint, avatarID uint) (models.ActiveTool, error) {
	query := tool.Active.DB
	query = query.Preload("Tool")
	query = query.Where("avatar_id = ? AND tool_id = ?", avatarID, ID)

	var activeTool models.ActiveTool

	result := query.First(&activeTool)
	if result.Error != nil {
		return models.ActiveTool{}, result.Error
	}

	return activeTool, nil
}

func (tool *ToolsRepository) GetAgentTools(agentID uint) []models.AgentTool {
	query := tool.Active.DB
	query = query.Preload("Tool")
	query = query.Where("agent_id = ?", agentID)

	var agentTools []models.AgentTool

	result := query.Find(&agentTools)
	if result.Error != nil {
		return []models.AgentTool{}
	}

	return agentTools
}

func (tool *ToolsRepository) GetAgentTool(agentID uint, toolID uint) (models.AgentTool, error) {
	query := tool.Active.DB
	query = query.Preload("Tool")
	query = query.Where("agent_id = ? AND tool_id = ?", agentID, toolID)

	var agentTool models.AgentTool

	result := query.First(&agentTool)
	if result.Error != nil {
		return models.AgentTool{}, result.Error
	}

	return agentTool, nil
}
