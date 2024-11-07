package database

import (
	"fmt"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=%s",
		config.Config.DB_HOST,
		config.Config.DB_USERNAME,
		config.Config.DB_PASSWORD,
		config.Config.DB_PORT,
		config.Config.DB_SSL_MODE,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	DBClient = db

	return nil
}
