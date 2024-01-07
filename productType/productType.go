package producttype

import "fabiloco/hotel-trivoli-api/model"

type Store interface {
	List() ([]model.ProductType, error)
	FindById(id int) (*model.ProductType, error)
	Create(data *model.CreateProductType) (*model.ProductType, error)
	Delete(id int) (*model.ProductType, error)
	Update(id int, data *model.CreateProductType) (*model.ProductType, error)
}
