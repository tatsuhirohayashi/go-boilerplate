package main

import (
	"go-boilerplate/internal/domain"
	"go-boilerplate/internal/pkg/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Start migrate")
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env.local file: %v", err)
		return
	}

	db, err := database.InitConnectDB()
	if err != nil {
		log.Fatalf("Error connect to database: %v", err)
		return
	}

	// UUID拡張機能を有効化
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	db.AutoMigrate(&domain.User{}, &domain.Todo{})

	log.Printf("Migration completed")
}
