package user

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/Karenmiano/vibe/internal/models"
)

var( 
	ErrUsernameTaken = errors.New("username is already taken")
	ErrEmailTaken = errors.New("email is already taken")
	ErrInvalidCredentials = errors.New("invalid username/email or password")
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	RegisterUser(ctx context.Context, fullName string, email string, username string, password string) error
	Authenticate(ctx context.Context, identifier string, password string) (*models.User, error) // identifier can be username or email
	GetUserByID(ctx context.Context, userId uuid.UUID) (*models.User, error)
}