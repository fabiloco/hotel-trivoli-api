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
	Receipts []uint `valid:"required" json:"receipts"`
}

type UpdateShift struct {
	Receipts []uint `valid:"optional" json:"receipts"`
}
