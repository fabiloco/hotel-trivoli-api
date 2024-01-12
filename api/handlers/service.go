package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/service"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ListServices   godoc
// @Summary       List services
// @Description   list avaliable services in the database
// @Tags          service
// @Accept        json
// @Produce       json
// @Success       200  {array}   entities.Service
// @Router        /api/v1/service [get]
func GetServices(service service.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		services, error := service.FetchServices()

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(services))
	}
}

// GetServiceById   godoc
// @Summary       Get a service
// @Description   Get a single service by its id
// @Tags          service
// @Accept        json
// @Param			    id  path  number  true  "id of the service to retrieve"
// @Produce       json
// @Success       200  {array}   entities.Service
// @Router        /api/v1/service/{id} [get]
func GetServiceById(service service.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		product, error := service.FetchServiceById(uint(id))

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

// PostService   godoc
// @Summary       Create a service
// @Description   Create new services
// @Tags          service
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test service"})
// @Produce       json
// @Success       200  {array}   entities.Service
// @Router        /api/v1/service [post]
func PostServices(service service.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body entities.CreateService

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}
		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ""))))
		}

		product, error := service.InsertService(&body)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

// PutService   godoc
// @Summary       Update service
// @Description   Edit existing service
// @Tags          service
// @Accept        json
// @Param			    body  body  string  true  "Body of the request" SchemaExample({\n"name": "test product"})
// @Param			    id  path  number  true  "id of the service to update"
// @Produce       json
// @Success       200  {array}   entities.Service
// @Router        /api/v1/service/{id} [put]
func PutService(service service.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    var body entities.UpdateService

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

		product, error := service.UpdateService(uint(id), &body)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}

// DeleteServiceById   godoc
// @Summary       Delete service
// @Description   Delete existing service
// @Tags          service
// @Accept        json
// @Param			    id  path  number  true  "id of the service to delete"
// @Produce       json
// @Success       200  {array}   entities.Service
// @Router        /api/v1/service/{id} [delete]
func DeleteServiceById(service service.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		product, error := service.RemoveService(uint(id))

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		return ctx.JSON(presenter.SuccessResponse(product))
	}
}
