package product

import "fabiloco/hotel-trivoli-api/model"

type Store interface {
  List() ([]model.Product, error)
  FindById(id int) (*model.Product, error)
  Create(data *model.CreateProduct) (*model.Product, error)
  Delete(id int) (*model.Product, error)
  Update(id int, data *model.CreateProduct) (*model.Product, error)
}
