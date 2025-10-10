package shift

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.Shift) (*entities.Shift, error)
	Read() (*[]entities.Shift, error)
	Update(id uint, data *entities.Shift) (*entities.Shift, error)
	Delete(id uint) (*entities.Shift, error)
	ReadById(id uint) (*entities.Shift, error)
	ReadUniqueShifts(limit, offset int) (*[]entities.Shift, int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Read() (*[]entities.Shift, error) {
	var shifts []entities.Shift

	r.db.Preload("Receipts").Find(&shifts)

	return &shifts, nil
}

func (r *repository) ReadById(id uint) (*entities.Shift, error) {
	var shift entities.Shift

	result := r.db.First(&shift, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &shift, nil
}

func (r *repository) Create(data *entities.Shift) (*entities.Shift, error) {
	var shift entities.Shift

	result := r.db.Create(&shift)

	if result.Error != nil {
		return nil, result.Error
	}

	return &shift, nil
}

func (r *repository) Update(id uint, data *entities.Shift) (*entities.Shift, error) {
	shift, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Model(&shift).Updates(
		entities.Shift{
			Model: data.Model,
		},
	)

	if result.Error != nil {
		return nil, result.Error
	}

	return shift, nil
}

func (r *repository) Delete(id uint) (*entities.Shift, error) {
	shift, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Delete(&shift, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return shift, nil
}

func (r *repository) ReadUniqueShifts(limit, offset int) (*[]entities.Shift, int64, error) {
	var shifts []entities.Shift
	var total int64

	if err := r.db.Raw(`
		SELECT COUNT(DISTINCT s.id) 
		FROM shifts s 
		WHERE EXISTS (
			SELECT 1 FROM receipts r WHERE r.shift_id = s.id
		) OR EXISTS (
			SELECT 1 FROM individual_receipts ir WHERE ir.shift_id = s.id
		)
	`).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	result := r.db.Raw(`
		SELECT DISTINCT s.* 
		FROM shifts s 
		WHERE EXISTS (
			SELECT 1 FROM receipts r WHERE r.shift_id = s.id
		) OR EXISTS (
			SELECT 1 FROM individual_receipts ir WHERE ir.shift_id = s.id
		)
		ORDER BY s.id DESC
		LIMIT ? OFFSET ?
	`, limit, offset).Scan(&shifts)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return &shifts, total, nil
}
