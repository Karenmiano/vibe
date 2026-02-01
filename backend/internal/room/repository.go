package room

import (
	"context"

	"github.com/Karenmiano/vibe/internal/models"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, newRoomData models.CreateRoomData) error
}