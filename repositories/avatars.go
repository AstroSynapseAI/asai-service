package repositories

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
)

type AvatarsRepository struct {
	Repo *gorm.Repository[models.Avatar]
}

func NewAvatarsRepository(db *database.Database) *AvatarsRepository {
	return &AvatarsRepository{
		Repo: gorm.NewRepository[models.Avatar](db, models.Avatar{}),
	}
}

func (avatar *AvatarsRepository) SaveAvatar(data *models.Avatar) error {
	return nil
}

func (avatar *AvatarsRepository) GetAgents(avatarID uint32) []models.Agent {
	return nil
}
