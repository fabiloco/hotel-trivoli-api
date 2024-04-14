package individualreceipt

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.IndividualReceipt) (*entities.IndividualReceipt, error)
	Read() (*[]entities.IndividualReceipt, error)
	Update(id uint, data *entities.IndividualReceipt) (*entities.IndividualReceipt, error)
	Delete(id uint) (*entities.IndividualReceipt, error)
	ReadById(id uint) (*entities.IndividualReceipt, error)
	ReadByDate(targetDate time.Time) (*[]entities.IndividualReceipt, error)
	ReadBetweenDates(startDate time.Time, endDate time.Time) (*[]entities.IndividualReceipt, error)
}


type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
    db: db,
	}
}

func (r *repository) Read() (*[]entities.IndividualReceipt, error) {
	var receipts []entities.IndividualReceipt

	r.db.Preload("Products").Preload("User").Find(&receipts)

	return &receipts, nil
}

func (r *repository) ReadById(id uint) (*entities.IndividualReceipt, error) {
	var receipt entities.IndividualReceipt

	result := r.db.Preload("Products").Preload("User").First(&receipt, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipt, nil
}

func (r *repository) ReadByDate(targetDate time.Time) (*[]entities.IndividualReceipt, error) {
	var receipts []entities.IndividualReceipt

	result := r.db.Where("DATE(created_at) = DATE(?)", targetDate).Preload("Products").Preload("User").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}


func (r *repository) ReadBetweenDates(startDate time.Time, endDate time.Time) (*[]entities.IndividualReceipt, error) {
	var receipts []entities.IndividualReceipt

	result := r.db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Preload("Products").Preload("User").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}

func (r *repository) Create(data *entities.IndividualReceipt) (*entities.IndividualReceipt, error) {
	var receipt entities.IndividualReceipt

	receipt = entities.IndividualReceipt{
    TotalPrice: data.TotalPrice,
    Products: data.Products,
    User: data.User,
	}

	result := r.db.Create(&receipt)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipt, nil
}

func (r *repository) Update(id uint, data *entities.IndividualReceipt) (*entities.IndividualReceipt, error) {
	receipt, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Model(&receipt).Updates(
    entities.IndividualReceipt{
      TotalPrice: data.TotalPrice,
      Products: data.Products,
      User: data.User,
    },
  )

	if result.Error != nil {
		return nil, result.Error
	}

	return receipt, nil
}

func (r *repository) Delete(id uint) (*entities.IndividualReceipt, error) {
	receipt, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Delete(&receipt, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return receipt, nil
}