package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/pkg/user"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ListUsers   godoc
// @Summary       List users
// @Description   list avaliable users in the database
// @Tags          user
// @Accept        json
// @Produce       json
// @Success       200  {array}   entities.User
// @Router        /api/v1/user [get]
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

// GetUserById   godoc
// @Summary       Get a user
// @Description   Get a single user by its id
// @Tags          user
// @Accept        json
// @Param			    id  path  number  true  "id of the user to retrieve" 
// @Produce       json
// @Success       200  {array}   entities.User
// @Router        /api/v1/user/{id} [get]
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

// DeleteUserById   godoc
// @Summary       Delete user
// @Description   Delete existing user
// @Tags          user
// @Accept        json
// @Param			    id  path  number  true  "id of the user to delete" 
// @Produce       json
// @Success       200  {array}   entities.User
// @Router        /api/v1/user/{id} [delete]
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

