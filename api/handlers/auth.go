package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/auth"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type RegisterUser struct {
  Username        string  `valid:"required,stringlength(3|100)" json:"username"`
	Password        string  `valid:"required,stringlength(5|40)" json:"password"`

  Firstname       string  `valid:"required,stringlength(3|100)" json:"firstname"`
  Lastname        string  `valid:"required,stringlength(3|100)" json:"lastname"`
  Identification  string  `valid:"required,numeric" json:"identification"`
  Birthday        string  `valid:"required" json:"birthday" json:"brithday"`

  Role            uint    `valid:"optional,numeric" json:"role"`
}

// Register   godoc
// @Summary       register new users in the system
// @Description   register
// @Tags          user, person
// @Accept        json
// @Produce       json
// @Success       200  {array}   entities.User
// @Router        /api/v1/register [get]
func Register(service auth.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body RegisterUser

    if err := ctx.BodyParser(&body); err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(err))
    }
    validationErrors := utils.ValidateInput(ctx, body)

    if validationErrors != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ", "))))
    }

    newPerson := entities.CreatePerson{
      Firstname: body.Firstname,
      Lastname: body.Lastname,
      Birthday: body.Birthday,
      Identification: body.Identification,
    }

    newUser := entities.CreateUser{
      Username: body.Username,
      Password: body.Password,
      Role: body.Role,
    }

    users, error := service.Register(&newUser, &newPerson)

    if error != nil {
      ctx.Status(http.StatusInternalServerError)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessRegisterResponse(*users))
  }
}
