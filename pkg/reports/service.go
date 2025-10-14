package reports

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	individualreceipt "fabiloco/hotel-trivoli-api/pkg/individual_receipt"
	product "fabiloco/hotel-trivoli-api/pkg/product"
	receipt "fabiloco/hotel-trivoli-api/pkg/receipt"
	"fmt"
	"sort"
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

	allReceipts, err := s.receiptRepository.ReadBetweenDates(sd, ed)
	if err != nil {
		return nil, 0, 0, err
	}

	allIndividualReceipts, err := s.individualReceiptRepository.ReadBetweenDates(sd, ed)
	if err != nil {
		return nil, 0, 0, err
	}

	totalReceipts := int64(len(*allReceipts))
	totalIndividualReceipts := int64(len(*allIndividualReceipts))

	all := make([]entities.GeneralReceiptItem, 0, len(*allReceipts)+len(*allIndividualReceipts))

	for _, r := range *allReceipts {
		all = append(all, entities.GeneralReceiptItem{
			Receipt:      r,
			IsIndividual: false,
		})
	}

	for _, ir := range *allIndividualReceipts {
		all = append(all, entities.GeneralReceiptItem{
			Receipt:      ir,
			IsIndividual: true,
		})
	}

	sort.Slice(all, func(i, j int) bool {
		var timeI, timeJ time.Time
		if all[i].IsIndividual {
			if individualReceipt, ok := all[i].Receipt.(entities.IndividualReceipt); ok {
				timeI = individualReceipt.CreatedAt
			}
		} else {
			if receipt, ok := all[i].Receipt.(entities.Receipt); ok {
				timeI = receipt.CreatedAt
			}
		}

		if all[j].IsIndividual {
			if individualReceipt, ok := all[j].Receipt.(entities.IndividualReceipt); ok {
				timeJ = individualReceipt.CreatedAt
			}
		} else {
			if receipt, ok := all[j].Receipt.(entities.Receipt); ok {
				timeJ = receipt.CreatedAt
			}
		}

		return timeI.After(timeJ)
	})

	totalCombined := totalReceipts + totalIndividualReceipts

	start := params.GetOffset()
	end := start + params.GetLimit()
	if start > len(all) {
		start = len(all)
	}
	if end > len(all) {
		end = len(all)
	}
	paged := all[start:end]

	var unifiedReceipts []interface{}

	for _, item := range paged {
		if item.IsIndividual {
			if individualReceipt, ok := item.Receipt.(entities.IndividualReceipt); ok {
				unifiedReceipts = append(unifiedReceipts, individualReceipt)
			}
		} else {
			if receipt, ok := item.Receipt.(entities.Receipt); ok {
				unifiedReceipts = append(unifiedReceipts, receipt)
			}
		}
	}

	combinedData := map[string]interface{}{
		"receipts": unifiedReceipts,
	}

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
