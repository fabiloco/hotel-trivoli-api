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
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func ReceiptsToShiftsResponse(receipts *[]entities.Receipt, individual_receipts *[]entities.IndividualReceipt) *[]ShiftResponse {
	var shiftsResponse []ShiftResponse
	shiftsMap := make(map[null.Int]*ShiftResponse)

	for _, receipt := range *receipts {
		if _, ok := shiftsMap[receipt.ShiftID]; !ok {
			shiftsMap[receipt.ShiftID] = &ShiftResponse{
				ShiftID:   receipt.ShiftID,
				CreatedAt: receipt.Shift.CreatedAt,
				UpdatedAt: receipt.Shift.UpdatedAt,
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

	return &shiftsResponse
}
