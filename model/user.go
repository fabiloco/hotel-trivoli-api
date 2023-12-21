package model

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Username  string  `gorm:"uniqueIndex;not null" json:"username"`
	Password  string  `gorm:"not null" json:"password"`
	Firstname  string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Identification  string  `json:"identification"`
}
