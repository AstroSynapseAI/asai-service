package models

import "gorm.io/gorm"

type LLM struct {
	gorm.Model
	ID   		int
	Name 		string
	Slug  	string
	Avatars []Avatar 
	Agents 	[]Agent 
}