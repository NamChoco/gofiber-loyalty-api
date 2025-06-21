package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/NamChoco/go-database/database"
	"github.com/NamChoco/go-database/model"
)

func GetRewards(c *fiber.Ctx) error {
	var rewards []model.Reward
	database.DB.Find(&rewards)
	return c.JSON(rewards)
}

func CreateReward(c *fiber.Ctx) error {
	var reward model.Reward
	if err := c.BodyParser(&reward); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&reward)
	return c.JSON(reward)
}