package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	"fabiloco/hotel-trivoli-api/pkg/room"

	"github.com/gofiber/fiber/v2"
)

func RoomRouter(app fiber.Router, service room.Service) {
	productGroup := app.Group("/room")
	productGroup.Get("/", handlers.GetRooms(service))
	productGroup.Get("/:id", handlers.GetRoomById(service))
	productGroup.Post("/", handlers.PostRooms(service))
	productGroup.Put("/:id", handlers.PutRoom(service))
	productGroup.Delete("/:id", handlers.DeleteRoomById(service))
}
