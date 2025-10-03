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

	// --- FASE 1: PRE-MAPEO Y CONVERSIÓN DE DATOS (Una pasada por slice) ---

	// 1. Convertir todos los recibos principales a sus DTOs de respuesta
	mappedReceipts := make([]receipt_presenter.ReceiptResponse, len(*receipts))
	for i, receipt := range *receipts {
		// Asumiendo que esta es una función de presentación eficiente que
		// retorna el DTO por valor o puntero
		mappedReceipts[i] = *receipt_presenter.ReceiptToReceiptResponse(&receipt)
	}

	// 2. Convertir todos los recibos individuales a sus DTOs de respuesta
	mappedIndividualReceipts := make([]receipt_presenter.IndividualReceiptResponse, len(*individual_receipts))
	for i, ir := range *individual_receipts {
		mappedIndividualReceipts[i] = *receipt_presenter.IndividualReceiptToIndividualReceiptResponse(&ir)
	}

	// --- FASE 2: AGRUPACIÓN RÁPIDA (Solo acceso a memoria) ---

	shiftsMap := make(map[null.Int]*ShiftResponse)

	// 3. Agrupar Recibos Principales
	for i := range mappedReceipts {
		// Usar la referencia a la struct pre-mapeada
		receipt := &mappedReceipts[i]
		shiftID := (*receipts)[i].ShiftID // Obtener el ShiftID de la entidad original

		if _, ok := shiftsMap[shiftID]; !ok {
			// Inicializar el ShiftResponse
			shiftsMap[shiftID] = &ShiftResponse{
				ShiftID: shiftID,
				// Usar la entidad original para datos del Shift/User que no requieren conversión
				CreatedAt: (*receipts)[i].Shift.CreatedAt,
				UpdatedAt: (*receipts)[i].Shift.UpdatedAt,
				User:      (*receipts)[i].User,
				Receipts:  make([]receipt_presenter.ReceiptResponse, 0), // Inicializa la slice
			}
		}
		// Añadir el DTO pre-mapeado a la slice
		shiftsMap[shiftID].Receipts = append(shiftsMap[shiftID].Receipts, *receipt)
	}

	// 4. Agrupar Recibos Individuales
	for i := range mappedIndividualReceipts {
		// Usar la referencia a la struct pre-mapeada
		ir := &mappedIndividualReceipts[i]
		shiftID := (*individual_receipts)[i].ShiftID // Obtener el ShiftID de la entidad original

		if _, ok := shiftsMap[shiftID]; !ok {
			// Inicializar si el Shift no existía (basado en recibos individuales)
			shiftsMap[shiftID] = &ShiftResponse{
				ShiftID:            shiftID,
				CreatedAt:          (*individual_receipts)[i].Shift.CreatedAt,
				UpdatedAt:          (*individual_receipts)[i].Shift.UpdatedAt,
				User:               (*individual_receipts)[i].User,
				IndividualReceipts: make([]receipt_presenter.IndividualReceiptResponse, 0),
			}
		}
		// Añadir el DTO pre-mapeado a la slice
		shiftsMap[shiftID].IndividualReceipts = append(shiftsMap[shiftID].IndividualReceipts, *ir)
	}

	// --- FASE 3: CONVERSIÓN FINAL ---
	// 5. Convertir el mapa a slice (igual que antes)
	shiftsResponse := make([]ShiftResponse, 0, len(shiftsMap))
	for _, shiftResponse := range shiftsMap {
		shiftsResponse = append(shiftsResponse, *shiftResponse)
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
