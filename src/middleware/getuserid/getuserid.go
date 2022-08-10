package getuserid

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Header is the header key where to get/set the unique request ID
	//
	// Optional. Default: "X-Request-ID"
	Header string

	// Generator defines a function to generate the unique identifier.
	//
	// Optional. Default: utils.UUID
	Generator func() string

	// ContextKey defines the key used when storing the request ID in
	// the locals for a specific request.
	//
	// Optional. Default: requestid
	ContextKey string
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:       nil,
	Header:     fiber.HeaderXRequestID,
	Generator:  utils.UUID,
	ContextKey: "requestid",
}

func New(config ...Config) fiber.Handler {

	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(string)

		costumerId, _ := uuid.Parse(id)

		c.Locals("costumerId", costumerId)

		return c.Next()
	}

}
