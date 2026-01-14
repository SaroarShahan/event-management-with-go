package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SaroarShahan/event-management/api/handlers"
	"github.com/SaroarShahan/event-management/utils"
)

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(context *gin.Context) {
	var req SignupRequest
	
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid request payload",
			"data": nil,
		})
		return
	}

	user := handlers.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := user.SaveUserHandler(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Failed to create user",
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status": true,
		"message": "User has been created successfully!",
		"data": user,
	})
}

func Login(context *gin.Context) {
	var req LoginRequest
	
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid request payload",
			"data": nil,
		})
		return
	}

	if err := handlers.ValidateCredentialsHanlder(req.Email, req.Password); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status": false,
			"message": "Invalid email or password",
			"data": nil,
		})
		return
	}

	user, err := handlers.GetUserByEmailHandler(req.Email)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status": false,
			"message": "Invalid email or password",
			"data": nil,
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Could not authenticate user.",
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Login successful!",
		"data": gin.H{
			"token": token,
		},
	})
}