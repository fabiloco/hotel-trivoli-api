package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// this middleware handles all the resposes to format them so the output json looks consistent
func FormatResponse() fiber.Handler {
  return func(c *fiber.Ctx) error {
    if err := c.Next(); err != nil {
      return err
    }

    data := c.Locals("data")

    response := c.Response()

    formattedResponse := fiber.Map{
      "success": response.StatusCode() >= 200 && response.StatusCode() < 400,
      // "message": response.Status(),
      "data": data,
    }

    return c.Status(response.StatusCode()).JSON(formattedResponse)
  }
}
