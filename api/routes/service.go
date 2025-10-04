package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	"fabiloco/hotel-trivoli-api/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func ServiceRouter(app fiber.Router, service service.Service) {
	productGroup := app.Group("/service")
	productGroup.Get("/paginated", handlers.GetServicesPaginated(service))
	productGroup.Get("/", handlers.GetServices(service))
	productGroup.Get("/:id", handlers.GetServiceById(service))
	productGroup.Post("/", handlers.PostServices(service))
	productGroup.Put("/:id", handlers.PutService(service))
	productGroup.Delete("/:id", handlers.DeleteServiceById(service))
}
