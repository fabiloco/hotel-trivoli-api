package entities

import "gorm.io/gorm"

// User model info
// @Description User registered in the system
type User struct {
	gorm.Model
	Username    string  `gorm:"not null;unique" json:"username"`
  Password    string  `gorm:"not null" json:"-"`
  Role        Role    `gorm:"not null" json:"role"`
  RoleID      uint    `gorm:"not null"`
  Person      Person  `gorm:"not null" json:"person"`
  PersonID    uint    `gorm:"not null"`
}

type CreateUser struct {
  Username    string  `valid:"required,stringlength(3|100)" json:"username"`
	Password    string  `valid:"required,stringlength(3|100)" json:"password"`
  Person      uint    `valid:"required,numeric" json:"person"`
  Role        uint    `valid:"required,numeric" json:"role"`
}

type UpdateUser struct {
  Username    string  `valid:"optional,stringlength(3|100)" json:"username"`
	Password    string  `valid:"optional,stringlength(3|40)" json:"password"`
  Person      uint    `valid:"optional,numeric" json:"person"`
  Role        uint    `valid:"optional,numeric" json:"role"`
}
