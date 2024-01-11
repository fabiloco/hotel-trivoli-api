package producttype

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/room"
	serviceModule "fabiloco/hotel-trivoli-api/pkg/service"
	"fmt"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertRoomHistory (roomHistory *entities.CreateRoomHistory) (*entities.RoomHistory, error)
	FetchRoomHistorys () (*[]entities.RoomHistory, error)
  FetchRoomHistoryById (id uint) (*entities.RoomHistory, error)
	UpdateRoomHistory (id uint, roomHistory *entities.UpdateRoomHistory) (*entities.RoomHistory, error)
	RemoveRoomHistory (id uint) (*entities.RoomHistory, error)
}

type service struct {
	repository Repository
	roomRepository room.Repository
	serviceRepository serviceModule.Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertRoomHistory(roomHistory *entities.CreateRoomHistory) (*entities.RoomHistory, error) {

  room, error := s.roomRepository.ReadById(roomHistory.Room)

  if error != nil {
    return nil, errors.New(fmt.Sprintf("no room with id %d", roomHistory.Room))
  }

  service, error := s.serviceRepository.ReadById(roomHistory.Service)

  if error != nil {
    return nil, errors.New(fmt.Sprintf("no room with id %d", roomHistory.Room))
  }

  newRoomHistory := entities.RoomHistory {
    StartDate: roomHistory.StartDate,
    EndDate: roomHistory.EndDate,
    Room: *room,
    Service: *service,
  }


	return s.repository.Create(&newRoomHistory)
}

func (s *service) FetchRoomHistorys() (*[]entities.RoomHistory, error) {
	return s.repository.Read()
}

func (s *service) UpdateRoomHistory(id uint, roomHistory *entities.UpdateRoomHistory) (*entities.RoomHistory, error) {
  room, error := s.roomRepository.ReadById(roomHistory.Room)

  if error != nil {
    return nil, errors.New(fmt.Sprintf("no room with id %d", roomHistory.Room))
  }

  service, error := s.serviceRepository.ReadById(roomHistory.Service)

  if error != nil {
    return nil, errors.New(fmt.Sprintf("no room with id %d", roomHistory.Room))
  }

  newRoomHistory := entities.RoomHistory {
    StartDate: roomHistory.StartDate,
    EndDate: roomHistory.EndDate,
    Room: *room,
    Service: *service,
  }

	return s.repository.Update(id, &newRoomHistory)
}

func (s *service) RemoveRoomHistory(ID uint) (*entities.RoomHistory, error) {
	return s.repository.Delete(ID)
}

func (s *service) FetchRoomHistoryById(ID uint) (*entities.RoomHistory, error) {
	return s.repository.ReadById(ID)
}
