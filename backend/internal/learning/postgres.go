package learning

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

var _ Store = (*PostgresStore)(nil)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) Tracks(ctx context.Context) (TracksResponse, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT
			t.id,
			t.title,
			t.description,
			COALESCE(
				json_agg(l.id ORDER BY l.order_index) FILTER (WHERE l.id IS NOT NULL),
				'[]'::json
			)::text AS level_ids
		FROM tracks t
		LEFT JOIN levels l ON l.track_id = t.id
		GROUP BY t.id, t.title, t.description
		ORDER BY t.id
	`)
	if err != nil {
		return TracksResponse{}, fmt.Errorf("query tracks: %w", err)
	}
	defer rows.Close()

	tracks := make([]Track, 0)
	for rows.Next() {
		var track Track
		var levelIDsJSON string

		if err := rows.Scan(&track.ID, &track.Title, &track.Description, &levelIDsJSON); err != nil {
			return TracksResponse{}, fmt.Errorf("scan track: %w", err)
		}

		track.LevelIDs, err = decodeStringSlice(levelIDsJSON)
		if err != nil {
			return TracksResponse{}, fmt.Errorf("decode track level IDs: %w", err)
		}

		tracks = append(tracks, track)
	}
	if err := rows.Err(); err != nil {
		return TracksResponse{}, fmt.Errorf("iterate tracks: %w", err)
	}

	return TracksResponse{Tracks: tracks}, nil
}

func (s *PostgresStore) Levels(ctx context.Context) (LevelsResponse, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT
			id,
			title,
			status,
			COALESCE((unlock_rules->>'completedBlocks')::int, 0) AS completed_blocks,
			COALESCE((unlock_rules->>'totalBlocks')::int, 0) AS total_blocks,
			required_xp
		FROM levels
		ORDER BY order_index
	`)
	if err != nil {
		return LevelsResponse{}, fmt.Errorf("query levels: %w", err)
	}
	defer rows.Close()

	levels := make([]CareerLevel, 0)
	for rows.Next() {
		var level CareerLevel
		if err := rows.Scan(
			&level.ID,
			&level.Title,
			&level.Status,
			&level.CompletedBlocks,
			&level.TotalBlocks,
			&level.RequiredXP,
		); err != nil {
			return LevelsResponse{}, fmt.Errorf("scan level: %w", err)
		}

		levels = append(levels, level)
	}
	if err := rows.Err(); err != nil {
		return LevelsResponse{}, fmt.Errorf("iterate levels: %w", err)
	}

	return LevelsResponse{Levels: levels}, nil
}

func (s *PostgresStore) Lesson(ctx context.Context, id string) (Lesson, error) {
	var lesson Lesson
	var orderIndex int
	var tagsJSON string
	var nextChallengeID sql.NullString

	err := s.db.QueryRowContext(ctx, `
		SELECT
			l.id,
			l.title,
			l.order_index,
			l.xp_reward,
			COALESCE(array_to_json(l.tags), '[]'::json)::text AS tags,
			l.question,
			l.explanation,
			lv.checkpoint_challenge_id
		FROM lessons l
		JOIN levels lv ON lv.id = l.level_id
		WHERE l.id = $1
	`, id).Scan(
		&lesson.ID,
		&lesson.Title,
		&orderIndex,
		&lesson.XPReward,
		&tagsJSON,
		&lesson.Question.Prompt,
		&lesson.Question.Explanation,
		&nextChallengeID,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return Lesson{}, ErrLessonNotFound
	}
	if err != nil {
		return Lesson{}, fmt.Errorf("query lesson %q: %w", id, err)
	}

	tags, err := decodeStringSlice(tagsJSON)
	if err != nil {
		return Lesson{}, fmt.Errorf("decode lesson tags: %w", err)
	}

	lesson.Tags = tags
	lesson.Step = fmt.Sprintf("Урок %d", orderIndex)
	lesson.Question.Feedback = "Верно"
	if nextChallengeID.Valid {
		lesson.NextChallengeID = nextChallengeID.String
	}

	if fallback, ok := mockLessons[id]; ok {
		lesson.Step = fallback.Step
		lesson.ProgressPercent = fallback.ProgressPercent
		lesson.Question.Feedback = fallback.Question.Feedback
		lesson.Sidebar = fallback.Sidebar
		if lesson.NextChallengeID == "" {
			lesson.NextChallengeID = fallback.NextChallengeID
		}
	}

	options, correctID, err := s.lessonOptions(ctx, id)
	if err != nil {
		return Lesson{}, err
	}

	lesson.Question.Options = options
	lesson.Question.CorrectID = correctID

	return lesson, nil
}

func (s *PostgresStore) Progress(ctx context.Context, userID string) (Progress, error) {
	var progress Progress
	var level string
	var nextLevelXP int
	var currentLessonID sql.NullString

	err := s.db.QueryRowContext(ctx, `
		SELECT
			up.user_id,
			COALESCE(l.title, ''),
			up.xp,
			COALESCE(l.required_xp, 0),
			up.streak_days,
			up.current_lesson_id,
			up.completed_lessons,
			up.completed_checkpoints
		FROM user_progress up
		LEFT JOIN levels l ON l.id = up.current_level_id
		WHERE up.user_id = $1
	`, userID).Scan(
		&progress.UserID,
		&level,
		&progress.XP,
		&nextLevelXP,
		&progress.StreakDays,
		&currentLessonID,
		&progress.CompletedLessons,
		&progress.CompletedCheckpoints,
	)
	if err != nil {
		return Progress{}, fmt.Errorf("query progress for user %q: %w", userID, err)
	}

	progress.Level = level
	progress.NextLevelXP = nextLevelXP
	if currentLessonID.Valid {
		progress.CurrentLessonID = currentLessonID.String
	}

	progress.PracticeTasks = mockProgress.PracticeTasks
	progress.TotalCheckpoints = mockProgress.TotalCheckpoints
	progress.DailyGoal = mockProgress.DailyGoal
	progress.WeeklyXP = append([]DailyXP(nil), mockProgress.WeeklyXP...)

	return progress, nil
}

func (s *PostgresStore) Profile(ctx context.Context, userID string) (Profile, error) {
	var profile Profile
	var joinedAt time.Time
	var level string
	var nextLevelXP int

	err := s.db.QueryRowContext(ctx, `
		SELECT
			u.id,
			u.display_name,
			u.handle,
			u.created_at,
			COALESCE(l.title, ''),
			COALESCE(up.xp, 0),
			COALESCE(l.required_xp, 0),
			COALESCE(up.streak_days, 0)
		FROM users u
		LEFT JOIN user_progress up ON up.user_id = u.id
		LEFT JOIN levels l ON l.id = up.current_level_id
		WHERE u.id = $1
	`, userID).Scan(
		&profile.User.ID,
		&profile.User.Name,
		&profile.User.Handle,
		&joinedAt,
		&level,
		&profile.XP,
		&nextLevelXP,
		&profile.StreakDays,
	)
	if err != nil {
		return Profile{}, fmt.Errorf("query profile for user %q: %w", userID, err)
	}

	profile.User.Initials = initials(profile.User.Name)
	profile.User.JoinedAt = formatJoinedAt(joinedAt)
	profile.Level = level
	profile.NextLevelXP = nextLevelXP
	profile.Skills, err = s.profileSkills(ctx, userID)
	if err != nil {
		return Profile{}, err
	}

	profile.Readiness = mockProfile.Readiness
	profile.Recommendation = mockProfile.Recommendation

	return profile, nil
}

func (s *PostgresStore) lessonOptions(ctx context.Context, lessonID string) ([]LessonOption, string, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT
			id,
			label,
			body,
			is_correct
		FROM lesson_options
		WHERE lesson_id = $1
		ORDER BY order_index
	`, lessonID)
	if err != nil {
		return nil, "", fmt.Errorf("query lesson options for %q: %w", lessonID, err)
	}
	defer rows.Close()

	options := make([]LessonOption, 0)
	correctID := ""
	for rows.Next() {
		var option LessonOption
		var isCorrect bool

		if err := rows.Scan(&option.ID, &option.Label, &option.Text, &isCorrect); err != nil {
			return nil, "", fmt.Errorf("scan lesson option: %w", err)
		}

		if isCorrect {
			correctID = option.ID
		}
		options = append(options, option)
	}
	if err := rows.Err(); err != nil {
		return nil, "", fmt.Errorf("iterate lesson options: %w", err)
	}

	return options, correctID, nil
}

func (s *PostgresStore) profileSkills(ctx context.Context, userID string) ([]ProfileSkill, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT
			skill,
			score
		FROM skill_progress
		WHERE user_id = $1
		ORDER BY
			CASE skill
				WHEN 'Go' THEN 0
				WHEN 'REST API' THEN 1
				WHEN 'PostgreSQL' THEN 2
				WHEN 'Redis' THEN 3
				WHEN 'Docker' THEN 4
				WHEN 'Kafka' THEN 5
				WHEN 'Kubernetes' THEN 6
				WHEN 'Prometheus' THEN 7
				ELSE 100
			END,
			skill
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("query profile skills for user %q: %w", userID, err)
	}
	defer rows.Close()

	skills := make([]ProfileSkill, 0)
	for rows.Next() {
		var skill ProfileSkill
		if err := rows.Scan(&skill.Name, &skill.Score); err != nil {
			return nil, fmt.Errorf("scan profile skill: %w", err)
		}

		skills = append(skills, skill)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate profile skills: %w", err)
	}

	return skills, nil
}

func decodeStringSlice(value string) ([]string, error) {
	var items []string
	if err := json.Unmarshal([]byte(value), &items); err != nil {
		return nil, err
	}

	return items, nil
}

func initials(name string) string {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return ""
	}

	var builder strings.Builder
	for _, part := range parts {
		builder.WriteString(strings.ToUpper(string([]rune(part)[0])))
	}

	return builder.String()
}

func formatJoinedAt(value time.Time) string {
	months := [...]string{
		"янв.",
		"февр.",
		"мар.",
		"апр.",
		"мая",
		"июн.",
		"июл.",
		"авг.",
		"сент.",
		"окт.",
		"нояб.",
		"дек.",
	}

	month := months[int(value.Month())-1]

	return fmt.Sprintf("%d %s %d", value.Day(), month, value.Year())
}
