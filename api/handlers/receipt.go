package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/receipt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ListReceipts   godoc
// @Summary       List product types
// @Description   list avaliable product types in the database
// @Tags          product type
// @Accept        json
// @Produce       json
// @Success       200  {array}   model.Receipt
// @Router        /product-type [get]
func GetReceipts(service receipt.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    receipts, error := service.FetchReceipts()

    if error != nil {
      ctx.Status(http.StatusInternalServerError)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(receipts))
  }
}

// GetReceiptById   godoc
// @Summary       Get a product type
// @Description   Get a single product type by its id
// @Tags          product type
// @Accept        json
// @Param			    id  path  number  true  "id of the product type to retrieve" 
// @Produce       json
// @Success       200  {array}   model.Receipt
// @Router        /product-type/{id} [get]
func GetReceiptById(service receipt.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    id, err := ctx.ParamsInt("id")
    if err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
    }

    product, error := service.FetchReceiptById(uint(id))

    if error != nil {
      ctx.Status(http.StatusInternalServerError)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}

// PostReceipt   godoc
// @Summary       Create a product type
// @Description   Create new product types
// @Tags          product type
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product type"})
// @Produce       json
// @Success       200  {array}   model.Receipt
// @Router        /product-type [post]
func PostReceipts(service receipt.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body entities.CreateReceipt

    if err := ctx.BodyParser(&body); err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(err))
    }
    validationErrors := utils.ValidateInput(ctx, body)

    if validationErrors != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ", "))))
    }

    product, error := service.InsertReceipt(&body)

    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}


// PutReceipt   godoc
// @Summary       Update product type
// @Description   Edit existing product type
// @Tags          product type
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product"})
// @Param			    id  path  number  true  "id of the product type to update" 
// @Produce       json
// @Success       200  {array}   model.Receipt
// @Router        /product-type/{id} [put]
func PutReceipt(service receipt.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body entities.UpdateReceipt

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

    product, error := service.UpdateReceipt(uint(id), &body)

    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}




// DeleteReceiptById   godoc
// @Summary       Delete product type
// @Description   Delete existing product type
// @Tags          product type
// @Accept        json
// @Param			    id  path  number  true  "id of the product type to delete" 
// @Produce       json
// @Success       200  {array}   model.Receipt
// @Router        /product-type/{id} [delete]
func DeleteReceiptById(service receipt.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    id, err := ctx.ParamsInt("id")
    if err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
    }
    
    product, error := service.RemoveReceipt(uint(id))

    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}

