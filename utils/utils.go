package utils

import (
	"icecreambash/flika-backend/database"
	"icecreambash/flika-backend/models"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func LoadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
}
