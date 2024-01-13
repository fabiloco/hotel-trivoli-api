package middleware

import (
	"fabiloco/hotel-trivoli-api/api/utils"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract the token from the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}

		tokenString := authHeader[len("Bearer "):]

		// Parse the access token
		claims := utils.ParseAccessToken(tokenString)
		if claims == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}

		// You can now access the claims (user ID, username, role) in your route handlers
		c.Locals("userClaims", claims)

		// Proceed to the next middleware or route handler
		return c.Next()
	}
}
