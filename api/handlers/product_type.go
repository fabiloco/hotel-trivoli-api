package handlers

import (
	productType "fabiloco/hotel-trivoli-api/pkg/product_type"
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
func GetProductTypes(service productType.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    products, error := service.FetchProductTypes()

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
func GetProductTypeById(service productType.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    id, err := ctx.ParamsInt("id")
    if err != nil {
      ctx.Locals("data", fiber.Map{
        "errors": err.Error(),
      })
      return ctx.SendStatus(fiber.StatusBadRequest)
    }

    product, error := service.FetchProductTypeById(uint(id))

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
func PostProductTypes(service productType.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
  //
  // var body entities.CreateProductType
  // 
  // if err := ctx.BodyParser(&body); err != nil {
  //   ctx.Locals("data", fiber.Map{
		// 	"errors": err.Error(),
		// })
  //   return ctx.SendStatus(fiber.StatusBadRequest)
  // }
  //
  // validationErrors := utils.ValidateInput(ctx, body)
  //
  // if validationErrors != nil {
  //   ctx.Locals("data", fiber.Map{
		// 	"errors": validationErrors,
		// })
  //   return ctx.SendStatus(fiber.StatusBadRequest)
  // }
  //
  // var productTypesSlice []entities.ProductTypeType
  //
  // for i := 0; i < len(body.Type); i++{
  //   productType, error := h.productTypeStore.FindById(body.Type[i])
  //
  //   if error != nil {
  //     ctx.Locals("data", fiber.Map{
  //       "errors": error.Error() + fmt.Sprintf(": product type with id: %d", body.Type[i]),
  //     })
  //     return ctx.SendStatus(fiber.StatusNotFound)
  //   }
  //
  //   productTypesSlice = append(productTypesSlice, *productType)
  // }
  //
  // newProductType := model.ProductType {
  //   Name: body.Name,
  //   Price: body.Price,
  //   Stock: body.Stock,
  //   Type: productTypesSlice,
  // }
  //
  // product, error := h.productStore.Create(&newProductType)
  //
  // if error != nil {
  //   ctx.Locals("data", fiber.Map{
		// 	"errors": error.Error(),
		// })
  //   return ctx.SendStatus(fiber.StatusNotFound)
  // }
  //
  // ctx.Locals("data", fiber.Map{
  //   "product": product,
  // })
  return ctx.SendStatus(fiber.StatusCreated)
  }
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
func PutProductType(service productType.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    // var body entities.CreateProductType
    //
    // if err := ctx.BodyParser(&body); err != nil {
    //   ctx.Locals("data", fiber.Map{
    //     "errors": err.Error(),
    //   })
    //   return ctx.SendStatus(fiber.StatusBadRequest)
    // }
    //
    // validationErrors := utils.ValidateInput(ctx, body)
    //
    // if validationErrors != nil {
    //   ctx.Locals("data", fiber.Map{
    //     "errors": validationErrors,
    //   })
    //   return ctx.SendStatus(fiber.StatusBadRequest)
    // }
    // 
    // id, err := ctx.ParamsInt("id")
    //
    // if err != nil {
    //   ctx.Locals("data", fiber.Map{
    //     "errors": err.Error(),
    //   })
    //   return ctx.SendStatus(fiber.StatusBadRequest)
    // }
    //
    // var productTypesSlice []model.ProductTypeType
    //
    // for i := 0; i < len(body.Type); i++{
    //   productType, error := h.productTypeStore.FindById(body.Type[i])
    //
    //   if error != nil {
    //     ctx.Locals("data", fiber.Map{
    //       "errors": error.Error() + fmt.Sprintf(": product type with id: %d", body.Type[i]),
    //     })
    //     return ctx.SendStatus(fiber.StatusNotFound)
    //   }
    //
    //   productTypesSlice = append(productTypesSlice, *productType)
    // }
    //
    // newProductType := model.ProductType {
    //   Name: body.Name,
    //   Price: body.Price,
    //   Stock: body.Stock,
    //   Type: productTypesSlice,
    // }
    //
    //
    // product, error := h.productStore.Update(uint(id), &newProductType)
    //
    // if error != nil {
    //   ctx.Locals("data", fiber.Map{
    //     "errors": error.Error(),
    //   })
    //   return ctx.SendStatus(fiber.StatusBadRequest)
    // }
    //
    // ctx.Locals("data", fiber.Map{
    //   "product": product,
    // })
    //
    return ctx.SendStatus(fiber.StatusCreated)
  }
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
func DeleteProductTypeById(service productType.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    id, err := ctx.ParamsInt("id")

    if err != nil {
      ctx.Locals("data", fiber.Map{
        "errors": err.Error(),
      })
      return ctx.SendStatus(fiber.StatusBadRequest)
    }
    
    product, error := service.RemoveProductType(uint(id))

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
}

