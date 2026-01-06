package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func InitializeEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to initialize environment")
		return
	}
}
