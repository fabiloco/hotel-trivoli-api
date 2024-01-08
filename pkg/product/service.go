package product

import "fabiloco/hotel-trivoli-api/pkg/entities"

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertProduct (book *entities.Product) (*entities.Product, error)
	FetchProducts () (*[]entities.Product, error)
  FetchProductById (id uint) (*entities.Product, error)
	UpdateProduct (id uint, product *entities.Product) (*entities.Product, error)
	RemoveProduct (id uint) (*entities.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertProduct(product *entities.Product) (*entities.Product, error) {
	return s.repository.Create(product)
}

func (s *service) FetchProducts() (*[]entities.Product, error) {
	return s.repository.Read()
}

func (s *service) UpdateProduct(id uint, product *entities.Product) (*entities.Product, error) {
	return s.repository.Update(id, product)
}

func (s *service) RemoveProduct(ID uint) (*entities.Product, error) {
	return s.repository.Delete(ID)
}

func (s *service) FetchProductById(ID uint) (*entities.Product, error) {
	return s.repository.ReadById(ID)
}
