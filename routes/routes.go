package routes

import (
	"example.com/m/v2/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.GET("/events/user", middleware.Authenticate, getUserEvents)
	server.POST("/events", middleware.Authenticate, createEvent)
	server.PUT("/events/:id", middleware.Authenticate, editEvent)
	server.PUT("/user/update-password", middleware.Authenticate, updatePassword)
	server.DELETE("/events/:id", middleware.Authenticate, deleteEvent)
	server.POST("/register", register)
	server.POST("/login", login)
}
