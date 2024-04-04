package gorm

import "time"

type DBSeeder struct {
	ID         uint32 `json:"id" gorm:"primary_key"`
	SeederName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

