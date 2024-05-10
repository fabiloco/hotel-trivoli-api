package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	receipt_presenter "fabiloco/hotel-trivoli-api/api/presenter/receipt"
	shift_presenter "fabiloco/hotel-trivoli-api/api/presenter/shift"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/shift"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ListShifts   godoc
// @Summary       List Shifts
// @Description   list avaliable Shifts in the database
// @Tags          Shift
// @Accept        json
// @Produce       json
// @Success       200  {array}   entities.Shift
// @Router        /api/v1/shift [get]
func GetShifts(service shift.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		receipts, individual_receipts, error := service.FetchAllShifts()

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		// return ctx.JSON(receipt_presenter.SuccessReceiptsResponse(receipts))
		return ctx.JSON(presenter.SuccessResponse(shift_presenter.ReceiptsToShiftsResponse(receipts, individual_receipts)))
	}
}

// ListShifts   godoc
// @Summary       Get a Shift
// @Description   Get a single Shift by its id
// @Tags          Shift
// @Accept        json
// @Param			    id  path  number  true  "id of the Shift to retrieve"
// @Produce       json
// @Success       200  {array}   entities.Shift
// @Router        /api/v1/shift/{id} [get]
func GetShiftById(service shift.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		receipts, individual_receipts, error := service.FetchShiftsById(uint(id))

		fmt.Println(individual_receipts)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(receipt_presenter.SuccessReceiptsResponse(receipts))
	}
}

// ListShifts   godoc
// @Summary       Create Shifts
// @Description   Create new Shifts
// @Tags          Shift
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test Shift",\n"price": 2000,\n"stock": 100,\n"type": "test type"\n})
// @Produce       json
// @Success       200  {array}   entities.Shift
// @Router        /api/v1/shift [post]
func PostShifts(service shift.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body entities.CreateShiftWithCode

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		var receiptIds []uint

		for _, receiptId := range body.Receipts {
			if !utils.ReceiptIdPatternMatch(receiptId) {
				return ctx.JSON(presenter.ErrorResponse(errors.New("An id does not match with the pattern")))
			}

			idParsed, error := utils.ConvertReceiptsId(receiptId)
			if error != nil {
				return ctx.JSON(presenter.ErrorResponse(error))
			}
			receiptIds = append(receiptIds, idParsed)
		}

		var individualReceiptIds []uint

		for _, individualReceiptId := range body.IndividualReceipts {
			if !utils.IndividualReceiptIdPatternMatch(individualReceiptId) {
				return ctx.JSON(presenter.ErrorResponse(errors.New("An id does not match with the pattern")))
			}

			idParsed, error := utils.ConvertReceiptsId(individualReceiptId)
			if error != nil {
				return ctx.JSON(presenter.ErrorResponse(error))
			}
			individualReceiptIds = append(individualReceiptIds, idParsed)
		}

		bodyParsed := entities.CreateShift{
			Receipts:           receiptIds,
			IndividualReceipts: individualReceiptIds,
		}

		shift, error := service.InsertShift(&bodyParsed)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(shift))
	}
}

// ListShifts   godoc
// @Summary       Update Shifts
// @Description   Edit existing Shift
// @Tags          Shift
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test Shift",\n"price": 2000,\n"stock": 100,\n"type": "test type"\n})
// @Param			    id  path  number  true  "id of the Shift to update"
// @Produce       json
// @Success       200  {array}   entities.Shift
// @Router        /api/v1/shift/{id} [put]
func PutShift(service shift.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body entities.UpdateShift

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		shift, error := service.UpdateShift(uint(id), &body)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(shift))
	}
}

// ListShifts   godoc
// @Summary       Delete Shifts
// @Description   Delete existing Shift
// @Tags          Shift
// @Accept        json
// @Param			    id  path  number  true  "id of the Shift to delete"
// @Produce       json
// @Success       200  {array}   entities.Shift
// @Router        /api/v1/shift/{id} [delete]
func DeleteShiftById(service shift.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		shift, error := service.RemoveShift(uint(id))

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(shift))
	}
}
