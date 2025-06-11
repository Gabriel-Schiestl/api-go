package server

import (
	"github.com/Gabriel-Schiestl/api-go/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.New()

	Router.Use(gin.Recovery())
	Router.Use(middlewares.AuthMiddleware())
}