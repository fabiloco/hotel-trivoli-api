package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (h *Handler) Register(app *fiber.App) {
  // middleware
  api := app.Group("/api/v1", logger.New())
  // api.Get("/", handler.Hello)

  // Product
  product := api.Group("/product")
  product.Get("/", h.ListProducts)
  product.Get("/:id", h.GetProductById)
  product.Post("/", h.PostProducts)
  product.Put("/:id", h.PutProduct)
  product.Delete("/:id", h.DeleteProductById)

  // Auth
	// auth := api.Group("/auth")
	// auth.Post("/login", handler.Login)
}
