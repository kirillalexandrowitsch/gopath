# API Contract Notes

Этот документ фиксирует текущий backend API snapshot для GoPath. Это не финальный production contract, а описание фактически существующего API после bootstrap commits.

Product requirements живут в `docs/gopath-technical-spec.md`, порядок дальнейших изменений - в `docs/delivery-roadmap.md`.

## Base URL

Локальный backend по умолчанию:

```text
http://localhost:8080
```

Address можно изменить через `GOPATH_API_ADDR`.

## Response Format

Текущие API responses возвращаются как JSON:

```http
Content-Type: application/json
```

Текущий error response:

```json
{
  "error": "internal server error"
}
```

Текущие user-facing error strings:

- `lesson not found`
- `challenge not found`
- `invalid JSON`
- `code is required`
- `internal server error`

## Health

### `GET /healthz`

Проверяет, что HTTP server отвечает.

Response:

```json
{
  "status": "ok"
}
```

## Learning Endpoints

Learning endpoints сейчас read-only.

### `GET /api/v1/tracks`

Возвращает список learning tracks.

Response shape:

```json
{
  "tracks": [
    {
      "id": "go-backend",
      "title": "Go Backend",
      "description": "Карьерный путь Go backend-разработчика от основ до production reliability.",
      "levelIds": ["trainee-backend", "junior-backend"]
    }
  ]
}
```

### `GET /api/v1/levels`

Возвращает карьерные уровни.

Response shape:

```json
{
  "levels": [
    {
      "id": "junior-backend",
      "title": "Junior Backend",
      "status": "active",
      "completedBlocks": 2,
      "totalBlocks": 6,
      "requiredXp": 2000
    }
  ]
}
```

### `GET /api/v1/lessons/{id}`

Возвращает урок по id.

Current known lesson id:

```text
go-errors-tests
```

Response shape:

```json
{
  "id": "go-errors-tests",
  "title": "Go: ошибки и тесты",
  "step": "Урок 4 из 12",
  "progressPercent": 67,
  "xpReward": 40,
  "tags": ["Go", "Testing", "Errors"],
  "question": {
    "prompt": "Почему в Go ошибку обычно возвращают явно?",
    "options": [
      {
        "id": "caller-controls-error",
        "label": "A",
        "text": "Чтобы caller сам решил, как обработать сбой"
      }
    ],
    "correctId": "caller-controls-error",
    "feedback": "Верно",
    "explanation": "В Go ошибки обрабатываются явно, чтобы вызывающая сторона понимала контекст и могла выбрать стратегию обработки сбоя."
  },
  "sidebar": {
    "conceptTitle": "Концепция",
    "conceptText": [],
    "codeSnippet": "func GetUser(id int) (User, error) { ... }",
    "practiceNotes": [],
    "relatedTopics": []
  },
  "nextChallengeId": "retry-context"
}
```

Errors:

- unknown lesson id -> `404 {"error":"lesson not found"}`

Note:

- `correctId` currently exists for bootstrap lesson UI. It should be revisited before real assessment flow, because production answer checking should not rely on exposing the correct answer to the client.

## Progress Endpoints

### `GET /api/v1/progress`

Возвращает summary прогресса текущего demo user.

Response shape:

```json
{
  "userId": "demo-user",
  "level": "Junior Backend",
  "xp": 1420,
  "nextLevelXp": 2000,
  "streakDays": 7,
  "currentLessonId": "go-errors-tests",
  "completedLessons": 24,
  "practiceTasks": 18,
  "completedCheckpoints": 2,
  "totalCheckpoints": 6,
  "dailyGoal": {
    "earnedXp": 150,
    "targetXp": 200,
    "completedTasks": 2,
    "targetTasks": 3
  },
  "weeklyXp": [
    {
      "day": "Пн",
      "xp": 120
    }
  ]
}
```

Temporary behavior:

- endpoint uses `demo-user` until auth and sessions are implemented;
- XP, streak, daily goal and progression writes are not implemented yet.

### `GET /api/v1/profile`

Возвращает profile/progress data для текущего demo user.

Response shape:

```json
{
  "user": {
    "id": "demo-user",
    "name": "Alex Kim",
    "handle": "@alexkim",
    "initials": "AK",
    "joinedAt": "15 апр. 2024"
  },
  "level": "Junior Backend",
  "xp": 1420,
  "nextLevelXp": 2000,
  "streakDays": 7,
  "skills": [
    {
      "name": "Go",
      "score": 84
    }
  ],
  "readiness": {
    "targetLevel": "Middle Backend",
    "progressPercent": 67,
    "items": []
  },
  "recommendation": {
    "title": "Рекомендация для вас",
    "description": "Попробуйте практику по теме Concurrency.",
    "path": "/challenge/retry-context"
  }
}
```

Temporary behavior:

- endpoint uses `demo-user` until auth and sessions are implemented;
- some profile/readiness fields are still temporary static defaults in PostgreSQL-backed mode.

## Challenge Endpoints

Current known challenge id:

```text
retry-context
```

### `POST /api/v1/challenges/{id}/run`

Runs public tests for a challenge.

Request body:

```json
{
  "code": "package main\n\nfunc FetchUser(...) { ... }"
}
```

Response shape:

```json
{
  "status": "passed",
  "stdout": "go test ./... ok",
  "stderr": "",
  "testsPassed": 3,
  "testsTotal": 3,
  "durationMs": 148,
  "hint": ""
}
```

### `POST /api/v1/challenges/{id}/submit`

Runs public and hidden tests for a challenge.

Request body:

```json
{
  "code": "package main\n\nfunc FetchUser(...) { ... }"
}
```

Response shape is the same as `run`, but `testsTotal` includes hidden tests.

### Challenge Statuses

Current terminal statuses:

- `passed`
- `failed`
- `compile_error`
- `timeout`
- `internal_error`

Future statuses from the product spec, such as `queued`, `running` and `blocked`, are not implemented yet.

### Challenge Errors

- invalid JSON -> `400 {"error":"invalid JSON"}`
- empty or whitespace-only `code` -> `400 {"error":"code is required"}`
- unknown challenge id -> `404 {"error":"challenge not found"}`
- unexpected runner failure -> `500 {"error":"internal server error"}`

### Challenge Limitations

- `run` and `submit` execute a local prototype runner.
- Attempts are not persisted.
- XP and progress are not updated.
- Redis job queue is not used.
- Docker sandbox isolation is not implemented.
- Network, syscall, CPU and memory isolation are future sandbox-hardening work.

## Current Known IDs

- Track: `go-backend`
- Levels:
  - `trainee-backend`
  - `junior-backend`
  - `middle-backend`
  - `senior-backend`
- Lesson: `go-errors-tests`
- Challenge: `retry-context`
- Demo user: `demo-user`

## Current Contract Limitations

- Frontend is not integrated with the API yet.
- Auth endpoints are not implemented yet.
- `GET /api/v1/me` is not implemented yet.
- Progress/profile endpoints use `demo-user`.
- Attempts persistence is not wired into API flows.
- XP, streak, daily goal and progression writes are not implemented yet.
- Challenge runner is a local prototype, not a production sandbox.
- This document must be updated whenever API behavior changes.
