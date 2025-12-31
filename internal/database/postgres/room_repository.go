package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Karenmiano/vibe/internal/models"
)

type RoomRepository struct {
	db *pgxpool.Pool
}

func NewRoomRepository(db *pgxpool.Pool) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (r *RoomRepository) CreateRoom(ctx context.Context, newRoomData models.CreateRoomData) error {
	roomId := uuid.New()
	query := `INSERT INTO rooms (id, name, description, private) VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(ctx, query, roomId, newRoomData.Name, newRoomData.Description, newRoomData.Private)
	return err
}