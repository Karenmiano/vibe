package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var( 
	ErrUsernameTaken = errors.New("username is already taken")
	ErrEmailTaken = errors.New("email is already taken")
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type UserRepository interface {
	RegisterUser(ctx context.Context, fullName string, email string, username string, password string) error
	Authenticate(ctx context.Context, username string, password string) (uuid.UUID, error)
}