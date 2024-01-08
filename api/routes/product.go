package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	"fabiloco/hotel-trivoli-api/pkg/product"
	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app fiber.Router, service product.Service) {
  productGroup := app.Group("/product")
	productGroup.Get("/", handlers.GetProducts(service))
  productGroup.Get("/:id", handlers.GetProductById(service))
	productGroup.Post("/", handlers.PostProducts(service))
  productGroup.Put("/:id", handlers.PutProduct(service))
  productGroup.Delete("/:id", handlers.DeleteProductById(service))
}
