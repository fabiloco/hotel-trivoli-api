package product

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	productType "fabiloco/hotel-trivoli-api/pkg/product_type"
	"fmt"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertProduct(productType *entities.CreateProduct, file string) (*entities.Product, error)
	RestockProduct(id uint, productRestock *entities.RestockProduct) (*entities.Product, error)
	FetchProducts() (*[]entities.Product, error)
	FetchProductsPaginated(params *entities.PaginationParams) (*entities.PaginatedResponse, error)
	FetchProductById(id uint) (*entities.Product, error)
	UpdateProduct(id uint, product *entities.UpdateProduct, file string) (*entities.Product, error)
	RemoveProduct(id uint) (*entities.Product, error)
}

type service struct {
	productRepository     Repository
	productTypeRepository productType.Repository
}

func NewService(pr Repository, ptr productType.Repository) Service {
	return &service{
		productRepository:     pr,
		productTypeRepository: ptr,
	}
}

func (s *service) InsertProduct(product *entities.CreateProduct, file string) (*entities.Product, error) {
	var productTypesSlice []entities.ProductType

	for i := 0; i < len(product.Type); i++ {
		productType, error := s.productTypeRepository.ReadById(product.Type[i])

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no product type with id %d", product.Type[i]))
		}

		productTypesSlice = append(productTypesSlice, *productType)
	}

	newProduct := entities.Product{
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
		Type:  productTypesSlice,
		Img:   file,
	}

	return s.productRepository.Create(&newProduct)
}

func (s *service) RestockProduct(id uint, productRestock *entities.RestockProduct) (*entities.Product, error) {
	productWithId, error := s.productRepository.ReadById(id)

	if error != nil {
		return nil, errors.New(fmt.Sprintf("no product with id %d", id))
	}

	productWithId.Stock += productRestock.Stock

	return s.productRepository.Update(id, productWithId)
}

func (s *service) FetchProducts() (*[]entities.Product, error) {
	return s.productRepository.Read()
}

func (s *service) FetchProductsPaginated(params *entities.PaginationParams) (*entities.PaginatedResponse, error) {
	products, total, err := s.productRepository.ReadPaginated(params)
	if err != nil {
		return nil, err
	}

	return entities.NewPaginatedResponse(products, total, params.Page, params.PageSize), nil
}

func (s *service) UpdateProduct(id uint, product *entities.UpdateProduct, file string) (*entities.Product, error) {
	var productTypesSlice []entities.ProductType

	for i := 0; i < len(product.Type); i++ {
		productType, error := s.productTypeRepository.ReadById(product.Type[i])

		if error != nil {
			return nil, errors.New(fmt.Sprintf("no product type with id %d", product.Type[i]))
		}

		productTypesSlice = append(productTypesSlice, *productType)
	}

	var newProduct entities.Product
	if file == "no file" {
		newProduct = entities.Product{
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
			Type:  productTypesSlice,
		}
	} else {
		newProduct = entities.Product{
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
			Type:  productTypesSlice,
			Img:   file,
		}
	}

	return s.productRepository.Update(id, &newProduct)
}

func (s *service) RemoveProduct(ID uint) (*entities.Product, error) {
	return s.productRepository.Delete(ID)
}

func (s *service) FetchProductById(ID uint) (*entities.Product, error) {
	return s.productRepository.ReadById(ID)
}
