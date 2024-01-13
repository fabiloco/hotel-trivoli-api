package role

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.Role) (*entities.Role, error)
	Read() (*[]entities.Role, error)
	Update(id uint, data *entities.Role) (*entities.Role, error)
	Delete(id uint) (*entities.Role, error)
	ReadById(id uint) (*entities.Role, error)
}


type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
    db: db,
	}
}

func (r *repository) Read() (*[]entities.Role, error) {
	var roles []entities.Role

	r.db.Find(&roles)

	return &roles, nil
}

func (r *repository) ReadById(id uint) (*entities.Role, error) {
	var role entities.Role

	result := r.db.First(&role, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (r *repository) Create(data *entities.Role) (*entities.Role, error) {
	var role entities.Role

	role = entities.Role{
    Name: data.Name,
	}

	result := r.db.Create(&role)

	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (r *repository) Update(id uint, data *entities.Role) (*entities.Role, error) {
	role, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Model(&role).Updates(
    entities.Role{
      Name: data.Name,
    },
  )

	if result.Error != nil {
		return nil, result.Error
	}

	return role, nil
}

func (r *repository) Delete(id uint) (*entities.Role, error) {
	role, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Delete(&role, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return role, nil
}
