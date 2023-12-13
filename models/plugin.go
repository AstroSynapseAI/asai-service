package models

import "gorm.io/gorm"

type Plugins struct {
	gorm.Model
	ID      int
	Name    string
	Slug    string
	Avatars []Avatar `gorm:"many2many:avatar_plugins;"`
}
