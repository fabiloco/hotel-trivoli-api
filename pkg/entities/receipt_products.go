package entities

import "gorm.io/gorm"

// ReceiptProduct model info
// @Description Relation between Receipt and Product
type ReceiptProduct struct {
  gorm.Model
  ReceiptID      uint 
  ProductID      uint         
}
