package shift

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	individualreceipt "fabiloco/hotel-trivoli-api/pkg/individual_receipt"
	receipt "fabiloco/hotel-trivoli-api/pkg/receipt"
	"fmt"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertShift(receipt *entities.CreateShift) (*entities.Shift, error)
	FetchShiftsById(id uint) (*[]entities.Receipt, *[]entities.IndividualReceipt, error)
	FetchAllShifts() (*[]entities.Receipt, *[]entities.IndividualReceipt, error)
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

		_, error := s.receiptRepository.Update(receipt.ID, &receipt)

		if error != nil {
			return nil, errors.New(fmt.Sprintf("error editing receipt with id %d", receipt.ID))
		}
	}

	for _, individual_receipt := range individual_receipts {
		individual_receipt.Shift = *shiftCreated

		_, error := s.individualReceiptRepository.Update(individual_receipt.ID, &individual_receipt)

		if error != nil {
			return nil, errors.New(fmt.Sprintf("error editing individual receipt with id %d", individual_receipt.ID))
		}
	}

	return shiftCreated, err
}

func (s *service) FetchShiftsById(id uint) (*[]entities.Receipt, *[]entities.IndividualReceipt, error) {
	receipts, error := s.receiptRepository.ReadAllByShiftId(id)

	if error != nil {
		return nil, nil, error
	}

	individual_receipts, error := s.individualReceiptRepository.ReadAllByShiftId(id)

	if error != nil {
		return nil, nil, error
	}

	return receipts, individual_receipts, nil
}

func (s *service) FetchAllShifts() (*[]entities.Receipt, *[]entities.IndividualReceipt, error) {
	receipts, error := s.receiptRepository.ReadByShiftNotNull()

	if error != nil {
		return nil, nil, error
	}

	individual_receipts, error := s.individualReceiptRepository.ReadByShiftNotNull()

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
