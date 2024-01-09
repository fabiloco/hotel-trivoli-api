package entities

import "gorm.io/gorm"

// Product type model info
// @Description Product type information
type ProductType struct {
	gorm.Model
	Name  string  `gorm:"not null" json:"name"`  // type name
}

type CreateProductType struct {
	Name  string  `valid:"required,stringlength(3|100)"`
}

type UpdateProductType struct {
	Name  string  `valid:"optional,stringlength(3|100)"`
}
