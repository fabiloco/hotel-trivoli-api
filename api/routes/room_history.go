package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	roomHistory "fabiloco/hotel-trivoli-api/pkg/room_history"

	"github.com/gofiber/fiber/v2"
)

func RoomHistoryRouter(app fiber.Router, service roomHistory.Service) {
  productGroup := app.Group("/room-history")
	productGroup.Get("/", handlers.GetRoomHistorys(service))
  productGroup.Get("/:id", handlers.GetRoomHistoryById(service))
	productGroup.Post("/", handlers.PostRoomHistorys(service))
  productGroup.Put("/:id", handlers.PutRoomHistory(service))
  productGroup.Delete("/:id", handlers.DeleteRoomHistoryById(service))
}
