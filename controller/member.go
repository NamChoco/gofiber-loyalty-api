package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/NamChoco/go-database/database"
	"github.com/NamChoco/go-database/model"
)

func GetMembers(c *fiber.Ctx) error {
	var members []model.Member
	database.DB.Find(&members)
	return c.JSON(members)
}

func CreateMember(c *fiber.Ctx) error {
	var member model.Member
	if err := c.BodyParser(&member); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&member)
	return c.JSON(member)
}