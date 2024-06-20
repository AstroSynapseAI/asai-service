package models

import (
	"time"

	"github.com/GoLangWebSDK/crud/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username                 string       `json:"username,omitempty"`
	Password                 string       `json:"password,omitempty"`
	ApiToken                 string       `json:"api_token,omitempty"`
	InviteToken              string       `json:"invite_token,omitempty"`
	IsAdmin                  bool         `json:"is_admin,omitempty"`
	Accounts                 []Account    `json:"accounts,omitempty"`
	Roles                    []AvatarRole `gorm:"foreignKey:UserID;" json:"roles,omitempty"`
	PasswordResetToken       string       `json:"password_reset_token,omitempty"`
	PasswordResetTokenExpiry time.Time    `json:"password_reset_token_expiry,omitempty"`
}

type AvatarRole struct {
	gorm.Model
	RoleID   uint   `json:"role_id,omitempty"`
	UserID   uint   `json:"user_id,omitempty"`
	AvatarID uint   `json:"avatar_id,omitempty"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role,omitempty"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Avatar   Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar,omitempty"`
}

func (*User) SeedModel(db *database.Database) []database.SeedAction {
	return []database.SeedAction{
		{
			ID: "seed_users",
			Execute: func(db *database.Database) error {
				users := []User{
					{
						Username:    "SuperAdmin",
						Password:    "admin_admin",
						ApiToken:    "tmp_token_superadmin_123",
						InviteToken: "",
						IsAdmin:     true,
					},
				}

				return db.Adapter.Gorm().Create(&users).Error
			},
		},
	}
}
