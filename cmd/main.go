package main

import (
	_ "github.com/Gabriel-Schiestl/api-go/internal/controllers"
	"github.com/Gabriel-Schiestl/api-go/internal/server"
	"github.com/Gabriel-Schiestl/go-clarch/presentation/controller"
)

func main() {
	controller.SetupRoutes()

	server.Router.Run(":8080")
}