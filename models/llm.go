package models

import (
	"github.com/GoLangWebSDK/crud/database"
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
						Name:        "GPT-4o",
						Slug:        "gpt-4o",
						Description: "OpenAI is a large-scale, open-source AI research project.",
						Provider:    "OpenAI",
					},
					{
						Name:        "GPT-4 Turbo",
						Slug:        "gpt-4-turbo",
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
		{
			ID: "seed_mistral",
			Execute: func(db *database.Database) error {
				llms := []LLM{
					{
						Name:        "Mistral-7B",
						Slug:        "open-mistral-7b",
						Description: "Mistral is developed by a creative team with high scientific standards",
						Provider:    "Mistral",
					},
					{
						Name:        "Mixtral-8x7B",
						Slug:        "open-mixtral-8x7b",
						Description: "Mistral is developed by a creative team with high scientific standards",
						Provider:    "Mistral",
					},
					{
						Name:        "Mistral Small Latest",
						Slug:        "mistral-small-latest",
						Description: "Mistral is developed by a creative team with high scientific standards",
						Provider:    "Mistral",
					},
					{
						Name:        "Mistral Medium Latest",
						Slug:        "mistral-medium-latest",
						Description: "Mistral is developed by a creative team with high scientific standards",
						Provider:    "Mistral",
					},
					{
						Name:        "Mistral Large Latest",
						Slug:        "mistral-large-latest",
						Description: "Mistral is developed by a creative team with high scientific standards",
						Provider:    "Mistral",
					},
					{
						Name:        "Mistral Embeder",
						Slug:        "mistral-embed",
						Description: "Mistral is developed by a creative team with high scientific standards",
						Provider:    "Mistral",
					},
				}

				return db.Adapter.Gorm().Create(&llms).Error
			},
		},
	}
}
