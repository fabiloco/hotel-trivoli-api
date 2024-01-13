package user

import "fabiloco/hotel-trivoli-api/pkg/entities"

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	FetchUsers () (*[]entities.User, error)
  FetchUserById (id uint) (*entities.User, error)
	RemoveUser (id uint) (*entities.User, error)
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
