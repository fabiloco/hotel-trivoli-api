package entities

import "gorm.io/gorm"

// IndividualReceiptProduct model info
// @Description Relation between Receipt and Product
type IndividualReceiptProduct struct {
  gorm.Model
  IndividualReceiptID      uint 
  ProductID      uint         
}
