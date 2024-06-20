package repositories

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/GoLangWebSDK/crud/database"
	"github.com/GoLangWebSDK/crud/orms/gorm"
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
		if err.Error() == "record not found" {
			err = fmt.Errorf("invalid username or password")
		}
		return models.User{}, err
	}

	query := user.Repo.DB.Preload("Roles").Preload("Roles.Role")
	query = query.Preload("Roles.Avatar")
	query = query.Preload("Accounts")

	err = query.First(&record, record.ID).Error
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

func (user *UsersRepository) CreateInvite(username string) (models.User, error) {
	token, err := user.GenerateToken(64)
	if err != nil {
		return models.User{}, err
	}

	record := models.User{
		InviteToken: token,
		Username:    username,
	}

	userRecord, err := user.Repo.Create(record)
	if err != nil {
		return models.User{}, err
	}

	return userRecord, nil
}

func (user *UsersRepository) CreateAndSendRecoveryToken(email string) (models.User, error) {

	account, err := user.GetAccountByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	recoveryToken, err := user.GenerateToken(64)
	if err != nil {
		return models.User{}, err
	}

	userRecord, err := user.GetUserByAccountID(account.UserID)
	if err != nil {
		return models.User{}, err
	}

	updatedUserRecord, err := user.InsertPasswordResetToken(userRecord.ID, recoveryToken, time.Now().Add(24*time.Hour))
	if err != nil {
		return models.User{}, err
	}

	return updatedUserRecord, nil
}

func (user *UsersRepository) CreateAndSendEmailConfirmation(id uint, email string) (string, error) {

	account, err := user.GetAccountByID(id)
	if err != nil {
		return "", err
	}

	token, err := user.GenerateToken(64)
	if err != nil {
		return "", err
	}

	account.Email = token
	user.SaveAccount(account)

	return token, err
}

func (user *UsersRepository) ConfirmInvite(username string, password string, token string) (models.User, error) {
	invitedUser, err := user.GetByInviteToken(token)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid invite token")
	}

	existingUser, err := user.GetByUsername(username)
	if err == nil && existingUser.ID != invitedUser.ID {
		return models.User{}, fmt.Errorf("user already exists")
	}

	fmt.Println("username is not taken")

	apiToken, err := user.GenerateToken(64)
	if err != nil {
		return models.User{}, err
	}

	invitedUser.Username = username
	invitedUser.Password = password
	invitedUser.ApiToken = apiToken

	err = user.Repo.DB.Save(invitedUser).Error
	if err != nil {
		return models.User{}, err
	}

	return invitedUser, nil
}

func (user *UsersRepository) FetchUser(id uint) (models.User, error) {
	var record models.User
	query := user.Repo.DB
	query = query.Preload("Roles").Preload("Roles.Role")
	query = query.Preload("Roles.Avatar")
	query = query.Preload("Accounts")

	err := query.First(&record, id).Error
	if err != nil {
		return models.User{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetAll() ([]models.User, error) {
	var records []models.User

	query := user.Repo.DB
	query = query.Preload("Roles").Preload("Roles.Role")
	query = query.Preload("Roles.Avatar")
	query = query.Preload("Accounts")

	err := query.Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (user *UsersRepository) GetByUsername(username string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("username = ?", username).First(&record).Error
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}

func (user *UsersRepository) GetUserByAccountID(id uint) (models.User, error) {

	var record models.User
	err := user.Repo.DB.Where("id = ?", id).First(&record).Error
	if err != nil {
		return models.User{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetAccountByUserID(id uint) (models.Account, error) {
	var record models.Account
	err := user.Repo.DB.Where("user_id = ?", id).First(&record).Error
	if err != nil {
		return models.Account{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetAccountByID(id uint) (models.Account, error) {
	var record models.Account
	err := user.Repo.DB.Where("id = ?", id).First(&record).Error
	if err != nil {
		return models.Account{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetAccountByEmail(email string) (models.Account, error) {
	var record models.Account
	err := user.Repo.DB.Where("email = ?", email).First(&record).Error
	if err != nil {
		return models.Account{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetByResetToken(token string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("password_reset_token = ?", token).First(&record).Error
	if err != nil {
		return models.User{}, fmt.Errorf("invalid reset token")
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

func (user *UsersRepository) GetAccounts(userID uint) ([]models.Account, error) {
	var userRecord models.User
	query := user.Repo.DB.Preload("Accounts")

	err := query.Find(&userRecord, userID).Error
	if err != nil {
		return nil, err
	}
	return userRecord.Accounts, nil
}

func (user *UsersRepository) GetAccount(userID uint, accountID uint) (models.Account, error) {
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

func (user *UsersRepository) Update(userData models.User) (models.User, error) {
	result := user.Repo.DB.Save(&userData)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	var updatedUser models.User
	query := user.Repo.DB.Preload("Roles").Preload("Roles.Role")
	query = query.Preload("Roles.Avatar")
	query = query.Preload("Accounts")
	query = query.First(&updatedUser, userData.ID)

	if query.Error != nil {
		return models.User{}, query.Error
	}

	return updatedUser, nil
}

func (user *UsersRepository) SaveAccount(accountData models.Account) (models.Account, error) {
	result := user.Repo.DB.Save(&accountData)
	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return accountData, nil

}

func (user *UsersRepository) UpdateUsername(userID uint, username string) (models.User, error) {
	var record models.User

	query := user.Repo.DB
	query = query.Preload("Roles").Preload("Roles.Role")
	query = query.Preload("Roles.Avatar")
	query = query.Preload("Accounts")

	err := query.Where("id = ?", userID).First(&record).Error
	if err != nil {
		return models.User{}, err
	}

	record.Username = username
	_, err = user.Repo.Update(userID, record)
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}

func (user *UsersRepository) UpdatePassword(userID uint, password string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("id = ?", userID).First(&record).Error
	if err != nil {
		return models.User{}, err
	}

	record.Password = password

	_, err = user.Repo.Update(userID, record)
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}

func (user *UsersRepository) InsertPasswordResetToken(userID uint, resetToken string, resetTokenExpiry time.Time) (models.User, error) {

	var record models.User
	err := user.Repo.DB.Where("id = ?", userID).First(&record).Error
	if err != nil {
		return models.User{}, err
	}

	record.PasswordResetToken = resetToken
	record.PasswordResetTokenExpiry = resetTokenExpiry

	_, err = user.Repo.Update(userID, record)
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}

func (user *UsersRepository) RemovePasswordResetToken(userID uint) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("id = ?", userID).First(&record).Error
	if err != nil {
		return models.User{}, err
	}

	zeroTime := time.Time{}
	desiredTime := zeroTime.Add(24 * time.Hour)

	record.PasswordResetToken = ""
	record.PasswordResetTokenExpiry = desiredTime

	_, err = user.Repo.Update(userID, record)
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}
