package main

import (
	"net/http"

	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"testing": "working"})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error parsing data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event Box})
}
