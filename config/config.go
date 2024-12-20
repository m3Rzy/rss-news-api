package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() (string) {
	err := godotenv.Load("../../.env") 
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурационного файла .env: %v", err)
	}

	api := os.Getenv("API")
	
	if api == "" {
		log.Fatal("API key can't be empty")
	}

	return api
}