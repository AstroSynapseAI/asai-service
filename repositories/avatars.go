package repositories

import (
	"database/sql"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"

	db "gorm.io/gorm"
)

type AvatarsRepository struct {
	Repo *gorm.Repository[models.Avatar]
}

func NewAvatarsRepository(db *database.Database) *AvatarsRepository {
	return &AvatarsRepository{
		Repo: gorm.NewRepository[models.Avatar](db, models.Avatar{}),
	}
}

func (avatar *AvatarsRepository) Create(userID uint, data models.Avatar) (models.Avatar, error) {
	avatarRecord, err := avatar.Repo.Create(data)
	if err != nil {
		return models.Avatar{}, err
	}

	var user models.User
	result := avatar.Repo.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return models.Avatar{}, result.Error
	}

	var role models.Role
	result = avatar.Repo.DB.Where("permission = ?", "owner").First(&role)
	if result.Error != nil {
		return models.Avatar{}, result.Error
	}

	userRole := models.AvatarRole{
		User:   user,
		Role:   role,
		Avatar: avatarRecord,
	}

	result = avatar.Repo.DB.Create(&userRole)
	if result.Error != nil {
		return models.Avatar{}, result.Error
	}

	record, err := avatar.Fetch(avatarRecord.ID)
	if err != nil {
		return models.Avatar{}, err
	}

	return record, nil
}

func (avatar *AvatarsRepository) Update(ID uint, data models.Avatar) (models.Avatar, error) {
	avatarRecord, err := avatar.Repo.Update(ID, data)
	if err != nil {
		return models.Avatar{}, err
	}

	record, err := avatar.Fetch(avatarRecord.ID)
	if err != nil {
		return models.Avatar{}, err
	}

	return record, nil
}

func (avatar *AvatarsRepository) Fetch(ID uint) (models.Avatar, error) {
	query := avatar.Repo.DB
	query = query.Preload("LLM")
	query = query.Preload("Documents")
	query = query.Preload("Roles").Preload("Roles.Role").Preload("Roles.User")
	query = query.Preload("ActiveAgents").Preload("ActiveAgents.Agent")
	query = query.Preload("ActiveTools").Preload("ActiveTools.Tool")
	query = query.Preload("ActivePlugins").Preload("ActivePlugins.Plugin")

	var record models.Avatar
	result := query.First(&record, ID)
	if result.Error != nil {
		return models.Avatar{}, result.Error
	}

	return record, nil
}

func (avatar *AvatarsRepository) SaveActiveAgent(avatarData models.ActiveAgent) error {
	result := avatar.Repo.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (avatar *AvatarsRepository) ToggleActiveAgent(avatarID uint, agentID uint, active bool) error {
	var activeAgent models.ActiveAgent

	result := avatar.Repo.DB.Where("avatar_id = ? AND agent_id = ?", avatarID, agentID).First(&activeAgent)
	if result.Error == db.ErrRecordNotFound {
		activeAgent.AgentID = sql.NullInt64{Int64: int64(agentID), Valid: true}
		activeAgent.AvatarID = sql.NullInt64{Int64: int64(avatarID), Valid: true}
	}

	activeAgent.IsActive = active
	result = avatar.Repo.DB.Save(&activeAgent)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (avatar *AvatarsRepository) GetActiveAgents(avatarID uint) []models.ActiveAgent {
	query := avatar.Repo.DB
	query = query.Preload("Agent").Preload("LLM")
	query = query.Where("avatar_id = ?", avatarID)

	var activeAgents []models.ActiveAgent

	result := query.Find(&activeAgents)
	if result.Error != nil {
		return []models.ActiveAgent{}
	}

	return activeAgents
}

func (avatar *AvatarsRepository) GetActiveAgent(avatarID uint, agentID uint) (models.ActiveAgent, error) {
	query := avatar.Repo.DB
	query = query.Preload("Agent").Preload("LLM")
	query = query.Where("avatar_id = ? AND agent_id = ?", avatarID, agentID)

	var activeAgent models.ActiveAgent
	result := query.First(&activeAgent)
	if result.Error != nil {
		return models.ActiveAgent{}, result.Error
	}

	return activeAgent, nil
}

func (avatar *AvatarsRepository) SaveActiveTool(avatarData models.ActiveTool) error {
	result := avatar.Repo.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (avatar *AvatarsRepository) ToggleActiveTool(avatarID uint, toolID uint, active bool) error {
	var activeTool models.ActiveTool

	result := avatar.Repo.DB.Where("avatar_id = ? AND tool_id = ?", avatarID, toolID).First(&activeTool)
	if result.Error == db.ErrRecordNotFound {
		activeTool.ToolID = sql.NullInt64{Int64: int64(toolID), Valid: true}
		activeTool.AvatarID = sql.NullInt64{Int64: int64(avatarID), Valid: true}
	}

	activeTool.IsActive = active

	result = avatar.Repo.DB.Save(&activeTool)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (avatar *AvatarsRepository) GetActiveTools(avatarID uint) []models.ActiveTool {
	query := avatar.Repo.DB
	query = query.Preload("Tool")
	query = query.Where("avatar_id = ?", avatarID)

	var activeTools []models.ActiveTool

	result := query.Find(&activeTools)
	if result.Error != nil {
		return []models.ActiveTool{}
	}

	return activeTools
}

func (avatar *AvatarsRepository) GetActiveTool(avatarID uint, toolID uint) (models.ActiveTool, error) {
	query := avatar.Repo.DB
	query = query.Preload("Tool")
	query = query.Where("avatar_id = ? AND tool_id = ?", avatarID, toolID)

	var activeTool models.ActiveTool

	result := query.First(&activeTool)
	if result.Error != nil {
		return models.ActiveTool{}, result.Error
	}

	return activeTool, nil
}

func (avatar *AvatarsRepository) SaveActivePlugin(avatarData models.ActivePlugin) error {
	result := avatar.Repo.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (avatar *AvatarsRepository) ToggleActivePlugin(avatarID uint, pluginID uint, active bool) error {
	var activePlugin models.ActivePlugin

	result := avatar.Repo.DB.Where("avatar_id = ? AND plugin_id = ?", avatarID, pluginID).First(&activePlugin)
	if result.Error == db.ErrRecordNotFound {
		activePlugin.PluginID = sql.NullInt64{Int64: int64(pluginID), Valid: true}
		activePlugin.AvatarID = sql.NullInt64{Int64: int64(avatarID), Valid: true}
	}

	activePlugin.IsActive = active
	result = avatar.Repo.DB.Save(&activePlugin)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (avatar *AvatarsRepository) GetActivePlugins(avatarID uint) []models.ActivePlugin {
	query := avatar.Repo.DB
	query = query.Preload("Plugin")
	query = query.Where("avatar_id = ?", avatarID)

	var activePlugins []models.ActivePlugin
	result := query.Find(&activePlugins)
	if result.Error != nil {
		return []models.ActivePlugin{}
	}

	return activePlugins
}

func (avatar *AvatarsRepository) GetActivePlugin(avatarID uint, pluginID uint) (models.ActivePlugin, error) {
	query := avatar.Repo.DB
	query = query.Preload("Plugin")
	query = query.Where("avatar_id = ? AND plugin_id = ?", avatarID, pluginID)

	var activePlugin models.ActivePlugin

	result := query.First(&activePlugin)
	if result.Error != nil {
		return models.ActivePlugin{}, result.Error
	}

	return activePlugin, nil
}

func (avatar *AvatarsRepository) SaveActiveLLM(avatarData models.ActiveLLM) error {
	result := avatar.Repo.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (avatar *AvatarsRepository) ToggleActiveLLM(avatarID uint, LLMID uint, active bool) error {
	var activeLLM models.ActiveLLM

	result := avatar.Repo.DB.Where("avatar_id = ? AND LLM_id = ?", avatarID, LLMID).First(&activeLLM)
	if result.Error == db.ErrRecordNotFound {
		activeLLM.LLMID = sql.NullInt64{Int64: int64(LLMID), Valid: true}
		activeLLM.AvatarID = sql.NullInt64{Int64: int64(avatarID), Valid: true}
	}

	activeLLM.IsActive = active

	result = avatar.Repo.DB.Save(&activeLLM)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (avatar *AvatarsRepository) GetActiveLLMs(avatarID uint) []models.ActiveLLM {
	query := avatar.Repo.DB
	query = query.Preload("LLM")
	query = query.Where("avatar_id = ?", avatarID)

	var activeLLMs []models.ActiveLLM

	result := query.Find(&activeLLMs)
	if result.Error != nil {
		return []models.ActiveLLM{}
	}

	return activeLLMs
}

func (avatar *AvatarsRepository) GetActiveLLM(avatarID uint, LLMID uint) (models.ActiveLLM, error) {
	query := avatar.Repo.DB
	query = query.Preload("LLM")
	query = query.Where("avatar_id = ? AND LLM_id = ?", avatarID, LLMID)

	var activeLLM models.ActiveLLM

	result := query.First(&activeLLM)
	if result.Error != nil {
		return models.ActiveLLM{}, result.Error
	}

	return activeLLM, nil
}

func (avatar *AvatarsRepository) GetDocuments(avatarID uint) []models.Document {
	var docs []models.Document
	result := avatar.Repo.DB.Where("avatar_id = ?", avatarID).Find(&docs)
	if result.Error != nil {
		return []models.Document{}
	}

	return docs
}
