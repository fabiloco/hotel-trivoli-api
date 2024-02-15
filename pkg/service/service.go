package service

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertService (productType *entities.CreateService) (*entities.Service, error)
	FetchServices () (*[]entities.Service, error)
  FetchServiceById (id uint) (*entities.Service, error)
	UpdateService (id uint, product *entities.UpdateService) (*entities.Service, error)
	RemoveService (id uint) (*entities.Service, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertService(service *entities.CreateService) (*entities.Service, error) {
  newService := entities.Service {
    Name: service.Name,
    Price: service.Price,
    Details: service.Details,
  }

	return s.repository.Create(&newService)
}

func (s *service) FetchServices() (*[]entities.Service, error) {
	return s.repository.Read()
}

func (s *service) UpdateService(id uint, service *entities.UpdateService) (*entities.Service, error) {
  newService := entities.Service {
    Name: service.Name,
    Price: service.Price,
    Details: service.Details,
  }

	return s.repository.Update(id, &newService)
}

func (s *service) RemoveService(ID uint) (*entities.Service, error) {
	return s.repository.Delete(ID)
}

func (s *service) FetchServiceById(ID uint) (*entities.Service, error) {
	return s.repository.ReadById(ID)
}
