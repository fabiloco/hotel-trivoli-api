package room

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.Room) (*entities.Room, error)
	Read() (*[]entities.Room, error)
	Update(id uint, data *entities.Room) (*entities.Room, error)
	Delete(id uint) (*entities.Room, error)
	ReadById(id uint) (*entities.Room, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Read() (*[]entities.Room, error) {
	var rooms []entities.Room

	r.db.Preload("Type").Find(&rooms)

	return &rooms, nil
}

func (r *repository) ReadById(id uint) (*entities.Room, error) {
	var room entities.Room

	result := r.db.Preload("Type").First(&room, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &room, nil
}

func (r *repository) Create(data *entities.Room) (*entities.Room, error) {
	var room entities.Room

	room = entities.Room{
		Number: data.Number,
	}

	result := r.db.Create(&room)

	if result.Error != nil {
		return nil, result.Error
	}

	return &room, nil
}

func (r *repository) Update(id uint, data *entities.Room) (*entities.Room, error) {
	room, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	room.Number = data.Number

	result := r.db.Save(&room)

	r.db.Model(&room).Association("Type").Replace(data.Number)

	if result.Error != nil {
		return nil, result.Error
	}

	return room, nil
}

func (r *repository) Delete(id uint) (*entities.Room, error) {
	room, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Delete(&room, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return room, nil
}
