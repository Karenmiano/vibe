package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
	Username string
	CreatedAt time.Time
}

