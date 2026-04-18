package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var( 
	ErrUserExists = errors.New("username already exists")
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type UserRepository interface {
	RegisterUser(ctx context.Context, username string, password string) error
	Authenticate(ctx context.Context, username string, password string) (uuid.UUID, error)
}