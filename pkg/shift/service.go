package shift

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	individualreceipt "fabiloco/hotel-trivoli-api/pkg/individual_receipt"
	receipt "fabiloco/hotel-trivoli-api/pkg/receipt"
	"fmt"
	"sort"
	"time"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertShift(receipt *entities.CreateShift) (*entities.Shift, error)
	FetchShiftsById(id uint) (*[]entities.Receipt, *[]entities.IndividualReceipt, error)
	//FetchAllShifts() (*[]entities.Receipt, *[]entities.IndividualReceipt, error)
	FetchAllShifts(limit, offset int) (*[]entities.Receipt, *[]entities.IndividualReceipt, int64, error)
	FetchShiftsBetweenDate(startDate string, endDate string) (*[]entities.Receipt, *[]entities.IndividualReceipt, error)
	UpdateShift(id uint, receipt *entities.UpdateShift) (*entities.Shift, error)
	RemoveShift(id uint) (*entities.Shift, error)
}

type service struct {
	repository                  Repository
	receiptRepository           receipt.Repository
	individualReceiptRepository individualreceipt.Repository
}

func NewService(r Repository, rr receipt.Repository, irr individualreceipt.Repository) Service {
	return &service{
		repository:                  r,
		receiptRepository:           rr,
		individualReceiptRepository: irr,
	}
}

func (s *service) InsertShift(shift *entities.CreateShift) (*entities.Shift, error) {
	var receipts []entities.Receipt

	for i := 0; i < len(shift.Receipts); i++ {
		receiptWithId, error := s.receiptRepository.ReadById(shift.Receipts[i])

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no receipt with id %d", shift.Receipts[i]))
		}

		receipts = append(receipts, *receiptWithId)
	}

	var individual_receipts []entities.IndividualReceipt

	for i := 0; i < len(shift.IndividualReceipts); i++ {
		inidividual_receiptWithId, error := s.individualReceiptRepository.ReadById(shift.IndividualReceipts[i])

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no individual receipt with id %d", shift.IndividualReceipts[i]))
		}

		individual_receipts = append(individual_receipts, *inidividual_receiptWithId)
	}

	newShift := entities.Shift{}

	shiftCreated, err := s.repository.Create(&newShift)

	if err != nil {
		return nil, err
	}

	for _, receipt := range receipts {
		receipt.Shift = *shiftCreated

		_, error := s.receiptRepository.UpdateShift(receipt.ID, &receipt)

		if error != nil {
			return nil, errors.New(fmt.Sprintf("error editing receipt with id %d", receipt.ID))
		}
	}

	for _, individual_receipt := range individual_receipts {
		individual_receipt.Shift = *shiftCreated

		_, error := s.individualReceiptRepository.UpdateShift(individual_receipt.ID, &individual_receipt)

		if error != nil {
			return nil, errors.New(fmt.Sprintf("error editing individual receipt with id %d", individual_receipt.ID))
		}
	}

	return shiftCreated, err
}

func (s *service) FetchShiftsById(id uint) (*[]entities.Receipt, *[]entities.IndividualReceipt, error) {
	receipts, _ := s.receiptRepository.ReadAllByShiftId(id)

	individual_receipts, _ := s.individualReceiptRepository.ReadAllByShiftId(id)

	return receipts, individual_receipts, nil
}

// func (s *service) FetchAllShifts(limit, offset int) (*[]entities.Receipt, *[]entities.IndividualReceipt, error) {
func (s *service) FetchAllShifts(limit, offset int) (*[]entities.Receipt, *[]entities.IndividualReceipt, int64, error) {

	// 1️⃣ Traer todos los shiftIDs existentes en ambas tablas
	allShiftIDs, err := s.receiptRepository.FindAllShiftIDs()
	if err != nil {
		return nil, nil, 0, err
	}

	individualShiftIDs, err := s.individualReceiptRepository.FindAllShiftIDs()
	if err != nil {
		return nil, nil, 0, err
	}

	// 2️⃣ Unir y eliminar duplicados
	shiftMap := make(map[int64]bool)
	for _, id := range allShiftIDs {
		shiftMap[id] = true
	}
	for _, id := range individualShiftIDs {
		shiftMap[id] = true
	}

	// 3️⃣ Convertir el mapa a slice y ordenar
	shiftIDs := make([]int64, 0, len(shiftMap))
	for id := range shiftMap {
		shiftIDs = append(shiftIDs, id)
	}
	sort.Slice(shiftIDs, func(i, j int) bool {
		return shiftIDs[i] > shiftIDs[j]
	})

	total := len(shiftIDs)

	// 4️⃣ Paginar sobre los shiftIDs únicos
	start := offset
	end := offset + limit
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	pagedShiftIDs := shiftIDs[start:end]

	// 5️⃣ Obtener los receipts e individualReceipts solo de esos shifts
	receipts, err := s.receiptRepository.FindByShiftIDs(pagedShiftIDs)
	if err != nil {
		return nil, nil, 0, err
	}

	individualReceipts, err := s.individualReceiptRepository.FindByShiftIDs(pagedShiftIDs)
	if err != nil {
		return nil, nil, 0, err
	}

	return &receipts, &individualReceipts, int64(total), nil

	/* receipts, totalReceipts, err := s.receiptRepository.ReadByShiftNotNull(limit, offset)
	if err != nil {
		return nil, nil, 0, err
	}

	individualReceipts, totalIndividuals, err := s.individualReceiptRepository.ReadByShiftNotNull(limit, offset)
	if err != nil {
		return nil, nil, 0, err
	}

	// si quieres paginar sobre los dos juntos, definamos qué total devolver
	total := totalReceipts
	if totalReceipts < totalIndividuals {
		total = totalIndividuals
	}

	return receipts, individualReceipts, total, nil */

	/* receipts, error := s.receiptRepository.ReadByShiftNotNull()

	if error != nil {
		return nil, nil, error
	}

	individual_receipts, error := s.individualReceiptRepository.ReadByShiftNotNull()

	if error != nil {
		return nil, nil, error
	}

	return receipts, individual_receipts, nil */
}

func (s *service) FetchShiftsBetweenDate(startDate string, endDate string) (*[]entities.Receipt, *[]entities.IndividualReceipt, error) {
	sd, error := time.Parse(time.RFC3339, startDate)
	if error != nil {
		return nil, nil, errors.New(fmt.Sprintf("error parsing Date %s", startDate))
	}

	ed, error := time.Parse(time.RFC3339, endDate)
	if error != nil {
		return nil, nil, errors.New(fmt.Sprintf("error parsing Date %s", endDate))
	}

	receipts, error := s.receiptRepository.ReadByShiftBetweenDatesNotNull(sd, ed)

	if error != nil {
		return nil, nil, error
	}

	individual_receipts, error := s.individualReceiptRepository.ReadByShiftBetweenDatesNotNull(sd, ed)

	if error != nil {
		return nil, nil, error
	}

	return receipts, individual_receipts, nil
}

func (s *service) UpdateShift(id uint, shift *entities.UpdateShift) (*entities.Shift, error) {
	var receipts []entities.Receipt

	for i := 0; i < len(shift.Receipts); i++ {
		receiptWithId, error := s.receiptRepository.ReadById(shift.Receipts[i])

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no receipt with id %d", shift.Receipts[i]))
		}

		receipts = append(receipts, *receiptWithId)
	}

	var individual_receipts []entities.IndividualReceipt

	for i := 0; i < len(shift.Receipts); i++ {
		individual_receiptWithId, error := s.individualReceiptRepository.ReadById(shift.IndividualReceipts[i])

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no individual receipt with id %d", shift.IndividualReceipts[i]))
		}

		individual_receipts = append(individual_receipts, *individual_receiptWithId)
	}

	newReceipt := entities.Shift{}

	return s.repository.Update(id, &newReceipt)
}

func (s *service) RemoveShift(ID uint) (*entities.Shift, error) {
	return s.repository.Delete(ID)
}
