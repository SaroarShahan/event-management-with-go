package routes

import (
	"github.com/SaroarShahan/event-management/api/middlewares"
	"github.com/SaroarShahan/event-management/api/responses"
	"github.com/gin-gonic/gin"
)

func eventRoutes(server *gin.Engine) {
	server.GET("/events", responses.GetEvents)
	server.GET("/events/:id", responses.GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", responses.CreateEvent)
	authenticated.PUT("/events/:id", responses.UpdateEvent)
	authenticated.DELETE("/events/:id", responses.DeleteEvent)
}