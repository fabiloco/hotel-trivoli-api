package shift_presenter

import (
	receipt_presenter "fabiloco/hotel-trivoli-api/api/presenter/receipt"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"sort"
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

func ReceiptsToShiftsResponse2(receipts *[]entities.Receipt, individual_receipts *[]entities.IndividualReceipt) *[]ShiftResponse {

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

	sort.Slice(shiftsResponse, func(i, j int) bool {
		return shiftsResponse[i].ShiftID.Int64 > shiftsResponse[j].ShiftID.Int64
	})

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

func ReceiptsToShiftsResponse(
	receipts *[]entities.Receipt,
	individual_receipts *[]entities.IndividualReceipt,
	limit int,
	offset int,
) *[]ShiftResponse {
	//) []receipt_presenter.ReceiptResponse {

	var shiftsResponse []ShiftResponse
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

	// end := limit
	// if end > len(shiftsResponse) {
	// 	end = len(shiftsResponse)
	// }

	// paged := shiftsResponse[0:end]

	return &shiftsResponse

	/* 	// --- FASE 1: PRE-MAPEO Y CONVERSIÓN DE DATOS ---
	   	// --- FASE 1: PRE-MAPEO ---
	   	 mappedReceipts := make([]receipt_presenter.ReceiptResponse, len(*receipts))
	   	for i, receipt := range *receipts {
	   		mappedReceipts[i] = *receipt_presenter.ReceiptToReceiptResponse(&receipt)
	   	}

	   	mappedIndividualReceipts := make([]receipt_presenter.IndividualReceiptResponse, len(*individual_receipts))
	   	for i, ir := range *individual_receipts {
	   		mappedIndividualReceipts[i] = *receipt_presenter.IndividualReceiptToIndividualReceiptResponse(&ir)
	   	}

	   	mappedReceipts := receipt_presenter.ReceiptsToReceiptsResponses(*receipts)
	   	mappedIndividualReceipts := receipt_presenter.IndividualReceiptsToIndividualReceiptsResponses(*individual_receipts)

	   	fmt.Printf("AL COMENZAR Rec %d, IndRec %d, mappedReceipts %d.  mappedIndividualReceipts %d .", len(*receipts), len(*individual_receipts), len(mappedReceipts), len(mappedIndividualReceipts))
	   	fmt.Println(" ")
	   	// --- FASE 2: AGRUPACIÓN ---
	   	shiftsMap := make(map[int64]*ShiftResponse)

	   	// Agrupar recibos normales
	   	 for i := range mappedReceipts {
	   		receipt := &mappedReceipts[i]
	   		shiftEntity := (*receipts)[i]
	   		shiftID := shiftEntity.ShiftID.Int64

	   		if _, ok := shiftsMap[shiftID]; !ok {
	   			// Inicializa solo si no existe
	   			shiftsMap[shiftID] = &ShiftResponse{
	   				ShiftID:   shiftEntity.ShiftID,
	   				CreatedAt: shiftEntity.Shift.CreatedAt,
	   				UpdatedAt: shiftEntity.Shift.UpdatedAt,
	   				User:      shiftEntity.User,
	   				Receipts:  make([]receipt_presenter.ReceiptResponse, 0),
	   			}
	   		}
	   		// Acumula
	   		shiftsMap[shiftID].Receipts = append(shiftsMap[shiftID].Receipts, *receipt)
	   	}

	   	// Agrupar recibos individuales
	   	for i := range mappedIndividualReceipts {
	   		ir := &mappedIndividualReceipts[i]
	   		shiftEntity := (*individual_receipts)[i]
	   		shiftID := shiftEntity.ShiftID.Int64

	   		if _, ok := shiftsMap[shiftID]; !ok {
	   			shiftsMap[shiftID] = &ShiftResponse{
	   				ShiftID:            shiftEntity.ShiftID,
	   				CreatedAt:          shiftEntity.Shift.CreatedAt,
	   				UpdatedAt:          shiftEntity.Shift.UpdatedAt,
	   				User:               shiftEntity.User,
	   				IndividualReceipts: make([]receipt_presenter.IndividualReceiptResponse, 0),
	   			}
	   		}
	   		shiftsMap[shiftID].IndividualReceipts = append(shiftsMap[shiftID].IndividualReceipts, *ir)
	   	}

	   	// --- FASE 3: CONVERSIÓN A SLICE ---
	   	shiftsResponse := make([]ShiftResponse, 0, len(shiftsMap))
	   	for _, shiftResponse := range shiftsMap {
	   		shiftsResponse = append(shiftsResponse, *shiftResponse)
	   	}

	   	// --- FASE 4: ORDEN DESCENDENTE ---
	   	 sort.Slice(shiftsResponse, func(i, j int) bool {
	   		// 🔸 Cambia por CreatedAt si prefieres:
	   		return shiftsResponse[i].CreatedAt.After(shiftsResponse[j].CreatedAt)
	   		//return shiftsResponse[i].ShiftID.Int64 > shiftsResponse[j].ShiftID.Int64
	   	})

	   	// --- FASE 5: PAGINACIÓN FINAL ---
	   	 	start := offset
	   	   	end := offset + limit

	   	fmt.Printf("Rec %d, IndRec %d, mappedReceipts %d.  mappedIndividualReceipts %d .", len(*receipts), len(*individual_receipts), len(mappedReceipts), len(mappedIndividualReceipts))
	   	fmt.Println(" ")
	   	fmt.Printf("Response %d, shiftMap %d. ", len(shiftsResponse), len(shiftsMap))
	   	fmt.Println(" ")

	   	 if start > len(shiftsResponse) {
	   		start = len(shiftsResponse)
	   	}
	   	if end > len(shiftsResponse) {
	   		end = len(shiftsResponse)
	   	}

	   	//paged := shiftsResponse[0:limit]

	   	return mappedReceipts */
}
