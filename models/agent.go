package models

import "gorm.io/gorm"

type Agent struct {
  gorm.Model
  ID    				int
  Name  				string
	Slug  				string
	LLMID   			int
	LLM   				*LLM
	DefaultPrimer string
	Primer 				string
	IsActive 			bool
  Tools 				[]Tool `gorm:"many2many:agent_tools;"`
}