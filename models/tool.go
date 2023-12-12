package models

import "gorm.io/gorm"

type Tool struct {
  gorm.Model
  ID      int
  Name    string
	Slug  	string
	Token   string
	Avatars []Avatar `gorm:"many2many:avatar_tools;"`
  Agents  []Agent  `gorm:"many2many:agent_tools;"`
}