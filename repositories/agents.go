package repositories

import (
	"github.com/AstroSynapseAI/asai-service/models"
	"github.com/GoLangWebSDK/crud/database"
	"github.com/GoLangWebSDK/crud/orms/gorm"
	db "gorm.io/gorm"
)

type AgentsRepository struct {
	Repo   *gorm.Repository[models.Agent]
	Active *gorm.Repository[models.ActiveAgent]
}

func NewAgentsRepository(db *database.Database) *AgentsRepository {
	return &AgentsRepository{
		Repo:   gorm.NewRepository[models.Agent](db, models.Agent{}),
		Active: gorm.NewRepository[models.ActiveAgent](db, models.ActiveAgent{}),
	}
}

func (agent *AgentsRepository) SaveActiveAgent(agentData models.ActiveAgent) (models.ActiveAgent, error) {
	result := agent.Active.DB.Save(&agentData)
	if result.Error != nil {
		return models.ActiveAgent{}, result.Error
	}
	return agentData, nil
}

func (agent *AgentsRepository) ToggleActiveAgent(avatarID uint, agentID uint, active bool) error {
	var activeAgent models.ActiveAgent

	result := agent.Active.DB.Where("avatar_id = ? AND agent_id = ?", avatarID, agentID).First(&activeAgent)
	if result.Error == db.ErrRecordNotFound {
		activeAgent.AgentID = agentID
		activeAgent.AvatarID = avatarID
	}

	activeAgent.IsActive = active
	result = agent.Active.DB.Save(&activeAgent)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (agent *AgentsRepository) GetActiveAgents(avatarID uint) []models.ActiveAgent {
	query := agent.Active.DB
	query = query.Preload("Agent").Preload("LLM")
	query = query.Where("avatar_id = ?", avatarID)

	var activeAgents []models.ActiveAgent

	result := query.Find(&activeAgents)
	if result.Error != nil {
		return []models.ActiveAgent{}
	}

	return activeAgents
}

func (agent *AgentsRepository) GetActiveAgent(avatarID uint, agentID uint) (models.ActiveAgent, error) {
	query := agent.Active.DB
	query = query.Preload("Agent").Preload("LLM")
	query = query.Where("avatar_id = ? AND agent_id = ?", avatarID, agentID)

	var activeAgent models.ActiveAgent
	result := query.First(&activeAgent)
	if result.Error != nil {
		return models.ActiveAgent{}, result.Error
	}

	return activeAgent, nil
}

func (avatar *AvatarsRepository) SaveAgentTool(avatarData models.ActiveTool) error {
	result := avatar.Repo.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
