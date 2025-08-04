package routes

import (
	"net/http"
	"strconv"

	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error retrieving data", "error": err.Error()})
		return
	}

	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error retrieving data", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error retrieving data", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindBodyWithJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error parsing data", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	println(userId)
	event.UserId = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error saving data", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func editEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error parsing param", "error": err.Error()})
		return
	}

	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error retrieving data", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	eventId := event.ID
	if eventId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized to edit"})
		return
	}

	var editedEvent models.Event
	err = context.ShouldBindBodyWithJSON(&editedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error binding incoming JSON", "error": err.Error()})
		return
	}

	editedEvent.ID = id
	err = editedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error updating", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "update successful"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error parsing param", "error": err.Error()})
		return
	}

	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error retrieving data", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized to delete"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error deleting data", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}
