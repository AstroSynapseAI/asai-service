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
	Avatars      []Avatar      `gorm:"foreignKey:LLMID;" json:"avatars,omitempty"`
	ActiveAgents []ActiveAgent `gorm:"foreignKey:LLMID;" json:"active_agents,omitempty"`
}

func (*LLM) SeedModel(db *database.Database) []database.SeedAction {
	return []database.SeedAction{
		{
			ID: "seed_llms",
			Execute: func(db *database.Database) error {
				llms := []LLM{
					{
						Name:        "GPT-4 Turbo Preview",
						Slug:        "gpt-4-turbo-preview",
						Description: "OpenAI is a large-scale, open-source AI research project.",
						Provider:    "OpenAI",
					},
					{
						Name:        "GPT-4",
						Slug:        "gpt-4",
						Description: "OpenAI is a large-scale, open-source AI research project.",
						Provider:    "OpenAI",
					},
					{
						Name:        "GPT-4-0613",
						Slug:        "gpt-4-0613",
						Description: "OpenAI is a large-scale, open-source AI research project.",
						Provider:    "OpenAI",
					},
					{
						Name:        "GPT-3.5",
						Slug:        "gpt-3.5",
						Description: "OpenAI is a large-scale, open-source AI research project.",
						Provider:    "OpenAI",
					},
				}

				return db.Adapter.Gorm().Create(&llms).Error
			},
		},
	}
}
