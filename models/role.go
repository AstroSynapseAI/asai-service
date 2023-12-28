package models

import (
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Permission  string       `json:"permission,omitempty"`
	Slug        string       `json:"slug,omitempty"`
	AvatarRoles []AvatarRole `gorm:"foreignKey:RoleID;" json:"avatar_roles,omitempty"`
}

func (*Role) SeedModel(db *database.Database) error {
	seeder := "seed_roles"

	result := db.Adapter.Gorm().Where("seeder_name = ?", seeder).First(&DBSeeder{})
	if result.Error == gorm.ErrRecordNotFound {
		var roles []Role = []Role{
			{
				Name:        "Avatar Owner",
				Slug:        "role-owner",
				Description: "Owner Role for Avatar",
				Permission:  "owner",
			},
			{
				Name:        "Admin",
				Slug:        "admin",
				Description: "Platform Admin Role",
				Permission:  "admin",
			},
		}

		if result := db.Adapter.Gorm().Create(&roles); result.Error != nil {
			return result.Error
		}

		if result := db.Adapter.Gorm().Create(&DBSeeder{SeederName: seeder}); result.Error != nil {
			return result.Error
		}
	}
	return nil
}
