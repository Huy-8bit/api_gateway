package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvrionment(variableName string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(variableName)

}
