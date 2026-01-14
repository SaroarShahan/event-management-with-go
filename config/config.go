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
	HttpPort int
}

var configurations Config

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	host := os.Getenv("HOST")
	portStr := os.Getenv("PORT")
	

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

	if portStr == "" {
		fmt.Println("HTTP_PORT not set in .env file")
		os.Exit(1)
	}

	httpPort, err := strconv.ParseInt(portStr, 10, 64)
	
	if err != nil {
		fmt.Println("Invalid HTTP_PORT value")
		os.Exit(1)
	}

	configurations = Config{
		Version: version,
		ServiceName: serviceName,
		Host: host,
		HttpPort: int(httpPort),
	}
}

func GetConfig() Config {
	loadConfig()
	
	return configurations
}