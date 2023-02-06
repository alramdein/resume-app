package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error(err.Error())
		panic(fmt.Errorf("fatal can't load env file: %w", err))
	}
}

func HTTPHost() string {
	return os.Getenv("HOST")
}

func HTTPPort() string {
	return os.Getenv("PORT")
}

func DBHost() string {
	return os.Getenv("DB_HOST")
}

func DBPort() string {
	return os.Getenv("DB_PORT")
}

func DBName() string {
	return os.Getenv("DB_NAME")
}

func DBUsername() string {
	return os.Getenv("DB_USERNAME")
}

func DBPassword() string {
	return os.Getenv("DB_PASSWORD")
}
