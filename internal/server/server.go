package server

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init() {
	Router = gin.New()

	Router.Use(gin.Recovery())
}