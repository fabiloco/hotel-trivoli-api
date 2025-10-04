package individualreceipt

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.IndividualReceipt) (*entities.IndividualReceipt, error)
	Read(limit, offset int) (*[]entities.IndividualReceipt, int64, error)

	Update(id uint, data *entities.IndividualReceipt) (*entities.IndividualReceipt, error)
	UpdateShift(id uint, data *entities.IndividualReceipt) (*entities.IndividualReceipt, error)

	Delete(id uint) (*entities.IndividualReceipt, error)
	ReadById(id uint) (*entities.IndividualReceipt, error)

	ReadByShiftBetweenDatesNotNull(startDate time.Time, endDate time.Time) (*[]entities.IndividualReceipt, error)
	//ReadByShiftNotNull() (*[]entities.IndividualReceipt, error)
	ReadByShiftNotNull(limit, offset int) (*[]entities.IndividualReceipt, int64, error)
	ReadAllByShiftId(id uint) (*[]entities.IndividualReceipt, error)

	ReadByDate(targetDate time.Time, limit, offset int) (*[]entities.IndividualReceipt, int64, error)
	ReadBetweenDates(startDate time.Time, endDate time.Time) (*[]entities.IndividualReceipt, error)
	ReadBetweenDatesPaginated(startDate time.Time, endDate time.Time, params *entities.PaginationParams) (*[]entities.IndividualReceipt, int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Read(limit, offset int) (*[]entities.IndividualReceipt, int64, error) {
	var receipts []entities.IndividualReceipt
	var total int64

	if err := r.db.Model(&entities.IndividualReceipt{}).
		Where("shift_id IS NOT NULL").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	r.db.Preload("Products").Preload("User").Preload("User.Person").Preload("Shift").Limit(limit).
		Offset(offset).Find(&receipts)

	return &receipts, total, nil
}

func (r *repository) ReadById(id uint) (*entities.IndividualReceipt, error) {
	var receipt entities.IndividualReceipt

	result := r.db.Preload("Products").Preload("User").Preload("User.Person").Preload("Shift").First(&receipt, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipt, nil
}

func (r *repository) ReadByDate(targetDate time.Time, limit, offset int) (*[]entities.IndividualReceipt, int64, error) {
	var receipts []entities.IndividualReceipt
	var total int64

	if err := r.db.Model(&entities.IndividualReceipt{}).
		Where("shift_id IS NOT NULL AND DATE(created_at) = DATE(?)", targetDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	result := r.db.Where("DATE(created_at) = DATE(?)", targetDate).Preload("Products").Preload("User").Preload("User.Person").Preload("Shift").Limit(limit).
		Offset(offset).Find(&receipts)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return &receipts, total, nil
}

func (r *repository) ReadBetweenDates(startDate time.Time, endDate time.Time) (*[]entities.IndividualReceipt, error) {
	var receipts []entities.IndividualReceipt

	result := r.db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Preload("Products").Preload("User").Preload("User.Person").Preload("Shift").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}

func (r *repository) ReadBetweenDatesPaginated(startDate time.Time, endDate time.Time, params *entities.PaginationParams) (*[]entities.IndividualReceipt, int64, error) {
	var receipts []entities.IndividualReceipt
	var total int64

	params.Normalize()

	if err := r.db.Model(&entities.IndividualReceipt{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := params.GetOffset()
	limit := params.GetLimit()

	if err := r.db.Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Preload("Products").
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

func (r *repository) Create(data *entities.IndividualReceipt) (*entities.IndividualReceipt, error) {
	var receipt entities.IndividualReceipt

	receipt = entities.IndividualReceipt{
		TotalPrice: data.TotalPrice,
		Products:   data.Products,
		User:       data.User,
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
		},
	)

	receipt.User = data.User
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

func (r *repository) UpdateShift(id uint, data *entities.IndividualReceipt) (*entities.IndividualReceipt, error) {
	receipt, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	receipt.Shift = data.Shift

	r.db.Save(receipt)

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

func (r *repository) ReadByShiftNotNull(limit, offset int) (*[]entities.IndividualReceipt, int64, error) {

	var receipts []entities.IndividualReceipt
	var total int64

	if err := r.db.Model(&entities.IndividualReceipt{}).
		Where("shift_id IS NOT NULL").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	result := r.db.Preload("User").Preload("User.Person").Preload("Products").Preload("Shift").Where("shift_id IS NOT NULL").Limit(limit).Offset(offset).Find(&receipts)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return &receipts, total, nil

	/* var receipts []entities.IndividualReceipt

	result := r.db.Preload("User").Preload("User.Person").Preload("Products").Preload("Shift").Where("shift_id IS NOT NULL").Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil */
}

func (r *repository) ReadByShiftBetweenDatesNotNull(startDate time.Time, endDate time.Time) (*[]entities.IndividualReceipt, error) {
	var receipts []entities.IndividualReceipt

	result := r.db.Preload("Products").Preload("User").Preload("User.Person").Preload("Shift").Where("shift_id IS NOT NULL and created_at BETWEEN ? AND ?", startDate, endDate).Find(&receipts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &receipts, nil
}

func (r *repository) ReadAllByShiftId(id uint) (*[]entities.IndividualReceipt, error) {
	var receipt []entities.IndividualReceipt

	result := r.db.Preload("User").Preload("User.Person").Preload("Products").Preload("Shift").Where("shift_id = ?", id).Find(&receipt)

	if result.Error != nil {
		empty := []entities.IndividualReceipt{}
		return &empty, result.Error
	}

	return &receipt, nil
}
