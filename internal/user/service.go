package user

import (
	"context"

	"github.com/google/uuid"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, username string, password string) (uuid.UUID, error) {
	return s.repo.RegisterUser(ctx, username, password)
}

func (s *UserService) LoginUser(ctx context.Context, username string, password string) (uuid.UUID, error) {
	return s.repo.Authenticate(ctx, username, password)
}