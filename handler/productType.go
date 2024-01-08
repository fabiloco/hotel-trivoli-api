package handler

import (
	"fabiloco/hotel-trivoli-api/model"
	"fabiloco/hotel-trivoli-api/utils"
	"github.com/gofiber/fiber/v2"
)

// ListProductTypes   godoc
// @Summary       List product types
// @Description   list avaliable product types in the database
// @Tags          product type
// @Accept        json
// @Produce       json
// @Success       200  {array}   model.ProductType
// @Router        /product-type [get]
func (h *Handler) ListProductTypes(ctx *fiber.Ctx) error {
  productTypes, error := h.productTypeStore.List()

  if error != nil {
    ctx.Locals("data", fiber.Map{
			"errors": error.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  ctx.Locals("data", fiber.Map{
    "productTypes": productTypes,
  })

  return ctx.SendStatus(fiber.StatusOK)
}

// GetProductTypeById   godoc
// @Summary       Get a product type
// @Description   Get a single product type by its id
// @Tags          product type
// @Accept        json
// @Param			    id  path  number  true  "id of the product type to retrieve" 
// @Produce       json
// @Success       200  {array}   model.ProductType
// @Router        /product-type/{id} [get]
func (h *Handler) GetProductTypeById(ctx *fiber.Ctx) error {
  id, err := ctx.ParamsInt("id")
  if err != nil {
    ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  productType, error := h.productTypeStore.FindById(uint(id))

  if error != nil {
    ctx.Locals("data", fiber.Map{
			"errors": error.Error(),
		})
    return ctx.SendStatus(fiber.StatusNotFound)
  }

  ctx.Locals("data", fiber.Map{
    "productType": productType,
  })

  return ctx.SendStatus(fiber.StatusOK)
}

// PostProductType   godoc
// @Summary       Create a product type
// @Description   Create new product types
// @Tags          product type
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product type"})
// @Produce       json
// @Success       200  {array}   model.ProductType
// @Router        /product-type [post]
func (h *Handler) PostProductType(ctx *fiber.Ctx) error {
  var body model.CreateProductType
  
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

  productType, error := h.productTypeStore.Create(&body)

  if error != nil {
    ctx.Locals("data", fiber.Map{
			"errors": error.Error(),
		})
    return ctx.SendStatus(fiber.StatusNotFound)
  }

  ctx.Locals("data", fiber.Map{
    "productType": productType,
  })

  return ctx.SendStatus(fiber.StatusCreated)
}


// PutProductType   godoc
// @Summary       Update product type
// @Description   Edit existing product type
// @Tags          product type
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product"})
// @Param			    id  path  number  true  "id of the product type to update" 
// @Produce       json
// @Success       200  {array}   model.ProductType
// @Router        /product-type/{id} [put]
func (h *Handler) PutProductType(ctx *fiber.Ctx) error {
  var body model.CreateProductType

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
  
  id, err := ctx.ParamsInt("id")

  if err != nil {
    ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  productType, error := h.productTypeStore.Update(uint(id), &body)

  if error != nil {
    ctx.Locals("data", fiber.Map{
			"errors": error.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  ctx.Locals("data", fiber.Map{
    "productType": productType,
  })

  return ctx.SendStatus(fiber.StatusCreated)
}




// DeleteProductTypeById   godoc
// @Summary       Delete product type
// @Description   Delete existing product type
// @Tags          product type
// @Accept        json
// @Param			    id  path  number  true  "id of the product type to delete" 
// @Produce       json
// @Success       200  {array}   model.ProductType
// @Router        /product-type/{id} [delete]
func (h *Handler) DeleteProductTypeById(ctx *fiber.Ctx) error {
  id, err := ctx.ParamsInt("id")

  if err != nil {
    ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }
  
  productType, error := h.productTypeStore.Delete(uint(id))

  if error != nil {
    ctx.Locals("data", fiber.Map{
			"errors": error.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  ctx.Locals("data", fiber.Map{
    "productType": productType,
  })

  return ctx.SendStatus(fiber.StatusOK)
}

