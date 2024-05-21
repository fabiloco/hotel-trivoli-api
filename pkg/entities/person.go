package entities

import "gorm.io/gorm"

// Person model info
// @Description Person information in the system
type Person struct {
	gorm.Model
	Firstname      string `gorm:"not null" json:"firstname"`
	Lastname       string `gorm:"not null" json:"lastname"`
	Identification string `gorm:"not null" json:"identification"`
	Birthday       string `gorm:"not null" json:"birthday"`
}

type CreatePerson struct {
	Firstname      string `valid:"required,stringlength(3|100)" json:"firstname"`
	Lastname       string `valid:"required,stringlength(3|100)" json:"lastname"`
	Identification string `valid:"required,numeric" json:"identification"`
	Birthday       string `valid:"required" json:"brithday"`
}

type UpdatePerson struct {
	Firstname      string `valid:"optional,stringlength(3|100)" json:"firstname"`
	Lastname       string `valid:"optional,stringlength(3|100)" json:"lastname"`
	Identification string `valid:"optional,numeric" json:"identification"`
	Birthday       string `gorm:"not null" json:"birthday"`
}
