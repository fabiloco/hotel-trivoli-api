package reports

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	product "fabiloco/hotel-trivoli-api/pkg/product"
	receipt "fabiloco/hotel-trivoli-api/pkg/receipt"
	"fmt"
	"time"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	ReceiptByTargetDate (targetDate string) (*[]entities.Receipt, error)
	ReceiptsBetweenDates (startDate string, endDate string) (*[]entities.Receipt, error)
}

type service struct {
	productRepository product.Repository
	receiptRepository receipt.Repository
}

func NewService(pr product.Repository, rr receipt.Repository) Service {
	return &service{
    productRepository: pr,
    receiptRepository: rr,
	}
}

func (s *service) ReceiptByTargetDate(targetDate string) (*[]entities.Receipt, error) {
  date, error := time.Parse(time.RFC3339, targetDate)
  if error != nil {
    return nil, errors.New(fmt.Sprintf("error parsing Date %s", targetDate))
  }
   
	return s.receiptRepository.ReadByDate(date)
}

func (s *service) ReceiptsBetweenDates(startDate string, endDate string) (*[]entities.Receipt, error) {
  sd, error := time.Parse(time.RFC3339, startDate)
  if error != nil {
    return nil, errors.New(fmt.Sprintf("error parsing Date %s", startDate))
  }

  ed, error := time.Parse(time.RFC3339, endDate)
  if error != nil {
    return nil, errors.New(fmt.Sprintf("error parsing Date %s", endDate))
  }
   
	return s.receiptRepository.ReadBetweenDates(sd, ed)
}

// func (s *service) FetchProducts() (*[]entities.Product, error) {
// 	return s.productRepository.Read()
// }
//
// func (s *service) UpdateProduct(id uint, product *entities.UpdateProduct) (*entities.Product, error) {
//   var productTypesSlice []entities.ProductType
//
//   for i := 0; i < len(product.Type); i++{
//     productType, error := s.productTypeRepository.ReadById(product.Type[i])
//
//     if error != nil {
//       return nil, errors.New(fmt.Sprintf("no product type with id %d", product.Type[i]))
//     }
//
//     productTypesSlice = append(productTypesSlice, *productType)
//   }
//
//   newProduct := entities.Product {
//     Name: product.Name,
//     Stock: product.Stock,
//     Price: product.Price,
//   }
//
// 	return s.productRepository.Update(id, &newProduct)
// }
//
// func (s *service) RemoveProduct(ID uint) (*entities.Product, error) {
// 	return s.productRepository.Delete(ID)
// }
//
// func (s *service) FetchProductById(ID uint) (*entities.Product, error) {
// 	return s.productRepository.ReadById(ID)
// }
