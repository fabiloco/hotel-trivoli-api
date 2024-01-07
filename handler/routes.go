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

	// Product type
	productType := api.Group("/product-type")
	productType.Get("/", h.ListProductTypes)
	productType.Get("/:id", h.GetProductTypeById)
	productType.Post("/", h.PostProductType)
	productType.Put("/:id", h.PutProductType)
	productType.Delete("/:id", h.DeleteProductTypeById)

	// User
	user := api.Group("/user")
	user.Get("/", h.ListUsers)
	user.Get("/:id", h.GetUserById)
	user.Post("/", h.CreateUser)
	user.Put("/:id", h.UpdateUser)
	user.Delete("/:id", h.DeleteUserById)

	// Auth
	// auth := api.Group("/auth")
	// auth.Post("/login", handler.Login)

}
