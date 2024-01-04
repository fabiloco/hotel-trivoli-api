package store

import (
	"fabiloco/hotel-trivoli-api/model"

	"gorm.io/gorm"
)

// UserStore maneja operaciones CRUD para el modelo User
type UserStore struct {
	db *gorm.DB
}

// NewUserStore crea una nueva instancia de UserStore
func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

// List retorna la lista de usuarios
func (store *UserStore) List() ([]model.User, error) {
	var users []model.User

	store.db.Find(&users)

	return users, nil
}

// FindById retorna un usuario por su ID
func (store *UserStore) FindById(id int) (*model.User, error) {
	var user model.User

	result := store.db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// Create crea un nuevo usuario
func (store *UserStore) Create(data *model.CreateUser) (*model.User, error) {
	var user model.User

	user = model.User{
		Username:       data.Username,
		Password:       data.Password,
		Firstname:      data.Firstname,
		Lastname:       data.Lastname,
		Identification: data.Identification,
	}

	result := store.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// Update actualiza un usuario existente
func (store *UserStore) Update(id int, data *model.CreateUser) (*model.User, error) {
	user, err := store.FindById(id)

	if err != nil {
		return nil, err
	}

	user.Username = data.Username
	user.Password = data.Password
	user.Firstname = data.Firstname
	user.Lastname = data.Lastname
	user.Identification = data.Identification

	result := store.db.Save(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// Delete elimina un usuario por su ID
func (store *UserStore) Delete(id int) (*model.User, error) {
	user, err := store.FindById(id)

	if err != nil {
		return nil, err
	}

	result := store.db.Delete(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
