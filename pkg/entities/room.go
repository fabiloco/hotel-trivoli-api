package entities

import "gorm.io/gorm"

// Room model info
// @Description Rooms information in stock
type Room struct {
	gorm.Model
	Number_room int `gorm:"not null" json:"Numer_room"` // room name

}

type CreateRoom struct {
	Number_room int `valid:"required"`
}
