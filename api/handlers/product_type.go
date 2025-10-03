package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	productType "fabiloco/hotel-trivoli-api/pkg/product_type"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ListProductTypes   godoc
// @Summary       List product types
// @Description   list avaliable product types in the database
// @Tags          product type
// @Accept        json
// @Produce       json
// @Success       200  {array}   entities.ProductType
// @Router        /api/v1/product-type [get]
func GetProductTypes(service productType.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		productTypes, error := service.FetchProductTypes()

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(productTypes))
	}
}

// ListProductTypesPaginated   godoc
// @Summary       List product types paginated
// @Description   list avaliable product types in the database with pagination
// @Tags          product type
// @Accept        json
// @Produce       json
// @Param         page      query  int  false  "Page number (default: 1)"
// @Param         pageSize  query  int  false  "Page size (default: 10, max: 100)"
// @Success       200  {object}   entities.PaginatedResponse
// @Router        /api/v1/product-type/paginated [get]
func GetProductTypesPaginated(service productType.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params entities.PaginationParams

		if err := ctx.QueryParser(&params); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}

		paginatedProductTypes, error := service.FetchProductTypesPaginated(&params)

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(paginatedProductTypes))
	}
}

// GetProductTypeById   godoc
// @Summary       Get a product type
// @Description   Get a single product type by its id
// @Tags          product type
// @Accept        json
// @Param			    id  path  number  true  "id of the product type to retrieve"
// @Produce       json
// @Success       200  {array}   entities.ProductType
// @Router        /api/v1/product-type/{id} [get]
func GetProductTypeById(service productType.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		product, error := service.FetchProductTypeById(uint(id))

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

// PostProductType   godoc
// @Summary       Create a product type
// @Description   Create new product types
// @Tags          product type
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product type"})
// @Produce       json
// @Success       200  {array}   entities.ProductType
// @Router        /api/v1/product-type [post]
func PostProductTypes(service productType.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body entities.CreateProductType

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		product, error := service.InsertProductType(&body)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
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
// @Success       200  {array}   entities.ProductType
// @Router        /api/v1/product-type/{id} [put]
func PutProductType(service productType.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body entities.UpdateProductType

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

		product, error := service.UpdateProductType(uint(id), &body)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

// DeleteProductTypeById   godoc
// @Summary       Delete product type
// @Description   Delete existing product type
// @Tags          product type
// @Accept        json
// @Param			    id  path  number  true  "id of the product type to delete"
// @Produce       json
// @Success       200  {array}   entities.ProductType
// @Router        /api/v1/product-type/{id} [delete]
func DeleteProductTypeById(service productType.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		product, error := service.RemoveProductType(uint(id))

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}
