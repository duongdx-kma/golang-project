package initializers

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver         string `mapstructure:"DB_DRIVER"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBUserName       string `mapstructure:"DB_USER"`
	DBUserPassword   string `mapstructure:"DB_PASSWORD"`
	DBName           string `mapstructure:"DB_DATABASE"`
	DBPort           int    `mapstructure:"DB_PORT"`
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	ClientOrigin     string `mapstructure:"CLIENT_ORIGIN"`
	AppPort          int    `mapstructure:"APP_PORT"`
	Region           string `mapstructure:"AWS_REGION"`
	SecretManagerKey string `mapstructure:"SECRET_MANAGER_KEY"`
	AppEnvironment   string `mapstructure:"APP_ENV"`
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	// database environment
	config.DBDriver = os.Getenv("DB_DRIVER")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBUserName = os.Getenv("DB_USER")
	config.DBUserPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_DATABASE")
	config.DBPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Println("Error converting value of DB_PORT to integer")
	}

	// app environment
	config.ClientOrigin = os.Getenv("CLIENT_ORIGIN")
	config.JWTSecret = os.Getenv("JWT_SECRET")
	config.AppEnvironment = os.Getenv("APP_ENV")
	config.AppPort, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Println("Error converting value of APP_PORT to integer")
	}

	// aws environment
	config.Region = os.Getenv("AWS_REGION")
	config.SecretManagerKey = os.Getenv("SECRET_MANAGER_KEY")

	return
}
