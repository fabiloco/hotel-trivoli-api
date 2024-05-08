package entities

import (
	"time"

	"github.com/guregu/null/v5"
	"gorm.io/gorm"
)

// Receipt model info
// @Description Receipt information in stock
type Receipt struct {
	gorm.Model
	TotalPrice float32          `gorm:"not null" json:"total_price"`
	TotalTime  time.Duration    `gorm:"not null" json:"total_time"`
	Products   []ReceiptProduct `json:"products"`
	Service    Service          `gorm:"not null" json:"service"`
	ServiceID  uint             `gorm:"not null"`
	Room       Room             `gorm:"not null" json:"room"`
	RoomID     uint             `gorm:"not null"`
	User       User             `gorm:"not null" json:"user"`
	UserID     uint             `gorm:"not null"`
	Shift      Shift            ` json:"shift"`
	ShiftID    null.Int
}

type CreateReceipt struct {
	TotalPrice float32 `valid:"required" json:"total_price"`
	TotalTime  uint    `valid:"required,numeric" json:"total_time"`
	Products   []uint  `valid:"optional" json:"products"`
	Service    uint    `valid:"required,numeric" json:"service"`
	Room       uint    `valid:"required,numeric" json:"room"`
	User       uint    `valid:"required,numeric" json:"user"`
	Shift      uint    `valid:"optional,numeric" json:"shift"`
}

type UpdateReceipt struct {
	TotalPrice float32 `valid:"optional" json:"total_price"`
	TotalTime  uint    `valid:"optional,numeric" json:"total_time"`
	Products   []uint  `valid:"optional" json:"products"`
	Service    uint    `valid:"optional,numeric" json:"service"`
	Room       uint    `valid:"optional,numeric" json:"room"`
	User       uint    `valid:"optional,numeric" json:"user"`
	Shift      uint    `valid:"optional,numeric" json:"shift"`
}
