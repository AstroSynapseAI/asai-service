package repositories

import (
	"github.com/AstroSynapseAI/asai-service/models"
	"github.com/GoLangWebSDK/crud/database"
	"github.com/GoLangWebSDK/crud/orms/gorm"
	db "gorm.io/gorm"
)

type ToolsRepository struct {
	Repo   *gorm.Repository[models.Tool]
	Active *gorm.Repository[models.ActiveTool]
	Agent  *gorm.Repository[models.ActiveAgentTool]
}

func NewToolsRepository(db *database.Database) *ToolsRepository {
	return &ToolsRepository{
		Repo:   gorm.NewRepository[models.Tool](db, models.Tool{}),
		Active: gorm.NewRepository[models.ActiveTool](db, models.ActiveTool{}),
		Agent:  gorm.NewRepository[models.ActiveAgentTool](db, models.ActiveAgentTool{}),
	}
}

func (tool *ToolsRepository) SaveActiveTool(avatarData models.ActiveTool) error {
	result := tool.Active.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tool *ToolsRepository) SaveAgentTool(avatarData models.ActiveAgentTool) error {
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

func (tool *ToolsRepository) ToggleAgentTool(activeAgentID uint, toolID uint, active bool) error {
	var agentTool models.ActiveAgentTool

	result := tool.Agent.DB.Where("agent_id = ? AND tool_id = ?", activeAgentID, toolID).First(&agentTool)
	if result.Error == db.ErrRecordNotFound {
		agentTool.ToolID = toolID
		agentTool.ActiveAgentID = activeAgentID
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

func (tool *ToolsRepository) GetAgentTools(agentID uint) []models.ActiveAgentTool {
	query := tool.Active.DB
	query = query.Preload("Tool")
	query = query.Where("active_agent_id = ?", agentID)

	var agentTools []models.ActiveAgentTool

	result := query.Find(&agentTools)
	if result.Error != nil {
		return []models.ActiveAgentTool{}
	}

	return agentTools
}

func (tool *ToolsRepository) GetAgentTool(agentID uint, toolID uint) (models.ActiveAgentTool, error) {
	query := tool.Active.DB
	query = query.Preload("Tool")
	query = query.Where("active_agent_id = ? AND tool_id = ?", agentID, toolID)

	var agentTool models.ActiveAgentTool

	result := query.First(&agentTool)
	if result.Error != nil {
		return models.ActiveAgentTool{}, result.Error
	}

	return agentTool, nil
}
