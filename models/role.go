package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID 		int
	Name 	string
	Slug  string
	Users []User
}