package room

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
)

// Room is an interface from which our api module can access our repository of all our models
type Room interface {
	InsertRoom(productType *entities.CreateRoom) (*entities.Room, error)
	FetchRooms() (*[]entities.Room, error)
	FetchRoomById(id uint) (*entities.Room, error)
	UpdateRoom(id uint, product *entities.CreateRoom) (*entities.Room, error)
	RemoveRoom(id uint) (*entities.Room, error)
}

type room struct {
	repository Repository
}

func NewRoom(r Repository) Room {
	return &room{
		repository: r,
	}
}

func (s *room) InsertRoom(room *entities.CreateRoom) (*entities.Room, error) {
	newRoom := entities.Room{
		Number_room: room.Number_room,
	}

	return s.repository.Create(&newRoom)
}

func (s *room) FetchRooms() (*[]entities.Room, error) {
	return s.repository.Read()
}

func (s *room) UpdateRoom(id uint, room *entities.CreateRoom) (*entities.Room, error) {
	newRoom := entities.Room{
		Number_room: room.Number_room,
	}

	return s.repository.Update(id, &newRoom)
}

func (s *room) RemoveRoom(ID uint) (*entities.Room, error) {
	return s.repository.Delete(ID)
}

func (s *room) FetchRoomById(ID uint) (*entities.Room, error) {
	return s.repository.ReadById(ID)
}
