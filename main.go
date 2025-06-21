package main

import (
	"log"


	"github.com/NamChoco/go-database/database"
	"github.com/NamChoco/go-database/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	database.Connect()
	// Mock data
	// database.SeedData(database.DB)

	app := fiber.New()

	// Api
	router.SetUpRouter(app)

	app.Listen(":8080")

}
