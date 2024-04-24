package entities

import (
	"mime/multipart"

	"gorm.io/gorm"
)

// Product model info
// @Description Products information in stock
type Product struct {
	gorm.Model
	Name  string  `gorm:"not null" json:"name"`   // product name
	Stock int  `json:"stock"`   // product stock avaliable
  Type  []ProductType  `gorm:"many2many:product_types_product" json:"type"`  // product type
	Price float32 `gorm:"not null" json:"price"`  // product price
	Img   string  `gorm:"not null" json:"img"`    // product price
}

type RestockProduct struct {
	Stock int   `valid:"required,numeric"`
}

type CreateProduct struct {
	Name  string  `valid:"required,stringlength(3|100)"`
	Stock int   `valid:"required,numeric"`
  Type  []uint  `valid:"required"`
	Price float32 `valid:"required"`
  Img   multipart.File  `valid:"required"`
}

type UpdateProduct struct {
	Name  string          `valid:"stringlength(3|100),optional"`
	Stock int            `valid:"numeric,optional"`
  Type  []uint          `valid:"optional"`
	Price float32         `valid:"optional"`
  Img   multipart.File  `valid:"optional"`
}
