package main

import (
	"fmt"
	persistence_gorm "go-boilerplate/internal/infrastructure/persistence/gorm"
	"go-boilerplate/internal/interfaces/handler"
	"go-boilerplate/internal/pkg/database"
	"go-boilerplate/internal/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Printf("Start server")
	db, err := database.InitConnectDB()
	if err != nil {
		log.Fatalf("Error connect to database: %v", err)
		return
	}

	r := mux.NewRouter()
	userRepository := persistence_gorm.NewUserRepository(db)
	todoRepository := persistence_gorm.NewTodoRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	userUsecase := usecase.NewUserUseCase(userRepository)
	todoUsecase := usecase.NewTodoUseCase(todoRepository)
	authHandler := handler.NewAuthHandler(authUsecase)
	todoHandler := handler.NewTodoHandler(todoUsecase, userUsecase)

	authHandler.RegisterAuthHandlers(r)
	todoHandler.RegisterTodoHandlers(r)

	c := cors.New(cors.Options{
		// AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")}, // フロントエンドのオリジン
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("BACKEND_CONTAINER_POST")), handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
