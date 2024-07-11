package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MainSQLitePath string
)

func Initial() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MainSQLitePath = os.Getenv("MAINSQLITEPATH")
}
