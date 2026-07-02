package sessions

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var ErrNotFound = errors.New("session not found")

type Session struct {
	ID         string
	UserID     string
	TokenHash  string
	ExpiresAt  time.Time
	LastUsedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CreateParams struct {
	ID         string
	UserID     string
	TokenHash  string
	ExpiresAt  time.Time
	LastUsedAt *time.Time
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, params CreateParams) (Session, error) {
	var lastUsedAt any
	if params.LastUsedAt != nil {
		lastUsedAt = *params.LastUsedAt
	}

	session, err := scanSession(r.db.QueryRowContext(ctx, `
		INSERT INTO sessions (
			id,
			user_id,
			token_hash,
			expires_at,
			last_used_at
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id,
			user_id,
			token_hash,
			expires_at,
			last_used_at,
			created_at,
			updated_at
	`, params.ID, params.UserID, params.TokenHash, params.ExpiresAt, lastUsedAt))
	if err != nil {
		return Session{}, fmt.Errorf("create session: %w", err)
	}

	return session, nil
}

func (r *Repository) FindByTokenHash(ctx context.Context, tokenHash string) (Session, error) {
	session, err := scanSession(r.db.QueryRowContext(ctx, `
		SELECT
			id,
			user_id,
			token_hash,
			expires_at,
			last_used_at,
			created_at,
			updated_at
		FROM sessions
		WHERE token_hash = $1
	`, tokenHash))
	if err != nil {
		return Session{}, fmt.Errorf("find session by token hash: %w", err)
	}

	return session, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM sessions WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete session: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("delete session rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("delete session: %w", ErrNotFound)
	}

	return nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanSession(row scanner) (Session, error) {
	var session Session
	var lastUsedAt sql.NullTime

	if err := row.Scan(
		&session.ID,
		&session.UserID,
		&session.TokenHash,
		&session.ExpiresAt,
		&lastUsedAt,
		&session.CreatedAt,
		&session.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Session{}, ErrNotFound
		}

		return Session{}, err
	}

	if lastUsedAt.Valid {
		session.LastUsedAt = &lastUsedAt.Time
	}

	return session, nil
}
