package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/NamChoco/go-database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB instance
var DB *gorm.DB

func Connect() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error 
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	err = DB.AutoMigrate(
		&model.Member{},
		&model.Transaction{},
		&model.Reward{},
		&model.Redemption{},
	)
	if err != nil {
		panic("failed to migrate database schemas")
	}
	fmt.Println("Database Migrated")
}
