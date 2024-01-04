package user

import "fabiloco/hotel-trivoli-api/model"

type UserStore interface {
	List() ([]model.User, error)
	FindById(id int) (*model.User, error)
	Create(data *model.CreateUser) (*model.User, error)
	Delete(id int) (*model.User, error)
	Update(id int, data *model.CreateUser) (*model.User, error)
}
