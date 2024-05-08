package receipt

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(data *entities.Receipt) (*entities.Receipt, error)
	Read() (*[]entities.Receipt, error)
	Update(id uint, data *entities.Receipt) (*entities.Receipt, error)
	Delete(id uint) (*entities.Receipt, error)
	ReadById(id uint) (*entities.Receipt, error)
	ReadByShiftNotNull() (*[]entities.Receipt, error)
	ReadAllByShiftId(id uint) (*[]entities.Receipt, error)
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

	r.db.Preload("Products").Preload("Type").Preload("Service").Preload("Room").Preload("User").Preload("User.Person").Find(&receipts)

	return &receipts, nil
}

func (r *repository) ReadById(id uint) (*entities.Receipt, error) {
	var receipt entities.Receipt

	result := r.db.Preload("Products").Preload(clause.Associations).Preload("Service").Preload("Room").Preload("User").Preload("User.Person").First(&receipt, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipt, nil
}

func (r *repository) ReadByShiftNotNull() (*[]entities.Receipt, error) {
	var receipts []entities.Receipt

	result := r.db.Preload("Products").Preload(clause.Associations).Preload("Service").Preload("Room").Preload("User").Preload("User.Person").Where("shift_id IS NOT NULL").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}

func (r *repository) ReadAllByShiftId(id uint) (*[]entities.Receipt, error) {
	var receipt []entities.Receipt

	result := r.db.Preload("Products").Preload(clause.Associations).Preload("Service").Preload("Room").Preload("User").Preload("User.Person").Where("shift_id = ?", id).First(&receipt)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipt, nil
}

func (r *repository) ReadByDate(targetDate time.Time) (*[]entities.Receipt, error) {
	var receipts []entities.Receipt

	result := r.db.Where("DATE(created_at) = DATE(?)", targetDate).Preload("Products").Preload("Service").Preload("User").Preload("User.Person").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}

func (r *repository) ReadBetweenDates(startDate time.Time, endDate time.Time) (*[]entities.Receipt, error) {
	var receipts []entities.Receipt

	result := r.db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Preload("Products").Preload("Service").Preload("User").Preload("User.Person").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}

func (r *repository) Create(data *entities.Receipt) (*entities.Receipt, error) {
	var receipt entities.Receipt

	receipt = entities.Receipt{
		TotalPrice: data.TotalPrice,
		TotalTime:  data.TotalTime,
		Service:    data.Service,
		Room:       data.Room,
		Products:   data.Products,
		User:       data.User,
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

	fmt.Println("antes de editar")
	fmt.Println(data.Shift)

	result := r.db.Model(&receipt).Updates(
		entities.Receipt{
			TotalPrice: data.TotalPrice,
			TotalTime:  data.TotalTime,
			Service:    data.Service,
			Room:       data.Room,
			Products:   data.Products,
			User:       data.User,
		},
	)

	receipt.Shift = data.Shift

	r.db.Save(receipt)

	fmt.Println("luego de editar")
	fmt.Println(receipt.Shift)

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
