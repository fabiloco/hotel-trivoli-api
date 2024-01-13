package auth

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/role"
	"fabiloco/hotel-trivoli-api/pkg/user"
	"fmt"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	Register (user *entities.CreateUser, person *entities.CreatePerson) (*entities.User, error)
}

type service struct {
	userRepository user.Repository
	roleRepository role.Repository
}

func NewService(ur user.Repository, rr role.Repository) Service {
	return &service{
		userRepository: ur,
		roleRepository: rr,
	}
}

func (s *service) Register(user *entities.CreateUser, person *entities.CreatePerson) (*entities.User, error) {
	newPerson := entities.Person{
    Firstname: person.Firstname,
    Lastname: person.Lastname,
    Identification: person.Identification,
    Birthday: person.Birthday,
	}

	newUser := entities.User{
    Username: user.Username,
    Password: user.Password,
    Person: newPerson,
	}

  role, error := s.roleRepository.ReadById(user.Role)

  if error != nil {
    role, error := s.roleRepository.ReadById(1)
    if error != nil {
      return nil, errors.New(fmt.Sprintf("No default roles setup. Please, seed the database with the default values and roles"))
    }
    newUser.Role = *role
  } else {
    newUser.Role = *role
  }

  return s.userRepository.Create(&newUser)
}
