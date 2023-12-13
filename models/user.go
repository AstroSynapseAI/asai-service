package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int
	Email    string
	Password string
	Username string
	AvatarID int
	Avatar   *Avatar
	Avatars  []Avatar `gorm:"many2many:user_avatars;"`
	RoleID   int
	Role     *Role
}
