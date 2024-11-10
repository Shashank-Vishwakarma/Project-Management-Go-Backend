package database

import (
	"fmt"
	"log"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func createTables() {
	err := DBClient.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{})
	if err != nil {
		log.Fatalf("Error creating tables: %s", err)
	}

	log.Println("Tables migration successful...")
}

func InitDB() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Config.DB_HOST,
		config.Config.DB_USERNAME,
		config.Config.DB_PASSWORD,
		config.Config.DB_NAME,
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

	createTables()

	return nil
}
