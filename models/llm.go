package models

import (
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/gorm"
)

type LLM struct {
	gorm.Model
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	Provider     string        `json:"provider,omitempty"`
	Slug         string        `json:"slug,omitempty"`
	Avatars      []Avatar      `gorm:"foreignKey:LLMID;" json:"avatars"`
	ActiveAgents []ActiveAgent `gorm:"foreignKey:LLMID;" json:"active_agents"`
}

func (*LLM) SeedModel(db *database.Database) error {
	seeder := "seed_llms"

	result := db.Adapter.Gorm().Where("seeder_name = ?", seeder).First(&DBSeeder{})
	if result.Error == gorm.ErrRecordNotFound {
		var llms []LLM = []LLM{
			{
				Name:        "gpt-4",
				Slug:        "gpt-4",
				Description: "OpenAI is a large-scale, open-source AI research project.",
				Provider:    "OpenAI",
			},
		}

		if result := db.Adapter.Gorm().Create(&llms); result.Error != nil {
			return result.Error
		}

		if result := db.Adapter.Gorm().Create(&DBSeeder{SeederName: seeder}); result.Error != nil {
			return result.Error
		}
	}
	return nil
}
