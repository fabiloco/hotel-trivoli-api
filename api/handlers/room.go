package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/room"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetRooms(service room.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		rooms, err := service.FetchRooms()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(err))
		}

		return ctx.JSON(presenter.SuccessResponse(rooms))
	}
}

func GetRoomById(service room.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		room, err := service.FetchRoomById(uint(id))

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(err))
		}

		return ctx.JSON(presenter.SuccessResponse(room))
	}
}

func PostRooms(service room.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body entities.CreateRoom

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		product, err := service.InsertRoom(&body)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

func PutRoom(service room.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body entities.UpdateRoom

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

		product, err := service.UpdateRoom(uint(id), &body)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

func DeleteRoomById(service room.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		room, err := service.RemoveRoom(uint(id))

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}

		return ctx.JSON(presenter.SuccessResponse(room))
	}
}
