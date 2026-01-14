package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SaroarShahan/event-management/api/handlers"
	"github.com/SaroarShahan/event-management/internals"
)

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Signup(context *gin.Context) {
	var req SignupRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	user := handlers.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := user.SaveUserHandler(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User has been created successfully.",
		"user": gin.H{
			"id": user.ID,
			"username": user.Username,
			"email": user.Email,
		},
	})
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request payload",
			"data":    nil,
		})
		return
	}


	user, err := handlers.ValidateCredentialsHandler(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid email or password",
			"data":    nil,
		})
		return
	}

	token, err := internals.GenerateToken(user.Email, int64(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Could not authenticate user.",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Login successful!",
		"data": gin.H{
			"token": token,
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
			},
		},
	})
}
