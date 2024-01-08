package entities

import "gorm.io/gorm"

// Product model info
// @Description Products information in stock
type Product struct {
	gorm.Model
	Name  string  `gorm:"not null" json:"name"`  // product name
	Stock uint8  `gorm:"not null" json:"stock"` // product stock avaliable
  Type  []ProductType  `gorm:"many2many:product_types_product" json:"type"`  // product type
	Price float32 `gorm:"not null" json:"price"` // product price
}

type CreateProduct struct {
	Name  string  `valid:"required,stringlength(3|100)"`
	Stock uint8   `valid:"required,numeric"`
  Type  []uint   `valid:"required"`
	Price float32 `valid:"required,numeric"`
}
