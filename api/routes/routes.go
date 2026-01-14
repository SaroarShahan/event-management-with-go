package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine) {
	eventRoutes(server)
	loginRoutes(server)
}