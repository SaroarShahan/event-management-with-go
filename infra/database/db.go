package database

import (
	"fmt"

	"github.com/SaroarShahan/event-management/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnectionString() string{
	cfg := config.GetConfig()

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Dhaka",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)
}

func NewConnection() {
	dbSource := GetConnectionString()
	db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})

	if err != nil {
		fmt.Println("eeeeee ~~>", err)
		panic("Could not connect to database.")
	}

	DB = db
}