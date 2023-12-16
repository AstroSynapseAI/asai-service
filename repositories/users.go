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

func (user *UsersRepository) Login(username string, password string) bool {
	var result *models.User
	user.Repo.DB.Where("username = ? AND password = ?", username, password).First(&result)
	if result != nil {
		return true
	}
	return false
}

func (user *UsersRepository) Register(username string, password string) bool {
	return true
}

func (user *UsersRepository) GetByUsername(username string) *models.User {
	var result *models.User
	user.Repo.DB.Where("username = ?", username).First(&result)

	return result
}

func (user *UsersRepository) GetUserAccount(userID uint32, accountID uint32) *models.Account {
	var result *models.Account
	user.Repo.DB.Where("user_id = ? AND id = ?", userID, accountID).First(&result)

	return result
}

func (user *UsersRepository) GetUserAvatar(userID uint32) *models.Avatar {
	var role *models.AvatarRole
	
	err := user.Repo.DB.Where("user_id = ? and name = ?", userID, "owner").First(&role).Error
	if err != nil {
		return nil
	}

	return &role.Avatar

}
