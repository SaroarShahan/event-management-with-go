package routes

import (
	"github.com/SaroarShahan/event-management/api/responses"
	"github.com/gin-gonic/gin"
)

func loginRoutes(server *gin.Engine) {
	server.POST("/signup", responses.Signup)
	server.POST("/login", responses.Login)
}