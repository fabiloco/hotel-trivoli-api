package producttype

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.ProductType) (*entities.ProductType, error)
	Read() (*[]entities.ProductType, error)
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
		Name:  data.Name,
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

	productType.Name = data.Name

	result := r.db.Save(&productType)

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
