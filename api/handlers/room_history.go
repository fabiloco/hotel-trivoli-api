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
// @Summary       List room historys
// @Description   list avaliable room historys in the database
// @Tags          room history
// @Accept        json
// @Produce       json
// @Success       200  {array}   entities.RoomHistory
// @Router        /api/v1/room-history [get]
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
// @Summary       Get a room history
// @Description   Get a single room history by its id
// @Tags          room history
// @Accept        json
// @Param			    id  path  number  true  "id of the room history to retrieve" 
// @Produce       json
// @Success       200  {array}   entities.RoomHistory
// @Router        /api/v1/room-history/{id} [get]
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
// @Summary       Create a room history
// @Description   Create new room historys
// @Tags          room history
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test room history"})
// @Produce       json
// @Success       200  {array}   entities.RoomHistory
// @Router        /api/v1/room-history [post]
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
      return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ", ") )))
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
// @Summary       Update room history
// @Description   Edit existing room history
// @Tags          room history
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product"})
// @Param			    id  path  number  true  "id of the room history to update" 
// @Produce       json
// @Success       200  {array}   entities.RoomHistory
// @Router        /api/v1/room-history/{id} [put]
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
// @Summary       Delete room history
// @Description   Delete existing room history
// @Tags          room history
// @Accept        json
// @Param			    id  path  number  true  "id of the room history to delete" 
// @Produce       json
// @Success       200  {array}   entities.RoomHistory
// @Router        /api/v1/room-history/{id} [delete]
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

