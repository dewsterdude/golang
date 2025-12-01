package routes

import (
	"example.com/REST-API/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // Get, POST, PUT, PATCH, DELETE ("/ping", func(c *gin.Context) {
	server.GET("/events/:id", getEvent) // get specific event by id - dynamic parameter via gin using colin + identifier

	authenticated := server.Group("/")         // path they all have in common
	authenticated.Use(middleware.Authenticate) // will always run middleware handlers for the group
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup) // signup new users
	server.POST("/login", login)   // ...
}

/*	server.POST("/events", middleware.Authenticate, createEvent)
	server.PUT("/events/:id", updateEvent)    // used for updates
	server.DELETE("/events/:id", deleteEvent) // used for updates
*/
