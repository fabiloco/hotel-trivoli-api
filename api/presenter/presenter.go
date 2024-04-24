package presenter

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthResponse struct {
	ID          uint                `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Username    string              `gorm:"not null" json:"username"`
  Role        entities.Role       `gorm:"not null" json:"role"`
  Person      entities.Person     `gorm:"not null" json:"person"`
}

type TokenResponse struct {
  Token       string              `json:"token"` 
  Claims      jwt.StandardClaims  `json:"claims"`
}

func SuccessResponse(data interface{}) *fiber.Map {
  return &fiber.Map {
    "status": true,
    "error": nil,
    "data": data,
  }
}

func SuccessRegisterResponse(data entities.User) *fiber.Map {
  userResponse := AuthResponse{
    ID: data.ID,
    CreatedAt: data.CreatedAt,
    UpdatedAt: data.UpdatedAt,

    Username: data.Username,
    Role: data.Role,
    Person: data.Person,
  }

  return &fiber.Map {
    "status": true,
    "error": nil,
    "data": userResponse,
  }
}

func SuccessLoginResponse(token string, claims jwt.StandardClaims, data entities.User) *fiber.Map {
  userResponse := AuthResponse{
    ID: data.ID,
    CreatedAt: data.CreatedAt,
    UpdatedAt: data.UpdatedAt,

    Username: data.Username,
    Role: data.Role,
    Person: data.Person,
  }

  jwtToken := TokenResponse {
    Token: token,
    Claims: claims, 
  }

  return &fiber.Map {
    "status": true,
    "error": nil,
    "data": userResponse,
    "jwt": jwtToken,
  }
}


func ErrorResponse(error error) *fiber.Map {
  return &fiber.Map {
    "status": false,
    "error": strings.Split(error.Error(), ", "),
    "data": "",
  }
}
