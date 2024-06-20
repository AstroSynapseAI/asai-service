package repositories

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/GoLangWebSDK/crud/database"
	"gorm.io/gorm"
)

type ApiRepository struct {
	DB   *database.Database
	Gorm *gorm.DB
}

func NewApiRepository(db *database.Database) *ApiRepository {
	return &ApiRepository{
		DB:   db,
		Gorm: db.Adapter.Gorm(),
	}
}

func (repo *ApiRepository) GetChatHistory(ID string) *models.ChatHistory {
	var history *models.ChatHistory

	err := repo.Gorm.Where(models.ChatHistory{SessionID: ID}).Find(&history).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return history
}
