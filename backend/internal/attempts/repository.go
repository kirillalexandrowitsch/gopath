package attempts

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

const defaultListLimit = 20

var ErrNotFound = errors.New("attempt not found")

type Status string

const (
	StatusQueued        Status = "queued"
	StatusRunning       Status = "running"
	StatusPassed        Status = "passed"
	StatusFailed        Status = "failed"
	StatusTimeout       Status = "timeout"
	StatusCompileError  Status = "compile_error"
	StatusBlocked       Status = "blocked"
	StatusInternalError Status = "internal_error"
)

type Attempt struct {
	ID          string
	UserID      string
	LessonID    *string
	ChallengeID *string
	Answer      string
	IsCorrect   bool
	XPAwarded   int
	Status      Status
	Stdout      string
	Stderr      string
	TestsPassed int
	TestsTotal  int
	DurationMS  int
	CreatedAt   time.Time
}

type CreateParams struct {
	ID          string
	UserID      string
	LessonID    *string
	ChallengeID *string
	Answer      string
	IsCorrect   bool
	XPAwarded   int
	Status      Status
	Stdout      string
	Stderr      string
	TestsPassed int
	TestsTotal  int
	DurationMS  int
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, params CreateParams) (Attempt, error) {
	attempt, err := scanAttempt(r.db.QueryRowContext(ctx, `
		INSERT INTO attempts (
			id,
			user_id,
			lesson_id,
			challenge_id,
			answer,
			is_correct,
			xp_awarded,
			status,
			stdout,
			stderr,
			tests_passed,
			tests_total,
			duration_ms
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING
			id,
			user_id,
			lesson_id,
			challenge_id,
			answer,
			is_correct,
			xp_awarded,
			status,
			stdout,
			stderr,
			tests_passed,
			tests_total,
			duration_ms,
			created_at
	`,
		params.ID,
		params.UserID,
		nullableString(params.LessonID),
		nullableString(params.ChallengeID),
		params.Answer,
		params.IsCorrect,
		params.XPAwarded,
		params.Status,
		params.Stdout,
		params.Stderr,
		params.TestsPassed,
		params.TestsTotal,
		params.DurationMS,
	))
	if err != nil {
		return Attempt{}, fmt.Errorf("create attempt: %w", err)
	}

	return attempt, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (Attempt, error) {
	attempt, err := scanAttempt(r.db.QueryRowContext(ctx, `
		SELECT
			id,
			user_id,
			lesson_id,
			challenge_id,
			answer,
			is_correct,
			xp_awarded,
			status,
			stdout,
			stderr,
			tests_passed,
			tests_total,
			duration_ms,
			created_at
		FROM attempts
		WHERE id = $1
	`, id))
	if err != nil {
		return Attempt{}, fmt.Errorf("find attempt by id: %w", err)
	}

	return attempt, nil
}

func (r *Repository) ListByUserID(ctx context.Context, userID string, limit int) ([]Attempt, error) {
	if limit <= 0 {
		limit = defaultListLimit
	}

	rows, err := r.db.QueryContext(ctx, `
		SELECT
			id,
			user_id,
			lesson_id,
			challenge_id,
			answer,
			is_correct,
			xp_awarded,
			status,
			stdout,
			stderr,
			tests_passed,
			tests_total,
			duration_ms,
			created_at
		FROM attempts
		WHERE user_id = $1
		ORDER BY created_at DESC, id DESC
		LIMIT $2
	`, userID, limit)
	if err != nil {
		return nil, fmt.Errorf("list attempts by user id: %w", err)
	}
	defer rows.Close()

	attempts := make([]Attempt, 0)
	for rows.Next() {
		attempt, err := scanAttempt(rows)
		if err != nil {
			return nil, fmt.Errorf("scan attempt: %w", err)
		}

		attempts = append(attempts, attempt)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate attempts: %w", err)
	}

	return attempts, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAttempt(row scanner) (Attempt, error) {
	var attempt Attempt
	var lessonID sql.NullString
	var challengeID sql.NullString

	if err := row.Scan(
		&attempt.ID,
		&attempt.UserID,
		&lessonID,
		&challengeID,
		&attempt.Answer,
		&attempt.IsCorrect,
		&attempt.XPAwarded,
		&attempt.Status,
		&attempt.Stdout,
		&attempt.Stderr,
		&attempt.TestsPassed,
		&attempt.TestsTotal,
		&attempt.DurationMS,
		&attempt.CreatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Attempt{}, ErrNotFound
		}

		return Attempt{}, err
	}

	if lessonID.Valid {
		attempt.LessonID = &lessonID.String
	}
	if challengeID.Valid {
		attempt.ChallengeID = &challengeID.String
	}

	return attempt, nil
}

func nullableString(value *string) any {
	if value == nil {
		return nil
	}

	return *value
}
