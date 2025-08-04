package routes

import (
	"net/http"

	"example.com/m/v2/models"
	"example.com/m/v2/utils"
	"github.com/gin-gonic/gin"
)

func register(context *gin.Context) {
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error parsing data", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error saving data", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error parsing data", "error": err.Error()})
		return
	}

	err = user.Verify()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.CreateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldnt create jwt", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}

func getUserEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	userEvents, err := models.GetUserEvents(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error retrieving users events", "error": err.Error()})
		return
	}

}
