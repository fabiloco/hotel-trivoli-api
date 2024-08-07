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
	ReadByUsername(username string) (*entities.User, error)
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

	result := r.db.Preload("Person").Find(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *repository) ReadByUsername(username string) (*entities.User, error) {
	var user entities.User

	result := r.db.Where("username = ?", username).Preload("Role").Preload("Person").First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *repository) Create(data *entities.User) (*entities.User, error) {
	var user entities.User

	user = entities.User{
		Username: data.Username,
		Password: data.Password,
		Role:     data.Role,
		Person:   data.Person,
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

	result := r.db.Model(&user).Updates(
		entities.User{
			Password: data.Password,
			Username: data.Username,
			Role:     data.Role,
			Person:   data.Person,
		},
	)

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
	r.db.Unscoped().Delete(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
