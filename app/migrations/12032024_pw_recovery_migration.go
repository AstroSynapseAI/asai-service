package main

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=asai-admin password=asai-password dbname=asai-db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
		panic(err)
	}

	db.AutoMigrate(&models.User{})
}
