package main

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/app"
	"github.com/AstroSynapseAI/app-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	cnf := app.NewConfig().DSN
	db, err := gorm.Open(postgres.Open(cnf), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
		panic(err)
	}

	db.AutoMigrate(&models.User{})
}
