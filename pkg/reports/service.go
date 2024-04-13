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
	ReceiptByUser (userId uint) (*[]entities.Receipt, error)
	ReceiptTodayByUser (userId uint) (*[]entities.Receipt, error)
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

func (s *service) ReceiptTodayByUser(userId uint) (*[]entities.Receipt, error) {
  // Obtener la fecha de inicio de hoy
	startOfToday := time.Now().Truncate(24 * time.Hour)

  receipts, error := s.receiptRepository.ReadByDate(startOfToday)

  if error != nil {
    return nil, error    
  }

  var userReceipts []entities.Receipt

  for i := 0; i < len(*receipts); i++{
    if (*receipts)[i].UserID == userId {
      userReceipts = append(userReceipts, (*receipts)[i])
    }
  }

   
	return &userReceipts, nil
}

func (s *service) ReceiptByUser(userId uint) (*[]entities.Receipt, error) {
  // date, error := time.Parse(time.RFC3339, targetDate)
  // if error != nil {
  //   return nil, errors.New(fmt.Sprintf("error parsing Date %s", targetDate))
  // }

  receipts, error := s.receiptRepository.Read()

  if error != nil {
    return nil, error    
  }

  var userReceipts []entities.Receipt

  for i := 0; i < len(*receipts); i++{
    if (*receipts)[i].UserID == userId {
      userReceipts = append(userReceipts, (*receipts)[i])
    }
  }

   
	return &userReceipts, nil
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
