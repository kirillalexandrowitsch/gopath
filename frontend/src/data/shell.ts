export type CareerLevel = {
  title: string
  status: 'completed' | 'active' | 'locked'
}

export type LearningBlockStatus = 'completed' | 'active' | 'locked' | 'checkpoint'

export type LearningBlock = {
  id: string
  title: string
  status: LearningBlockStatus
  tags: string[]
  progressPercent?: number
  tasksCompleted?: number
  tasksTotal?: number
}

export type LearningStage = {
  title: string
  subtitle: string
  completedBlocks: number
  totalBlocks: number
  blocks: LearningBlock[]
}

export type ActiveLearningBlock = {
  id: string
  title: string
  description: string
  progressPercent: number
  skills: string[]
  requirements: string[]
  careerOutcomes: string[]
}

export type ChallengeTestStatus = 'passed' | 'failed' | 'locked'

export type ChallengeTest = {
  name: string
  status: ChallengeTestStatus
  durationMs?: number
}

export type ChallengeAttempt = {
  timestamp: string
  summary: string
  xp: number
  status: 'passed' | 'failed'
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
    attemptsRemaining: number
    totalAttempts: number
    question: {
      prompt: string
      options: {
        id: string
        label: string
        text: string
        isSelected: boolean
        isCorrect: boolean
      }[]
      feedbackTitle: string
      feedbackText: string
      explanation: string
    }
    sidebar: {
      conceptTitle: string
      conceptText: string[]
      codeSnippet: string
      practiceNotes: string[]
      usefulNote: string
      relatedTopics: string[]
    }
  }
  currentChallenge: {
    id: string
    title: string
    difficulty: string
    rewardXp: number
    description: string
    requirements: string[]
    signature: string
    types: string
    hint: string
    starterCode: string
    terminalOutput: string[]
    publicTests: ChallengeTest[]
    hiddenTests: ChallengeTest[]
    result: {
      title: string
      description: string
      passedTests: number
      totalTests: number
    }
    sandbox: {
      status: string
      network: string
      memoryLimit: string
      timeout: string
      goVersion: string
    }
    attempts: ChallengeAttempt[]
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
  learningPath: {
    nextLevelProgressPercent: number
    activeBlock: ActiveLearningBlock
    stages: LearningStage[]
  }
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
    attemptsRemaining: 2,
    totalAttempts: 3,
    question: {
      prompt: 'Почему в Go ошибку обычно возвращают явно?',
      options: [
        {
          id: 'caller-controls-error',
          label: 'A',
          text: 'Чтобы caller сам решил, как обработать сбой',
          isSelected: true,
          isCorrect: true,
        },
        {
          id: 'hide-problem',
          label: 'B',
          text: 'Чтобы скрыть проблему от пользователя',
          isSelected: false,
          isCorrect: false,
        },
        {
          id: 'replace-tests',
          label: 'C',
          text: 'Чтобы заменить unit tests',
          isSelected: false,
          isCorrect: false,
        },
        {
          id: 'speed-runtime',
          label: 'D',
          text: 'Чтобы ускорить runtime',
          isSelected: false,
          isCorrect: false,
        },
      ],
      feedbackTitle: 'Верно',
      feedbackText:
        'Явное возвращение ошибки позволяет вызывающей стороне принять правильное решение.',
      explanation:
        'В Go ошибки обрабатываются явно, чтобы не скрывать сбои и не прерывать выполнение неожиданно. Вызывающая сторона понимает контекст и может выбрать стратегию: повторить операцию, вернуть пользователю сообщение или выполнить компенсацию.',
    },
    sidebar: {
      conceptTitle: 'Концепция',
      conceptText: [
        'В Go функции, которые могут завершиться ошибкой, возвращают значение типа error.',
        'Такой подход держит обработку сбоя рядом с местом, где есть достаточно контекста.',
      ],
      codeSnippet:
        'func GetUser(id int) (User, error) {\n' +
        '    user, err := repo.Find(id)\n' +
        '    if err != nil {\n' +
        '        return User{}, fmt.Errorf("find user: %w", err)\n' +
        '    }\n' +
        '\n' +
        '    return user, nil\n' +
        '}',
      practiceNotes: [
        'Ошибки не игнорируются.',
        'Ошибки оборачиваются с контекстом.',
        'На границе API ошибки маппятся в корректные HTTP-статусы.',
        'Логи содержат детали, но не утечки.',
      ],
      usefulNote:
        'Не возвращайте nil error вместе с некорректным значением. Проверяйте ошибку перед использованием результата функции.',
      relatedTopics: ['Errors', 'Best Practices', 'Error Wrapping', 'Go Conventions'],
    },
  },
  currentChallenge: {
    id: 'retry-context',
    title: 'Напиши безопасный retry',
    difficulty: 'Junior Backend',
    rewardXp: 120,
    description:
      'Реализуй функцию FetchUser с механизмом retry для временно неудачных запросов. Повтор должен выполняться с экспоненциальной задержкой и уважать context.Context.',
    requirements: [
      'Используй context для отмены операций.',
      'Повторяй запрос до 3 попыток включительно.',
      'Начальная задержка 100ms, множитель 2.0 между попытками.',
      'Повторяй только для временных ошибок.',
      'Не повторяй для остальных ошибок.',
      'Убедись, что функция потокобезопасна.',
    ],
    signature: 'func FetchUser(ctx context.Context, id int) (User, error)',
    types:
      'type User struct {\n' +
      '    ID   int    `json:"id"`\n' +
      '    Name string `json:"name"`\n' +
      '}\n' +
      '\n' +
      'type Doer interface {\n' +
      '    Do(req *http.Request) (*http.Response, error)\n' +
      '}',
    hint:
      'Используй time.NewTimer и обязательно останавливай timer, чтобы избежать утечек goroutine.',
    starterCode:
      'package main\n' +
      '\n' +
      'import (\n' +
      '    "context"\n' +
      '    "errors"\n' +
      '    "io"\n' +
      '    "net/http"\n' +
      '    "time"\n' +
      ')\n' +
      '\n' +
      'type User struct {\n' +
      '    ID   int    `json:"id"`\n' +
      '    Name string `json:"name"`\n' +
      '}\n' +
      '\n' +
      'type Doer interface {\n' +
      '    Do(req *http.Request) (*http.Response, error)\n' +
      '}\n' +
      '\n' +
      'var client Doer\n' +
      '\n' +
      'func FetchUser(ctx context.Context, id int) (User, error) {\n' +
      '    // TODO: реализуй функцию согласно требованиям\n' +
      '    return User{}, errors.New("not implemented")\n' +
      '}\n' +
      '\n' +
      'func closeBody(resp *http.Response) {\n' +
      '    if resp != nil && resp.Body != nil {\n' +
      '        _, _ = io.Copy(io.Discard, resp.Body)\n' +
      '        _ = resp.Body.Close()\n' +
      '    }\n' +
      '}',
    terminalOutput: [
      '$ go test ./...',
      '=== RUN   TestFetchUser_Success',
      '--- PASS: TestFetchUser_Success (0.15s)',
      '=== RUN   TestFetchUser_RetryOnTemporary',
      '--- PASS: TestFetchUser_RetryOnTemporary (0.45s)',
      '=== RUN   TestFetchUser_ContextCancel',
      '--- PASS: TestFetchUser_ContextCancel (0.12s)',
      '=== RUN   TestFetchUser_NonRetryable',
      '--- PASS: TestFetchUser_NonRetryable (0.08s)',
      'PASS',
      'ok      challenge       0.812s',
    ],
    publicTests: [
      { name: 'TestFetchUser_Success', status: 'passed', durationMs: 95 },
      { name: 'TestFetchUser_RetryOnTemporary', status: 'passed', durationMs: 418 },
      { name: 'TestFetchUser_ContextCancel', status: 'passed', durationMs: 112 },
    ],
    hiddenTests: [
      { name: 'TestFetchUser_NonRetryable', status: 'passed' },
      { name: 'TestFetchUser_MaxRetries', status: 'passed' },
    ],
    result: {
      title: 'Все тесты пройдены!',
      description: 'Решение корректно обрабатывает повторы, отмену context и типы ошибок.',
      passedTests: 5,
      totalTests: 5,
    },
    sandbox: {
      status: 'Sandbox готов',
      network: 'Сеть запрещена',
      memoryLimit: '256 MB',
      timeout: '2s timeout',
      goVersion: 'Go 1.26',
    },
    attempts: [
      { timestamp: 'Сегодня, 12:34', summary: 'Все тесты пройдены', xp: 120, status: 'passed' },
      { timestamp: 'Сегодня, 12:21', summary: '2 / 3 теста пройдены', xp: 60, status: 'failed' },
      { timestamp: 'Сегодня, 11:58', summary: '1 / 3 теста пройдены', xp: 30, status: 'failed' },
    ],
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
  learningPath: {
    nextLevelProgressPercent: 67,
    activeBlock: {
      id: 'go-core',
      title: 'Go Core',
      description:
        'Активный блок про idiomatic error handling, table-driven tests, mocks и coverage.',
      progressPercent: 35,
      skills: ['Go', 'Testing', 'Errors', 'Table Tests', 'Mocking', 'Coverage'],
      requirements: ['Основы Go', 'Алгоритмы и структуры данных'],
      careerOutcomes: ['Go Developer (Junior)', 'Backend Developer', 'API Developer'],
    },
    stages: [
      {
        title: 'Стажер Backend',
        subtitle: '2 / 2 блока',
        completedBlocks: 2,
        totalBlocks: 2,
        blocks: [
          {
            id: 'go-basics',
            title: 'Основы Go',
            status: 'completed',
            tags: ['Go', 'Basics'],
            progressPercent: 100,
          },
          {
            id: 'algorithms-data-structures',
            title: 'Алгоритмы и структуры данных',
            status: 'completed',
            tags: ['Go'],
            progressPercent: 100,
          },
        ],
      },
      {
        title: 'Junior Backend',
        subtitle: '2 / 6 блоков',
        completedBlocks: 2,
        totalBlocks: 6,
        blocks: [
          {
            id: 'go-core',
            title: 'Go Core',
            status: 'active',
            tags: ['Go', 'Testing'],
            progressPercent: 35,
          },
          {
            id: 'rest-api-checkpoint',
            title: 'Checkpoint: REST API',
            status: 'checkpoint',
            tags: ['Go', 'REST API'],
            tasksCompleted: 0,
            tasksTotal: 8,
          },
          {
            id: 'postgresql',
            title: 'PostgreSQL',
            status: 'locked',
            tags: ['PostgreSQL', 'SQL'],
          },
          {
            id: 'concurrency',
            title: 'Concurrency',
            status: 'locked',
            tags: ['Go', 'Concurrency'],
          },
          {
            id: 'message-brokers',
            title: 'Message Brokers',
            status: 'locked',
            tags: ['Kafka', 'RabbitMQ', 'Redis'],
          },
        ],
      },
      {
        title: 'Middle Backend',
        subtitle: '0 / 6 блоков',
        completedBlocks: 0,
        totalBlocks: 6,
        blocks: [
          {
            id: 'production-reliability',
            title: 'Production Reliability',
            status: 'locked',
            tags: ['Kubernetes', 'Prometheus', 'Grafana'],
          },
          {
            id: 'architecture-scaling',
            title: 'Архитектура и масштабирование',
            status: 'locked',
            tags: ['System Design', 'Performance'],
          },
        ],
      },
      {
        title: 'Senior Backend',
        subtitle: '0 / 6 блоков',
        completedBlocks: 0,
        totalBlocks: 6,
        blocks: [
          {
            id: 'technical-leadership',
            title: 'Technical Leadership',
            status: 'locked',
            tags: ['SLO', 'Review', 'Mentoring'],
          },
        ],
      },
    ],
  },
  levels: [
    { title: 'Стажер Backend', status: 'completed' },
    { title: 'Junior Backend', status: 'active' },
    { title: 'Middle Backend', status: 'locked' },
    { title: 'Senior Backend', status: 'locked' },
  ],
}
