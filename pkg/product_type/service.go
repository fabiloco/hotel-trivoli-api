package producttype

import "fabiloco/hotel-trivoli-api/pkg/entities"

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertProductType (productType *entities.CreateProductType) (*entities.ProductType, error)
	FetchProductTypes () (*[]entities.ProductType, error)
  FetchProductTypeById (id uint) (*entities.ProductType, error)
	UpdateProductType (id uint, productType *entities.CreateProductType) (*entities.ProductType, error)
	RemoveProductType (id uint) (*entities.ProductType, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertProductType(productType *entities.CreateProductType) (*entities.ProductType, error) {
  newProductType := entities.ProductType {
    Name: productType.Name,
  }

	return s.repository.Create(&newProductType)
}

func (s *service) FetchProductTypes() (*[]entities.ProductType, error) {
	return s.repository.Read()
}

func (s *service) UpdateProductType(id uint, productType *entities.CreateProductType) (*entities.ProductType, error) {
  newProductType := entities.ProductType {
    Name: productType.Name,
  }

	return s.repository.Update(id, &newProductType)
}

func (s *service) RemoveProductType(ID uint) (*entities.ProductType, error) {
	return s.repository.Delete(ID)
}

func (s *service) FetchProductTypeById(ID uint) (*entities.ProductType, error) {
	return s.repository.ReadById(ID)
}
