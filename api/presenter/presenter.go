package presenter

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserResponse struct {
	ID          uint            `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Username    string          `gorm:"not null" json:"username"`
  Role        entities.Role   `gorm:"not null" json:"role"`
  RoleID      uint            `gorm:"not null"`
  Person      entities.Person `gorm:"not null" json:"person"`
  PersonID    uint            `gorm:"not null"`
}

func SuccessResponse(data interface{}) *fiber.Map {
  return &fiber.Map {
    "status": true,
    "error": nil,
    "data": data,
  }
}

func SuccessRegisterResponse(data entities.User) *fiber.Map {
  userResponse := UserResponse{
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


func ErrorResponse(error error) *fiber.Map {
  return &fiber.Map {
    "status": false,
    "error": strings.Split(error.Error(), ", "),
    "data": "",
  }
}
