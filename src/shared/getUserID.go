package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func New() fiber.Handler {

	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(string)

		costumerId, _ := uuid.Parse(id)

		c.Locals("costumerId", costumerId)

		return c.Next()
	}

}
