package entities

import (
	"time"
	"gorm.io/gorm"
)

// RoomHistory model info
// @Description Rooms information in stock
type RoomHistory struct {
	gorm.Model
  StartDate   time.Time   `gorm:"not null" json:"start_date"`
  EndDate     time.Time   `gorm:"optional" json:"end_date"`
  Room        Room        `gorm:"not null" json:"room"`
  RoomID      uint        `gorm:"not null"`
  Service     Service     `gorm:"not null" json:"service"`
  ServiceID   uint        `gorm:"not null"`
}

type CreateRoomHistory struct {
  StartDate   string   `valid:"required,rfc3339"`
  EndDate     string   `valid:"optional,rfc3339"`
  Room        uint     `valid:"required,numeric"`
  Service     uint     `valid:"required,numeric"`
}

type UpdateRoomHistory struct {
  StartDate   string   `valid:"optional,rfc3339"`
  EndDate     string   `valid:"optional,rfc3339"`
  Room        uint     `valid:"optional,numeric"`
  Service     uint     `valid:"optional,numeric"`
}
