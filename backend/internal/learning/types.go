package learning

type TracksResponse struct {
	Tracks []Track `json:"tracks"`
}

type Track struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	LevelIDs    []string `json:"levelIds"`
}

type LevelsResponse struct {
	Levels []CareerLevel `json:"levels"`
}

type CareerLevel struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Status          string `json:"status"`
	CompletedBlocks int    `json:"completedBlocks"`
	TotalBlocks     int    `json:"totalBlocks"`
	RequiredXP      int    `json:"requiredXp"`
}

type Lesson struct {
	ID              string         `json:"id"`
	Title           string         `json:"title"`
	Step            string         `json:"step"`
	ProgressPercent int            `json:"progressPercent"`
	XPReward        int            `json:"xpReward"`
	Tags            []string       `json:"tags"`
	Question        LessonQuestion `json:"question"`
	Sidebar         LessonSidebar  `json:"sidebar"`
	NextChallengeID string         `json:"nextChallengeId"`
}

type LessonQuestion struct {
	Prompt      string         `json:"prompt"`
	Options     []LessonOption `json:"options"`
	CorrectID   string         `json:"correctId"`
	Feedback    string         `json:"feedback"`
	Explanation string         `json:"explanation"`
}

type LessonOption struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Text  string `json:"text"`
}

type LessonSidebar struct {
	ConceptTitle  string   `json:"conceptTitle"`
	ConceptText   []string `json:"conceptText"`
	CodeSnippet   string   `json:"codeSnippet"`
	PracticeNotes []string `json:"practiceNotes"`
	RelatedTopics []string `json:"relatedTopics"`
}

type Progress struct {
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
	DailyGoal            DailyGoal `json:"dailyGoal"`
	WeeklyXP             []DailyXP `json:"weeklyXp"`
}

type DailyGoal struct {
	EarnedXP       int `json:"earnedXp"`
	TargetXP       int `json:"targetXp"`
	CompletedTasks int `json:"completedTasks"`
	TargetTasks    int `json:"targetTasks"`
}

type DailyXP struct {
	Day string `json:"day"`
	XP  int    `json:"xp"`
}

type Profile struct {
	User           ProfileUser    `json:"user"`
	Level          string         `json:"level"`
	XP             int            `json:"xp"`
	NextLevelXP    int            `json:"nextLevelXp"`
	StreakDays     int            `json:"streakDays"`
	Skills         []ProfileSkill `json:"skills"`
	Readiness      Readiness      `json:"readiness"`
	Recommendation Recommendation `json:"recommendation"`
}

type ProfileUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Handle   string `json:"handle"`
	Initials string `json:"initials"`
	JoinedAt string `json:"joinedAt"`
}

type ProfileSkill struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type Readiness struct {
	TargetLevel     string          `json:"targetLevel"`
	ProgressPercent int             `json:"progressPercent"`
	Items           []ReadinessItem `json:"items"`
}

type ReadinessItem struct {
	Title      string `json:"title"`
	Current    int    `json:"current"`
	Target     int    `json:"target"`
	IsComplete bool   `json:"isComplete"`
}

type Recommendation struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Path        string `json:"path"`
}
