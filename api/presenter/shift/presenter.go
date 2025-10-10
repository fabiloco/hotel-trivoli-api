package shift_presenter

import (
	receipt_presenter "fabiloco/hotel-trivoli-api/api/presenter/receipt"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"time"

	"github.com/guregu/null/v5"
)

type ShiftResponse struct {
	ShiftID            null.Int                                      `json:"shift_id"`
	Receipts           []receipt_presenter.ReceiptResponse           `json:"receipts"`
	IndividualReceipts []receipt_presenter.IndividualReceiptResponse `json:"individual_receipts"`
	User               entities.User                                 `json:"user"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func ReceiptsToShiftsResponse(receipts *[]entities.Receipt, individual_receipts *[]entities.IndividualReceipt) *[]ShiftResponse {
	shiftsMap := make(map[null.Int]*ShiftResponse)

	for _, receipt := range *receipts {
		if _, ok := shiftsMap[receipt.ShiftID]; !ok {
			shiftsMap[receipt.ShiftID] = &ShiftResponse{
				ShiftID:   receipt.ShiftID,
				CreatedAt: receipt.Shift.CreatedAt,
				UpdatedAt: receipt.Shift.UpdatedAt,
				User:      receipt.User,
				Receipts:  []receipt_presenter.ReceiptResponse{},
			}
		}
		shiftsMap[receipt.ShiftID].Receipts = append(shiftsMap[receipt.ShiftID].Receipts, *receipt_presenter.ReceiptToReceiptResponse(&receipt))
	}

	for _, individual_receipt := range *individual_receipts {
		if _, ok := shiftsMap[individual_receipt.ShiftID]; !ok {
			shiftsMap[individual_receipt.ShiftID] = &ShiftResponse{
				ShiftID:            individual_receipt.ShiftID,
				CreatedAt:          individual_receipt.Shift.CreatedAt,
				UpdatedAt:          individual_receipt.Shift.UpdatedAt,
				User:               individual_receipt.User,
				IndividualReceipts: []receipt_presenter.IndividualReceiptResponse{},
			}
		}
		shiftsMap[individual_receipt.ShiftID].IndividualReceipts = append(shiftsMap[individual_receipt.ShiftID].IndividualReceipts, *receipt_presenter.IndividualReceiptToIndividualReceiptResponse(&individual_receipt))
	}

	shiftsResponse := make([]ShiftResponse, 0, len(shiftsMap))
	for _, shiftResponse := range shiftsMap {
		shiftsResponse = append(shiftsResponse, *shiftResponse)
	}

	return &shiftsResponse
}

func ShiftsToShiftsResponse(shifts *[]entities.Shift) *[]ShiftResponse {
	shiftsResponse := make([]ShiftResponse, 0, len(*shifts))

	for _, shift := range *shifts {
		receipts := make([]receipt_presenter.ReceiptResponse, 0, len(shift.Receipts))
		for _, receipt := range shift.Receipts {
			receipts = append(receipts, *receipt_presenter.ReceiptToReceiptResponse(&receipt))
		}

		individualReceipts := make([]receipt_presenter.IndividualReceiptResponse, 0, len(shift.IndividualReceipts))
		for _, individualReceipt := range shift.IndividualReceipts {
			individualReceipts = append(individualReceipts, *receipt_presenter.IndividualReceiptToIndividualReceiptResponse(&individualReceipt))
		}

		var user entities.User
		if len(shift.Receipts) > 0 {
			user = shift.Receipts[0].User
		} else if len(shift.IndividualReceipts) > 0 {
			user = shift.IndividualReceipts[0].User
		}

		shiftResponse := ShiftResponse{
			ShiftID:            null.IntFrom(int64(shift.ID)),
			CreatedAt:          shift.CreatedAt,
			UpdatedAt:          shift.UpdatedAt,
			User:               user,
			Receipts:           receipts,
			IndividualReceipts: individualReceipts,
		}
		shiftsResponse = append(shiftsResponse, shiftResponse)
	}

	return &shiftsResponse

	/* var shiftsResponse []ShiftResponse
	shiftsMap := make(map[null.Int]*ShiftResponse)

	for _, receipt := range *receipts {
		if _, ok := shiftsMap[receipt.ShiftID]; !ok {
			shiftsMap[receipt.ShiftID] = &ShiftResponse{
				ShiftID:   receipt.ShiftID,
				CreatedAt: receipt.Shift.CreatedAt,
				UpdatedAt: receipt.Shift.UpdatedAt,
				User:      receipt.User,
				Receipts:  []receipt_presenter.ReceiptResponse{},
			}
			shiftsMap[receipt.ShiftID].Receipts = append(shiftsMap[receipt.ShiftID].Receipts, *receipt_presenter.ReceiptToReceiptResponse(&receipt))
		} else {
			shiftsMap[receipt.ShiftID].Receipts = append(shiftsMap[receipt.ShiftID].Receipts, *receipt_presenter.ReceiptToReceiptResponse(&receipt))
		}
	}

	for _, individual_receipt := range *individual_receipts {
		if _, ok := shiftsMap[individual_receipt.ShiftID]; !ok {
			shiftsMap[individual_receipt.ShiftID] = &ShiftResponse{
				ShiftID:            individual_receipt.ShiftID,
				CreatedAt:          individual_receipt.Shift.CreatedAt,
				UpdatedAt:          individual_receipt.Shift.UpdatedAt,
				User:               individual_receipt.User,
				IndividualReceipts: []receipt_presenter.IndividualReceiptResponse{},
			}
			shiftsMap[individual_receipt.ShiftID].IndividualReceipts = append(shiftsMap[individual_receipt.ShiftID].IndividualReceipts, *receipt_presenter.IndividualReceiptToIndividualReceiptResponse(&individual_receipt))
		} else {
			shiftsMap[individual_receipt.ShiftID].IndividualReceipts = append(shiftsMap[individual_receipt.ShiftID].IndividualReceipts, *receipt_presenter.IndividualReceiptToIndividualReceiptResponse(&individual_receipt))
		}
	}

	for _, shiftResponse := range shiftsMap {
		shiftsResponse = append(shiftsResponse, *shiftResponse)
	}

	return &shiftsResponse */
}
