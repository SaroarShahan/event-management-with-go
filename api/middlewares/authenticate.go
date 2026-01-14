package middlewares

import (
	"net/http"

	"github.com/SaroarShahan/event-management/internals"
	"github.com/gin-gonic/gin"
)



func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": false,
			"message": "Unauthorized access",
			"data": nil,
		})
		return
	}

	userId, err := internals.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": false,
			"message": "Unauthorized access",
			"data": nil,
		})
		return
	}

	context.Set("userId", userId)
	context.Next()
}