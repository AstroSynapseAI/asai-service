package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	Path        string `json:"path,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AvatarID    uint   `json:"avatar_id"`
	Avatar      Avatar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"avatar"`
}
