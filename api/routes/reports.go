package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	"fabiloco/hotel-trivoli-api/pkg/reports"
	"github.com/gofiber/fiber/v2"
)

func ReportsRouter(app fiber.Router, service reports.Service) {
  reportsGroup := app.Group("/reports")
	reportsGroup.Get("/receipt-by-date", handlers.GetReceiptsByDate(service))
	reportsGroup.Get("/receipt-between-dates", handlers.GetReceiptsBetweenDates(service))
}
