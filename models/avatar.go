package models

import "gorm.io/gorm"

type Avatar struct {
  gorm.Model
  ID      			int
	Name    			string
	Slug   				string
	LLMID   			int
	LLM     			*LLM
	DefaultPrimer string
	Primer 				string
	IsPublic 			bool
  UserID  			int
  Users  	 			[]User  	`gorm:"many2many:user_avatars;"`
  Agents  			[]Agent 	`gorm:"many2many:avatar_agents;"`
  Tools   			[]Tool	  `gorm:"many2many:avatar_tools;"`
	Plugins 			[]Plugins `gorm:"many2many:avatar_plugins;"`
}



