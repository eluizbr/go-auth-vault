package address

import (
	"github.com/eluizbr/go_pagamentos/auth/src/configs"
	"github.com/eluizbr/go_pagamentos/auth/src/costumers/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateAddress(c *fiber.Ctx) error {

	address := models.Address{}
	db := configs.DB

	if err := c.BodyParser(&address); err != nil {
		return c.Status(401).SendString("Invalid request body")
	}

	address.CostomerId = c.Locals("costumerId").(uuid.UUID)

	if err := db.Create(&address).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error creating address",
			"error":   err,
		})
	}

	return c.Status(201).JSON(address)

}
