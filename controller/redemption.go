package controller

import (
	"github.com/NamChoco/go-database/database"
	"github.com/NamChoco/go-database/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetRedemptions(c *fiber.Ctx) error {
	var redemptions []model.Redemption
	database.DB.Preload("Member").Preload("Reward").Find(&redemptions)
	return c.JSON(redemptions)
}

func CreateRedemption(c *fiber.Ctx) error {
	var redeem model.Redemption
	if err := c.BodyParser(&redeem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// ลด reward quantity
	database.DB.Model(&model.Reward{}).
		Where("id = ?", redeem.RewardID).
		Update("quantity", gorm.Expr("quantity - ?", 1))

	database.DB.Create(&redeem)
	return c.JSON(redeem)
}