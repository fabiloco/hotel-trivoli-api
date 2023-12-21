package handler

import (
	"fabiloco/hotel-trivoli-api/database"
	"fabiloco/hotel-trivoli-api/model"
	"github.com/gofiber/fiber/v2"
)

// ListProducts   godoc
// @Summary       List products
// @Description   list avaliable products in the database
// @Tags          product
// @Accept        json
// @Produce       json
// @Success       200  {array}   model.Product
// @Router        /product [get]
func GetProducts(ctx *fiber.Ctx) error {
  var products []model.Product
  
  result := database.DB.Find(&products)

  if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errorMessage": "Error trying retreive products from the database.",
      "errors":  result.Error.Error(),
		})
  }

  return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
    "products": products,
    "rowsAffected": result.RowsAffected,
  })
}
