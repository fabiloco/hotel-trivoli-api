package routes

import (
	"fabiloco/hotel-trivoli-api/api/handlers"
	"fabiloco/hotel-trivoli-api/pkg/receipt"
	"fabiloco/hotel-trivoli-api/pkg/shift"

	"github.com/gofiber/fiber/v2"
)

func ReceiptRouter(app fiber.Router, service receipt.Service, shiftService shift.Service) {
	receiptGroup := app.Group("/receipt")
	receiptGroup.Get("/", handlers.GetReceipts(service))
	receiptGroup.Get("/:id", handlers.GetReceiptById(service))
	receiptGroup.Post("/", handlers.PostReceipts(service))
	receiptGroup.Put("/:id", handlers.PutReceipt(service))
	receiptGroup.Delete("/:id", handlers.DeleteReceiptById(service))

	receiptGroup.Post("/generate", handlers.GenerateReceipts(service))
	receiptGroup.Post("/generate-individual", handlers.GenerateIndividualReceipts(service))

	receiptGroup.Post("/print-receipts", handlers.PrintReceipts(service, shiftService))

	receiptGroup.Post("/print-shift/:id", handlers.PrintShift(service, shiftService))
	receiptGroup.Post("/print-receipt/:id", handlers.PrintReceipt(service))
	receiptGroup.Post("/print-individual-receipt/:id", handlers.PrintIndividualReceipt(service))
}
