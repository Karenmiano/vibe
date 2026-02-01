package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateRoomData struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	Description string `json:"description" validate:"required,min=2"`
	Private bool `json:"private"`
}