package user

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, username string, password string) error
	Authenticate(ctx context.Context, username string, password string) (uuid.UUID, error)
}