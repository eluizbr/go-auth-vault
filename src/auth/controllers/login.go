package controllers

import (
	"context"

	"github.com/eluizbr/go_pagamentos/auth/src/auth/models"
	"github.com/eluizbr/go_pagamentos/auth/src/configs"
	"github.com/eluizbr/go_pagamentos/auth/src/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Login(c *fiber.Ctx) error {

	body := models.Login{}
	var user models.User
	db := configs.DB

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	db.Find(&user, "email = ?", body.Email)

	if user.ID == uuid.Nil {
		return c.Status(400).JSON(fiber.Map{
			"Code":    400,
			"message": user,
		})
	}

	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	secret, err := configs.VaultConn.KVv2("secret").Get(context.Background(), user.ID.String())
	if err != nil {
		return c.Status(400).SendString("Erro ao buscar a senha do usuário")
	}

	if secret.Data["password"] != body.Password {
		return c.Status(400).SendString("Senha inválida")
	}

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Code":    400,
			"message": "Error: Invalid password",
		})
	}

	token, err := shared.CreateJWT(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Code":    400,
			"message": "Error: Invalid token",
		})
	}

	return c.Status(200).JSON(token)
}
