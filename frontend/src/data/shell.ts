export type CareerLevel = {
  title: string
  status: 'completed' | 'active' | 'locked'
}

export type ShellSummary = {
  userName: string
  level: string
  xp: number
  nextLevelXp: number
  streakDays: number
  dailyGoal: {
    earnedXp: number
    targetXp: number
    completedTasks: number
    targetTasks: number
  }
  currentLesson: {
    id: string
    title: string
    step: string
    progressPercent: number
    xpReward: number
  }
  weakTopics: {
    title: string
    confidencePercent: number
    action: string
  }[]
  checkpoint: {
    title: string
    description: string
    tasksTotal: number
    rewardXp: number
    difficulty: string
  }
  recentActivity: {
    title: string
    timestamp: string
    xp: number
  }[]
  weeklyActivity: {
    day: string
    xp: number
  }[]
  levels: CareerLevel[]
}

export const shellSummary: ShellSummary = {
  userName: 'Alex Kim',
  level: 'Junior Backend',
  xp: 1420,
  nextLevelXp: 2000,
  streakDays: 7,
  dailyGoal: {
    earnedXp: 150,
    targetXp: 200,
    completedTasks: 2,
    targetTasks: 3,
  },
  currentLesson: {
    id: 'go-errors-tests',
    title: 'Go: ошибки и тесты',
    step: 'Урок 4 из 12',
    progressPercent: 67,
    xpReward: 40,
  },
  weakTopics: [
    { title: 'PostgreSQL', confidencePercent: 45, action: 'Тренировать' },
    { title: 'context.Context', confidencePercent: 50, action: 'Повторить' },
    { title: 'Docker', confidencePercent: 40, action: 'Тренировать' },
  ],
  checkpoint: {
    title: 'Checkpoint: REST API',
    description: 'Реализуйте RESTful API с validation, middleware и тестами.',
    tasksTotal: 7,
    rewardXp: 300,
    difficulty: 'Средний',
  },
  recentActivity: [
    {
      title: 'Завершён урок «Go: структуры и методы»',
      timestamp: 'Сегодня, 10:24',
      xp: 60,
    },
    {
      title: 'Решён практический блок «HTTP handlers»',
      timestamp: 'Сегодня, 09:15',
      xp: 40,
    },
    {
      title: 'Пройден checkpoint «Go Basics»',
      timestamp: 'Вчера, 18:42',
      xp: 200,
    },
  ],
  weeklyActivity: [
    { day: 'Пн', xp: 80 },
    { day: 'Вт', xp: 120 },
    { day: 'Ср', xp: 160 },
    { day: 'Чт', xp: 220 },
    { day: 'Пт', xp: 150 },
    { day: 'Сб', xp: 180 },
    { day: 'Вс', xp: 110 },
  ],
  levels: [
    { title: 'Стажер Backend', status: 'completed' },
    { title: 'Junior Backend', status: 'active' },
    { title: 'Middle Backend', status: 'locked' },
    { title: 'Senior Backend', status: 'locked' },
  ],
}
