package main

import (
	"go-boilerplate/internal/domain"
	"go-boilerplate/internal/pkg/auth"
	"go-boilerplate/internal/pkg/database"
	"go-boilerplate/internal/pkg/pointer"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Start seed")
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

	userID1 := uuid.New()
	userID2 := uuid.New()

	pass, err := auth.HashPassword("password")
	if err != nil {
		log.Fatalf("Error hash password: %v", err)
		return
	}

	insertUserList := []*domain.User{
		{
			ID:       userID1,
			Name:     "user1",
			Email:    "user1@test.com",
			Password: pass,
		},
		{
			ID:       userID2,
			Name:     "user2",
			Email:    "user2@test.com",
			Password: pass,
		},
	}

	todoID1 := uuid.New()
	todoID2 := uuid.New()
	todoID3 := uuid.New()
	todoID4 := uuid.New()

	insertTodoList := []*domain.Todo{
		{
			ID:      todoID1,
			UserID:  userID1,
			Title:   "title1",
			Content: pointer.String("content1"),
		},
		{
			ID:      todoID2,
			UserID:  userID1,
			Title:   "title2",
			Content: pointer.String("content2"),
		},
		{
			ID:      todoID3,
			UserID:  userID2,
			Title:   "title3",
			Content: pointer.String("content3"),
		},
		{
			ID:      todoID4,
			UserID:  userID1,
			Title:   "title4",
			Content: pointer.String("content4"),
		},
	}

	db.Create(insertUserList)
	db.Create(insertTodoList)

	users := []*domain.User{}
	todos := []*domain.Todo{}

	db.Find(&users)
	db.Find(&todos)

	log.Printf("Successfully inserted %d users", len(users))
	log.Printf("Successfully inserted %d todos", len(todos))
}
