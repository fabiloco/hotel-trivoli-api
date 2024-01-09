package entities

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Username       string `gorm:"uniqueIndex;not null" json:"username"`
	Password       string `gorm:"not null" json:"password"`
	Firstname      string `gorm:"not null" json:"firstname"`
	Lastname       string `gorm:"not null" json:"lastname"`
	Identification string `gorm:"uniqueIndex;not null" json:"identification"`
}

type CreateUser struct {
	Username       string `valid:"required,stringlength(3|100)"`
	Password       string `valid:"required,stringlength(5|40)"`
	Firstname      string `valid:"required,stringlength(3|100)"`
	Lastname       string `valid:"required,stringlength(3|100)"`
	Identification string `valid:"required,numeric"`
}

type UpdateUser struct {
	Username       string `valid:"optional,stringlength(3|100)"`
	Password       string `valid:"optional,stringlength(5|40)"`
	Firstname      string `valid:"optional,stringlength(3|100)"`
	Lastname       string `valid:"optional,stringlength(3|100)"`
	Identification string `valid:"optional,numeric"`
}
