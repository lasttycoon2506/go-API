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

		if context.Request.Method == "OPTIONS" && origin == "http://localhost:3000" {
			context.Header("Access-Control-Allow-Origin", "http://localhost:3000")
			context.Header("Access-Control-Allow-Headers", "Content-Type")
			context.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT")
			context.Status(http.StatusOK)
			context.Abort()
			return
		}

		if (origin != "http://localhost:3000" && origin != "http://localhost:8080") || origin == "" {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: Invalid Origin"})
			return
		}

		context.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		context.Header("Access-Control-Allow-Headers", "Content-Type")
		context.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT")
	})

	routes.RegisterRoutes(server)
	server.Run(":8080")
}
