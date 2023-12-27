package models

import (
	"database/sql"

	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string       `json:"username,omitempty"`
	Password    string       `json:"password"`
	ApiToken    string       `json:"api_token"`
	InviteToken string       `json:"invite_token"`
	Accounts    []Account    `json:"accounts"`
	Roles       []AvatarRole `gorm:"foreignKey:UserID;" json:"roles"`
}

type AvatarRole struct {
	gorm.Model
	RoleID   sql.NullInt64 `json:"role_id"`
	UserID   sql.NullInt64 `json:"user_id"`
	AvatarID sql.NullInt64 `json:"avatar_id"`
	Role     Role          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	User     User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Avatar   Avatar        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar"`
}

func (*User) SeedModel(db *database.Database) error {
	seeder := "seed_users"

	result := db.Adapter.Gorm().Where("seeder_name = ?", seeder).First(&DBSeeder{})
	if result.Error == gorm.ErrRecordNotFound {
		var users []User = []User{
			{
				Username:    "SuperAdmin",
				Password:    "admin_admin",
				ApiToken:    "tmp_token_superadmin_123",
				InviteToken: "",
			},
		}

		if result := db.Adapter.Gorm().Create(&users); result.Error != nil {
			return result.Error
		}

		if result := db.Adapter.Gorm().Create(&DBSeeder{SeederName: seeder}); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
