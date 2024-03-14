package main

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/app"
	"github.com/AstroSynapseAI/app-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	app := app.NewConfig()
	app.LoadEnvironment()
	if app.DSN == "" {
		app.DSN = "host=localhost user=asai-admin password=asai-password dbname=asai-db port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(app.DSN), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
		panic(err)
	}

	db.AutoMigrate(&models.User{})

}
