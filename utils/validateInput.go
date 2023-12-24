package utils

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func ValidateInput(ctx *fiber.Ctx, input interface{}) []string {
  _, err := govalidator.ValidateStruct(input)
  if err != nil {
    return strings.Split(err.Error(), ";")
  }
  return nil
}
