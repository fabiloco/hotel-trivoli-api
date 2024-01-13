package receipt

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.Receipt) (*entities.Receipt, error)
	Read() (*[]entities.Receipt, error)
	Update(id uint, data *entities.Receipt) (*entities.Receipt, error)
	Delete(id uint) (*entities.Receipt, error)
	ReadById(id uint) (*entities.Receipt, error)
	ReadByDate(targetDate time.Time) (*[]entities.Receipt, error)
	ReadBetweenDates(startDate time.Time, endDate time.Time) (*[]entities.Receipt, error)
}


type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
    db: db,
	}
}

func (r *repository) Read() (*[]entities.Receipt, error) {
	var receipts []entities.Receipt

	r.db.Preload("Products").Find(&receipts)

	return &receipts, nil
}

func (r *repository) ReadById(id uint) (*entities.Receipt, error) {
	var receipt entities.Receipt

	result := r.db.Preload("Products").First(&receipt, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipt, nil
}

func (r *repository) ReadByDate(targetDate time.Time) (*[]entities.Receipt, error) {
	var receipts []entities.Receipt

	result := r.db.Where("DATE(created_at) = DATE(?)", targetDate).Preload("Products").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}


func (r *repository) ReadBetweenDates(startDate time.Time, endDate time.Time) (*[]entities.Receipt, error) {
	var receipts []entities.Receipt

	result := r.db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Preload("Products").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}

func (r *repository) Create(data *entities.Receipt) (*entities.Receipt, error) {
	var receipt entities.Receipt

	receipt = entities.Receipt{
    TotalPrice: data.TotalPrice,
    TotalTime: data.TotalTime,
    Service: data.Service,
    Room: data.Room,
    Products: data.Products,
	}

	result := r.db.Create(&receipt)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipt, nil
}

func (r *repository) Update(id uint, data *entities.Receipt) (*entities.Receipt, error) {
	receipt, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Model(&receipt).Updates(
    entities.Receipt{
      TotalPrice: data.TotalPrice,
      TotalTime: data.TotalTime,
      Service: data.Service,
      Room: data.Room,
      Products: data.Products,
    },
  )

	if result.Error != nil {
		return nil, result.Error
	}

	return receipt, nil
}

func (r *repository) Delete(id uint) (*entities.Receipt, error) {
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
