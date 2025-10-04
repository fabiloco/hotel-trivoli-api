package handlers

import (
	"errors"
	"fabiloco/hotel-trivoli-api/api/presenter"
	receipt_presenter "fabiloco/hotel-trivoli-api/api/presenter/receipt"
	"fabiloco/hotel-trivoli-api/printer"
	"fmt"
	"time"

	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/receipt"
	"fabiloco/hotel-trivoli-api/pkg/shift"
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

		printer.GetESCPOSPrinter().Print(receipt_presenter.ReceiptToReceiptResponse(receipt))

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

		receipt, error := service.GenerateIndividualReceipt(&body)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		printer.GetESCPOSPrinter().PrintIndividual(receipt_presenter.IndividualReceiptToIndividualReceiptResponse(receipt))

		return ctx.JSON(receipt_presenter.SuccessIndividualReceiptResponse(receipt))
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

		_, limit, offset := utils.GetPaginationParams(ctx)

		receipts, total, err := service.FetchReceipts(limit, offset)
		individualReceipts, total2, error := service.FetchIndividualReceipts(limit, offset)

		if error != nil || err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		totalTotal := total

		if total < total2 {
			totalTotal = total2
		}

		response := fiber.Map{
			"receipts":           receipt_presenter.SuccessReceiptsResponse(receipts),
			"individualReceipts": receipt_presenter.SuccessIndividualReceiptsResponse(individualReceipts),
		}

		// return ctx.JSON(receipt_presenter.SuccessReceiptsResponse(receipts))
		//return ctx.JSON(response)
		return ctx.JSON(utils.Paginate(ctx, totalTotal, response))
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
		idStr := ctx.Params("id")
		id, error := utils.ConvertReceiptsId(idStr)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(fmt.Sprint("Error converting param to id: ", error))))
		}

		if utils.ReceiptIdPatternMatch(idStr) {
			receipt, error := service.FetchReceiptById(uint(id))

			if error != nil {
				ctx.Status(http.StatusInternalServerError)
				return ctx.JSON(presenter.ErrorResponse(error))
			}

			return ctx.JSON(receipt_presenter.SuccessReceiptResponse(receipt))
		} else if utils.IndividualReceiptIdPatternMatch(idStr) {
			receipt, error := service.FetchIndividualReceiptById(uint(id))
			fmt.Println("here")

			if error != nil {
				ctx.Status(http.StatusInternalServerError)
				return ctx.JSON(presenter.ErrorResponse(error))
			}

			return ctx.JSON(receipt_presenter.SuccessIndividualReceiptResponse(receipt))
		} else {

			return ctx.JSON(presenter.ErrorResponse(errors.New("Id does not match with the pattern")))
		}
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

		idStr := ctx.Params("id")
		id, error := utils.ConvertReceiptsId(idStr)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(fmt.Sprint("Error converting param to id: ", error))))
		}

		if utils.ReceiptIdPatternMatch(idStr) {
			product, error := service.UpdateReceipt(uint(id), &body)

			if error != nil {
				ctx.Status(http.StatusBadRequest)
				return ctx.JSON(presenter.ErrorResponse(error))
			}

			return ctx.JSON(presenter.SuccessResponse(product))
		} else if utils.IndividualReceiptIdPatternMatch(idStr) {
			product, error := service.UpdateIndividualReceipt(uint(id), &entities.UpdateIndividualReceipt{
				TotalPrice: body.TotalPrice,
				Products:   body.Products,
				User:       body.User,
				Shift:      body.Shift,
			})

			if error != nil {
				ctx.Status(http.StatusBadRequest)
				return ctx.JSON(presenter.ErrorResponse(error))
			}

			return ctx.JSON(presenter.SuccessResponse(product))
		} else {
			return ctx.JSON(presenter.ErrorResponse(errors.New("Id does not match with the pattern")))
		}
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
		idStr := ctx.Params("id")
		id, error := utils.ConvertReceiptsId(idStr)

		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(fmt.Sprint("Error converting param to id: ", error))))
		}

		if utils.ReceiptIdPatternMatch(idStr) {
			product, error := service.RemoveReceipt(uint(id))

			if error != nil {
				ctx.Status(http.StatusBadRequest)
				return ctx.JSON(presenter.ErrorResponse(error))
			}

			return ctx.JSON(presenter.SuccessResponse(product))
		} else if utils.IndividualReceiptIdPatternMatch(idStr) {
			product, error := service.RemoveIndividualReceipt(uint(id))

			if error != nil {
				ctx.Status(http.StatusBadRequest)
				return ctx.JSON(presenter.ErrorResponse(error))
			}

			return ctx.JSON(presenter.SuccessResponse(product))
		} else {
			return ctx.JSON(presenter.ErrorResponse(errors.New("Id does not match with the pattern")))
		}
	}
}

type PrintReceiptType struct {
	Receipts           []uint `valid:"optional" json:"receipts"`
	IndividualReceipts []uint `valid:"optional" json:"individual_receipts"`
}

func PrintReceipts(service receipt.Service, shiftService shift.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body entities.CreateShiftWithCode

		if err := ctx.BodyParser(&body); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(err))
		}

		if len(body.Receipts) <= 0 && len(body.IndividualReceipts) <= 0 {
			return ctx.JSON(presenter.SuccessResponse(fiber.Map{
				"message": "no print",
			}))
		}

		validationErrors := utils.ValidateInput(ctx, body)

		if validationErrors != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New(strings.Join(validationErrors, ", "))))
		}

		var receiptIds []uint

		for _, receiptId := range body.Receipts {
			if !utils.ReceiptIdPatternMatch(receiptId) {
				return ctx.JSON(presenter.ErrorResponse(errors.New("An id does not match with the pattern")))
			}

			idParsed, error := utils.ConvertReceiptsId(receiptId)
			if error != nil {
				return ctx.JSON(presenter.ErrorResponse(error))
			}

			receiptIds = append(receiptIds, idParsed)
		}

		var individualReceiptIds []uint

		for _, individualReceiptId := range body.IndividualReceipts {
			if !utils.IndividualReceiptIdPatternMatch(individualReceiptId) {
				return ctx.JSON(presenter.ErrorResponse(errors.New("An id does not match with the pattern")))
			}

			idParsed, error := utils.ConvertReceiptsId(individualReceiptId)
			if error != nil {
				return ctx.JSON(presenter.ErrorResponse(error))
			}
			individualReceiptIds = append(individualReceiptIds, idParsed)
		}

		var user entities.User

		var receipts []receipt_presenter.ReceiptResponse
		var existingReceiptIds []uint
		for _, receiptId := range receiptIds {
			receipt, error := service.FetchReceiptById(receiptId)

			if error != nil {
				// ctx.Status(http.StatusInternalServerError)
				// return ctx.JSON(presenter.ErrorResponse(error))
			} else {

				existingReceiptIds = append(existingReceiptIds, receiptId)

				if user.ID == 0 {
					user = receipt.User
				}
				receipts = append(receipts, *receipt_presenter.ReceiptToReceiptResponse(receipt))
			}
		}

		var individualReceipts []receipt_presenter.IndividualReceiptResponse
		var existingIndividualReceiptIds []uint
		for _, receiptId := range individualReceiptIds {
			individualReceipt, error := service.FetchIndividualReceiptById(receiptId)

			if error != nil {
				// ctx.Status(http.StatusInternalServerError)
				// return ctx.JSON(presenter.ErrorResponse(error))
			} else {
				existingIndividualReceiptIds = append(existingIndividualReceiptIds, receiptId)
				individualReceipts = append(individualReceipts, *receipt_presenter.IndividualReceiptToIndividualReceiptResponse(individualReceipt))
				if user.ID == 0 {
					user = individualReceipt.User
				}
			}
		}

		bodyParsed := entities.CreateShift{
			Receipts:           existingReceiptIds,
			IndividualReceipts: existingIndividualReceiptIds,
		}

		_, error := shiftService.InsertShift(&bodyParsed)
		if error != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		var services []entities.Service
		var products []receipt_presenter.ProductResponse

		var totalServices float32
		var totalProducts float32

		var totalPrice float32

		for _, receipt := range receipts {
			for _, product := range receipt.Products {
				products = append(products, receipt_presenter.ProductResponse{
					ID:        product.ID,
					CreatedAt: product.CreatedAt,
					UpdatedAt: product.UpdatedAt,
					Name:      product.Name,
					Quantity:  product.Quantity,
					Type:      product.Type,
					Price:     product.Price,
					Img:       product.Img,
				})

				totalProducts += product.Price * float32(product.Quantity)
			}

			services = append(services, receipt.Service)
			totalServices += receipt.Service.Price
			totalPrice += receipt.TotalPrice
		}

		for _, individualReceipt := range individualReceipts {
			for _, product := range individualReceipt.Products {
				products = append(products, receipt_presenter.ProductResponse{
					ID:        product.ID,
					CreatedAt: product.CreatedAt,
					UpdatedAt: product.UpdatedAt,
					Name:      product.Name,
					Quantity:  product.Quantity,
					Type:      product.Type,
					Price:     product.Price,
					Img:       product.Img,
				})

				totalProducts += product.Price * float32(product.Quantity)
			}

			totalPrice += individualReceipt.TotalPrice
		}

		productsMap := make(map[uint]*receipt_presenter.ProductResponse)

		for _, product := range products {
			if existingProduct, ok := productsMap[product.ID]; ok {
				existingProduct.Quantity += product.Quantity
			} else {
				productsMap[product.ID] = &receipt_presenter.ProductResponse{
					ID:       product.ID,
					Name:     product.Name,
					Type:     product.Type,
					Price:    product.Price,
					Img:      product.Img,
					Quantity: product.Quantity,
				}
			}
		}

		var productsResponseList []receipt_presenter.ProductResponse
		for _, product := range productsMap {
			productsResponseList = append(productsResponseList, *product)
		}

		servicesMap := make(map[uint]*receipt_presenter.ServiceResponse)

		for _, service := range services {
			if existingService, ok := servicesMap[service.ID]; ok {
				existingService.Quantity += 1
			} else {
				servicesMap[service.ID] = &receipt_presenter.ServiceResponse{
					Name:     service.Name,
					Price:    service.Price,
					Quantity: 1,
				}
			}
		}

		var serviceResponseList []receipt_presenter.ServiceResponse
		for _, service := range servicesMap {
			serviceResponseList = append(serviceResponseList, *service)
		}

		printer.GetESCPOSPrinter().PrintReport(
			productsResponseList,
			totalProducts,
			serviceResponseList,
			totalServices,
			user,
			time.Now(),
			totalPrice,
		)

		return ctx.JSON(presenter.SuccessResponse(fiber.Map{
			"products":      productsResponseList,
			"totalProducts": totalProducts,
			"services":      serviceResponseList,
			"totalServices": totalServices,
			"totalPrice":    totalPrice,
		}))
	}
}

func PrintShift(service receipt.Service, shiftService shift.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		id, err := ctx.ParamsInt("id")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.ErrorResponse(errors.New("param id not valid")))
		}

		receipts_from_shifts, individual_receipts_from_shifts, error := shiftService.FetchShiftsById(uint(id))

		if error != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		var user entities.User

		var receipts []receipt_presenter.ReceiptResponse
		for _, receipt := range *receipts_from_shifts {
			receipts = append(receipts, *receipt_presenter.ReceiptToReceiptResponse(&receipt))
			if user.ID == 0 {
				user = receipt.User
			}
		}

		var individualReceipts []receipt_presenter.IndividualReceiptResponse
		for _, receipt_individual := range *individual_receipts_from_shifts {
			individualReceipts = append(individualReceipts, *receipt_presenter.IndividualReceiptToIndividualReceiptResponse(&receipt_individual))
			if user.ID == 0 {
				user = receipt_individual.User
			}
		}

		var services []entities.Service
		var products []receipt_presenter.ProductResponse

		var totalServices float32
		var totalProducts float32

		var totalPrice float32

		for _, receipt := range receipts {
			for _, product := range receipt.Products {
				products = append(products, receipt_presenter.ProductResponse{
					ID:        product.ID,
					CreatedAt: product.CreatedAt,
					UpdatedAt: product.UpdatedAt,
					Name:      product.Name,
					Quantity:  product.Quantity,
					Type:      product.Type,
					Price:     product.Price,
					Img:       product.Img,
				})

				totalProducts += product.Price * float32(product.Quantity)
			}

			services = append(services, receipt.Service)
			totalServices += receipt.Service.Price
			totalPrice += receipt.TotalPrice
		}

		for _, individualReceipt := range individualReceipts {
			for _, product := range individualReceipt.Products {
				products = append(products, receipt_presenter.ProductResponse{
					ID:        product.ID,
					CreatedAt: product.CreatedAt,
					UpdatedAt: product.UpdatedAt,
					Name:      product.Name,
					Quantity:  product.Quantity,
					Type:      product.Type,
					Price:     product.Price,
					Img:       product.Img,
				})

				totalProducts += product.Price * float32(product.Quantity)
			}

			totalPrice += individualReceipt.TotalPrice
		}

		productsMap := make(map[uint]*receipt_presenter.ProductResponse)

		for _, product := range products {
			if existingProduct, ok := productsMap[product.ID]; ok {
				existingProduct.Quantity += product.Quantity
			} else {
				productsMap[product.ID] = &receipt_presenter.ProductResponse{
					ID:       product.ID,
					Name:     product.Name,
					Type:     product.Type,
					Price:    product.Price,
					Img:      product.Img,
					Quantity: product.Quantity,
				}
			}
		}

		var productsResponseList []receipt_presenter.ProductResponse
		for _, product := range productsMap {
			productsResponseList = append(productsResponseList, *product)
		}

		servicesMap := make(map[uint]*receipt_presenter.ServiceResponse)

		for _, service := range services {
			if existingService, ok := servicesMap[service.ID]; ok {
				existingService.Quantity += 1
			} else {
				servicesMap[service.ID] = &receipt_presenter.ServiceResponse{
					Name:     service.Name,
					Price:    service.Price,
					Quantity: 1,
				}
			}
		}

		var serviceResponseList []receipt_presenter.ServiceResponse
		for _, service := range servicesMap {
			serviceResponseList = append(serviceResponseList, *service)
		}

		// return ctx.JSON(presenter.SuccessResponse(shift_presenter.ReceiptsToShiftsResponse(receipts, individual_receipts)))

		printer.GetESCPOSPrinter().PrintReport(
			productsResponseList,
			totalProducts,
			serviceResponseList,
			totalServices,
			user,
			time.Now(),
			totalPrice,
		)

		return ctx.JSON(presenter.SuccessResponse(fiber.Map{
			"products":      productsResponseList,
			"totalProducts": totalProducts,
			"services":      serviceResponseList,
			"totalServices": totalServices,
			"totalPrice":    totalPrice,
		}))
	}
}

// endpoint to just print an existing receipt, without creating a new one
func PrintReceipt(service receipt.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		idStr := ctx.Params("id")
		if !utils.ReceiptIdPatternMatch(idStr) {
			return ctx.JSON(presenter.ErrorResponse(errors.New("An id does not match with the pattern")))
		}

		idParsed, error := utils.ConvertReceiptsId(idStr)
		if error != nil {
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		receipt, error := service.FetchReceiptById(idParsed)
		if error != nil {
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		printer.GetESCPOSPrinter().Print(receipt_presenter.ReceiptToReceiptResponse(receipt))

		return ctx.JSON(receipt_presenter.SuccessReceiptResponse(receipt))
	}
}

// endpoint to just print an existing individual receipt, without creating a new one
func PrintIndividualReceipt(service receipt.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		idStr := ctx.Params("id")
		if !utils.IndividualReceiptIdPatternMatch(idStr) {
			return ctx.JSON(presenter.ErrorResponse(errors.New("An id does not match with the pattern")))
		}

		idParsed, error := utils.ConvertReceiptsId(idStr)
		if error != nil {
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		receipt, error := service.FetchIndividualReceiptById(idParsed)
		if error != nil {
			return ctx.JSON(presenter.ErrorResponse(error))
		}

		printer.GetESCPOSPrinter().PrintIndividual(receipt_presenter.IndividualReceiptToIndividualReceiptResponse(receipt))

		return ctx.JSON(receipt_presenter.SuccessIndividualReceiptResponse(receipt))
	}
}
