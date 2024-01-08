package producttype

import "fabiloco/hotel-trivoli-api/model"

type Store interface {
	List() ([]model.ProductType, error)
	FindById(id uint) (*model.ProductType, error)
	Create(data *model.CreateProductType) (*model.ProductType, error)
	Delete(id uint) (*model.ProductType, error)
	Update(id uint, data *model.CreateProductType) (*model.ProductType, error)
}
