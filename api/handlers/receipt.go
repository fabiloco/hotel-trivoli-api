package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	receipt_presenter "fabiloco/hotel-trivoli-api/api/presenter/receipt"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/receipt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PostReceipt   godoc
// @Summary       Create a receipt
// @Description   Create new receipts
// @Tags          receipt
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test receipt"})
// @Produce       json
// @Success       200  {array}   entities.Receipt
// @Router        /api/v1/receipt [post]
func GenerateReceipts(service receipt.Service) fiber.Handler {
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


    receipt, error := service.GenerateReceipt(&body)



    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(receipt_presenter.SuccessReceiptResponse(receipt))
  }
}

// PostReceipt   godoc
// @Summary       Create a receipt
// @Description   Create new receipts
// @Tags          receipt
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test receipt"})
// @Produce       json
// @Success       200  {array}   entities.Receipt
// @Router        /api/v1/receipt [post]
func GenerateIndividualReceipts(service receipt.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body entities.CreateIndividualReceipt

    if err := ctx.BodyParser(&body); err != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(err))
    }
    validationErrors := utils.ValidateInput(ctx, body)

    if validationErrors != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ", "))))
    }

    // return ctx.JSON(&fiber.Map{
    //   "test":"test",
    // })

    product, error := service.GenerateIndividualReceipt(&body)


    if error != nil {
      ctx.Status(http.StatusBadRequest)
      return ctx.JSON(presenter.ErrorResponse(error))
    }

    return ctx.JSON(presenter.SuccessResponse(product))
  }
}


// ListReceipts   godoc
// @Summary       List receipts
// @Description   list avaliable receipts in the database
// @Tags          receipt
// @Accept        json
// @Produce       json
// @Success       200  {array}   entities.Receipt
// @Router        /api/v1/receipt [get]
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
// @Summary       Get a receipt
// @Description   Get a single receipt by its id
// @Tags          receipt
// @Accept        json
// @Param			    id  path  number  true  "id of the receipt to retrieve" 
// @Produce       json
// @Success       200  {array}   entities.Receipt
// @Router        /api/v1/receipt/{id} [get]
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
// @Summary       Create a receipt
// @Description   Create new receipts
// @Tags          receipt
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test receipt"})
// @Produce       json
// @Success       200  {array}   entities.Receipt
// @Router        /api/v1/receipt [post]
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
// @Summary       Update receipt
// @Description   Edit existing receipt
// @Tags          receipt
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product"})
// @Param			    id  path  number  true  "id of the receipt to update" 
// @Produce       json
// @Success       200  {array}   entities.Receipt
// @Router        /api/v1/receipt/{id} [put]
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
// @Summary       Delete receipt
// @Description   Delete existing receipt
// @Tags          receipt
// @Accept        json
// @Param			    id  path  number  true  "id of the receipt to delete" 
// @Produce       json
// @Success       200  {array}   entities.Receipt
// @Router        /api/v1/receipt/{id} [delete]
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

