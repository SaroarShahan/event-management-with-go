package models

import (
	"errors"

	"github.com/SaroarShahan/event-management/infra/database"
	"github.com/SaroarShahan/event-management/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (usr *User) Save() error {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(usr.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(usr.Username, usr.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	usr.ID = id

	return err
}

func GetUserByEmail(email string) (*User, error) {
	query := `SELECT * FROM users WHERE email = ?`
	row := database.DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func ValidateCredentials(email string, password string)  error {
	query := `SELECT password FROM users WHERE email = ?`
	row := database.DB.QueryRow(query, email)

	var storedHashedPassword string
	
	if err := row.Scan(&storedHashedPassword); err != nil {
		return errors.New("Invalid email or password")
	}

	if err := utils.ComparePassword(storedHashedPassword, password); err != nil {
		return errors.New("Invalid email or password")
	}

	return nil
}