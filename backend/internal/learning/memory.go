package learning

import "context"

type MemoryStore struct{}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Tracks(ctx context.Context) (TracksResponse, error) {
	return TracksResponse{Tracks: mockTracks}, nil
}

func (s *MemoryStore) Levels(ctx context.Context) (LevelsResponse, error) {
	return LevelsResponse{Levels: mockLevels}, nil
}

func (s *MemoryStore) Lesson(ctx context.Context, id string) (Lesson, error) {
	lesson, ok := mockLessons[id]
	if !ok {
		return Lesson{}, ErrLessonNotFound
	}

	return lesson, nil
}

func (s *MemoryStore) Progress(ctx context.Context, userID string) (Progress, error) {
	return mockProgress, nil
}

func (s *MemoryStore) Profile(ctx context.Context, userID string) (Profile, error) {
	return mockProfile, nil
}

var mockTracks = []Track{
	{
		ID:          "go-backend",
		Title:       "Go Backend",
		Description: "Карьерный путь Go backend-разработчика от основ до production reliability.",
		LevelIDs:    []string{"trainee-backend", "junior-backend", "middle-backend", "senior-backend"},
	},
}

var mockLevels = []CareerLevel{
	{
		ID:              "trainee-backend",
		Title:           "Стажер Backend",
		Status:          "completed",
		CompletedBlocks: 2,
		TotalBlocks:     2,
		RequiredXP:      600,
	},
	{
		ID:              "junior-backend",
		Title:           "Junior Backend",
		Status:          "active",
		CompletedBlocks: 2,
		TotalBlocks:     6,
		RequiredXP:      2000,
	},
	{
		ID:              "middle-backend",
		Title:           "Middle Backend",
		Status:          "locked",
		CompletedBlocks: 0,
		TotalBlocks:     6,
		RequiredXP:      5000,
	},
	{
		ID:              "senior-backend",
		Title:           "Senior Backend",
		Status:          "locked",
		CompletedBlocks: 0,
		TotalBlocks:     6,
		RequiredXP:      9000,
	},
}

var mockLessons = map[string]Lesson{
	"go-errors-tests": {
		ID:              "go-errors-tests",
		Title:           "Go: ошибки и тесты",
		Step:            "Урок 4 из 12",
		ProgressPercent: 67,
		XPReward:        40,
		Tags:            []string{"Go", "Testing", "Errors"},
		Question: LessonQuestion{
			Prompt:    "Почему в Go ошибку обычно возвращают явно?",
			CorrectID: "caller-controls-error",
			Feedback:  "Верно",
			Explanation: "В Go ошибки обрабатываются явно, чтобы вызывающая сторона понимала контекст " +
				"и могла выбрать стратегию обработки сбоя.",
			Options: []LessonOption{
				{
					ID:    "caller-controls-error",
					Label: "A",
					Text:  "Чтобы caller сам решил, как обработать сбой",
				},
				{
					ID:    "hide-problem",
					Label: "B",
					Text:  "Чтобы скрыть проблему от пользователя",
				},
				{
					ID:    "replace-tests",
					Label: "C",
					Text:  "Чтобы заменить unit tests",
				},
				{
					ID:    "speed-runtime",
					Label: "D",
					Text:  "Чтобы ускорить runtime",
				},
			},
		},
		Sidebar: LessonSidebar{
			ConceptTitle: "Концепция",
			ConceptText: []string{
				"В Go функции, которые могут завершиться ошибкой, возвращают значение типа error.",
				"Такой подход держит обработку сбоя рядом с местом, где есть достаточно контекста.",
			},
			CodeSnippet: "func GetUser(id int) (User, error) {\n" +
				"    user, err := repo.Find(id)\n" +
				"    if err != nil {\n" +
				"        return User{}, fmt.Errorf(\"find user: %w\", err)\n" +
				"    }\n\n" +
				"    return user, nil\n" +
				"}",
			PracticeNotes: []string{
				"Ошибки не игнорируются.",
				"Ошибки оборачиваются с контекстом.",
				"На границе API ошибки маппятся в корректные HTTP-статусы.",
			},
			RelatedTopics: []string{"Errors", "Best Practices", "Error Wrapping", "Go Conventions"},
		},
		NextChallengeID: "retry-context",
	},
}

var mockProgress = Progress{
	UserID:               "demo-user",
	Level:                "Junior Backend",
	XP:                   1420,
	NextLevelXP:          2000,
	StreakDays:           7,
	CurrentLessonID:      "go-errors-tests",
	CompletedLessons:     24,
	PracticeTasks:        18,
	CompletedCheckpoints: 2,
	TotalCheckpoints:     6,
	DailyGoal: DailyGoal{
		EarnedXP:       150,
		TargetXP:       200,
		CompletedTasks: 2,
		TargetTasks:    3,
	},
	WeeklyXP: []DailyXP{
		{Day: "Пн", XP: 120},
		{Day: "Вт", XP: 180},
		{Day: "Ср", XP: 260},
		{Day: "Чт", XP: 320},
		{Day: "Пт", XP: 220},
		{Day: "Сб", XP: 200},
		{Day: "Вс", XP: 120},
	},
}

var mockProfile = Profile{
	User: ProfileUser{
		ID:       "demo-user",
		Name:     "Alex Kim",
		Handle:   "@alexkim",
		Initials: "AK",
		JoinedAt: "15 апр. 2024",
	},
	Level:       "Junior Backend",
	XP:          1420,
	NextLevelXP: 2000,
	StreakDays:  7,
	Skills: []ProfileSkill{
		{Name: "Go", Score: 84},
		{Name: "REST API", Score: 78},
		{Name: "PostgreSQL", Score: 72},
		{Name: "Redis", Score: 68},
		{Name: "Docker", Score: 75},
		{Name: "Kafka", Score: 60},
		{Name: "Kubernetes", Score: 58},
		{Name: "Prometheus", Score: 50},
	},
	Readiness: Readiness{
		TargetLevel:     "Middle Backend",
		ProgressPercent: 67,
		Items: []ReadinessItem{
			{Title: "Завершите 30 уроков", Current: 24, Target: 30},
			{Title: "Пройдите 3 checkpoint", Current: 2, Target: 3},
			{Title: "Решите 20 практических задач", Current: 18, Target: 20},
			{Title: "Создайте 2 проекта", Current: 1, Target: 2},
			{Title: "Наберите 2 000 XP", Current: 1420, Target: 2000},
		},
	},
	Recommendation: Recommendation{
		Title:       "Рекомендация для вас",
		Description: "Попробуйте практику по теме Concurrency.",
		Path:        "/challenge/retry-context",
	},
}
