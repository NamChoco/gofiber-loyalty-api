package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/NamChoco/go-database/database"
	"github.com/NamChoco/go-database/model"
)

func GetTransactions(c *fiber.Ctx) error {
	var txs []model.Transaction
	database.DB.Preload("Member").Find(&txs)
	return c.JSON(txs)
}

func CreateTransaction(c *fiber.Ctx) error {
	var tx model.Transaction
	if err := c.BodyParser(&tx); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&tx)
	return c.JSON(tx)
}