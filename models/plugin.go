package models

import (
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/gorm"
)

type Plugin struct {
	gorm.Model
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	Slug          string         `json:"slug,omitempty"`
	ActivePlugins []ActivePlugin `gorm:"foreignKey:PluginID;" json:"active_plugins"`
}

func (*Plugin) SeedModel(db *database.Database) error {
	seeder := "seed_plugins"

	result := db.Adapter.Gorm().Where("seeder_name = ?", seeder).First(&DBSeeder{})
	if result.Error == gorm.ErrRecordNotFound {
		var plugins []Plugin = []Plugin{
			{
				Name:        "Slack",
				Slug:        "slack",
				Description: `Empower your conversations with your AI Avatar on Slack! This dynamic feature translates easily into an assistant for your team, answering inquiries, simplifying tasks, and streamlining information flow like a pro. All this, right within your existing team chats!`,
			},
			{
				Name:        "Discord",
				Slug:        "discord",
				Description: `Welcome your AI Avatar into your Discord channels, ready to redefine your community interactions! Providing instant responses and meaningful insights, this feature seamlessly merges with your discussions, adding a layer of efficiency to your community engagements!`,
			},
		}

		if result := db.Adapter.Gorm().Create(&plugins); result.Error != nil {
			return result.Error
		}

		if result := db.Adapter.Gorm().Create(&DBSeeder{SeederName: seeder}); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
