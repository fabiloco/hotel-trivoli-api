package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/reports"
	"net/http"
	"strings"
	"github.com/gofiber/fiber/v2"
)

type ReceiptsByDate struct {
  Date  string  `valid:"required,rfc3339" json:"date"`
}

type ReceiptsBetweenDates struct {
	StartDate string  `valid:"required,rfc3339" json:"start_date"`
	EndDate   string  `valid:"required,rfc3339" json:"end_date"`
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

		return ctx.JSON(presenter.SuccessResponse(receipts))
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

		return ctx.JSON(presenter.SuccessResponse(receipts))
	}
}

