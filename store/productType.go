package store

import (
	"fabiloco/hotel-trivoli-api/model"

	"gorm.io/gorm"
)

type ProductTypeStore struct {
	db *gorm.DB
}

func NewProductTypeStore(db *gorm.DB) *ProductTypeStore {
	return &ProductTypeStore{
		db: db,
	}
}

func (store *ProductTypeStore) List() ([]model.ProductType, error) {
	var productType []model.ProductType

	store.db.Find(&productType)

	return productType, nil
}

func (store *ProductTypeStore) FindById(id int) (*model.ProductType, error) {
	var productType model.ProductType

	result := store.db.First(&productType, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &productType, nil
}

func (store *ProductTypeStore) Create(data *model.CreateProductType) (*model.ProductType, error) {
	var productType model.ProductType

	productType = model.ProductType{
		Name:  data.Name,
	}

	result := store.db.Create(&productType)

	if result.Error != nil {
		return nil, result.Error
	}

	return &productType, nil
}

func (store *ProductTypeStore) Update(id int, data *model.CreateProductType) (*model.ProductType, error) {
	product, error := store.FindById(id)

	if error != nil {
		return nil, error
	}

	product.Name = data.Name

	result := store.db.Save(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (store *ProductTypeStore) Delete(id int) (*model.ProductType, error) {
	productType, error := store.FindById(id)

	if error != nil {
		return nil, error
	}

	result := store.db.Delete(&productType, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return productType, nil
}
