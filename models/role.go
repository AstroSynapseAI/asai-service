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

func (*Role) SeedModel(db *database.Database) []database.SeedAction {
	return []database.SeedAction{
		{
			ID: "seed_roles",
			Execute: func(db *database.Database) error {
				roles := []Role{
					{
						Name:        "Avatar Owner",
						Slug:        "role-owner",
						Description: "Owner Role for Avatar",
						Permission:  "owner",
					},
				}

				return db.Adapter.Gorm().Create(&roles).Error
			},
		},
	}
}
