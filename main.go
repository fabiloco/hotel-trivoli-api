
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  app := fiber.New()

  app.Use(cors.New())

  // Serve static assets
  app.Static("/public", "./public")

  // database.ConnectDB()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("hello world")
  })

  app.Listen("0.0.0.0:3001")
}
