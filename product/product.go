package product

import "fabiloco/hotel-trivoli-api/model"

type Store interface {
	List() ([]model.Product, error)
	FindById(id uint) (*model.Product, error)
	Create(data *model.Product) (*model.Product, error)
	Delete(id uint) (*model.Product, error)
	Update(id uint, data *model.Product) (*model.Product, error)
}
