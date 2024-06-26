package reports

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	individualreceipt "fabiloco/hotel-trivoli-api/pkg/individual_receipt"
	product "fabiloco/hotel-trivoli-api/pkg/product"
	receipt "fabiloco/hotel-trivoli-api/pkg/receipt"
	"fmt"
	"time"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	ReceiptByTargetDate(targetDate string) (*[]entities.Receipt, error)
	ReceiptByUser(userId uint) (*[]entities.Receipt, error)
	ReceiptTodayByUser(userId uint) (*[]entities.Receipt, error)
	ReceiptsBetweenDates(startDate string, endDate string) (*[]entities.Receipt, error)

	IndividualReceiptByTargetDate(targetDate string) (*[]entities.IndividualReceipt, error)
	IndividualReceiptByUser(userId uint) (*[]entities.IndividualReceipt, error)
	IndividualReceiptTodayByUser(userId uint) (*[]entities.IndividualReceipt, error)
	IndividualReceiptsBetweenDates(startDate string, endDate string) (*[]entities.IndividualReceipt, error)
}

type service struct {
	productRepository           product.Repository
	receiptRepository           receipt.Repository
	individualReceiptRepository individualreceipt.Repository
}

func NewService(pr product.Repository, rr receipt.Repository, irr individualreceipt.Repository) Service {
	return &service{
		productRepository:           pr,
		receiptRepository:           rr,
		individualReceiptRepository: irr,
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

	for i := 0; i < len(*receipts); i++ {
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

	for i := 0; i < len(*receipts); i++ {
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

func (s *service) IndividualReceiptTodayByUser(userId uint) (*[]entities.IndividualReceipt, error) {
	// Obtener la fecha de inicio de hoy
	startOfToday := time.Now().Truncate(24 * time.Hour)

	receipts, error := s.individualReceiptRepository.ReadByDate(startOfToday)

	if error != nil {
		return nil, error
	}

	var userReceipts []entities.IndividualReceipt

	for i := 0; i < len(*receipts); i++ {
		if (*receipts)[i].UserID == userId {
			userReceipts = append(userReceipts, (*receipts)[i])
		}
	}

	return &userReceipts, nil
}

func (s *service) IndividualReceiptByUser(userId uint) (*[]entities.IndividualReceipt, error) {
	// date, error := time.Parse(time.RFC3339, targetDate)
	// if error != nil {
	//   return nil, errors.New(fmt.Sprintf("error parsing Date %s", targetDate))
	// }

	receipts, error := s.individualReceiptRepository.Read()

	if error != nil {
		return nil, error
	}

	var userReceipts []entities.IndividualReceipt

	for i := 0; i < len(*receipts); i++ {
		if (*receipts)[i].UserID == userId {
			userReceipts = append(userReceipts, (*receipts)[i])
		}
	}

	return &userReceipts, nil
}

func (s *service) IndividualReceiptByTargetDate(targetDate string) (*[]entities.IndividualReceipt, error) {
	date, error := time.Parse(time.RFC3339, targetDate)
	if error != nil {
		return nil, errors.New(fmt.Sprintf("error parsing Date %s", targetDate))
	}

	return s.individualReceiptRepository.ReadByDate(date)
}

func (s *service) IndividualReceiptsBetweenDates(startDate string, endDate string) (*[]entities.IndividualReceipt, error) {
	sd, error := time.Parse(time.RFC3339, startDate)
	if error != nil {
		return nil, errors.New(fmt.Sprintf("error parsing Date %s", startDate))
	}

	ed, error := time.Parse(time.RFC3339, endDate)
	if error != nil {
		return nil, errors.New(fmt.Sprintf("error parsing Date %s", endDate))
	}

	return s.individualReceiptRepository.ReadBetweenDates(sd, ed)
}
