package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func NewServer() *gin.Engine {
	// create a default gin server
	server := gin.Default()

	// register routes on the server instance
	RegisterRoutes(server)
	return server
}
