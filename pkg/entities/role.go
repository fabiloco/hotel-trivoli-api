package entities

import "gorm.io/gorm"

type Role struct {
	gorm.Model
  Name  string  `gorm:"not null,default:'USER'" json:"role"`
}
