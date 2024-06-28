package models

import (
	"github.com/GoLangWebSDK/crud/database"
	"gorm.io/gorm"
)

type Plugin struct {
	gorm.Model
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	Slug          string         `json:"slug,omitempty"`
	ActivePlugins []ActivePlugin `gorm:"foreignKey:PluginID;" json:"active_plugins"`
}

func (*Plugin) SeedModel(db *database.Database) []database.SeedAction {
	return []database.SeedAction{
		{
			ID: "seed_plugins",
			Execute: func(db *database.Database) error {
				plugins := []Plugin{
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

				return db.Adapter.Gorm().Create(&plugins).Error
			},
		},
	}
}
