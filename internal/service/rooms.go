package service

import "github.com/lavatee/messenger/internal/repository"

type RoomsService struct {
	repo repository.Rooms
}

func NewRoomsService(repos repository.Rooms) *RoomsService {
	return &RoomsService{
		repo: repos,
	}
}
func (s *RoomsService) JoinRoom(userId int) (int, error) {
	return s.repo.JoinRoom(userId)
}
func (s *RoomsService) LeaveRoom(userId int, roomId int) (int, error) {
	return s.repo.LeaveRoom(userId, roomId)
}
func (s *RoomsService) LeaveMatchMaking(userId int, roomId int) error {
	return s.repo.LeaveMatchMaking(userId, roomId)
}
