package main

import (
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error parsing data"})
		return
	}

	event.ID = 1
	event.UserId = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error saving data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
