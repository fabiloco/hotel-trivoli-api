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

// ListRooms   godoc
// @Summary       List rooms
// @Description   list avaliable rooms in the database
// @Tags          room
// @Accept        json
// @Produce       json
// @Success       200  {array}   entities.Room
// @Router        /api/v1/room [get]
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

// GetRoomById   godoc
// @Summary       Get a room
// @Description   Get a single room by its id
// @Tags          room
// @Accept        json
// @Param			    id  path  number  true  "id of the room to retrieve" 
// @Produce       json
// @Success       200  {array}   entities.Room
// @Router        /api/v1/room/{id} [get]
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

// PostRoom   godoc
// @Summary       Create a room
// @Description   Create new rooms
// @Tags          room
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test room"})
// @Produce       json
// @Success       200  {array}   entities.Room
// @Router        /api/v1/room [post]
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

// PutRoom   godoc
// @Summary       Update room
// @Description   Edit existing room
// @Tags          room
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product"})
// @Param			    id  path  number  true  "id of the room to update" 
// @Produce       json
// @Success       200  {array}   entities.Room
// @Router        /api/v1/room/{id} [put]
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

// DeleteRoomById   godoc
// @Summary       Delete room
// @Description   Delete existing room
// @Tags          room
// @Accept        json
// @Param			    id  path  number  true  "id of the room to delete" 
// @Produce       json
// @Success       200  {array}   entities.Room
// @Router        /api/v1/room/{id} [delete]
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
