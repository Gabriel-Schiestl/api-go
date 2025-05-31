package main

import (
	"github.com/Gabriel-Schiestl/api-go/internal/controllers"
	"github.com/Gabriel-Schiestl/api-go/internal/server"
	_ "github.com/Gabriel-Schiestl/api-go/pkg"
)

func main() {
	controllers.Init()

	server.Router.Run(":8080")
}