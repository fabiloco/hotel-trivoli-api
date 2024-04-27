package receipt_presenter

import (
	"fabiloco/hotel-trivoli-api/api/database"
	"fabiloco/hotel-trivoli-api/api/presenter"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ProductResponse struct {
	ID          uint                
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Name  string                    `json:"name"`   
	Quantity int                    `json:"quantity"` 
  Type  []entities.ProductType    `json:"type"` 
	Price float32                   `json:"price"`
	Img   string                    `json:"img"` 
}

type ReceiptResponse struct {
	ID          uint                
	CreatedAt   time.Time
	UpdatedAt   time.Time

  TotalPrice  float32                 `json:"total_price"`
  TotalTime   time.Duration           `json:"total_time"`
  Products    []ProductResponse       `json:"products"` 
  Service     entities.Service        `json:"service"`
  Room        entities.Room           `json:"room"`
  User        entities.User           `json:"user"`
}

type IndividualReceiptResponse struct {
	ID          uint                
	CreatedAt   time.Time
	UpdatedAt   time.Time

  TotalPrice  float32                 `json:"total_price"`
  Products    []ProductResponse       `json:"products"` 
  User        entities.User           `json:"user"`
}

func SuccessReceiptResponse(receipt *entities.Receipt) *fiber.Map {
  var receiptResponse ReceiptResponse

  // Map to store products by their IDs
  productsMap := make(map[uint]*ProductResponse)

  for _, receipt_product := range receipt.Products {
    var product = entities.Product{}

    database.DB.Find(&product, receipt_product.ProductID)

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

  receiptResponse.ID = receipt.ID
  receiptResponse.CreatedAt = receipt.CreatedAt
  receiptResponse.UpdatedAt = receipt.UpdatedAt

  return presenter.SuccessResponse(receiptResponse)
}



func SuccessIndividualReceiptResponse(receipt *entities.IndividualReceipt) *fiber.Map {
  var individualReceiptResponse IndividualReceiptResponse

  // Map to store products by their IDs
  productsMap := make(map[uint]*ProductResponse)

  for _, product := range receipt.Products {
    if existingProduct, ok := productsMap[product.ID]; ok {
      existingProduct.Quantity++
    } else {
      productsMap[product.ID] = &ProductResponse{
        ID:        product.ID,
        Name:      product.Name,
        Type:      product.Type,
        Price:     product.Price,
        Img:       product.Img,
        Quantity:  1,
        CreatedAt: product.CreatedAt,
        UpdatedAt: product.UpdatedAt,
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

  individualReceiptResponse.ID = receipt.ID
  individualReceiptResponse.CreatedAt = receipt.CreatedAt
  individualReceiptResponse.UpdatedAt = receipt.UpdatedAt

  return presenter.SuccessResponse(individualReceiptResponse)
}

func SuccessReceiptsResponse(receipts *[]entities.Receipt) *fiber.Map {
  var receiptsResponse []ReceiptResponse

  for _, receipt := range(*receipts) {
    var receiptResponse ReceiptResponse

    // Map to store products by their IDs
    productsMap := make(map[uint]*ProductResponse)

    for _, receipt_product := range receipt.Products {

      var product = entities.Product{}

      database.DB.Find(&product, receipt_product.ProductID)

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

    receiptResponse.ID = receipt.ID
    receiptResponse.CreatedAt = receipt.CreatedAt
    receiptResponse.UpdatedAt = receipt.UpdatedAt

    receiptsResponse = append(receiptsResponse, receiptResponse)
  }

  return presenter.SuccessResponse(receiptsResponse)
}
