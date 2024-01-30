package entities

import "gorm.io/gorm"

// Service model info
// @Description Services information in stock
type Service struct {
	gorm.Model
	Name    string  `gorm:"not null" json:"name"`     // service name
	Price   float32 `gorm:"not null" json:"price"`    // service price
	Details string  `gorm:"not null" json:"details"`  // service price
}

type CreateService struct {
	Name    string  `valid:"required,stringlength(3|100)"`
	Price   float32 `valid:"required,numeric"`
  Details string  `valid:"required,stringlength(3|100)"`
}

type UpdateService struct {
	Name    string  `valid:"optional,stringlength(3|100)"`
	Price   float32 `valid:"optional,numeric"`
  Details string  `valid:"required,stringlength(3|100)"`
}
