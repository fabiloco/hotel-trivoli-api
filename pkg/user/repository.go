package user

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.User) (*entities.User, error)
	Read() (*[]entities.User, error)
	Update(id uint, data *entities.User) (*entities.User, error)
	Delete(id uint) (*entities.User, error)
	ReadById(id uint) (*entities.User, error)
}


type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
    db: db,
	}
}

func (r *repository) Read() (*[]entities.User, error) {
	var users []entities.User

	r.db.Find(&users)

	return &users, nil
}

func (r *repository) ReadById(id uint) (*entities.User, error) {
	var user entities.User

	result := r.db.Find(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *repository) Create(data *entities.User) (*entities.User, error) {
	var user entities.User

	user = entities.User{
    Username: data.Username,
    Firstname: data.Firstname,
    Lastname: data.Lastname,
    Password: data.Password,
    Identification: data.Identification,
	}

	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *repository) Update(id uint, data *entities.User) (*entities.User, error) {
	user, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

  user.Username = data.Username
  user.Firstname = data.Firstname
  user.Lastname = data.Lastname
  user.Password = data.Password
  user.Identification = data.Identification

	result := r.db.Save(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *repository) Delete(id uint) (*entities.User, error) {
	user, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Delete(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
