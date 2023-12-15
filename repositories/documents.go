package repositories

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
)

type DocumentsRepository struct {
	Repo *gorm.Repository[models.Document]
}

func NewDocumentsRepository(db *database.Database) *DocumentsRepository {
	return &DocumentsRepository{
		Repo: gorm.NewRepository[models.Document](db, models.Document{}),
	}
}
