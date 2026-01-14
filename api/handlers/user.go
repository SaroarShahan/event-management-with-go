package handlers

import (
	"errors"
	"fmt"

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
	// Check if email already exists
	var existing User
	if err := database.DB.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		return errors.New("email already registered")
	}

	// Check if username already exists
	if err := database.DB.Where("username = ?", user.Username).First(&existing).Error; err == nil {
		return errors.New("username already taken")
	}

	hashedPassword, err := internals.HashPassword(user.Password)

	if err != nil {
		return err
	}
	
	fmt.Println("Signup - Original password:", user.Password)
	fmt.Println("Signup - Hashed password:", string(hashedPassword))
	
	userWithHashPass := User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	// Insert into DB
	if err := database.DB.Create(&userWithHashPass).Error; err != nil {
		return err
	}

	fmt.Println("User saved with ID:", userWithHashPass.ID)
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

	// Get user by email (will also include Password for compare)
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare bcrypt hash
	if err := internals.ComparePassword(user.Password, password); err != nil {
		fmt.Println("Error fetching user:", err)
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}
