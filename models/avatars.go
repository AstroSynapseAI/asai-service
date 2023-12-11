package models

import "gorm.io/gorm"

type Avatar struct {
	gorm.Model
	ID     int
	LLM    *LLM
	Users  *[]User
	Agents *[]Agent
	Tools  *[]Tool
}



