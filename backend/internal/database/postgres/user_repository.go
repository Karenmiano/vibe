package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
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

func (r *UserRepository) RegisterUser(ctx context.Context, fullName string, email string, username string, password string) error {
	userId := uuid.New()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (id, full_name, email, username, password) VALUES ($1, $2, $3, $4, $5)`

	_, err = r.db.Exec(ctx, query, userId, fullName, email, username, passwordHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation{
			switch pgErr.ConstraintName {
				case "users_username_key":
					return user.ErrUsernameTaken
				case "users_email_key":
					return user.ErrEmailTaken
			}
		}

		return err
	}

	return nil
}

func (r *UserRepository) Authenticate(ctx context.Context, identifier string, password string) (uuid.UUID, error) { // identifier can be username or email
	var userId uuid.UUID
	var passwordHash []byte

	query := `SELECT id, password FROM users WHERE username = $1 OR email = $1`
	err := r.db.QueryRow(ctx, query, identifier).Scan(&userId, &passwordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.Nil, user.ErrInvalidCredentials
		}
		return uuid.Nil, err
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
	if err != nil {
		return uuid.Nil, user.ErrInvalidCredentials
	}

	return userId, nil
}
