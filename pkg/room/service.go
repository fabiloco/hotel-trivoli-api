package room

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertRoom(productType *entities.CreateRoom) (*entities.Room, error)
	FetchRooms() (*[]entities.Room, error)
	FetchRoomById(id uint) (*entities.Room, error)
	UpdateRoom(id uint, product *entities.CreateRoom) (*entities.Room, error)
	RemoveRoom(id uint) (*entities.Room, error)
}

type service struct {
	repository Repository
}

func NewRoom(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertRoom(room *entities.CreateRoom) (*entities.Room, error) {
	newRoom := entities.Room{
		Number_room: room.Number_room,
	}

	return s.repository.Create(&newRoom)
}

func (s *service) FetchRooms() (*[]entities.Room, error) {
	return s.repository.Read()
}

func (s *service) UpdateRoom(id uint, room *entities.CreateRoom) (*entities.Room, error) {
	newRoom := entities.Room{
		Number_room: room.Number_room,
	}

	return s.repository.Update(id, &newRoom)
}

func (s *service) RemoveRoom(ID uint) (*entities.Room, error) {
	return s.repository.Delete(ID)
}

func (s *service) FetchRoomById(ID uint) (*entities.Room, error) {
	return s.repository.ReadById(ID)
}
