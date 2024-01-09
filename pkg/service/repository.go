package service

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.Service) (*entities.Service, error)
	Read() (*[]entities.Service, error)
	Update(id uint, data *entities.Service) (*entities.Service, error)
	Delete(id uint) (*entities.Service, error)
	ReadById(id uint) (*entities.Service, error)
}


type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
    db: db,
	}
}

func (r *repository) Read() (*[]entities.Service, error) {
	var products []entities.Service

	r.db.Find(&products)

	return &products, nil
}

func (r *repository) ReadById(id uint) (*entities.Service, error) {
	var product entities.Service

	result := r.db.First(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *repository) Create(data *entities.Service) (*entities.Service, error) {
	var product entities.Service

	product = entities.Service{
		Name:  data.Name,
		Price: data.Price,
	}

	result := r.db.Create(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *repository) Update(id uint, data *entities.Service) (*entities.Service, error) {
	product, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	product.Name = data.Name
	product.Price = data.Price

	result := r.db.Save(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *repository) Delete(id uint) (*entities.Service, error) {
	product, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Delete(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}
