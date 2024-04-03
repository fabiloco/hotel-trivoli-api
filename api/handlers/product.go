package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/product"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ListProducts   godoc
// @Summary       List products
// @Description   list avaliable products in the database
// @Tags          product
// @Accept        json
// @Produce       json
// @Success       200  {array}   entities.Product
// @Router        /api/v1/product [get]
func GetProducts(service product.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		products, error := service.FetchProducts()

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(products))
	}
}

// ListProducts   godoc
// @Summary       Get a product
// @Description   Get a single product by its id
// @Tags          product
// @Accept        json
// @Param			    id  path  number  true  "id of the product to retrieve"
// @Produce       json
// @Success       200  {array}   entities.Product
// @Router        /api/v1/product/{id} [get]
func GetProductById(service product.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		product, error := service.FetchProductById(uint(id))

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

// ListProducts   godoc
// @Summary       Create products
// @Description   Create new products
// @Tags          product
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product",\n"price": 2000,\n"stock": 100,\n"type": "test type"\n})
// @Produce       json
// @Success       200  {array}   entities.Product
// @Router        /api/v1/product [post]
func PostProducts(service product.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body entities.CreateProduct

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		product, error := service.InsertProduct(&body)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

// ListProducts   godoc
// @Summary       Restock products
// @Description   Change the stock amount of a product b id
// @Tags          product
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product",\n"price": 2000,\n"stock": 100,\n"type": "test type"\n})
// @Produce       json
// @Success       200  {array}   entities.Product
// @Router        /api/v1/product [post]
func PostRestockProducts(service product.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body entities.RestockProduct

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

		product, error := service.RestockProduct(uint(id), &body)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

// ListProducts   godoc
// @Summary       Update products
// @Description   Edit existing product
// @Tags          product
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product",\n"price": 2000,\n"stock": 100,\n"type": "test type"\n})
// @Param			    id  path  number  true  "id of the product to update"
// @Produce       json
// @Success       200  {array}   entities.Product
// @Router        /api/v1/product/{id} [put]
func PutProduct(service product.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body entities.UpdateProduct

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

		product, error := service.UpdateProduct(uint(id), &body)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

// ListProducts   godoc
// @Summary       Delete products
// @Description   Delete existing product
// @Tags          product
// @Accept        json
// @Param			    id  path  number  true  "id of the product to delete"
// @Produce       json
// @Success       200  {array}   entities.Product
// @Router        /api/v1/product/{id} [delete]
func DeleteProductById(service product.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		product, error := service.RemoveProduct(uint(id))

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}
