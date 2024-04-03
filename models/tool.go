package models

import (
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/gorm"
)

type Tool struct {
	gorm.Model
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	Slug             string            `json:"slug,omitempty"`
	ActiveTools      []ActiveTool      `gorm:"foreignKey:ToolID;" json:"active_tools"`
	ActiveAgentTools []ActiveAgentTool `gorm:"foreignKey:ToolID;" json:"agent_tools"`
}

func (*Tool) SeedModel(db *database.Database) []database.SeedAction {
	return []database.SeedAction{
		{
			ID: "seed_tools",
			Execute: func(db *database.Database) error {
				tools := []Tool{
					{
						Name:        "Google Search",
						Slug:        "google-search",
						Description: "Google Search via SerpAPI, SerpAPI token required.",
					},
					// {
					// 	Name:        "DuckDuckGo Search",
					// 	Slug:        "ddg-search",
					// 	Description: "DuckDuckGo Search API, Free Service.",
					// },
					// {
					// 	Name:        "Metaphor Search",
					// 	Slug:        "metaphor-search",
					// 	Description: "Metaphor Search API, API key required.",
					// },
					{
						Name:        "PDF Reader",
						Slug:        "pdf-reader",
						Description: "Email tool enables sending emails.",
					},
				}

				return db.Adapter.Gorm().Create(&tools).Error
			},
		},
	}
}
