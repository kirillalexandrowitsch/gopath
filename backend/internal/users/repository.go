package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var ErrNotFound = errors.New("user not found")

type Role string

const (
	RoleLearner Role = "learner"
	RoleAdmin   Role = "admin"
)

type User struct {
	ID           string
	Email        string
	Handle       string
	DisplayName  string
	PasswordHash string
	Role         Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreateParams struct {
	ID           string
	Email        string
	Handle       string
	DisplayName  string
	PasswordHash string
	Role         Role
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, params CreateParams) (User, error) {
	role := params.Role
	if role == "" {
		role = RoleLearner
	}

	user, err := scanUser(r.db.QueryRowContext(ctx, `
		INSERT INTO users (
			id,
			email,
			handle,
			display_name,
			password_hash,
			role
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING
			id,
			email,
			handle,
			display_name,
			password_hash,
			role,
			created_at,
			updated_at
	`, params.ID, params.Email, params.Handle, params.DisplayName, params.PasswordHash, role))
	if err != nil {
		return User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (User, error) {
	user, err := scanUser(r.db.QueryRowContext(ctx, `
		SELECT
			id,
			email,
			handle,
			display_name,
			password_hash,
			role,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`, id))
	if err != nil {
		return User{}, fmt.Errorf("find user by id: %w", err)
	}

	return user, nil
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (User, error) {
	user, err := scanUser(r.db.QueryRowContext(ctx, `
		SELECT
			id,
			email,
			handle,
			display_name,
			password_hash,
			role,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`, email))
	if err != nil {
		return User{}, fmt.Errorf("find user by email: %w", err)
	}

	return user, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanUser(row scanner) (User, error) {
	var user User

	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Handle,
		&user.DisplayName,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrNotFound
		}

		return User{}, err
	}

	return user, nil
}
