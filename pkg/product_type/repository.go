package producttype

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.ProductType) (*entities.ProductType, error)
	Read() (*[]entities.ProductType, error)
	ReadPaginated(params *entities.PaginationParams) (*[]entities.ProductType, int64, error)
	Update(id uint, data *entities.ProductType) (*entities.ProductType, error)
	Delete(id uint) (*entities.ProductType, error)
	ReadById(id uint) (*entities.ProductType, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Read() (*[]entities.ProductType, error) {
	var productTypes []entities.ProductType

	r.db.Find(&productTypes)

	return &productTypes, nil
}

func (r *repository) ReadPaginated(params *entities.PaginationParams) (*[]entities.ProductType, int64, error) {
	var productTypes []entities.ProductType
	var total int64

	params.Normalize()

	if err := r.db.Model(&entities.ProductType{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := params.GetOffset()
	limit := params.GetLimit()

	if err := r.db.Offset(offset).Limit(limit).Find(&productTypes).Error; err != nil {
		return nil, 0, err
	}

	return &productTypes, total, nil
}

func (r *repository) ReadById(id uint) (*entities.ProductType, error) {
	var productType entities.ProductType

	result := r.db.First(&productType, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &productType, nil
}

func (r *repository) Create(data *entities.ProductType) (*entities.ProductType, error) {
	var productType entities.ProductType

	productType = entities.ProductType{
		Name: data.Name,
	}

	result := r.db.Create(&productType)

	if result.Error != nil {
		return nil, result.Error
	}

	return &productType, nil
}

func (r *repository) Update(id uint, data *entities.ProductType) (*entities.ProductType, error) {
	productType, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Model(&productType).Updates(
		entities.ProductType{Name: data.Name},
	)

	if result.Error != nil {
		return nil, result.Error
	}

	return productType, nil
}

func (r *repository) Delete(id uint) (*entities.ProductType, error) {
	productType, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Delete(&productType, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return productType, nil
}
