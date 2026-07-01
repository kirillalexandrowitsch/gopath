package httpserver

import "net/http"

type tracksResponse struct {
	Tracks []track `json:"tracks"`
}

type track struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	LevelIDs    []string `json:"levelIds"`
}

type levelsResponse struct {
	Levels []careerLevel `json:"levels"`
}

type careerLevel struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Status          string `json:"status"`
	CompletedBlocks int    `json:"completedBlocks"`
	TotalBlocks     int    `json:"totalBlocks"`
	RequiredXP      int    `json:"requiredXp"`
}

type lesson struct {
	ID              string         `json:"id"`
	Title           string         `json:"title"`
	Step            string         `json:"step"`
	ProgressPercent int            `json:"progressPercent"`
	XPReward        int            `json:"xpReward"`
	Tags            []string       `json:"tags"`
	Question        lessonQuestion `json:"question"`
	Sidebar         lessonSidebar  `json:"sidebar"`
	NextChallengeID string         `json:"nextChallengeId"`
}

type lessonQuestion struct {
	Prompt      string         `json:"prompt"`
	Options     []lessonOption `json:"options"`
	CorrectID   string         `json:"correctId"`
	Feedback    string         `json:"feedback"`
	Explanation string         `json:"explanation"`
}

type lessonOption struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Text  string `json:"text"`
}

type lessonSidebar struct {
	ConceptTitle  string   `json:"conceptTitle"`
	ConceptText   []string `json:"conceptText"`
	CodeSnippet   string   `json:"codeSnippet"`
	PracticeNotes []string `json:"practiceNotes"`
	RelatedTopics []string `json:"relatedTopics"`
}

type progressResponse struct {
	UserID               string    `json:"userId"`
	Level                string    `json:"level"`
	XP                   int       `json:"xp"`
	NextLevelXP          int       `json:"nextLevelXp"`
	StreakDays           int       `json:"streakDays"`
	CurrentLessonID      string    `json:"currentLessonId"`
	CompletedLessons     int       `json:"completedLessons"`
	PracticeTasks        int       `json:"practiceTasks"`
	CompletedCheckpoints int       `json:"completedCheckpoints"`
	TotalCheckpoints     int       `json:"totalCheckpoints"`
	DailyGoal            dailyGoal `json:"dailyGoal"`
	WeeklyXP             []dailyXP `json:"weeklyXp"`
}

type dailyGoal struct {
	EarnedXP       int `json:"earnedXp"`
	TargetXP       int `json:"targetXp"`
	CompletedTasks int `json:"completedTasks"`
	TargetTasks    int `json:"targetTasks"`
}

type dailyXP struct {
	Day string `json:"day"`
	XP  int    `json:"xp"`
}

type profileResponse struct {
	User           profileUser    `json:"user"`
	Level          string         `json:"level"`
	XP             int            `json:"xp"`
	NextLevelXP    int            `json:"nextLevelXp"`
	StreakDays     int            `json:"streakDays"`
	Skills         []profileSkill `json:"skills"`
	Readiness      readiness      `json:"readiness"`
	Recommendation recommendation `json:"recommendation"`
}

type profileUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Handle   string `json:"handle"`
	Initials string `json:"initials"`
	JoinedAt string `json:"joinedAt"`
}

type profileSkill struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type readiness struct {
	TargetLevel     string          `json:"targetLevel"`
	ProgressPercent int             `json:"progressPercent"`
	Items           []readinessItem `json:"items"`
}

type readinessItem struct {
	Title      string `json:"title"`
	Current    int    `json:"current"`
	Target     int    `json:"target"`
	IsComplete bool   `json:"isComplete"`
}

type recommendation struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Path        string `json:"path"`
}

var mockTracks = []track{
	{
		ID:          "go-backend",
		Title:       "Go Backend",
		Description: "Карьерный путь Go backend-разработчика от основ до production reliability.",
		LevelIDs:    []string{"trainee-backend", "junior-backend", "middle-backend", "senior-backend"},
	},
}

var mockLevels = []careerLevel{
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

var mockLessons = map[string]lesson{
	"go-errors-tests": {
		ID:              "go-errors-tests",
		Title:           "Go: ошибки и тесты",
		Step:            "Урок 4 из 12",
		ProgressPercent: 67,
		XPReward:        40,
		Tags:            []string{"Go", "Testing", "Errors"},
		Question: lessonQuestion{
			Prompt:    "Почему в Go ошибку обычно возвращают явно?",
			CorrectID: "caller-controls-error",
			Feedback:  "Верно",
			Explanation: "В Go ошибки обрабатываются явно, чтобы вызывающая сторона понимала контекст " +
				"и могла выбрать стратегию обработки сбоя.",
			Options: []lessonOption{
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
		Sidebar: lessonSidebar{
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

var mockProgress = progressResponse{
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
	DailyGoal: dailyGoal{
		EarnedXP:       150,
		TargetXP:       200,
		CompletedTasks: 2,
		TargetTasks:    3,
	},
	WeeklyXP: []dailyXP{
		{Day: "Пн", XP: 120},
		{Day: "Вт", XP: 180},
		{Day: "Ср", XP: 260},
		{Day: "Чт", XP: 320},
		{Day: "Пт", XP: 220},
		{Day: "Сб", XP: 200},
		{Day: "Вс", XP: 120},
	},
}

var mockProfile = profileResponse{
	User: profileUser{
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
	Skills: []profileSkill{
		{Name: "Go", Score: 84},
		{Name: "REST API", Score: 78},
		{Name: "PostgreSQL", Score: 72},
		{Name: "Redis", Score: 68},
		{Name: "Docker", Score: 75},
		{Name: "Kafka", Score: 60},
		{Name: "Kubernetes", Score: 58},
		{Name: "Prometheus", Score: 50},
	},
	Readiness: readiness{
		TargetLevel:     "Middle Backend",
		ProgressPercent: 67,
		Items: []readinessItem{
			{Title: "Завершите 30 уроков", Current: 24, Target: 30},
			{Title: "Пройдите 3 checkpoint", Current: 2, Target: 3},
			{Title: "Решите 20 практических задач", Current: 18, Target: 20},
			{Title: "Создайте 2 проекта", Current: 1, Target: 2},
			{Title: "Наберите 2 000 XP", Current: 1420, Target: 2000},
		},
	},
	Recommendation: recommendation{
		Title:       "Рекомендация для вас",
		Description: "Попробуйте практику по теме Concurrency.",
		Path:        "/challenge/retry-context",
	},
}

func tracksHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, tracksResponse{Tracks: mockTracks})
}

func levelsHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, levelsResponse{Levels: mockLevels})
}

func lessonHandler(w http.ResponseWriter, r *http.Request) {
	lessonID := r.PathValue("id")
	lesson, ok := mockLessons[lessonID]
	if !ok {
		writeJSON(w, http.StatusNotFound, errorResponse{Error: "lesson not found"})
		return
	}

	writeJSON(w, http.StatusOK, lesson)
}

func progressHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, mockProgress)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, mockProfile)
}
