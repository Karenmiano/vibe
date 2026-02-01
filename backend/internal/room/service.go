package room

import (
	"context"
	"github.com/Karenmiano/vibe/internal/models"
)

type RoomService struct {
	repo RoomRepository
}

func NewRoomService(repo RoomRepository) *RoomService {
	return &RoomService{
		repo: repo,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, newRoomData models.CreateRoomData) error {
	// add the currently authenticated user as first member(transaction?)
	return s.repo.CreateRoom(ctx, newRoomData)
}