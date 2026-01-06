package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var( 
	ErrUserExists = errors.New("Username already exists")
	ErrInvalidCredentials = errors.New("Invalid username or password")
)

type UserRepository interface {
	RegisterUser(ctx context.Context, username string, password string) (uuid.UUID, error)
	Authenticate(ctx context.Context, username string, password string) (uuid.UUID, error)
}