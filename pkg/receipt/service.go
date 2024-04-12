package receipt

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	individualreceipt "fabiloco/hotel-trivoli-api/pkg/individual_receipt"
	"fabiloco/hotel-trivoli-api/pkg/product"
	"fabiloco/hotel-trivoli-api/pkg/room"
	serviceModule "fabiloco/hotel-trivoli-api/pkg/service"
	"fabiloco/hotel-trivoli-api/pkg/user"
	"fmt"
	"time"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertReceipt (receipt *entities.CreateReceipt) (*entities.Receipt, error)
	GenerateReceipt (receipt *entities.CreateReceipt) (*entities.Receipt, error)
	GenerateIndividualReceipt (receipt *entities.CreateIndividualReceipt) (*entities.IndividualReceipt, error)
	FetchReceipts () (*[]entities.Receipt, error)
  FetchReceiptById (id uint) (*entities.Receipt, error)
	UpdateReceipt (id uint, receipt *entities.UpdateReceipt) (*entities.Receipt, error)
	RemoveReceipt (id uint) (*entities.Receipt, error)
}

type service struct {
	repository Repository
	individualReceiptRepository individualreceipt.Repository
	serviceRepository serviceModule.Repository
	roomRepository room.Repository
	productRepository product.Repository
	userRepository user.Repository
}

func NewService(r Repository, sr serviceModule.Repository, pr product.Repository, rr room.Repository, ur user.Repository, irr individualreceipt.Repository) Service {
	return &service{
		repository: r,
    serviceRepository: sr,
    productRepository: pr,
    roomRepository: rr,
    userRepository: ur,
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


  var products []entities.Product

  for i := 0; i < len(receipt.Products); i++{
    productWithId, error := s.productRepository.ReadById(receipt.Products[i])

    if error != nil {
      return nil, errors.New(fmt.Sprintf("no product with id %d", receipt.Products[i]))
    }

    product := entities.Product {
      Name:  productWithId.Name,
      Stock: productWithId.Stock - 1,
      Price: productWithId.Price,
      Type:  productWithId.Type,
    }

    productRestocked, error := s.productRepository.Update(receipt.Products[i], &product)

    if error != nil {
      return nil, errors.New(fmt.Sprintf("Error restocking product with id: %d", receipt.Products[i]))
    }

    products = append(products, *productRestocked)
  }

  newReceipt := entities.Receipt {
    TotalTime: time.Duration(receipt.TotalTime),
    TotalPrice: receipt.TotalPrice,
    Products: products,
    Service: *service,
    Room: *room,
    User: *user,
  }

	return s.repository.Create(&newReceipt)
}

func (s *service) GenerateIndividualReceipt(receipt *entities.CreateIndividualReceipt) (*entities.IndividualReceipt, error) {
  user, error := s.userRepository.ReadById(receipt.User)

  if error != nil {
    return nil, errors.New(fmt.Sprintf("no room with id %d", receipt.User))
  }

  var products []entities.Product

  for i := 0; i < len(receipt.Products); i++{
    productWithId, error := s.productRepository.ReadById(receipt.Products[i])

    if error != nil {
      return nil, errors.New(fmt.Sprintf("no product with id %d", receipt.Products[i]))
    }

    product := entities.Product {
      Name:  productWithId.Name,
      Stock: productWithId.Stock - 1,
      Price: productWithId.Price,
      Type:  productWithId.Type,
    }

    productRestocked, error := s.productRepository.Update(receipt.Products[i], &product)

    if error != nil {
      return nil, errors.New(fmt.Sprintf("Error restocking product with id: %d", receipt.Products[i]))
    }

    products = append(products, *productRestocked)
  }
  newReceipt := entities.IndividualReceipt {
    TotalPrice: receipt.TotalPrice,
    Products: products,
    User: *user,
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

  for i := 0; i < len(receipt.Products); i++{
    productWithId, error := s.productRepository.ReadById(receipt.Products[i])

    if error != nil {
      return nil, errors.New(fmt.Sprintf("no product with id %d", receipt.Products[i]))
    }

    products = append(products, *productWithId)
  }

  newReceipt := entities.Receipt {
    TotalTime: time.Duration(receipt.TotalTime),
    TotalPrice: receipt.TotalPrice,
    Products: products,
    Service: *service,
    Room: *room,
    User: *user,
  }

	return s.repository.Create(&newReceipt)
}

func (s *service) FetchReceipts() (*[]entities.Receipt, error) {
	return s.repository.Read()
}

func (s *service) UpdateReceipt(id uint, receipt *entities.UpdateReceipt) (*entities.Receipt, error) {
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

  for i := 0; i < len(receipt.Products); i++{
    product, error := s.productRepository.ReadById(receipt.Products[i])

    if error != nil {
      return nil, errors.New(fmt.Sprintf("no product with id %d", receipt.Products[i]))
    }

    products = append(products, *product)
  }

  newReceipt := entities.Receipt {
    TotalTime: time.Duration(receipt.TotalTime),
    TotalPrice: receipt.TotalPrice,
    Service: *service,
    Room: *room,
    User: *user,
    Products: products,
  }

	return s.repository.Update(id, &newReceipt)
}

func (s *service) RemoveReceipt(ID uint) (*entities.Receipt, error) {
	return s.repository.Delete(ID)
}

func (s *service) FetchReceiptById(ID uint) (*entities.Receipt, error) {
	return s.repository.ReadById(ID)
}
