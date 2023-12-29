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
  product := api.Group("/product")
  product.Get("/", handler.GetProducts)
  product.Get("/:id", handler.GetProductById)
  product.Post("/", handler.PostProducts)
  product.Put("/:id", handler.PutProduct)
  product.Delete("/:id", handler.DeleteProductById)

  // Auth
	// auth := api.Group("/auth")
	// auth.Post("/login", handler.Login)
}
