package main

import (
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.Use(func(context *gin.Context) {
		origin := context.GetHeader("Origin")
		if origin != "http://localhost:3000" || origin == "" {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: Invalid Origin"})
			return
		}
	})
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
