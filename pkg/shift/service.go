package shift

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	receipt "fabiloco/hotel-trivoli-api/pkg/receipt"
	"fmt"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertShift(receipt *entities.CreateShift) (*entities.Shift, error)
	FetchShiftsById(id uint) (*[]entities.Receipt, error)
	FetchAllShifts() (*[]entities.Receipt, error)
	UpdateShift(id uint, receipt *entities.UpdateShift) (*entities.Shift, error)
	RemoveShift(id uint) (*entities.Shift, error)
}

type service struct {
	repository        Repository
	receiptRepository receipt.Repository
}

func NewService(r Repository, rr receipt.Repository) Service {
	return &service{
		repository:        r,
		receiptRepository: rr,
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

	return shiftCreated, err
}

func (s *service) FetchShiftsById(id uint) (*[]entities.Receipt, error) {
	return s.receiptRepository.ReadAllByShiftId(id)
}

func (s *service) FetchAllShifts() (*[]entities.Receipt, error) {
	return s.receiptRepository.ReadByShiftNotNull()
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

	newReceipt := entities.Shift{}

	return s.repository.Update(id, &newReceipt)
}

func (s *service) RemoveShift(ID uint) (*entities.Shift, error) {
	return s.repository.Delete(ID)
}
