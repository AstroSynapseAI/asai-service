package repositories

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
)

type AgentsRepository struct {
	Repo *gorm.Repository[models.Agent]
}

func NewAgentsRepository(db *database.Database) *AgentsRepository {
	return &AgentsRepository{
		Repo: gorm.NewRepository[models.Agent](db, models.Agent{}),
	}
}