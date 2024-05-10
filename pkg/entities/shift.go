package entities

import (
	"gorm.io/gorm"
)

// Receipt model info
// @Description Receipt information in stock
type Shift struct {
	gorm.Model
}

type CreateShift struct {
	Receipts           []uint `valid:"optional" json:"receipts"`
	IndividualReceipts []uint `valid:"optional" json:"individual_receipts"`
}

type UpdateShift struct {
	Receipts           []uint `valid:"optional" json:"receipts"`
	IndividualReceipts []uint `valid:"optional" json:"individual_receipts"`
}

type CreateShiftWithCode struct {
	Receipts           []string `valid:"optional" json:"receipts"`
	IndividualReceipts []string `valid:"optional" json:"individual_receipts"`
}

type UpdateShiftWithCode struct {
	Receipts           []string `valid:"optional" json:"receipts"`
	IndividualReceipts []string `valid:"optional" json:"individual_receipts"`
}
