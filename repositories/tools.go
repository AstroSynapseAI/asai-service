package repositories

import (
	"database/sql"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	db "gorm.io/gorm"
)

type ToolsRepository struct {
	Repo   *gorm.Repository[models.Tool]
	Active *gorm.Repository[models.ActiveTool]
}

func NewToolsRepository(db *database.Database) *ToolsRepository {
	return &ToolsRepository{
		Repo:   gorm.NewRepository[models.Tool](db, models.Tool{}),
		Active: gorm.NewRepository[models.ActiveTool](db, models.ActiveTool{}),
	}
}

func (tool *ToolsRepository) SaveActiveTool(avatarData models.ActiveTool) error {
	result := tool.Active.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tool *ToolsRepository) ToggleActiveTool(avatarID uint, toolID uint, active bool) error {
	var activeTool models.ActiveTool

	result := tool.Active.DB.Where("avatar_id = ? AND tool_id = ?", avatarID, toolID).First(&activeTool)
	if result.Error == db.ErrRecordNotFound {
		activeTool.ToolID = sql.NullInt64{Int64: int64(toolID), Valid: true}
		activeTool.AvatarID = sql.NullInt64{Int64: int64(avatarID), Valid: true}
	}

	activeTool.IsActive = active

	result = tool.Active.DB.Save(&activeTool)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (tool *ToolsRepository) GetAvatarTools(avatarID uint) []models.ActiveTool {
	where := "avatar_id = ?"
	return tool.GetActiveTools(where, avatarID)
}

func (tool *ToolsRepository) GetAvatarTool(id uint, avatarId uint) (models.ActiveTool, error) {
	where := "avatar_id = ? AND tool_id = ?"
	return tool.GetActiveTool(where, id, avatarId)
}

func (tool *ToolsRepository) GetAgentTools(agentID uint) []models.ActiveTool {
	where := "agent_id = ?"
	return tool.GetActiveTools(where, agentID)
}

func (tool *ToolsRepository) GetAgentTool(agentID uint, toolID uint) (models.ActiveTool, error) {
	where := "agent_id = ? AND tool_id = ?"
	return tool.GetActiveTool(where, agentID, toolID)
}

func (tool *ToolsRepository) GetActiveTools(where string, ID uint) []models.ActiveTool {
	query := tool.Active.DB
	query = query.Preload("Tool")
	query = query.Where(where, ID)

	var activeTools []models.ActiveTool

	result := query.Find(&activeTools)
	if result.Error != nil {
		return []models.ActiveTool{}
	}

	return activeTools
}

func (tool *ToolsRepository) GetActiveTool(where string, ID uint, toolID uint) (models.ActiveTool, error) {
	query := tool.Active.DB
	query = query.Preload("Tool")
	query = query.Where(where, ID, toolID)

	var activeTool models.ActiveTool

	result := query.First(&activeTool)
	if result.Error != nil {
		return models.ActiveTool{}, result.Error
	}

	return activeTool, nil
}
