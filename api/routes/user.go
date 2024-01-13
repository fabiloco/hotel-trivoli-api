package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	"fabiloco/hotel-trivoli-api/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
  productGroup := app.Group("/user")
	productGroup.Get("/", handlers.GetUsers(service))
  productGroup.Get("/:id", handlers.GetUserById(service))
  productGroup.Delete("/:id", handlers.DeleteUserById(service))
}
