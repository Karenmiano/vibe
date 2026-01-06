package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"github.com/Karenmiano/vibe/internal/user"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) RegisterUser(ctx context.Context, username string, password string) (uuid.UUID, error) {
	userId := uuid.New()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return uuid.Nil, err
	}

	query := `INSERT INTO users (id, username, password) VALUES ($1, $2, $3)`

	_, err = r.db.Exec(ctx, query, userId, username, passwordHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation{
			return uuid.Nil, user.ErrUserExists
		}
		
		return uuid.Nil, err
	}

	return userId, nil
}

func (r *UserRepository) Authenticate(ctx context.Context, username string, password string) {}
