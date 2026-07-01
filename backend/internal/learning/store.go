package learning

import (
	"context"
	"errors"
)

var ErrLessonNotFound = errors.New("lesson not found")

type Store interface {
	Tracks(ctx context.Context) (TracksResponse, error)
	Levels(ctx context.Context) (LevelsResponse, error)
	Lesson(ctx context.Context, id string) (Lesson, error)
	Progress(ctx context.Context, userID string) (Progress, error)
	Profile(ctx context.Context, userID string) (Profile, error)
}
