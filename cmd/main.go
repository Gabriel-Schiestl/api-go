package main

import (
	"log"
	"os"

	_ "github.com/Gabriel-Schiestl/api-go/internal/controllers"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/database/connection"
	"github.com/Gabriel-Schiestl/api-go/internal/server"
	"github.com/Gabriel-Schiestl/go-clarch/presentation/controller"
	"github.com/joho/godotenv"
)

func main() {
	controller.SetupRoutes()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env: %v", err)
	}

	connection.SetupConfig(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	server.Router.Run(":8080")
}