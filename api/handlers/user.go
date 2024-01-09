package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/user"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(service user.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    users, error := service.FetchUsers()

    if error != nil {
      ctx.Status(http.StatusInternalServerError)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(users))
  }
}

func GetUserById(service user.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    id, err := ctx.ParamsInt("id")
    if err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
    }

    user, error := service.FetchUserById(uint(id))

    if error != nil {
      ctx.Status(http.StatusInternalServerError)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(user))
  }
}

func PostUsers(service user.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body entities.CreateUser

    if err := ctx.BodyParser(&body); err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(err))
    }
    validationErrors := utils.ValidateInput(ctx, body)

    if validationErrors != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
    }

    product, error := service.InsertUser(&body)
    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}

func PutUser(service user.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body entities.UpdateUser

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

    product, error := service.UpdateUser(uint(id), &body)

    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}

func DeleteUserById(service user.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    id, err := ctx.ParamsInt("id")
    if err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
    }
    
    user, error := service.RemoveUser(uint(id))

    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(user))
  }
}

