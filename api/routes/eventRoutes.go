package routes

import (
	"github.com/SaroarShahan/event-management/api/responses"
	"github.com/gin-gonic/gin"
)

func eventRoutes(server *gin.Engine) {
	server.GET("/events", responses.GetEvents)
	server.GET("/events/:id", responses.GetEvent)
	server.POST("/events", responses.CreateEvent)
	server.PUT("/events/:id", responses.UpdateEvent)
	server.DELETE("/events/:id", responses.DeleteEvent)
}