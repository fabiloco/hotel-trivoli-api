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

// PostUser   godoc
// @Summary       Create a user
// @Description   Create new users
// @Tags          user
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test user"})
// @Produce       json
// @Success       200  {array}   entities.User
// @Router        /api/v1/user [post]
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

// PutUser   godoc
// @Summary       Update user
// @Description   Edit existing user
// @Tags          user
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product"})
// @Param			    id  path  number  true  "id of the user to update" 
// @Produce       json
// @Success       200  {array}   entities.User
// @Router        /api/v1/user/{id} [put]
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

