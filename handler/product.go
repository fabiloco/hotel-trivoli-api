package handler

import (
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
func (h *Handler) ListProducts(ctx *fiber.Ctx) error {
  products, error := h.productStore.List()

  if error != nil {
    ctx.Locals("data", fiber.Map{
			"errors": error.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  ctx.Locals("data", fiber.Map{
    "products": products,
  })

  return ctx.SendStatus(fiber.StatusOK)
}

// ListProducts   godoc
// @Summary       Get a product
// @Description   Get a single product by its id
// @Tags          product
// @Accept        json
// @Param			    id  path  number  true  "id of the product to retrieve" 
// @Produce       json
// @Success       200  {array}   model.Product
// @Router        /product/{id} [get]
func (h *Handler) GetProductById(ctx *fiber.Ctx) error {
  id, err := ctx.ParamsInt("id")
  if err != nil {
    ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  product, error := h.productStore.FindById(id)

  if error != nil {
    ctx.Locals("data", fiber.Map{
			"errors": error.Error(),
		})
    return ctx.SendStatus(fiber.StatusNotFound)
  }

  ctx.Locals("data", fiber.Map{
    "product": product,
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
func (h *Handler) PostProducts(ctx *fiber.Ctx) error {
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

  product, error := h.productStore.Create(&body)

  if error != nil {
    ctx.Locals("data", fiber.Map{
			"errors": error.Error(),
		})
    return ctx.SendStatus(fiber.StatusNotFound)
  }

  ctx.Locals("data", fiber.Map{
    "product": product,
  })

  return ctx.SendStatus(fiber.StatusCreated)
}


// ListProducts   godoc
// @Summary       Update products
// @Description   Edit existing product
// @Tags          product
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product",\n"price": 2000,\n"stock": 100,\n"type": "test type"\n})
// @Param			    id  path  number  true  "id of the product to update" 
// @Produce       json
// @Success       200  {array}   model.Product
// @Router        /product/{id} [put]
func (h *Handler) PutProduct(ctx *fiber.Ctx) error {
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
  
  id, err := ctx.ParamsInt("id")

  if err != nil {
    ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  product, error := h.productStore.Update(id, &body)

  if error != nil {
    ctx.Locals("data", fiber.Map{
			"errors": error.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  ctx.Locals("data", fiber.Map{
    "product": product,
  })

  return ctx.SendStatus(fiber.StatusCreated)
}




// ListProducts   godoc
// @Summary       Delete products
// @Description   Delete existing product
// @Tags          product
// @Accept        json
// @Param			    id  path  number  true  "id of the product to delete" 
// @Produce       json
// @Success       200  {array}   model.Product
// @Router        /product/{id} [delete]
func (h *Handler) DeleteProductById(ctx *fiber.Ctx) error {
  id, err := ctx.ParamsInt("id")

  if err != nil {
    ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }
  
  product, error := h.productStore.Delete(id)

  if error != nil {
    ctx.Locals("data", fiber.Map{
			"errors": error.Error(),
		})
    return ctx.SendStatus(fiber.StatusBadRequest)
  }

  ctx.Locals("data", fiber.Map{
    "product": product,
  })

  return ctx.SendStatus(fiber.StatusOK)
}

