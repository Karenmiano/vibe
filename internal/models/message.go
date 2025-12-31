package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id       int64     `json:"id"`
	RoomId   uuid.UUID `json:"roomId"`
	SenderId uuid.UUID `json:"senderId"`
	Content  string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}