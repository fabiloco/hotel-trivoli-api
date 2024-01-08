package product

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.Product) (*entities.Product, error)
	Read() (*[]entities.Product, error)
	Update(id uint, data *entities.Product) (*entities.Product, error)
	Delete(id uint) (*entities.Product, error)
	ReadById(id uint) (*entities.Product, error)
}


type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
    db: db,
	}
}

func (r *repository) Read() (*[]entities.Product, error) {
	var products []entities.Product

	r.db.Preload("Type").Find(&products)

	return &products, nil
}

func (r *repository) ReadById(id uint) (*entities.Product, error) {
	var product entities.Product

	result := r.db.Preload("Type").First(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *repository) Create(data *entities.Product) (*entities.Product, error) {
	var product entities.Product

	product = entities.Product{
		Name:  data.Name,
		Stock: data.Stock,
		Price: data.Price,
		Type:  data.Type,
	}

	result := r.db.Create(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *repository) Update(id uint, data *entities.Product) (*entities.Product, error) {
	product, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	product.Name = data.Name
	product.Price = data.Price
	product.Stock = data.Stock
	product.Type = data.Type

	result := r.db.Save(&product)

  r.db.Model(&product).Association("Type").Replace(data.Type)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *repository) Delete(id uint) (*entities.Product, error) {
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
