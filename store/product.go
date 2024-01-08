package store

import (
	"fabiloco/hotel-trivoli-api/model"
	"gorm.io/gorm"
)

type ProductStore struct {
	db *gorm.DB
}

func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{
		db: db,
	}
}

func (store *ProductStore) List() ([]model.Product, error) {
	var products []model.Product

	store.db.Preload("Type").Find(&products)

	return products, nil
}

func (store *ProductStore) FindById(id uint) (*model.Product, error) {
	var product model.Product

	result := store.db.First(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (store *ProductStore) Create(data *model.Product) (*model.Product, error) {
	var product model.Product

	product = model.Product{
		Name:  data.Name,
		Stock: data.Stock,
		Price: data.Price,
		Type:  data.Type,
	}

	result := store.db.Create(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (store *ProductStore) Update(id uint, data *model.Product) (*model.Product, error) {
	product, error := store.FindById(id)

	if error != nil {
		return nil, error
	}

	product.Name = data.Name
	product.Price = data.Price
	product.Stock = data.Stock
	product.Type = data.Type

	result := store.db.Save(&product)

  store.db.Model(&product).Association("Type").Replace(data.Type)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (store *ProductStore) Delete(id uint) (*model.Product, error) {
	product, error := store.FindById(id)

	if error != nil {
		return nil, error
	}

	result := store.db.Delete(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}
