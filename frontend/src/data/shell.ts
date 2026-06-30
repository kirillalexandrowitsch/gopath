export type CareerLevel = {
  title: string
  status: 'completed' | 'active' | 'locked'
}

export type ShellSummary = {
  userName: string
  level: string
  xp: number
  streakDays: number
  dailyGoalXp: number
  dailyGoalTargetXp: number
  currentLesson: {
    id: string
    title: string
    step: string
    progressPercent: number
  }
  weakTopics: string[]
  levels: CareerLevel[]
}

export const shellSummary: ShellSummary = {
  userName: 'Alex Kim',
  level: 'Junior Backend',
  xp: 1420,
  streakDays: 7,
  dailyGoalXp: 150,
  dailyGoalTargetXp: 200,
  currentLesson: {
    id: 'go-errors-tests',
    title: 'Go: ошибки и тесты',
    step: 'Урок 4 из 12',
    progressPercent: 67,
  },
  weakTopics: ['PostgreSQL', 'context.Context', 'Docker'],
  levels: [
    { title: 'Стажер Backend', status: 'completed' },
    { title: 'Junior Backend', status: 'active' },
    { title: 'Middle Backend', status: 'locked' },
    { title: 'Senior Backend', status: 'locked' },
  ],
}
