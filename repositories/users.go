package repositories

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

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

func (user *UsersRepository) Login(username string, password string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("username = ? AND password = ?", username, password).First(&record).Error
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}

func (user *UsersRepository) Register(username string, password string) (models.User, error) {
	token, err := user.GenerateToken(64)
	if err != nil {
		return models.User{}, err
	}

	_, err = user.GetByUsername(username)
	if err == nil {
		return models.User{}, fmt.Errorf("user already exists")
	}

	record := models.User{
		Username: username,
		Password: password,
		ApiToken: token,
	}

	newUser, err := user.Repo.Create(record)
	if err != nil {
		return record, err
	}

	return newUser, nil
}

func (user *UsersRepository) CreateInvite() error {
	token, err := user.GenerateToken(64)
	if err != nil {
		return err
	}

	record := models.User{
		InviteToken: token,
	}

	_, err = user.Repo.Create(record)
	if err != nil {
		return err
	}

	return nil
}

func (user *UsersRepository) ConfirmInvite(username string, password string, token string) (models.User, error) {
	invitedUser, err := user.GetByInviteToken(token)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid invite token")
	}

	_, err = user.GetByUsername(username)
	if err == nil {
		return models.User{}, fmt.Errorf("user already exists")
	}

	apiToken, err := user.GenerateToken(64)
	if err != nil {
		return models.User{}, err
	}

	invitedUser.Username = username
	invitedUser.Password = password
	invitedUser.ApiToken = apiToken

	_, err = user.Repo.Update(invitedUser.ID, invitedUser)
	if err != nil {
		return models.User{}, err
	}

	return invitedUser, nil
}

func (user *UsersRepository) GetByUsername(username string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("username = ?", username).First(&record).Error
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}

func (user *UsersRepository) GetByAPIToken(token string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("api_token = ?", token).First(&record).Error
	if err != nil {
		return models.User{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetByInviteToken(token string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("invite_token = ?", token).First(&record).Error
	if err != nil {
		return models.User{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetUserAccount(userID uint, accountID uint) (models.Account, error) {
	var record models.Account
	err := user.Repo.DB.Where("user_id = ? AND id = ?", userID, accountID).First(&record).Error
	if err != nil {
		return models.Account{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetUserAvatar(userID uint) (models.Avatar, error) {
	query := user.Repo.DB
	query = query.Preload("Avatar").Preload("Avatar.Documents").Preload("Avatar.LLM")
	query = query.Preload("Avatar.ActiveAgents").Preload("Avatar.ActiveAgents.Agent")
	query = query.Preload("Avatar.ActiveTools").Preload("Avatar.ActiveTools.Tool")
	query = query.Preload("Avatar.ActivePlugins").Preload("Avatar.ActivePlugins.Plugin")
	query = query.Where("user_id = ? and role_id = ?", userID, 1)

	var record models.AvatarRole
	err := query.First(&record).Error
	if err != nil {
		return models.Avatar{}, err
	}

	return record.Avatar, nil
}

func (user *UsersRepository) GenerateToken(n int) (string, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}
