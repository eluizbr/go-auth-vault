package main

import (
	"log"

	"github.com/eluizbr/go_pagamentos/auth/src/auth/controllers"
	"github.com/eluizbr/go_pagamentos/auth/src/configs"
	address "github.com/eluizbr/go_pagamentos/auth/src/costumers/controllers/address"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	configs.ConnectDB()
	configs.ConnectVault()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(compress.New())

	// JWT Middleware
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("secret"),
	// }))

	// app.Use(getuserid.New()) // Get User ID from JWT Token

	controllers.InitAllRoutes(app)
	address.InitAllRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}
