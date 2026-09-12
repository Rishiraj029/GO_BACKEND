package repository

import (
	"database/sql"
	"errors"

	"devconnect/internal/apperror"
	"devconnect/internal/model"

	"github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user model.User) (int, error) {
	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int

	err := r.db.QueryRow(
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
	).Scan(&id)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return 0, apperror.ErrEmailAlreadyExists
			}
		}
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) FindUserByEmail(email string) (model.User, error) {
	query := `
	SELECT id, name, email, password_hash
	FROM users
	WHERE email = $1
	`

	var user model.User

	err := r.db.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, apperror.ErrUserNotFound
		}
		return model.User{}, err
	}

	return user, nil
}
