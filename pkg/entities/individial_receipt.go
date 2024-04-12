package entities

import (
	"gorm.io/gorm"
)

// Individual Receipt model info
// @Description Individual Receipt information in stock
type IndividualReceipt struct {
	gorm.Model
  TotalPrice  float32         `gorm:"not null" json:"total_price"`
  Products    []Product       `gorm:"many2many:individual_receipt_product" json:"products"` 
  User        User            `gorm:"not null" json:"user"`
  UserID      uint            `gorm:"not null"`
}

type CreateIndividualReceipt struct {
  TotalPrice  float32         `valid:"required,numeric" json:"total_price"`
  Products    []uint          `valid:"required" json:"products"`
  User        uint            `valid:"required,numeric" json:"user"`
}

type UpdateIndividualReceipt struct {
  TotalPrice  float32         `valid:"optional,numeric" json:"total_price"`
  Products    []uint          `valid:"optional" json:"products"`
  User        uint            `valid:"optional,numeric" json:"user"`
}
