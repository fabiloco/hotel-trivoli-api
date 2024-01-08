package user

import "fabiloco/hotel-trivoli-api/pkg/entities"

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertUser (user *entities.User) (*entities.User, error)
	FetchUsers () (*[]entities.User, error)
  FetchUserById (id uint) (*entities.User, error)
	UpdateUser (id uint, user *entities.User) (*entities.User, error)
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

func (s *service) InsertUser(user *entities.User) (*entities.User, error) {
	return s.repository.Create(user)
}

func (s *service) FetchUsers() (*[]entities.User, error) {
	return s.repository.Read()
}

func (s *service) UpdateUser(id uint, user *entities.User) (*entities.User, error) {
	return s.repository.Update(id, user)
}

func (s *service) RemoveUser(ID uint) (*entities.User, error) {
	return s.repository.Delete(ID)
}

func (s *service) FetchUserById(ID uint) (*entities.User, error) {
	return s.repository.ReadById(ID)
}
