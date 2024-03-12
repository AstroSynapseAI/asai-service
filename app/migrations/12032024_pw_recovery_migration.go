package main

import (
	"fmt"
	"os"

	"github.com/AstroSynapseAI/app-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	var dsn string
	//fix for local dev since os.Getenv returns empty string and not "LOCAL DEV"
	if os.Getenv("ENVIRONMENT") == "" {
		dsn = "host=localhost user=asai-admin password=asai-password dbname=asai-db port=5432 sslmode=disable"
	}

	if os.Getenv("ENVIRONMENT") == "HEROKU DEV" {
		dsn = os.Getenv("DATABASE_URL")
	}

	if os.Getenv("ENVIRONMENT") == "AWS DEV" {
		username := os.Getenv("RDS_USERNAME")
		password := os.Getenv("RDS_PASSWORD")
		database := os.Getenv("RDS_DB_NAME")
		host := os.Getenv("RDS_HOST")
		port := os.Getenv("RDS_PORT")
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
		panic(err)
	}

	db.AutoMigrate(&models.User{})
}
