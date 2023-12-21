package model

import "gorm.io/gorm"

// User struct
type Product struct {
	gorm.Model
	Name    string    `gorm:"not null" json:"name"`
	Stock   uint8     `gorm:"not null" json:"stock"`
	Type    string    `gorm:"not null" json:"type"`
	Price   float32   `gorm:"not null" json:"price"`
}
