package receipt_presenter

import (
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

func SuccessReceiptResponse(receipt *entities.Receipt) *fiber.Map {
  var receiptResponse ReceiptResponse

  var productsResponseList []ProductResponse

  // count quantity
  for _, product := range(receipt.Products) {
    var productResponse ProductResponse

    var quantity int = 0

    for _, productInside := range(receipt.Products) {
      if product.ID == productInside.ID {
        quantity += 1
      }
    }

    productResponse.Name = product.Name
    productResponse.Type = product.Type
    productResponse.Price = product.Price
    productResponse.Img = product.Img
    productResponse.Quantity = quantity

    productResponse.ID = product.ID
    productResponse.CreatedAt = product.CreatedAt
    productResponse.UpdatedAt = product.UpdatedAt

    productsResponseList = append(productsResponseList, productResponse)
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
