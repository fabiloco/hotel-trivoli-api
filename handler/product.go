package handler

import (
	"fabiloco/hotel-trivoli-api/database"
	"fabiloco/hotel-trivoli-api/model"
	"fabiloco/hotel-trivoli-api/utils"

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

    ctx.Locals("data", fiber.Map{
			"errors": result.Error.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)

		// return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 	"errorMessage": "Error trying retreive products from the database.",
  //     "errors":  result.Error.Error(),
		// })
  }

  ctx.Locals("data", fiber.Map{
    "products": products,
  })

  return ctx.SendStatus(fiber.StatusOK)
}

// ListProducts   godoc
// @Summary       Create products
// @Description   Create new products
// @Tags          product
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product",\n"price": 2000,\n"stock": 100,\n"type": "test type"\n})
// @Produce       json
// @Success       200  {array}   model.Product
// @Router        /product [post]
func PostProducts(ctx *fiber.Ctx) error {
  var body model.CreateProduct
  
  if err := ctx.BodyParser(&body); err != nil {
    ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  validationErrors := utils.ValidateInput(ctx, body)

  if validationErrors != nil {
    ctx.Locals("data", fiber.Map{
			"errors": validationErrors,
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  product := model.Product {
    Name: body.Name,
    Stock: body.Stock,
    Price: body.Price,
    Type: body.Type,
  }

  database.DB.Create(&product)

  ctx.Locals("data", fiber.Map{
    "product": body,
  })

  return ctx.SendStatus(fiber.StatusCreated)
}


func PutProduct(ctx *fiber.Ctx) error {
  var body model.UpdateProduct
  
  if err := ctx.BodyParser(&body); err != nil {
    ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  validationErrors := utils.ValidateInput(ctx, body)

  if validationErrors != nil {
    ctx.Locals("data", fiber.Map{
			"errors": validationErrors,
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  product := model.Product {}

  result := database.DB.First(&product, 2)

  if result.Error != nil {

  }

  ctx.Locals("data", fiber.Map{
    "product": product,
  })

  return ctx.SendStatus(fiber.StatusCreated)
}
