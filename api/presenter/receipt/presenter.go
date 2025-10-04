package receipt_presenter

import (
	"fabiloco/hotel-trivoli-api/api/database"
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

	// 1. RECOLECTAR TODOS LOS IDs DE PRODUCTOS ÚNICOS
	//    (De todos los recibos para la consulta por lotes)
	uniqueProductIDs := make(map[uint]bool)
	for _, receiptProduct := range receipt.Products {
		uniqueProductIDs[receiptProduct.ProductID] = true
	}

	// Convertir el mapa de IDs a un slice para la cláusula IN
	var productIDs []uint
	for id := range uniqueProductIDs {
		productIDs = append(productIDs, id)
	}

	// 2. HACER UNA ÚNICA CONSULTA POR LOTES (Batch Query)
	//    (Obtiene todos los detalles de productos de una vez)
	var products []entities.Product
	database.DB.Preload("Type").Where("id IN (?)", productIDs).Find(&products)

	// 3. CREAR UN MAPA DE ACCESO RÁPIDO PARA PRODUCTOS (O(1))
	productDetailMap := make(map[uint]entities.Product)
	for _, product := range products {
		productDetailMap[product.ID] = product
	}

	// 4. MAPEO FINAL (Ahora sin consultas a la DB dentro del bucle)

	var receiptResponse ReceiptResponse

	// Mapa para agrupar ítems *dentro* del recibo y contar la cantidad
	productsInReceiptMap := make(map[uint]*ProductResponse)

	for _, receiptProduct := range receipt.Products {
		productID := receiptProduct.ProductID

		// Usar el mapa global de productos (O(1) look-up)
		if product, ok := productDetailMap[productID]; ok {

			if existingProduct, ok := productsInReceiptMap[productID]; ok {
				existingProduct.Quantity++
			} else {
				// Inicializar el producto con los detalles obtenidos de la DB en el paso 2
				productsInReceiptMap[productID] = &ProductResponse{
					ID:        receiptProduct.ID,
					Name:      product.Name,
					Type:      product.Type,
					Price:     product.Price,
					Img:       product.Img,
					Quantity:  1, // Inicializar en 1
					CreatedAt: receiptProduct.CreatedAt,
					UpdatedAt: receiptProduct.UpdatedAt,
				}
			}
		}
	}

	// Convertir productsInReceiptMap a slice (igual que antes)
	var productsResponseList []ProductResponse
	for _, product := range productsInReceiptMap {
		productsResponseList = append(productsResponseList, *product)
	}

	// Mapear el resto de campos (igual que antes)
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

	/* var receiptResponse ReceiptResponse

	// Map to store products by their IDs
	productsMap := make(map[uint]*ProductResponse)

	for _, receipt_product := range receipt.Products {
		var product = entities.Product{}

		database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

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

	return &receiptResponse */
}

func IndividualReceiptToIndividualReceiptResponse(receipt *entities.IndividualReceipt) *IndividualReceiptResponse {
	// 1. RECOLECTAR TODOS LOS IDs DE PRODUCTOS ÚNICOS
	//    (De todos los recibos para la consulta por lotes)
	uniqueProductIDs := make(map[uint]bool)
	for _, receiptProduct := range receipt.Products {
		uniqueProductIDs[receiptProduct.ProductID] = true
	}

	// Convertir el mapa de IDs a un slice para la cláusula IN
	var productIDs []uint
	for id := range uniqueProductIDs {
		productIDs = append(productIDs, id)
	}

	// 2. HACER UNA ÚNICA CONSULTA POR LOTES (Batch Query)
	//    (Obtiene todos los detalles de productos de una vez)
	var products []entities.Product
	database.DB.Preload("Type").Where("id IN (?)", productIDs).Find(&products)

	// 3. CREAR UN MAPA DE ACCESO RÁPIDO PARA PRODUCTOS (O(1))
	productDetailMap := make(map[uint]entities.Product)
	for _, product := range products {
		productDetailMap[product.ID] = product
	}

	// 4. MAPEO FINAL (Ahora sin consultas a la DB dentro del bucle)

	var receiptResponse IndividualReceiptResponse

	// Mapa para agrupar ítems *dentro* del recibo y contar la cantidad
	productsInReceiptMap := make(map[uint]*ProductResponse)

	for _, receiptProduct := range receipt.Products {
		productID := receiptProduct.ProductID

		// Usar el mapa global de productos (O(1) look-up)
		if product, ok := productDetailMap[productID]; ok {

			if existingProduct, ok := productsInReceiptMap[productID]; ok {
				existingProduct.Quantity++
			} else {
				// Inicializar el producto con los detalles obtenidos de la DB en el paso 2
				productsInReceiptMap[productID] = &ProductResponse{
					ID:        receiptProduct.ID,
					Name:      product.Name,
					Type:      product.Type,
					Price:     product.Price,
					Img:       product.Img,
					Quantity:  1, // Inicializar en 1
					CreatedAt: receiptProduct.CreatedAt,
					UpdatedAt: receiptProduct.UpdatedAt,
				}
			}
		}
	}

	// Convertir productsInReceiptMap a slice (igual que antes)
	var productsResponseList []ProductResponse
	for _, product := range productsInReceiptMap {
		productsResponseList = append(productsResponseList, *product)
	}

	// Mapear el resto de campos (igual que antes)
	receiptResponse.User = receipt.User
	receiptResponse.TotalPrice = receipt.TotalPrice

	receiptResponse.Products = productsResponseList
	receiptResponse.Shift = receipt.Shift

	receiptResponse.ID = fmt.Sprint("ir-", receipt.ID)
	receiptResponse.CreatedAt = receipt.CreatedAt
	receiptResponse.UpdatedAt = receipt.UpdatedAt
	return &receiptResponse

	/* 	var receiptResponse IndividualReceiptResponse

	// Map to store products by their IDs
	productsMap := make(map[uint]*ProductResponse)

	for _, receipt_product := range receipt.Products {
		var product = entities.Product{}

		database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

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

	return &receiptResponse */
}

func SuccessReceiptResponse(receipt *entities.Receipt) *fiber.Map {

	// 1. RECOLECTAR TODOS LOS IDs DE PRODUCTOS ÚNICOS
	//    (De todos los recibos para la consulta por lotes)
	uniqueProductIDs := make(map[uint]bool)
	for _, receiptProduct := range receipt.Products {
		uniqueProductIDs[receiptProduct.ProductID] = true
	}

	// Convertir el mapa de IDs a un slice para la cláusula IN
	var productIDs []uint
	for id := range uniqueProductIDs {
		productIDs = append(productIDs, id)
	}

	// 2. HACER UNA ÚNICA CONSULTA POR LOTES (Batch Query)
	//    (Obtiene todos los detalles de productos de una vez)
	var products []entities.Product
	database.DB.Preload("Type").Where("id IN (?)", productIDs).Find(&products)

	// 3. CREAR UN MAPA DE ACCESO RÁPIDO PARA PRODUCTOS (O(1))
	productDetailMap := make(map[uint]entities.Product)
	for _, product := range products {
		productDetailMap[product.ID] = product
	}

	// 4. MAPEO FINAL (Ahora sin consultas a la DB dentro del bucle)

	var receiptResponse ReceiptResponse

	// Mapa para agrupar ítems *dentro* del recibo y contar la cantidad
	productsInReceiptMap := make(map[uint]*ProductResponse)

	for _, receiptProduct := range receipt.Products {
		productID := receiptProduct.ProductID

		// Usar el mapa global de productos (O(1) look-up)
		if product, ok := productDetailMap[productID]; ok {

			if existingProduct, ok := productsInReceiptMap[productID]; ok {
				existingProduct.Quantity++
			} else {
				// Inicializar el producto con los detalles obtenidos de la DB en el paso 2
				productsInReceiptMap[productID] = &ProductResponse{
					ID:        receiptProduct.ID,
					Name:      product.Name,
					Type:      product.Type,
					Price:     product.Price,
					Img:       product.Img,
					Quantity:  1, // Inicializar en 1
					CreatedAt: receiptProduct.CreatedAt,
					UpdatedAt: receiptProduct.UpdatedAt,
				}
			}
		}
	}

	// Convertir productsInReceiptMap a slice (igual que antes)
	var productsResponseList []ProductResponse
	for _, product := range productsInReceiptMap {
		productsResponseList = append(productsResponseList, *product)
	}

	// Mapear el resto de campos (igual que antes)
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

	/* var receiptResponse ReceiptResponse

	// Map to store products by their IDs
	productsMap := make(map[uint]*ProductResponse)

	for _, receipt_product := range receipt.Products {
		var product = entities.Product{}

		database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

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

	return presenter.SuccessResponse(receiptResponse) */
}

func SuccessIndividualReceiptResponse(receipt *entities.IndividualReceipt) *fiber.Map {

	// 1. RECOLECTAR TODOS LOS IDs DE PRODUCTOS ÚNICOS
	//    (De todos los recibos para la consulta por lotes)
	uniqueProductIDs := make(map[uint]bool)
	for _, receiptProduct := range receipt.Products {
		uniqueProductIDs[receiptProduct.ProductID] = true
	}

	// Convertir el mapa de IDs a un slice para la cláusula IN
	var productIDs []uint
	for id := range uniqueProductIDs {
		productIDs = append(productIDs, id)
	}

	// 2. HACER UNA ÚNICA CONSULTA POR LOTES (Batch Query)
	//    (Obtiene todos los detalles de productos de una vez)
	var products []entities.Product
	database.DB.Preload("Type").Where("id IN (?)", productIDs).Find(&products)

	// 3. CREAR UN MAPA DE ACCESO RÁPIDO PARA PRODUCTOS (O(1))
	productDetailMap := make(map[uint]entities.Product)
	for _, product := range products {
		productDetailMap[product.ID] = product
	}

	// 4. MAPEO FINAL (Ahora sin consultas a la DB dentro del bucle)

	var individualReceiptResponse IndividualReceiptResponse

	// Mapa para agrupar ítems *dentro* del recibo y contar la cantidad
	productsInReceiptMap := make(map[uint]*ProductResponse)

	for _, receiptProduct := range receipt.Products {
		productID := receiptProduct.ProductID

		// Usar el mapa global de productos (O(1) look-up)
		if product, ok := productDetailMap[productID]; ok {

			if existingProduct, ok := productsInReceiptMap[productID]; ok {
				existingProduct.Quantity++
			} else {
				// Inicializar el producto con los detalles obtenidos de la DB en el paso 2
				productsInReceiptMap[productID] = &ProductResponse{
					ID:        receiptProduct.ID,
					Name:      product.Name,
					Type:      product.Type,
					Price:     product.Price,
					Img:       product.Img,
					Quantity:  1, // Inicializar en 1
					CreatedAt: receiptProduct.CreatedAt,
					UpdatedAt: receiptProduct.UpdatedAt,
				}
			}
		}
	}

	// Convertir productsInReceiptMap a slice (igual que antes)
	var productsResponseList []ProductResponse
	for _, product := range productsInReceiptMap {
		productsResponseList = append(productsResponseList, *product)
	}

	// Mapear el resto de campos (igual que antes)
	individualReceiptResponse.User = receipt.User
	individualReceiptResponse.TotalPrice = receipt.TotalPrice

	individualReceiptResponse.Products = productsResponseList
	individualReceiptResponse.Shift = receipt.Shift

	individualReceiptResponse.ID = fmt.Sprint("ir-", receipt.ID)
	individualReceiptResponse.CreatedAt = receipt.CreatedAt
	individualReceiptResponse.UpdatedAt = receipt.UpdatedAt

	return presenter.SuccessResponse(individualReceiptResponse)

	/* var individualReceiptResponse IndividualReceiptResponse

	// Map to store products by their IDs
	productsMap := make(map[uint]*ProductResponse)

	for _, receipt_product := range receipt.Products {
		var product = entities.Product{}

		database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

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

	return presenter.SuccessResponse(individualReceiptResponse) */
}

func SuccessIndividualReceiptsResponse(individualReceipts *[]entities.IndividualReceipt) *fiber.Map {
	if individualReceipts == nil || len(*individualReceipts) == 0 {
		return presenter.SuccessResponse([]ReceiptResponse{})
	}

	// 1. RECOLECTAR TODOS LOS IDs DE PRODUCTOS ÚNICOS
	//    (De todos los recibos para la consulta por lotes)
	uniqueProductIDs := make(map[uint]bool)
	for _, receipt := range *individualReceipts {
		for _, receiptProduct := range receipt.Products {
			uniqueProductIDs[receiptProduct.ProductID] = true
		}
	}

	// Convertir el mapa de IDs a un slice para la cláusula IN
	var productIDs []uint
	for id := range uniqueProductIDs {
		productIDs = append(productIDs, id)
	}

	// 2. HACER UNA ÚNICA CONSULTA POR LOTES (Batch Query)
	//    (Obtiene todos los detalles de productos de una vez)
	var products []entities.Product
	database.DB.Preload("Type").Where("id IN (?)", productIDs).Find(&products)

	// 3. CREAR UN MAPA DE ACCESO RÁPIDO PARA PRODUCTOS (O(1))
	productDetailMap := make(map[uint]entities.Product)
	for _, product := range products {
		productDetailMap[product.ID] = product
	}

	// 4. MAPEO FINAL (Ahora sin consultas a la DB dentro del bucle)
	var individualReceiptsResponse []IndividualReceiptResponse

	for _, receipt := range *individualReceipts {
		var receiptResponse IndividualReceiptResponse

		// Mapa para agrupar ítems *dentro* del recibo y contar la cantidad
		productsInReceiptMap := make(map[uint]*ProductResponse)

		for _, receiptProduct := range receipt.Products {
			productID := receiptProduct.ProductID

			// Usar el mapa global de productos (O(1) look-up)
			if product, ok := productDetailMap[productID]; ok {

				if existingProduct, ok := productsInReceiptMap[productID]; ok {
					existingProduct.Quantity++
				} else {
					// Inicializar el producto con los detalles obtenidos de la DB en el paso 2
					productsInReceiptMap[productID] = &ProductResponse{
						ID:        receiptProduct.ID,
						Name:      product.Name,
						Type:      product.Type,
						Price:     product.Price,
						Img:       product.Img,
						Quantity:  1, // Inicializar en 1
						CreatedAt: receiptProduct.CreatedAt,
						UpdatedAt: receiptProduct.UpdatedAt,
					}
				}
			}
		}

		// Convertir productsInReceiptMap a slice (igual que antes)
		var productsResponseList []ProductResponse
		for _, product := range productsInReceiptMap {
			productsResponseList = append(productsResponseList, *product)
		}

		// Mapear el resto de campos (igual que antes)

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

	/* var individualReceiptsResponse []IndividualReceiptResponse

	for _, receipt := range *individualReceipts {
		var receiptResponse IndividualReceiptResponse

		// Map to store products by their IDs
		productsMap := make(map[uint]*ProductResponse)

		for _, receipt_product := range receipt.Products {
			var product = entities.Product{}

			database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

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

	return presenter.SuccessResponse(individualReceiptsResponse) */
}

func SuccessReceiptsResponse(receipts *[]entities.Receipt) *fiber.Map {
	if receipts == nil || len(*receipts) == 0 {
		return presenter.SuccessResponse([]ReceiptResponse{})
	}

	// 1. RECOLECTAR TODOS LOS IDs DE PRODUCTOS ÚNICOS
	//    (De todos los recibos para la consulta por lotes)
	uniqueProductIDs := make(map[uint]bool)
	for _, receipt := range *receipts {
		for _, receiptProduct := range receipt.Products {
			uniqueProductIDs[receiptProduct.ProductID] = true
		}
	}

	// Convertir el mapa de IDs a un slice para la cláusula IN
	var productIDs []uint
	for id := range uniqueProductIDs {
		productIDs = append(productIDs, id)
	}

	// 2. HACER UNA ÚNICA CONSULTA POR LOTES (Batch Query)
	//    (Obtiene todos los detalles de productos de una vez)
	var products []entities.Product
	database.DB.Preload("Type").Where("id IN (?)", productIDs).Find(&products)

	// 3. CREAR UN MAPA DE ACCESO RÁPIDO PARA PRODUCTOS (O(1))
	productDetailMap := make(map[uint]entities.Product)
	for _, product := range products {
		productDetailMap[product.ID] = product
	}

	// 4. MAPEO FINAL (Ahora sin consultas a la DB dentro del bucle)
	var receiptsResponse []ReceiptResponse

	for _, receipt := range *receipts {
		var receiptResponse ReceiptResponse

		// Mapa para agrupar ítems *dentro* del recibo y contar la cantidad
		productsInReceiptMap := make(map[uint]*ProductResponse)

		for _, receiptProduct := range receipt.Products {
			productID := receiptProduct.ProductID

			// Usar el mapa global de productos (O(1) look-up)
			if product, ok := productDetailMap[productID]; ok {

				if existingProduct, ok := productsInReceiptMap[productID]; ok {
					existingProduct.Quantity++
				} else {
					// Inicializar el producto con los detalles obtenidos de la DB en el paso 2
					productsInReceiptMap[productID] = &ProductResponse{
						ID:        receiptProduct.ID,
						Name:      product.Name,
						Type:      product.Type,
						Price:     product.Price,
						Img:       product.Img,
						Quantity:  1, // Inicializar en 1
						CreatedAt: receiptProduct.CreatedAt,
						UpdatedAt: receiptProduct.UpdatedAt,
					}
				}
			}
		}

		// Convertir productsInReceiptMap a slice (igual que antes)
		var productsResponseList []ProductResponse
		for _, product := range productsInReceiptMap {
			productsResponseList = append(productsResponseList, *product)
		}

		// Mapear el resto de campos (igual que antes)
		receiptResponse.User = receipt.User
		receiptResponse.Service = receipt.Service
		// ... (otros mapeos) ...
		receiptResponse.Products = productsResponseList

		// ... (mapeo de IDs y fechas) ...

		receiptsResponse = append(receiptsResponse, receiptResponse)
	}

	return presenter.SuccessResponse(receiptsResponse)

	/* var receiptsResponse []ReceiptResponse

	for _, receipt := range *receipts {
		var receiptResponse ReceiptResponse

		// Map to store products by their IDs
		productsMap := make(map[uint]*ProductResponse)

		for _, receipt_product := range receipt.Products {
			var product = entities.Product{}

			database.DB.Preload("Type").Find(&product, receipt_product.ProductID)

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

	return presenter.SuccessResponse(receiptsResponse) */
}
