package repositories

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/GoLangWebSDK/crud/database"
	"github.com/GoLangWebSDK/crud/orms/gorm"
)

type DocumentsRepository struct {
	Repo *gorm.Repository[models.Document]
}

func NewDocumentsRepository(db *database.Database) *DocumentsRepository {
	return &DocumentsRepository{
		Repo: gorm.NewRepository[models.Document](db, models.Document{}),
	}
}

func (document *DocumentsRepository) GetDocuments(avatarID uint) []models.Document {
	var docs []models.Document
	result := document.Repo.DB.Where("avatar_id = ?", avatarID).Find(&docs)
	if result.Error != nil {
		return []models.Document{}
	}

	return docs
}
