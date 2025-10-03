package user

import (
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	FetchUsers() (*[]entities.User, error)
	FetchUserById(id uint) (*entities.User, error)
	RemoveUser(id uint) (*entities.User, error)
	UpdateUserById(ID uint, data *entities.UserPatch) (*entities.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FetchUsers() (*[]entities.User, error) {
	return s.repository.Read()
}

func (s *service) RemoveUser(ID uint) (*entities.User, error) {
	return s.repository.Delete(ID)
}

func (s *service) FetchUserById(ID uint) (*entities.User, error) {
	return s.repository.ReadById(ID)
}

func (s *service) UpdateUserById(ID uint, data *entities.UserPatch) (*entities.User, error) {
	if data.Password != nil {
		passwordHashed, err := utils.HashPassword(*data.Password) // Asume que HashPassword acepta string
		if err != nil {
			return nil, err
		}
		// Reemplaza el puntero del DTO con el hash
		data.Password = &passwordHashed
	}

	// 2. Ejecutar la actualización en el repositorio
	// Aquí es donde el repositorio manejará las entidades anidadas.
	updatedUser, err := s.repository.Update(ID, data)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
	//return s.repository.Update(ID, data)
}
