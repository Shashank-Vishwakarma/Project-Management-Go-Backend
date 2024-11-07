package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/types"
	"github.com/joho/godotenv"
)

var Config *types.Config

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetConfig() error {
	loadEnv()

	Config = &types.Config{
		Port:           os.Getenv("PORT"),
		JWT_SECRET_KEY: os.Getenv("JWT_SECRET_KEY"),
		DB_USERNAME:    os.Getenv("DB_USERNAME"),
		DB_PASSWORD:    os.Getenv("DB_PASSWORD"),
		DB_PORT:        os.Getenv("DB_PORT"),
		DB_HOST:        os.Getenv("DB_HOST"),
		DB_SSL_MODE:    os.Getenv("DB_SSL_MODE"),
	}

	cfg := map[string]interface{}{
		"PORT":           Config.Port,
		"JWT_SECRET_KEY": Config.JWT_SECRET_KEY,
		"DB_USERNAME":        Config.DB_USERNAME,
		"DB_PASSWORD":        Config.DB_PASSWORD,
		"DB_PORT":            Config.DB_PORT,
		"DB_HOST":            Config.DB_HOST,
		"DB_SSL_MODE":        Config.DB_SSL_MODE,
	}

	for key, value := range cfg {
		if value == "" {
			return fmt.Errorf("%s is not set", key)
		}
	}

	return nil
}