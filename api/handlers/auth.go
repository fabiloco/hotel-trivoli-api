package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/auth"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

type LoginUser struct {
  Username        string  `valid:"required,stringlength(3|100)" json:"username"`
	Password        string  `valid:"required" json:"password"`
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

    fmt.Println("test")
    fmt.Println(body.Lastname)

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

    passwordHashed, error := utils.HashPassword(body.Password)

    if error != nil {
      ctx.Status(http.StatusInternalServerError)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    newUser := entities.CreateUser{
      Username: body.Username,
      Password: passwordHashed,
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

func Login(service auth.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body LoginUser

    if err := ctx.BodyParser(&body); err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(err))
    }
    validationErrors := utils.ValidateInput(ctx, body)

    if validationErrors != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ", "))))
    }

    user, error := service.Login(body.Username)

    if error != nil {
      ctx.Status(http.StatusInternalServerError)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    passwordMatch := utils.CheckPasswordHash(body.Password, user.Password)

    if passwordMatch == false {
      ctx.Status(http.StatusUnauthorized)
      return ctx.JSON(presenter.ErrorResponse(errors.New(fmt.Sprint("Incorrect password"))))
    }


    userClaims := utils.UserClaims{
      Id: user.ID,
      Username: user.Username,
      Role: user.Role.Name,
      StandardClaims: jwt.StandardClaims{
      IssuedAt:  time.Now().Unix(),
      ExpiresAt: time.Now().Add(time.Hour * 128).Unix(),
     },
    }

    fmt.Println(userClaims)

    signedAccessToken, err := utils.NewAccessToken(userClaims)
    if err != nil {
      log.Fatal("error creating access token")
    }

    return ctx.JSON(presenter.SuccessLoginResponse(signedAccessToken, userClaims.StandardClaims, *user))
  }
}

func Verify(service auth.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {

    var user = ctx.Locals("userClaims")

    userClaims, ok := user.(*utils.UserClaims)

    if !ok {
      ctx.Status(http.StatusUnauthorized)
      return ctx.JSON(presenter.ErrorResponse(errors.New(fmt.Sprint("Incorrect password"))))
    }

    userClaims.StandardClaims.IssuedAt = time.Now().Unix()
    userClaims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

    signedAccessToken, err := utils.NewAccessToken(*userClaims)
    if err != nil {
      log.Fatal("error creating access token")
    }

    userWithUsername, error := service.Login(userClaims.Username)

    if error != nil {
      ctx.Status(http.StatusInternalServerError)
      return ctx.JSON(presenter.ErrorResponse(error))
    }


    return ctx.JSON(presenter.SuccessLoginResponse(signedAccessToken, userClaims.StandardClaims, *userWithUsername))
  }
}

