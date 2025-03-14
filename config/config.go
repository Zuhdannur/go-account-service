package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zuhdannur/go-fiber-bank-api/prisma/db"
)

var DB *db.PrismaClient

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set correctly")
	}

	DB = db.NewClient()
	if err := DB.Connect(); err != nil {
		fmt.Print(databaseURL)
		log.Fatal("Failed to connect to database:", err)
	}
}
