package controllers

import (
	"context"
	"fmt"

	"github.com/eluizbr/go_pagamentos/auth/src/auth/models"
	"github.com/eluizbr/go_pagamentos/auth/src/configs"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Erro ao receber os dados do usuário")
	}

	// hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	// if err != nil {
	// 	return c.Status(400).SendString("Erro ao gerar o hash da senha")
	// }

	secretData := map[string]interface{}{
		"password": user.Password,
	}

	user.Password = ""

	if err := configs.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Code":    400,
			"Error":   user,
			"message": fmt.Sprintf("Erro ao criar o usuário: %v", err),
		})
	}

	_, err := configs.VaultConn.KVv2("secret").Put(context.Background(), user.ID.String(), secretData)
	if err != nil {
		return c.Status(400).SendString("Erro ao salvar a senha no vault")
	}

	return c.Status(200).JSON(user)
}
