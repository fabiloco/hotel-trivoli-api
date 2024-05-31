package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	"fabiloco/hotel-trivoli-api/pkg/shift"

	"github.com/gofiber/fiber/v2"
)

func ShiftRouter(app fiber.Router, service shift.Service) {
	shiftGroup := app.Group("/shift")
	shiftGroup.Get("/", handlers.GetShifts(service))
	shiftGroup.Get("/:id", handlers.GetShiftById(service))
	shiftGroup.Post("/", handlers.PostShifts(service))
	shiftGroup.Put("/:id", handlers.PutShift(service))
	shiftGroup.Delete("/:id", handlers.DeleteShiftById(service))

	shiftGroup.Post("/between-dates", handlers.GetShiftsBetweenDates(service))
}
