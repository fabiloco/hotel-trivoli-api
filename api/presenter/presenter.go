package presenter

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(data interface{}) *fiber.Map {
  return &fiber.Map {
    "status": true,
    "error": nil,
    "data": data,
  }
}


func ErrorResponse(error error) *fiber.Map {
  return &fiber.Map {
    "status": false,
    "error": strings.Split(error.Error(), ", "),
    "data": "",
  }
}
