package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	"fabiloco/hotel-trivoli-api/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service auth.Service) {
  authGroup := app.Group("/auth")
	authGroup.Post("/register", handlers.Register(service))
	authGroup.Post("/login", handlers.Login(service))
}
