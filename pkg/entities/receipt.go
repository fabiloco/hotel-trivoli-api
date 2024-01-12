package entities

import (
	"time"
	"gorm.io/gorm"
)

// Receipt model info
// @Description Receipt information in stock
type Receipt struct {
	gorm.Model
  TotalPrice  float32         `gorm:"not null" json:"total_price"`
  TotalTime   time.Duration   `gorm:"not null" json:"total_time"`
  Products    []Product       `gorm:"many2many:receipt_product" json:"products"` 
  Service     Service         `gorm:"not null" json:"service"`
  ServiceID   uint            `gorm:"not null"`
  Room        Room            `gorm:"not null" json:"room"`
  RoomID      uint            `gorm:"not null"`
}

type CreateReceipt struct {
  TotalPrice  float32         `valid:"required,numeric" json:"total_price"`
  TotalTime   uint            `valid:"required,numeric" json:"total_time"`
  Products    []uint          `valid:"required" json:"products"`
  Service     uint            `valid:"required,numeric" json:"service"`
  Room        uint            `valid:"required,numeric" json:"room"`
}

type UpdateReceipt struct {
  TotalPrice  float32         `valid:"optional,numeric" json:"total_price"`
  TotalTime   uint            `valid:"optional,numeric" json:"total_time"`
  Products    []uint          `valid:"optional" json:"products"`
  Service     uint            `valid:"optional,numeric" json:"service"`
  Room        uint            `valid:"optional,numeric" json:"room"`
}
