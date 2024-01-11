package entities

import "gorm.io/gorm"

// Room model info
// @Description Rooms information in stock
type Room struct {
	gorm.Model
	Number int `gorm:"not null" json:"number"` // room name

}

type CreateRoom struct {
	Number int `valid:"required"`
}

type UpdateRoom struct {
	Number int `valid:"optional"`
}
