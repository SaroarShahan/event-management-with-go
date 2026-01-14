package internals

import (
	"errors"
	"time"

	"github.com/SaroarShahan/event-management/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, userId int64) (string, error) {
	config := config.GetConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(config.SecretKey))
}

func VerifyToken(tokenString string) (int64, error) {
	config := config.GetConfig()

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method.")
		}

		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token.")
	}

	if isValidToken := parsedToken.Valid; !isValidToken {	
		return 0, errors.New("Invalid token.")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims.")
	}

	userIdFloat, ok := claims["userId"].(float64)
	
	if !ok {
		return 0, errors.New("Invalid userId in token.")
	}

	userId := int64(userIdFloat)

	return userId, nil
}