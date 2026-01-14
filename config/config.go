package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version string
	ServiceName string
	Host string
	SecretKey string
	HttpPort int
	DBName string
	DBUser string
	DBPassword string
	DBHost string
	DBPort int
}

var configurations Config

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	host := os.Getenv("HOST")
	secretKey := os.Getenv("SECRET_KEY")
	portStr := os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPortStr := os.Getenv("DB_PORT")

	if version == "" {
		fmt.Println("VERSION not set in .env file")
		os.Exit(1)
	}

	if serviceName == "" {
		fmt.Println("SERVICE_NAME not set in .env file")
		os.Exit(1)
	}
	if host == "" {
		fmt.Println("HOST not set in .env file")
		os.Exit(1)
	}

	if secretKey == "" {
		fmt.Println("SECRET_KEY not set in .env file")
		os.Exit(1)
	}

	if portStr == "" {
		fmt.Println("HTTP_PORT not set in .env file")
		os.Exit(1)
	}

	httpPort, err := strconv.ParseInt(portStr, 10, 64)
	
	if err != nil {
		fmt.Println("Invalid HTTP_PORT value")
		os.Exit(1)
	}

	if dbName == "" {
		fmt.Println("DB_NAME not set in .env file")
		os.Exit(1)
	}

	if dbUser == "" {
		fmt.Println("DB_USER not set in .env file")
		os.Exit(1)
	}

	if dbPassword == "" {
		fmt.Println("DB_PASSWORD not set in .env file")
		os.Exit(1)
	}

	if dbHost == "" {
		fmt.Println("DB_HOST not set in .env file")
		os.Exit(1)
	}

	httpDBPort, err := strconv.ParseInt(dbPortStr, 10, 64)

	if err != nil {
		fmt.Println("DB_PORT not set in .env file")
		os.Exit(1)
	}

	configurations = Config{
		Version: version,
		ServiceName: serviceName,
		Host: host,
		SecretKey: secretKey,
		HttpPort: int(httpPort),
		DBName: dbName,
		DBUser: dbUser,
		DBPassword: dbPassword,
		DBHost: dbHost,
		DBPort: int(httpDBPort),
	}
}

func GetConfig() Config {
	loadConfig()
	
	return configurations
}