package receipt

import (
	"errors"
	"fmt"
	"time"

	"fabiloco/hotel-trivoli-api/pkg/entities"
	individualreceipt "fabiloco/hotel-trivoli-api/pkg/individual_receipt"
	"fabiloco/hotel-trivoli-api/pkg/product"
	"fabiloco/hotel-trivoli-api/pkg/room"
	serviceModule "fabiloco/hotel-trivoli-api/pkg/service"
	"fabiloco/hotel-trivoli-api/pkg/user"

	"github.com/guregu/null/v5"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertReceipt(receipt *entities.CreateReceipt) (*entities.Receipt, error)

	GenerateReceipt(receipt *entities.CreateReceipt) (*entities.Receipt, error)
	GenerateIndividualReceipt(receipt *entities.CreateIndividualReceipt) (*entities.IndividualReceipt, error)

	FetchReceipts() (*[]entities.Receipt, error)
	FetchIndividualReceipts() (*[]entities.IndividualReceipt, error)

	FetchReceiptById(id uint) (*entities.Receipt, error)
	FetchIndividualReceiptById(id uint) (*entities.IndividualReceipt, error)

	UpdateReceipt(id uint, receipt *entities.UpdateReceipt) (*entities.Receipt, error)
	UpdateIndividualReceipt(id uint, receipt *entities.UpdateIndividualReceipt) (*entities.IndividualReceipt, error)

	RemoveReceipt(id uint) (*entities.Receipt, error)
	RemoveIndividualReceipt(id uint) (*entities.IndividualReceipt, error)
}

type service struct {
	repository                  Repository
	individualReceiptRepository individualreceipt.Repository
	serviceRepository           serviceModule.Repository
	roomRepository              room.Repository
	productRepository           product.Repository
	userRepository              user.Repository
}

func NewService(r Repository, sr serviceModule.Repository, pr product.Repository, rr room.Repository, ur user.Repository, irr individualreceipt.Repository) Service {
	return &service{
		repository:                  r,
		serviceRepository:           sr,
		productRepository:           pr,
		roomRepository:              rr,
		userRepository:              ur,
		individualReceiptRepository: irr,
	}
}

func (s *service) GenerateReceipt(receipt *entities.CreateReceipt) (*entities.Receipt, error) {
	service, error := s.serviceRepository.ReadById(receipt.Service)

	if error != nil {
		return nil, errors.New(fmt.Sprintf("no service with id %d", receipt.Service))
	}

	room, error := s.roomRepository.ReadById(receipt.Room)

	if error != nil {
		return nil, errors.New(fmt.Sprintf("no room with id %d", receipt.Service))
	}

	user, error := s.userRepository.ReadById(receipt.User)

	if error != nil {
		return nil, errors.New(fmt.Sprintf("no room with id %d", receipt.User))
	}

	receiptProducts, error := s.productToReceiptProduct(receipt)

	if error != nil {
		return nil, error
	}

	newReceipt := entities.Receipt{
		TotalTime:  time.Duration(receipt.TotalTime),
		TotalPrice: receipt.TotalPrice,
		Products:   receiptProducts,
		Service:    *service,
		Room:       *room,
		User:       *user,
		ShiftID:    null.IntFromPtr(nil),
	}

	return s.repository.Create(&newReceipt)
}

func (s *service) GenerateIndividualReceipt(receipt *entities.CreateIndividualReceipt) (*entities.IndividualReceipt, error) {
	user, error := s.userRepository.ReadById(receipt.User)

	if error != nil {
		return nil, errors.New(fmt.Sprintf("no room with id %d", receipt.User))
	}

	individualReceiptProducts, error := s.productToIndividualReceiptProduct(receipt)

	if error != nil {
		return nil, error
	}

	newReceipt := entities.IndividualReceipt{
		TotalPrice: receipt.TotalPrice,
		Products:   individualReceiptProducts,
		User:       *user,
	}

	return s.individualReceiptRepository.Create(&newReceipt)
}

func (s *service) InsertReceipt(receipt *entities.CreateReceipt) (*entities.Receipt, error) {
	service, error := s.serviceRepository.ReadById(receipt.Service)

	if error != nil {
		return nil, errors.New(fmt.Sprintf("no service with id %d", receipt.Service))
	}

	room, error := s.roomRepository.ReadById(receipt.Room)

	if error != nil {
		return nil, errors.New(fmt.Sprintf("no room with id %d", receipt.Service))
	}

	user, error := s.userRepository.ReadById(receipt.User)

	if error != nil {
		return nil, errors.New(fmt.Sprintf("no room with id %d", receipt.User))
	}

	var products []entities.Product

	for i := 0; i < len(receipt.Products); i++ {
		productWithId, error := s.productRepository.ReadById(receipt.Products[i])

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no product with id %d", receipt.Products[i]))
		}

		products = append(products, *productWithId)
	}

	newReceipt := entities.Receipt{
		TotalTime:  time.Duration(receipt.TotalTime),
		TotalPrice: receipt.TotalPrice,
		Service:    *service,
		Room:       *room,
		User:       *user,
		ShiftID:    null.IntFromPtr(nil),
	}

	return s.repository.Create(&newReceipt)
}

func (s *service) FetchReceipts() (*[]entities.Receipt, error) {
	return s.repository.Read()
}

func (s *service) FetchIndividualReceipts() (*[]entities.IndividualReceipt, error) {
	return s.individualReceiptRepository.Read()
}

func (s *service) UpdateReceipt(id uint, receipt *entities.UpdateReceipt) (*entities.Receipt, error) {
	var error error

	var service *entities.Service
	if receipt.Service != 0 {
		service, error = s.serviceRepository.ReadById(receipt.Service)

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no service with id %d", receipt.Service))
		}
	}

	var room *entities.Room
	if receipt.Room != 0 {
		room, error = s.roomRepository.ReadById(receipt.Room)

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no room with id %d", receipt.Room))
		}
	}

	var user *entities.User
	if receipt.User != 0 {
		user, error = s.userRepository.ReadById(receipt.User)

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no user with id %d", receipt.User))
		}
	}

	receiptProducts, error := s.productToReceiptProduct((*entities.CreateReceipt)(receipt))

	if error != nil {
		return nil, error
	}

	newReceipt := entities.Receipt{
		TotalTime:  time.Duration(receipt.TotalTime),
		TotalPrice: receipt.TotalPrice,
		Products:   receiptProducts,
		ShiftID:    null.IntFromPtr(nil),
	}

	if room == nil {
		newReceipt.Room = *&entities.Room{}
	} else {
		newReceipt.Room = *room
	}

	if service == nil {
		newReceipt.Service = *&entities.Service{}
	} else {
		newReceipt.Service = *service
	}

	if user == nil {
		newReceipt.User = *&entities.User{}
	} else {
		newReceipt.User = *user
	}

	return s.repository.Update(id, &newReceipt)
}

func (s *service) UpdateIndividualReceipt(id uint, receipt *entities.UpdateIndividualReceipt) (*entities.IndividualReceipt, error) {
	var error error

	var user *entities.User
	if receipt.User != 0 {
		user, error = s.userRepository.ReadById(receipt.User)

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no user with id %d", receipt.User))
		}
	}

	individualReceiptProducts, error := s.productToIndividualReceiptProduct((*entities.CreateIndividualReceipt)(receipt))

	if error != nil {
		return nil, error
	}

	newReceipt := entities.IndividualReceipt{
		TotalPrice: receipt.TotalPrice,
		Products:   individualReceiptProducts,
		ShiftID:    null.IntFromPtr(nil),
	}

	if user == nil {
		newReceipt.User = *&entities.User{}
	} else {
		newReceipt.User = *user
	}

	return s.individualReceiptRepository.Update(id, &newReceipt)
}

func (s *service) RemoveReceipt(ID uint) (*entities.Receipt, error) {
	return s.repository.Delete(ID)
}

func (s *service) RemoveIndividualReceipt(ID uint) (*entities.IndividualReceipt, error) {
	return s.individualReceiptRepository.Delete(ID)
}

func (s *service) FetchReceiptById(ID uint) (*entities.Receipt, error) {
	return s.repository.ReadById(ID)
}

func (s *service) FetchIndividualReceiptById(ID uint) (*entities.IndividualReceipt, error) {
	return s.individualReceiptRepository.ReadById(ID)
}

// func (s *service) productToReceiptProduct(receipt *entities.CreateReceipt) ([]entities.ReceiptProduct, error) {
// 	var products []entities.Product

// 	for i := 0; i < len(receipt.Products); i++ {
// 		productWithId, error := s.productRepository.ReadById(receipt.Products[i])

// 		if error != nil {
// 			return nil, errors.New(fmt.Sprintf("no product with id %d", receipt.Products[i]))
// 		}

// 		// check if the existing stock is enough to process the parchuse
// 		for j := 0; j < len(receipt.Products); j++ {
// 			var productIdTimes int = 0

// 			for _, productId := range receipt.Products {
// 				if productId == receipt.Products[j] {
// 					productIdTimes += 1
// 				}
// 			}

// 			if productWithId.Stock-productIdTimes < 0 {
// 				return nil, errors.New(
// 					fmt.Sprintf("not enough stock in the product with id %d to process the receipt. product stock: %d. product times in body: %d.",
// 						receipt.Products[j], productWithId.Stock, productIdTimes))
// 			}
// 		}

// 		product := entities.Product{
// 			Name:  productWithId.Name,
// 			Stock: productWithId.Stock - 1,
// 			Price: productWithId.Price,
// 			Type:  productWithId.Type,
// 		}

// 		productRestocked, error := s.productRepository.Update(receipt.Products[i], &product)

// 		if error != nil {
// 			return nil, errors.New(fmt.Sprintf("Error restocking product with id: %d", receipt.Products[i]))
// 		}

// 		products = append(products, *productRestocked)
// 	}

// 	var receiptProducts []entities.ReceiptProduct
// 	for _, product := range products {
// 		receiptProduct := entities.ReceiptProduct{
// 			ProductID: product.ID,
// 		}
// 		receiptProducts = append(receiptProducts, receiptProduct)
// 	}

// 	return receiptProducts, nil
// }

func (s *service) productToReceiptProduct(receipt *entities.CreateReceipt) ([]entities.ReceiptProduct, error) {
	productCounts := make(map[uint]int)
	var products []entities.Product

	for _, productId := range receipt.Products {
		productCounts[productId]++
	}

	for productId, count := range productCounts {
		productWithId, err := s.productRepository.ReadById(productId)
		if err != nil {
			return nil, fmt.Errorf("no product with id %d", productId)
		}

		if productWithId.Stock < count {
			return nil, fmt.Errorf("not enough stock for product with id %d. Available: %d, required: %d",
				productId, productWithId.Stock, count)
		}

		productWithId.Stock -= count

		productRestocked, err := s.productRepository.Update(productId, productWithId)
		if err != nil {
			return nil, fmt.Errorf("error updating product with id %d", productId)
		}

		products = append(products, *productRestocked)
	}

	var receiptProducts []entities.ReceiptProduct
	for productId, count := range productCounts {
		for i := 0; i < count; i++ {
			receiptProduct := entities.ReceiptProduct{
				ProductID: productId,
			}
			receiptProducts = append(receiptProducts, receiptProduct)
		}
	}

	return receiptProducts, nil
}

// func (s *service) productToIndividualReceiptProduct(receipt *entities.CreateIndividualReceipt) ([]entities.IndividualReceiptProduct, error) {
// 	var products []entities.Product

// 	for i := 0; i < len(receipt.Products); i++ {
// 		productWithId, error := s.productRepository.ReadById(receipt.Products[i])

// 		if error != nil {
// 			return nil, errors.New(fmt.Sprintf("no product with id %d", receipt.Products[i]))
// 		}

// 		// check if the existing stock is enough to process the parchuse
// 		for j := 0; j < len(receipt.Products); j++ {
// 			var productIdTimes int = 0

// 			for _, productId := range receipt.Products {
// 				if productId == receipt.Products[j] {
// 					productIdTimes += 1
// 				}
// 			}

// 			if productWithId.Stock-productIdTimes < 0 {
// 				return nil, errors.New(
// 					fmt.Sprintf("not enough stock in the product with id %d to process the receipt. product stock: %d. product times in body: %d.",
// 						receipt.Products[j], productWithId.Stock, productIdTimes))
// 			}
// 		}

// 		product := entities.Product{
// 			Name:  productWithId.Name,
// 			Stock: productWithId.Stock - 1,
// 			Price: productWithId.Price,
// 			Type:  productWithId.Type,
// 		}

// 		productRestocked, error := s.productRepository.Update(receipt.Products[i], &product)

// 		if error != nil {
// 			return nil, errors.New(fmt.Sprintf("Error restocking product with id: %d", receipt.Products[i]))
// 		}

// 		products = append(products, *productRestocked)
// 	}

// 	var receiptProducts []entities.IndividualReceiptProduct
// 	for _, product := range products {
// 		receiptProduct := entities.IndividualReceiptProduct{
// 			ProductID: product.ID,
// 		}
// 		receiptProducts = append(receiptProducts, receiptProduct)
// 	}

// 	return receiptProducts, nil
// }

func (s *service) productToIndividualReceiptProduct(receipt *entities.CreateIndividualReceipt) ([]entities.IndividualReceiptProduct, error) {
	productCounts := make(map[uint]int)
	var products []entities.Product

	for _, productId := range receipt.Products {
		productCounts[productId]++
	}

	for productId, count := range productCounts {
		productWithId, err := s.productRepository.ReadById(productId)
		if err != nil {
			return nil, fmt.Errorf("no product with id %d", productId)
		}

		if productWithId.Stock < count {
			return nil, fmt.Errorf("not enough stock for product with id %d. Available: %d, required: %d",
				productId, productWithId.Stock, count)
		}

		productWithId.Stock -= count

		productRestocked, err := s.productRepository.Update(productId, productWithId)
		if err != nil {
			return nil, fmt.Errorf("error updating product with id %d", productId)
		}

		products = append(products, *productRestocked)
	}

	var receiptProducts []entities.IndividualReceiptProduct
	for productId, count := range productCounts {
		for i := 0; i < count; i++ {
			receiptProduct := entities.IndividualReceiptProduct{
				ProductID: productId,
			}
			receiptProducts = append(receiptProducts, receiptProduct)
		}
	}

	return receiptProducts, nil
}
