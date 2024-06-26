package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	receipt_presenter "fabiloco/hotel-trivoli-api/api/presenter/receipt"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/reports"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ReceiptsByDate struct {
	Date string `valid:"required,rfc3339" json:"date"`
}

type ReceiptsByUser struct {
	UserID uint `valid:"required,numeric" json:"user_id"`
}

type ReceiptsBetweenDates struct {
	StartDate string `valid:"required,rfc3339" json:"start_date"`
	EndDate   string `valid:"required,rfc3339" json:"end_date"`
}

// ListReceiptsByDate   godoc
// @Summary       Receipts by date
// @Description   Report that shows the receipts created at a certain date
// @Tags          receipt
// @Accept        json
// @Produce       json
// @Success       200  {array}   ReceiptsByDate
// @Router        /api/v1/reports/receipt-by-date [get]
func GetReceiptsByDate(service reports.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body ReceiptsByDate

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		receipts, error := service.ReceiptByTargetDate(body.Date)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		individualReceipts, error := service.IndividualReceiptByTargetDate(body.Date)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		response := fiber.Map{
			"receipts":           receipt_presenter.SuccessReceiptsResponse(receipts),
			"individualReceipts": receipt_presenter.SuccessIndividualReceiptsResponse(individualReceipts),
		}

		return ctx.JSON(response)
	}
}

// ListReceiptsByDate   godoc
// @Summary       Receipts by date
// @Description   Report that shows the receipts created at a certain date
// @Tags          receipt
// @Accept        json
// @Produce       json
// @Success       200  {array}   ReceiptsByDate
// @Router        /api/v1/reports/receipt-by-date [get]
func GetReceiptsByUser(service reports.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body ReceiptsByUser

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		receipts, error := service.ReceiptByUser(body.UserID)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		individualReceipts, error := service.IndividualReceiptByUser(body.UserID)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		response := fiber.Map{
			"receipts":           receipt_presenter.SuccessReceiptsResponse(receipts),
			"individualReceipts": receipt_presenter.SuccessIndividualReceiptsResponse(individualReceipts),
		}

		return ctx.JSON(response)
	}
}

// ListReceiptsByDate   godoc
// @Summary       Receipts by date
// @Description   Report that shows the receipts created at a certain date
// @Tags          receipt
// @Accept        json
// @Produce       json
// @Success       200  {array}   ReceiptsByDate
// @Router        /api/v1/reports/receipt-by-date [get]
func GetReceiptsTodayByUser(service reports.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body ReceiptsByUser

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		receipts, error := service.ReceiptTodayByUser(body.UserID)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		individualReceipts, error := service.IndividualReceiptTodayByUser(body.UserID)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		response := fiber.Map{
			"receipts":           receipt_presenter.SuccessReceiptsResponse(receipts),
			"individualReceipts": receipt_presenter.SuccessIndividualReceiptsResponse(individualReceipts),
		}

		return ctx.JSON(response)
	}
}

// ListReceiptsBetweenDates   godoc
// @Summary       Receipts between dates
// @Description   Report that shows the receipts created between a range of dates
// @Tags          receipt
// @Accept        json
// @Produce       json
// @Success       200  {array}   ReceiptsBetweenDates
// @Router        /api/v1/reports/receipt-between-dates [get]
func GetReceiptsBetweenDates(service reports.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body ReceiptsBetweenDates

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		receipts, error := service.ReceiptsBetweenDates(body.StartDate, body.EndDate)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		individualReceipts, error := service.IndividualReceiptsBetweenDates(body.StartDate, body.EndDate)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		response := fiber.Map{
			"receipts":           receipt_presenter.SuccessReceiptsResponse(receipts),
			"individualReceipts": receipt_presenter.SuccessIndividualReceiptsResponse(individualReceipts),
		}

		return ctx.JSON(response)
	}
}

type TotalReceipt struct {
	TotalProduct  float64 `json:"total_products"`
	TotalServices float64 `json:"total_services"`
}

type TotalIndividualReceipt struct {
	TotalProduct float64 `json:"total_products"`
}

func GetTotalBetweenDates(service reports.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body ReceiptsBetweenDates

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		receipts, error := service.ReceiptsBetweenDates(body.StartDate, body.EndDate)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		individualReceipts, error := service.IndividualReceiptsBetweenDates(body.StartDate, body.EndDate)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		receiptsResponses := receipt_presenter.ReceiptsToReceiptsResponses(*receipts)
		var receiptsTotalProducts float64 = 0
		var receiptsTotalService float64 = 0

		for _, receiptResponse := range receiptsResponses {
			for _, product := range receiptResponse.Products {
				receiptsTotalProducts += float64(product.Price) * float64(product.Quantity)
			}

			receiptsTotalService += float64(receiptResponse.Service.Price)
		}

		receiptsTotalBetweenDates := TotalReceipt{
			TotalProduct:  receiptsTotalProducts,
			TotalServices: receiptsTotalService,
		}

		individualReceiptsResponses := receipt_presenter.IndividualReceiptsToIndividualReceiptsResponses(*individualReceipts)
		var individualReceiptsTotalProducts float64 = 0

		for _, receiptResponse := range individualReceiptsResponses {
			for _, product := range receiptResponse.Products {
				individualReceiptsTotalProducts += float64(product.Price) * float64(product.Quantity)
			}
		}

		individualReceiptsTotalBetweenDates := TotalIndividualReceipt{
			TotalProduct: individualReceiptsTotalProducts,
		}

		response := fiber.Map{
			"receipts":           receiptsTotalBetweenDates,
			"individualReceipts": individualReceiptsTotalBetweenDates,
		}

		return ctx.JSON(response)
	}
}
