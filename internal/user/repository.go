package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrUserExists = errors.New("Username already exists")

type UserRepository interface {
	RegisterUser(ctx context.Context, username string, password string) (uuid.UUID, error)
	Authenticate(ctx context.Context, username string, password string)
}