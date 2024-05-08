package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	"fabiloco/hotel-trivoli-api/pkg/reports"

	"github.com/gofiber/fiber/v2"
)

func ReportsRouter(app fiber.Router, service reports.Service) {
	reportsGroup := app.Group("/reports")
	reportsGroup.Post("/receipt-by-user", handlers.GetReceiptsByUser(service))
	reportsGroup.Post("/receipt-today-by-user", handlers.GetReceiptsTodayByUser(service))
	reportsGroup.Post("/receipt-by-date", handlers.GetReceiptsByDate(service))
	reportsGroup.Post("/receipt-between-dates", handlers.GetReceiptsBetweenDates(service))
}
