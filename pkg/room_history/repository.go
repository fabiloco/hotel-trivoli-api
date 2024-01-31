package producttype

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	Create(data *entities.RoomHistory) (*entities.RoomHistory, error)
	Read() (*[]entities.RoomHistory, error)
	Update(id uint, data *entities.RoomHistory) (*entities.RoomHistory, error)
	Delete(id uint) (*entities.RoomHistory, error)
	ReadById(id uint) (*entities.RoomHistory, error)
}


type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
    db: db,
	}
}

func (r *repository) Read() (*[]entities.RoomHistory, error) {
	var roomHistorys []entities.RoomHistory

	r.db.Preload("Room").Find(&roomHistorys)

	return &roomHistorys, nil
}

func (r *repository) ReadById(id uint) (*entities.RoomHistory, error) {
	var roomHistory entities.RoomHistory

	result := r.db.Preload("Room").First(&roomHistory, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &roomHistory, nil
}

func (r *repository) Create(data *entities.RoomHistory) (*entities.RoomHistory, error) {
	var roomHistory entities.RoomHistory

	roomHistory = entities.RoomHistory{
    StartDate: data.StartDate,
    EndDate: nil,
    Room: data.Room,
    Service: data.Service,
	}

	result := r.db.Create(&roomHistory)

	if result.Error != nil {
		return nil, result.Error
	}

	return &roomHistory, nil
}

func (r *repository) Update(id uint, data *entities.RoomHistory) (*entities.RoomHistory, error) {
	roomHistory, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Model(&roomHistory).Updates(
    entities.RoomHistory{
      StartDate: data.StartDate,
      EndDate: data.EndDate,
      Room: data.Room,
      Service: data.Service,
    },
  )

	if result.Error != nil {
		return nil, result.Error
	}

	return roomHistory, nil
}

func (r *repository) Delete(id uint) (*entities.RoomHistory, error) {
	roomHistory, error := r.ReadById(id)

	if error != nil {
		return nil, error
	}

	result := r.db.Delete(&roomHistory, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return roomHistory, nil
}
