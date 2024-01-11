package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	roomHistory "fabiloco/hotel-trivoli-api/pkg/room_history"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ListRoomHistorys   godoc
// @Summary       List product types
// @Description   list avaliable product types in the database
// @Tags          product type
// @Accept        json
// @Produce       json
// @Success       200  {array}   model.RoomHistory
// @Router        /room-history [get]
func GetRoomHistorys(service roomHistory.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    roomHistorys, error := service.FetchRoomHistorys()

    if error != nil {
      ctx.Status(http.StatusInternalServerError)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(roomHistorys))
  }
}

// GetRoomHistoryById   godoc
// @Summary       Get a product type
// @Description   Get a single product type by its id
// @Tags          product type
// @Accept        json
// @Param			    id  path  number  true  "id of the product type to retrieve" 
// @Produce       json
// @Success       200  {array}   model.RoomHistory
// @Router        /room-history/{id} [get]
func GetRoomHistoryById(service roomHistory.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    id, err := ctx.ParamsInt("id")
    if err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
    }

    product, error := service.FetchRoomHistoryById(uint(id))

    if error != nil {
      ctx.Status(http.StatusInternalServerError)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}

// PostRoomHistory   godoc
// @Summary       Create a product type
// @Description   Create new product types
// @Tags          product type
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product type"})
// @Produce       json
// @Success       200  {array}   model.RoomHistory
// @Router        /room-history [post]
func PostRoomHistorys(service roomHistory.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body entities.CreateRoomHistory

    if err := ctx.BodyParser(&body); err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(err))
    }
    validationErrors := utils.ValidateInput(ctx, body)

    if validationErrors != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
    }

    product, error := service.InsertRoomHistory(&body)

    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}


// PutRoomHistory   godoc
// @Summary       Update product type
// @Description   Edit existing product type
// @Tags          product type
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product"})
// @Param			    id  path  number  true  "id of the product type to update" 
// @Produce       json
// @Success       200  {array}   model.RoomHistory
// @Router        /room-history/{id} [put]
func PutRoomHistory(service roomHistory.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body entities.UpdateRoomHistory

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

    product, error := service.UpdateRoomHistory(uint(id), &body)

    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}




// DeleteRoomHistoryById   godoc
// @Summary       Delete product type
// @Description   Delete existing product type
// @Tags          product type
// @Accept        json
// @Param			    id  path  number  true  "id of the product type to delete" 
// @Produce       json
// @Success       200  {array}   model.RoomHistory
// @Router        /room-history/{id} [delete]
func DeleteRoomHistoryById(service roomHistory.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    id, err := ctx.ParamsInt("id")
    if err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
    }
    
    product, error := service.RemoveRoomHistory(uint(id))

    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}

