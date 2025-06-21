package router

import (
	"github.com/NamChoco/go-database/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App) {
	api := app.Group("/api")    // Group /api
	// book := api.Group("/books")
	// book.Get("/", controller.GetBooks) 

	api.Get("/members", controller.GetMembers)
	api.Post("/members", controller.CreateMember)

	api.Get("/rewards", controller.GetRewards)
	api.Post("/rewards", controller.CreateReward)

	api.Get("/transactions", controller.GetTransactions)
	api.Post("/transactions", controller.CreateTransaction)

	api.Get("/redemptions", controller.GetRedemptions)
	api.Post("/redemptions", controller.CreateRedemption)
}
