package receipt_presenter

import (
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ProductResponse struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time

	Name     string                 `json:"name"`
	Quantity int                    `json:"quantity"`
	Type     []entities.ProductType `json:"type"`
	Price    float32                `json:"price"`
	Img      string                 `json:"img"`
}

type ServiceResponse struct {
	Name     string  `json:"name"`  // service name
	Price    float32 `json:"price"` // service price
	Quantity int     `json:"quantity"`
}

type ReceiptResponse struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time

	TotalPrice float32           `json:"total_price"`
	TotalTime  time.Duration     `json:"total_time"`
	Products   []ProductResponse `json:"products"`
	Service    entities.Service  `json:"service"`
	Room       entities.Room     `json:"room"`
	User       entities.User     `json:"user"`
	Shift      entities.Shift    `json:"shift"`
}

type IndividualReceiptResponse struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time

	TotalPrice float32           `json:"total_price"`
	Products   []ProductResponse `json:"products"`
	User       entities.User     `json:"user"`
	Shift      entities.Shift    `json:"shift"`
}

func ReceiptsToReceiptsResponses(receipts []entities.Receipt) []ReceiptResponse {
	var receiptsResponses []ReceiptResponse

	for _, receipt := range receipts {
		receiptsResponses = append(receiptsResponses, *ReceiptToReceiptResponse(&receipt))
	}

	return receiptsResponses
}

func IndividualReceiptsToIndividualReceiptsResponses(individualReceipts []entities.IndividualReceipt) []IndividualReceiptResponse {
	var receiptsResponses []IndividualReceiptResponse

	for _, receipt := range individualReceipts {
		receiptsResponses = append(receiptsResponses, *IndividualReceiptToIndividualReceiptResponse(&receipt))
	}

	return receiptsResponses
}

func ReceiptToReceiptResponse(receipt *entities.Receipt) *ReceiptResponse {
	var receiptResponse ReceiptResponse

	// Map to store products by their IDs
	productsMap := make(map[uint]*ProductResponse)

	for _, receipt_product := range receipt.Products {
		var product = entities.Product{}

		// database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

		if existingProduct, ok := productsMap[receipt_product.ProductID]; ok {
			existingProduct.Quantity++
		} else {
			productsMap[receipt_product.ProductID] = &ProductResponse{
				ID:        receipt_product.ProductID,
				Name:      product.Name,
				Type:      product.Type,
				Price:     product.Price,
				Img:       product.Img,
				Quantity:  1,
				CreatedAt: receipt_product.CreatedAt,
				UpdatedAt: receipt_product.UpdatedAt,
			}
		}
	}

	// Convert productsMap back to slice
	var productsResponseList []ProductResponse
	for _, product := range productsMap {
		productsResponseList = append(productsResponseList, *product)
	}

	receiptResponse.User = receipt.User
	receiptResponse.Service = receipt.Service
	receiptResponse.Room = receipt.Room
	receiptResponse.TotalPrice = receipt.TotalPrice
	receiptResponse.TotalTime = receipt.TotalTime

	receiptResponse.Products = productsResponseList
	receiptResponse.Shift = receipt.Shift

	receiptResponse.ID = fmt.Sprint("r-", receipt.ID)
	receiptResponse.CreatedAt = receipt.CreatedAt
	receiptResponse.UpdatedAt = receipt.UpdatedAt

	return &receiptResponse
}

func IndividualReceiptToIndividualReceiptResponse(receipt *entities.IndividualReceipt) *IndividualReceiptResponse {
	var receiptResponse IndividualReceiptResponse

	// Map to store products by their IDs
	productsMap := make(map[uint]*ProductResponse)

	for _, receipt_product := range receipt.Products {
		var product = entities.Product{}

		// database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

		if existingProduct, ok := productsMap[receipt_product.ProductID]; ok {
			existingProduct.Quantity++
		} else {
			productsMap[receipt_product.ProductID] = &ProductResponse{
				ID:        receipt_product.ProductID,
				Name:      product.Name,
				Type:      product.Type,
				Price:     product.Price,
				Img:       product.Img,
				Quantity:  1,
				CreatedAt: receipt_product.CreatedAt,
				UpdatedAt: receipt_product.UpdatedAt,
			}
		}
	}

	// Convert productsMap back to slice
	var productsResponseList []ProductResponse
	for _, product := range productsMap {
		productsResponseList = append(productsResponseList, *product)
	}

	receiptResponse.User = receipt.User
	receiptResponse.TotalPrice = receipt.TotalPrice

	receiptResponse.Products = productsResponseList
	receiptResponse.Shift = receipt.Shift

	receiptResponse.ID = fmt.Sprint("ir-", receipt.ID)
	receiptResponse.CreatedAt = receipt.CreatedAt
	receiptResponse.UpdatedAt = receipt.UpdatedAt

	return &receiptResponse
}

func SuccessReceiptResponse(receipt *entities.Receipt) *fiber.Map {
	var receiptResponse ReceiptResponse

	// Map to store products by their IDs
	productsMap := make(map[uint]*ProductResponse)

	for _, receipt_product := range receipt.Products {
		var product = entities.Product{}

		// database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

		if existingProduct, ok := productsMap[receipt_product.ProductID]; ok {
			existingProduct.Quantity++
		} else {
			productsMap[receipt_product.ProductID] = &ProductResponse{
				ID:        receipt_product.ID,
				Name:      product.Name,
				Type:      product.Type,
				Price:     product.Price,
				Img:       product.Img,
				Quantity:  1,
				CreatedAt: receipt_product.CreatedAt,
				UpdatedAt: receipt_product.UpdatedAt,
			}
		}
	}

	// Convert productsMap back to slice
	var productsResponseList []ProductResponse
	for _, product := range productsMap {
		productsResponseList = append(productsResponseList, *product)
	}

	receiptResponse.User = receipt.User
	receiptResponse.Service = receipt.Service
	receiptResponse.Room = receipt.Room
	receiptResponse.TotalPrice = receipt.TotalPrice
	receiptResponse.TotalTime = receipt.TotalTime

	receiptResponse.Products = productsResponseList

	receiptResponse.Shift = receipt.Shift

	receiptResponse.ID = fmt.Sprint("r-", receipt.ID)
	receiptResponse.CreatedAt = receipt.CreatedAt
	receiptResponse.UpdatedAt = receipt.UpdatedAt

	return presenter.SuccessResponse(receiptResponse)
}

func SuccessIndividualReceiptResponse(receipt *entities.IndividualReceipt) *fiber.Map {
	var individualReceiptResponse IndividualReceiptResponse

	// Map to store products by their IDs
	productsMap := make(map[uint]*ProductResponse)

	for _, receipt_product := range receipt.Products {
		var product = entities.Product{}

		// database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

		if existingProduct, ok := productsMap[receipt_product.ProductID]; ok {
			existingProduct.Quantity++
		} else {
			productsMap[receipt_product.ProductID] = &ProductResponse{
				ID:        receipt_product.ID,
				Name:      product.Name,
				Type:      product.Type,
				Price:     product.Price,
				Img:       product.Img,
				Quantity:  1,
				CreatedAt: receipt_product.CreatedAt,
				UpdatedAt: receipt_product.UpdatedAt,
			}
		}
	}

	// Convert productsMap back to slice
	var productsResponseList []ProductResponse
	for _, product := range productsMap {
		productsResponseList = append(productsResponseList, *product)
	}

	individualReceiptResponse.User = receipt.User
	individualReceiptResponse.TotalPrice = receipt.TotalPrice

	individualReceiptResponse.Products = productsResponseList
	individualReceiptResponse.Shift = receipt.Shift

	individualReceiptResponse.ID = fmt.Sprint("ir-", receipt.ID)
	individualReceiptResponse.CreatedAt = receipt.CreatedAt
	individualReceiptResponse.UpdatedAt = receipt.UpdatedAt

	return presenter.SuccessResponse(individualReceiptResponse)
}

func SuccessIndividualReceiptsResponse(individualReceipts *[]entities.IndividualReceipt) *fiber.Map {
	var individualReceiptsResponse []IndividualReceiptResponse

	for _, receipt := range *individualReceipts {
		var receiptResponse IndividualReceiptResponse

		// Map to store products by their IDs
		productsMap := make(map[uint]*ProductResponse)

		for _, receipt_product := range receipt.Products {
			var product = entities.Product{}

			// database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

			if existingProduct, ok := productsMap[receipt_product.ProductID]; ok {
				existingProduct.Quantity++
			} else {
				productsMap[receipt_product.ProductID] = &ProductResponse{
					ID:        receipt_product.ID,
					Name:      product.Name,
					Type:      product.Type,
					Price:     product.Price,
					Img:       product.Img,
					Quantity:  1,
					CreatedAt: receipt_product.CreatedAt,
					UpdatedAt: receipt_product.UpdatedAt,
				}
			}
		}

		// Convert productsMap back to slice
		var productsResponseList []ProductResponse
		for _, product := range productsMap {
			productsResponseList = append(productsResponseList, *product)
		}

		receiptResponse.User = receipt.User
		receiptResponse.TotalPrice = receipt.TotalPrice

		receiptResponse.Products = productsResponseList
		receiptResponse.Shift = receipt.Shift

		receiptResponse.ID = fmt.Sprint("ir-", receipt.ID)
		receiptResponse.CreatedAt = receipt.CreatedAt
		receiptResponse.UpdatedAt = receipt.UpdatedAt

		individualReceiptsResponse = append(individualReceiptsResponse, receiptResponse)
	}

	return presenter.SuccessResponse(individualReceiptsResponse)
}

func SuccessReceiptsResponse(receipts *[]entities.Receipt) *fiber.Map {
	var receiptsResponse []ReceiptResponse

	for _, receipt := range *receipts {
		var receiptResponse ReceiptResponse

		// Map to store products by their IDs
		productsMap := make(map[uint]*ProductResponse)

		for _, receipt_product := range receipt.Products {
			var product = entities.Product{}

			// database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

			if existingProduct, ok := productsMap[receipt_product.ProductID]; ok {
				existingProduct.Quantity++
			} else {
				productsMap[receipt_product.ProductID] = &ProductResponse{
					ID:        receipt_product.ID,
					Name:      product.Name,
					Type:      product.Type,
					Price:     product.Price,
					Img:       product.Img,
					Quantity:  1,
					CreatedAt: receipt_product.CreatedAt,
					UpdatedAt: receipt_product.UpdatedAt,
				}
			}
		}

		// Convert productsMap back to slice
		var productsResponseList []ProductResponse
		for _, product := range productsMap {
			productsResponseList = append(productsResponseList, *product)
		}

		receiptResponse.User = receipt.User
		receiptResponse.Service = receipt.Service
		receiptResponse.Room = receipt.Room
		receiptResponse.TotalPrice = receipt.TotalPrice
		receiptResponse.TotalTime = receipt.TotalTime

		receiptResponse.Products = productsResponseList
		receiptResponse.Shift = receipt.Shift

		receiptResponse.ID = fmt.Sprint("r-", receipt.ID)
		receiptResponse.CreatedAt = receipt.CreatedAt
		receiptResponse.UpdatedAt = receipt.UpdatedAt

		receiptsResponse = append(receiptsResponse, receiptResponse)
	}

	return presenter.SuccessResponse(receiptsResponse)
}
