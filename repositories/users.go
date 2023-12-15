package repositories

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
)

type UsersRepository struct {
	Repo *gorm.Repository[models.User]
}

func NewUsersRepository(db *database.Database) *UsersRepository {
	return &UsersRepository{
		Repo: gorm.NewRepository[models.User](db, models.User{}),
	}
}

func (ctrl *UsersRepository) Login(username string, password string) bool {
	return true
}

func (ctrl *UsersRepository) Register(username string, password string) bool {
	return true
}

func (ctrl *UsersRepository) GetByUsername(username string) *models.User {
	return &models.User{}
}
