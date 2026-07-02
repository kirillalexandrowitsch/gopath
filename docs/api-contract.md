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

## Planned Auth Contract

Auth endpoints are planned but are not implemented yet. Current backend route behavior is unchanged; these notes exist to guide future auth implementation.

Planned endpoints:

```text
POST /api/v1/auth/register
POST /api/v1/auth/login
POST /api/v1/auth/logout
GET  /api/v1/me
```

Session direction:

- Auth will use cookie-based sessions for the web app.
- Planned cookie name: `gopath_session`.
- Planned cookie attributes: `HttpOnly`, `SameSite=Lax`, `Path=/`.
- Localhost and production `Secure` behavior will be finalized during implementation.
- OAuth, email verification, password reset, CSRF and auth rate limits are future security work.

### `POST /api/v1/auth/register`

Planned request body:

```json
{
  "email": "alex@example.com",
  "handle": "@alexkim",
  "displayName": "Alex Kim",
  "password": "demo-password"
}
```

Planned success response:

```json
{
  "user": {
    "id": "demo-user",
    "email": "alex@example.com",
    "handle": "@alexkim",
    "displayName": "Alex Kim",
    "role": "learner"
  }
}
```

Register will create a user and issue `gopath_session` when implemented.

### `POST /api/v1/auth/login`

Planned request body with email:

```json
{
  "email": "alex@example.com",
  "password": "demo-password"
}
```

Planned request body with handle:

```json
{
  "handle": "@alexkim",
  "password": "demo-password"
}
```

`email` or `handle` can identify the account. Exact validation rules will be finalized during implementation.

Planned success response uses the same auth user shape as register. Login will issue `gopath_session` when implemented.

### `POST /api/v1/auth/logout`

Planned success response:

```json
{
  "status": "ok"
}
```

Logout will delete the current session and clear `gopath_session` when implemented.

### `GET /api/v1/me`

Planned success response uses the same auth user shape as register.

Unauthenticated requests will return a JSON error response. Exact status code and error string will be finalized when auth middleware is implemented.

Current auth limitations:

- Auth endpoints are not implemented yet.
- Current progress/profile endpoints still use `demo-user`.
- No session table, session repository, auth middleware, frontend auth flow, CSRF handling or auth rate limits are wired yet.

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

Response fields:

- `tracks` - array of track objects.
- `tracks[].id` - stable track id. Current known value: `go-backend`.
- `tracks[].title` - display title.
- `tracks[].description` - short product-facing description.
- `tracks[].levelIds` - ordered career level ids in this track.

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

Response fields:

- `levels` - array of career level objects.
- `levels[].id` - stable level id. Current known values: `trainee-backend`, `junior-backend`, `middle-backend`, `senior-backend`.
- `levels[].title` - display title.
- `levels[].status` - current level status. Current values: `completed`, `active`, `locked`.
- `levels[].completedBlocks` - completed learning blocks in this level.
- `levels[].totalBlocks` - total learning blocks in this level.
- `levels[].requiredXp` - XP required to reach or unlock this level.

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

Response fields:

- `id` - stable lesson id. Current known value: `go-errors-tests`.
- `title` - lesson title.
- `step` - display progress label for the current lesson sequence.
- `progressPercent` - lesson or block progress percentage.
- `xpReward` - XP reward shown for completing the lesson.
- `tags` - technical topic tags.
- `question.prompt` - single-choice question text.
- `question.options` - ordered answer options.
- `question.options[].id` - stable option id.
- `question.options[].label` - short option label, for example `A`.
- `question.options[].text` - answer option text.
- `question.correctId` - id of the correct option.
- `question.feedback` - positive feedback text for the current bootstrap answer state.
- `question.explanation` - explanation shown after answer checking.
- `sidebar.conceptTitle` - sidebar concept heading.
- `sidebar.conceptText` - theory paragraphs.
- `sidebar.codeSnippet` - Go code snippet.
- `sidebar.practiceNotes` - backend practice notes.
- `sidebar.relatedTopics` - related topic tags.
- `nextChallengeId` - suggested challenge id after the lesson.

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

Response fields:

- `userId` - current user id. Current temporary value: `demo-user`.
- `level` - display title of the current career level.
- `xp` - current XP.
- `nextLevelXp` - XP target for the next level.
- `streakDays` - current streak length in days.
- `currentLessonId` - id of the lesson to continue.
- `completedLessons` - total completed lessons.
- `practiceTasks` - completed practice tasks count.
- `completedCheckpoints` - completed checkpoints count.
- `totalCheckpoints` - total checkpoints in the current progress scope.
- `dailyGoal.earnedXp` - XP earned toward today's goal.
- `dailyGoal.targetXp` - today's XP target.
- `dailyGoal.completedTasks` - completed tasks toward today's goal.
- `dailyGoal.targetTasks` - today's task target.
- `weeklyXp[].day` - short localized day label.
- `weeklyXp[].xp` - XP earned on that day.

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

Response fields:

- `user.id` - current user id. Current temporary value: `demo-user`.
- `user.name` - display name.
- `user.handle` - user handle.
- `user.initials` - initials for avatar placeholder.
- `user.joinedAt` - localized join date label.
- `level` - display title of the current career level.
- `xp` - current XP.
- `nextLevelXp` - XP target for the next level.
- `streakDays` - current streak length in days.
- `skills[].name` - skill or technology name.
- `skills[].score` - skill score from `0` to `100`.
- `readiness.targetLevel` - next target career level.
- `readiness.progressPercent` - readiness progress percentage.
- `readiness.items[].title` - readiness checklist item title.
- `readiness.items[].current` - current numeric progress for the item.
- `readiness.items[].target` - target numeric progress for the item.
- `readiness.items[].isComplete` - completion flag.
- `recommendation.title` - recommendation title.
- `recommendation.description` - recommendation body.
- `recommendation.path` - frontend path for the recommended next action.

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

Response fields:

- `status` - terminal challenge execution status for current sync endpoints.
- `stdout` - captured standard output from runner commands.
- `stderr` - captured standard error from runner commands.
- `testsPassed` - number of passed tests.
- `testsTotal` - total number of tests executed for this request.
- `durationMs` - execution duration in milliseconds.
- `hint` - optional learner-facing hint for failed, compile error or timeout results.

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

Reserved future statuses:

- `queued`
- `running`

Current sync `run` and `submit` endpoints do not return `queued` or `running`.

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
