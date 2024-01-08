package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	producttype "fabiloco/hotel-trivoli-api/pkg/product_type"

	"github.com/gofiber/fiber/v2"
)

func ProductTypeRouter(app fiber.Router, service producttype.Service) {
  productGroup := app.Group("/product-type")
	productGroup.Get("/", handlers.GetProductTypes(service))
  productGroup.Get("/:id", handlers.GetProductTypeById(service))
	productGroup.Post("/", handlers.PostProductTypes(service))
  productGroup.Put("/:id", handlers.PostProductTypes(service))
  productGroup.Delete("/:id", handlers.DeleteProductTypeById(service))
}
