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

	Read(limit, offset int) (*[]entities.Receipt, int64, error)

	Update(id uint, data *entities.Receipt) (*entities.Receipt, error)
	UpdateShift(id uint, data *entities.Receipt) (*entities.Receipt, error)

	Delete(id uint) (*entities.Receipt, error)
	ReadById(id uint) (*entities.Receipt, error)

	ReadByShiftBetweenDatesNotNull(startDate time.Time, endDate time.Time) (*[]entities.Receipt, error)
	//ReadByShiftNotNull() (*[]entities.Receipt, error)
	ReadByShiftNotNull(limit, offset int) (*[]entities.Receipt, int64, error)
	ReadAllByShiftId(id uint) (*[]entities.Receipt, error)

	ReadByDate(targetDate time.Time, limit, offset int) (*[]entities.Receipt, int64, error)
	ReadBetweenDates(startDate time.Time, endDate time.Time) (*[]entities.Receipt, error)
	ReadBetweenDatesPaginated(startDate time.Time, endDate time.Time, params *entities.PaginationParams) (*[]entities.Receipt, int64, error)

	FindAllShiftIDs() ([]int64, error)
	FindByShiftIDs(shiftIDs []int64) ([]entities.Receipt, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Read(limit, offset int) (*[]entities.Receipt, int64, error) {
	var receipts []entities.Receipt

	var total int64

	if err := r.db.Model(&entities.Receipt{}).
		Where("shift_id IS NOT NULL").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	r.db.Preload("Products").Preload("Type").
		Preload("Service").Preload("Room").Preload("User").
		Preload("User.Person").Preload("Shift").
		/* Limit(limit).Offset(offset). */ Find(&receipts)

	return &receipts, total, nil
}

func (r *repository) ReadById(id uint) (*entities.Receipt, error) {
	var receipt entities.Receipt

	result := r.db.Preload("Products").Preload(clause.Associations).Preload("Service").Preload("Room").Preload("User").Preload("User.Person").Preload("Shift").First(&receipt, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipt, nil
}

func (r *repository) ReadByShiftBetweenDatesNotNull(startDate time.Time, endDate time.Time) (*[]entities.Receipt, error) {
	var receipts []entities.Receipt

	result := r.db.Preload("Products").Preload(clause.Associations).Preload("Service").Preload("Room").Preload("User").Preload("User.Person").Preload("Shift").Where("shift_id IS NOT NULL and created_at BETWEEN ? AND ?", startDate, endDate).Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}

// func (r *repository) ReadByShiftNotNull() (*[]entities.Receipt, error) {
func (r *repository) ReadByShiftNotNull(limit, offset int) (*[]entities.Receipt, int64, error) {

	var receipts []entities.Receipt
	var total int64

	if err := r.db.Model(&entities.Receipt{}).
		Where("shift_id IS NOT NULL").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	result := r.db.Preload("Products").
		Preload("Service").
		Preload("Room").
		Preload("User").
		Preload("User.Person").
		Preload("Shift").
		Where("shift_id IS NOT NULL").
		Limit(limit).
		Offset(offset).
		Find(&receipts)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return &receipts, total, nil

	/* var receipts []entities.Receipt

	result := r.db.Preload("Products").Preload("Service").Preload("Room").Preload("User").Preload("User.Person").Preload("Shift").Where("shift_id IS NOT NULL").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil */
}

func (r *repository) ReadAllByShiftId(id uint) (*[]entities.Receipt, error) {
	var receipt []entities.Receipt

	result := r.db.Preload("Products").Preload(clause.Associations).Preload("Service").Preload("Room").Preload("User").Preload("User.Person").Preload("Shift").Where("shift_id = ?", id).Find(&receipt)

	if result.Error != nil {
		empty := []entities.Receipt{}
		return &empty, result.Error
	}

	return &receipt, nil
}

func (r *repository) ReadByDate(targetDate time.Time, limit, offset int) (*[]entities.Receipt, int64, error) {
	var receipts []entities.Receipt

	result := r.db.Where("DATE(created_at) = DATE(?)", targetDate).Preload("Products").Preload("Service").Preload("User").Preload("User.Person").Preload("Shift").Find(&receipts)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return &receipts, 0, nil
}

func (r *repository) ReadBetweenDates(startDate time.Time, endDate time.Time) (*[]entities.Receipt, error) {
	var receipts []entities.Receipt

	result := r.db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Preload("Products").Preload("Service").Preload("User").Preload("User.Person").Preload("Shift").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}

func (r *repository) ReadBetweenDatesPaginated(startDate time.Time, endDate time.Time, params *entities.PaginationParams) (*[]entities.Receipt, int64, error) {
	var receipts []entities.Receipt
	var total int64

	params.Normalize()

	if err := r.db.Model(&entities.Receipt{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := params.GetOffset()
	limit := params.GetLimit()

	if err := r.db.Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Preload("Products").
		Preload("Service").
		Preload("User").
		Preload("User.Person").
		Preload("Shift").
		Offset(offset).
		Limit(limit).
		Find(&receipts).Error; err != nil {
		return nil, 0, err
	}

	return &receipts, total, nil
}

type ReadTotalsResult struct {
	TotalPrice float64
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

	result := r.db.Model(&receipt).Updates(
		entities.Receipt{
			TotalPrice: data.TotalPrice,
			TotalTime:  data.TotalTime,
		},
	)

	receipt.Service = data.Service
	receipt.User = data.User
	receipt.Room = data.Room
	receipt.Shift = data.Shift

	r.db.Save(receipt)

	if len(data.Products) > 0 {
		r.db.Delete(receipt.Products)
		r.db.Save(receipt)
		receipt.Products = data.Products
		r.db.Save(receipt)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return receipt, nil
}

func (r *repository) UpdateShift(id uint, data *entities.Receipt) (*entities.Receipt, error) {
	receipt, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	receipt.Shift = data.Shift

	r.db.Save(receipt)

	fmt.Println(receipt)

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

// FindAllShiftIDs obtiene todos los shift_id distintos presentes en receipts
func (r *repository) FindAllShiftIDs() ([]int64, error) {
	var shiftIDs []int64
	err := r.db.
		Model(&entities.Receipt{}).
		Where("shift_id IS NOT NULL").
		Distinct().
		Order("shift_id DESC").
		Pluck("shift_id", &shiftIDs).
		Error

	if err != nil {
		return nil, err
	}

	return shiftIDs, nil
}

// FindByShiftIDs obtiene todos los receipts pertenecientes a los shift_id dados
func (r *repository) FindByShiftIDs(shiftIDs []int64) ([]entities.Receipt, error) {
	if len(shiftIDs) == 0 {
		return []entities.Receipt{}, nil
	}

	var receipts []entities.Receipt
	err := r.db.
		Preload("Shift").
		Preload("User").
		Preload("Service").
		Preload("Products").
		Where("shift_id IN ?", shiftIDs).
		Find(&receipts).
		Error

	if err != nil {
		return nil, err
	}

	return receipts, nil
}
