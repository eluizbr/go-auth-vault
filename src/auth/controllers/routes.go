package controllers

import "github.com/gofiber/fiber/v2"

func InitAllRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/login", Login)
	v1.Post("/register", Register)
}
