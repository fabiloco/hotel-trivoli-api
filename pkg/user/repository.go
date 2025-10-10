package user

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.User) (*entities.User, error)
	Read() (*[]entities.User, error)
	//Update(id uint, data *entities.User) (*entities.User, error)
	Update(id uint, data *entities.UserPatch) (*entities.User, error)
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

	r.db.Preload("Person").Preload("Role").Find(&users)

	return &users, nil
}

func (r *repository) ReadById(id uint) (*entities.User, error) {
	var user entities.User

	result := r.db.Preload("Person").Preload("Role").Find(&user, id)

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

/* func (r *repository) Update(id uint, data *entities.User) (*entities.User, error) {
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
*/

func (r *repository) Update(id uint, data *entities.UserPatch) (*entities.User, error) {

	// 1. Cargar la entidad base completa
	user, err := r.ReadById(id)
	if err != nil {
		return nil, err
	}

	// --- A. Actualizar la entidad Person ---
	if data.Person != nil {
		// Usa GORM para actualizar la entidad 'Person' del usuario.
		// 'user.Person' es un struct completo de la entidad Person.
		// 'data.Person' es el DTO con punteros, GORM lo maneja como PATCH.
		result := r.db.Model(&user.Person).Updates(data.Person)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	fmt.Printf("Rol %d.", data.RoleID)

	// --- B. Actualizar la entidad User ---
	// Usamos el mismo struct DTO para actualizar los campos del User principal.
	// Como 'data' (UserPatchDTO) tiene un campo 'Person *PersonPatchDTO',
	// GORM está dando error. Para solucionarlo, debemos usar un mapa o un struct
	// anónimo SÓLO con los campos de la tabla 'User'.

	updateData := make(map[string]interface{})

	if data.Username != nil {
		updateData["username"] = *data.Username
	}

	if data.RoleID != nil {
		updateData["role_id"] = *data.RoleID
	}

	if data.Password != nil {
		// El servicio ya hizo el Hash, solo actualiza el campo
		updateData["password"] = *data.Password
	}

	// Solo actualiza si hay algo que actualizar para el usuario
	if len(updateData) > 0 {
		result := r.db.Model(&user).Select("Username", "Password", "RoleID").Updates(updateData)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	// 3. Recargar y retornar el usuario completo
	return r.ReadById(id)
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
