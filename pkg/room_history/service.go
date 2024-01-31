package producttype

import (
	"errors"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fabiloco/hotel-trivoli-api/pkg/room"
	serviceModule "fabiloco/hotel-trivoli-api/pkg/service"
	"fmt"
	"time"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertRoomHistory (roomHistory *entities.CreateRoomHistory) (*entities.RoomHistory, error)
	FetchRoomHistorys () (*[]entities.RoomHistory, error)
  FetchRoomHistoryById (id uint) (*entities.RoomHistory, error)
	UpdateRoomHistory (id uint, roomHistory *entities.UpdateRoomHistory) (*entities.RoomHistory, error)
	RemoveRoomHistory (id uint) (*entities.RoomHistory, error)

	SetEndDate (id uint, roomHistory *entities.SetEndDateRoomHistory) (*entities.RoomHistory, error)
}

type service struct {
	repository Repository
	roomRepository room.Repository
	serviceRepository serviceModule.Repository
}

func NewService(r Repository, rr room.Repository, sr serviceModule.Repository) Service {
	return &service{
		repository: r,
    roomRepository: rr,
    serviceRepository: sr,
	}
}

func (s *service) InsertRoomHistory(roomHistory *entities.CreateRoomHistory) (*entities.RoomHistory, error) {

  room, error := s.roomRepository.ReadById(roomHistory.Room)

  if error != nil {
    return nil, errors.New(fmt.Sprintf("no room with id %d", roomHistory.Room))
  }

  service, error := s.serviceRepository.ReadById(roomHistory.Service)

  if error != nil {
    return nil, errors.New(fmt.Sprintf("no service with id %d", roomHistory.Service))
  }

  sd, error := time.Parse(time.RFC3339, roomHistory.StartDate)
  if error != nil {
    return nil, errors.New(fmt.Sprintf("error parsing StartDate %s", roomHistory.StartDate))
  }

  // ed, error := time.Parse(time.RFC3339, roomHistory.EndDate)
  // if error != nil {
  //   return nil, errors.New(fmt.Sprintf("error parsing EndDate %s", roomHistory.EndDate))
  // }

  newRoomHistory := entities.RoomHistory {
    StartDate: sd,
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
    return nil, errors.New(fmt.Sprintf("no service with id %d", roomHistory.Service))
  }

  sd, error := time.Parse(time.RFC3339, roomHistory.StartDate)
  if error != nil {
    return nil, errors.New(fmt.Sprintf("error parsing time %s", roomHistory.StartDate))
  }

  ed, error := time.Parse(time.RFC3339, roomHistory.EndDate)
  if error != nil {
    return nil, errors.New(fmt.Sprintf("error parsing time %s", roomHistory.EndDate))
  }

  newRoomHistory := entities.RoomHistory {
    StartDate: sd,
    EndDate: &ed,
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


func (s *service) SetEndDate(id uint, roomHistory *entities.SetEndDateRoomHistory) (*entities.RoomHistory, error) {
  ed, error := time.Parse(time.RFC3339, roomHistory.EndDate)
  if error != nil {
    return nil, errors.New(fmt.Sprintf("error parsing time %s", roomHistory.EndDate))
  }

  newRoomHistory := entities.RoomHistory {
    EndDate: &ed,
  }

	return s.repository.Update(id, &newRoomHistory)
}
