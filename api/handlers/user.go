package handlers

import (
	"errors"

	"github.com/SaroarShahan/event-management/infra/database"
	"github.com/SaroarShahan/event-management/internals"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
}


func (user *User) SaveUserHandler() error {
	var existing User
	if err := database.DB.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		return errors.New("email already registered")
	}

	if err := database.DB.Where("username = ?", user.Username).First(&existing).Error; err == nil {
		return errors.New("username already taken")
	}

	hashedPassword, err := internals.HashPassword(user.Password)

	if err != nil {
		return err
	}
	
	userWithHashPass := User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&userWithHashPass).Error; err != nil {
		return err
	}

	user.ID = userWithHashPass.ID

	return nil
}

func GetUserByEmailHandler(email string) (*User, error) {
	var user User

	err := database.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func ValidateCredentialsHandler(email, password string) (*User, error) {
	var user User

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := internals.ComparePassword(user.Password, password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}
