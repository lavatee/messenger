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
