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
	ReceiptByTargetDate(targetDate string, limit, offset int) (*[]entities.Receipt, int64, error)

	//ReceiptByUser(userId uint) (*[]entities.Receipt, error)
	ReceiptByUser(userId uint, limit, offset int) (*[]entities.Receipt, int64, error)
	ReceiptTodayByUser(userId uint) (*[]entities.Receipt, error)
	ReceiptsBetweenDates(startDate string, endDate string) (*[]entities.Receipt, error)
	ReceiptsBetweenDatesPaginated(startDate string, endDate string, params *entities.PaginationParams) (*entities.PaginatedResponse, int64, int64, error)

	IndividualReceiptByTargetDate(targetDate string, limit, offset int) (*[]entities.IndividualReceipt, int64, error)

	//IndividualReceiptByUser(userId uint) (*[]entities.IndividualReceipt, error)
	IndividualReceiptByUser(userId uint, limit, offset int) (*[]entities.IndividualReceipt, int64, error)
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

	receipts, _, error := s.receiptRepository.ReadByDate(startOfToday, 0, 0)

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

func (s *service) ReceiptByUser(userId uint, limit, offset int) (*[]entities.Receipt, int64, error) {
	// date, error := time.Parse(time.RFC3339, targetDate)
	// if error != nil {
	//   return nil, errors.New(fmt.Sprintf("error parsing Date %s", targetDate))
	// }

	receipts, total, error := s.receiptRepository.Read(limit, offset)

	if error != nil {
		return nil, 0, error
	}

	var userReceipts []entities.Receipt

	for i := 0; i < len(*receipts); i++ {
		if (*receipts)[i].UserID == userId {
			userReceipts = append(userReceipts, (*receipts)[i])
		}
	}

	return &userReceipts, total, nil
}

func (s *service) ReceiptByTargetDate(targetDate string, limit, offset int) (*[]entities.Receipt, int64, error) {
	date, error := time.Parse(time.RFC3339, targetDate)
	if error != nil {
		return nil, 0, errors.New(fmt.Sprintf("error parsing Date %s", targetDate))
	}

	return s.receiptRepository.ReadByDate(date, limit, offset)
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

func (s *service) ReceiptsBetweenDatesPaginated(startDate string, endDate string, params *entities.PaginationParams) (*entities.PaginatedResponse, int64, int64, error) {
	sd, error := time.Parse(time.RFC3339, startDate)
	if error != nil {
		return nil, 0, 0, errors.New(fmt.Sprintf("error parsing Date %s", startDate))
	}

	ed, error := time.Parse(time.RFC3339, endDate)
	if error != nil {
		return nil, 0, 0, errors.New(fmt.Sprintf("error parsing Date %s", endDate))
	}

	receipts, totalReceipts, err := s.receiptRepository.ReadBetweenDatesPaginated(sd, ed, params)
	if err != nil {
		return nil, 0, 0, err
	}

	individualReceipts, totalIndividualReceipts, err := s.individualReceiptRepository.ReadBetweenDatesPaginated(sd, ed, params)
	if err != nil {
		return nil, 0, 0, err
	}

	combinedData := map[string]interface{}{
		"receipts":           receipts,
		"individualReceipts": individualReceipts,
	}

	totalCombined := totalReceipts + totalIndividualReceipts

	return entities.NewPaginatedResponse(combinedData, totalCombined, params.Page, params.PageSize), totalReceipts, totalIndividualReceipts, nil
}

func (s *service) IndividualReceiptTodayByUser(userId uint) (*[]entities.IndividualReceipt, error) {
	// Obtener la fecha de inicio de hoy
	startOfToday := time.Now().Truncate(24 * time.Hour)

	receipts, _, error := s.individualReceiptRepository.ReadByDate(startOfToday, 0, 0)

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

func (s *service) IndividualReceiptByUser(userId uint, limit, offset int) (*[]entities.IndividualReceipt, int64, error) {
	// date, error := time.Parse(time.RFC3339, targetDate)
	// if error != nil {
	//   return nil, errors.New(fmt.Sprintf("error parsing Date %s", targetDate))
	// }

	receipts, total, error := s.individualReceiptRepository.Read(limit, offset)

	if error != nil {
		return nil, 0, error
	}

	var userReceipts []entities.IndividualReceipt

	for i := 0; i < len(*receipts); i++ {
		if (*receipts)[i].UserID == userId {
			userReceipts = append(userReceipts, (*receipts)[i])
		}
	}

	return &userReceipts, total, nil
}

func (s *service) IndividualReceiptByTargetDate(targetDate string, limit, offset int) (*[]entities.IndividualReceipt, int64, error) {
	date, error := time.Parse(time.RFC3339, targetDate)
	if error != nil {
		return nil, 0, errors.New(fmt.Sprintf("error parsing Date %s", targetDate))
	}

	return s.individualReceiptRepository.ReadByDate(date, limit, offset)
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
