package router

import (
	"fabiloco/hotel-trivoli-api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
  // middleware
  api := app.Group("/api/v1", logger.New())
  // api.Get("/", handler.Hello)

  // Product
  image := api.Group("/product")
  // image.Post("/", handler.PostImage)
  image.Get("/", handler.GetProducts)

  // Auth
	// auth := api.Group("/auth")
	// auth.Post("/login", handler.Login)
}
