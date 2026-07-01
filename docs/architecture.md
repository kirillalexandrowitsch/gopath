# GoPath Architecture Notes

This document records the current technical architecture for GoPath. Product requirements live in `docs/gopath-technical-spec.md`; delivery order lives in `docs/delivery-roadmap.md`.

## Repository Layout

Current directories:

- `docs/` - product specification, delivery roadmap, bootstrap history, and architecture notes.
- `frontend/` - React, TypeScript, Vite, Tailwind CSS, and React Router single-page application.
- `backend/` - Go REST API module.
- `infra/` - local Docker Compose configuration for PostgreSQL and Redis.
- `scripts/` - local operational scripts, currently including migration application.

Planned directories:

- `content/` - source learning content for tracks, levels, lessons, challenges, tests, hints, and seed generation.
- `.github/workflows/` - GitHub Actions CI checks.

## Frontend Boundary

The frontend is a single-page application built with:

- React
- TypeScript
- Vite
- Tailwind CSS
- React Router

Current application routes:

```text
/app
/learn
/lesson/:lessonId
/challenge/:challengeId
/profile
```

The frontend currently provides the static MVP shell and five mock screens:

- Dashboard
- Learning Path
- Lesson
- Code Challenge
- Profile

Current data source:

- Screen data is local typed mock data in the frontend source.
- Frontend routes do not yet call the backend API.
- Loading, error, auth, and real submission states are future work.

Future frontend boundaries:

- API client and environment-based API base URL.
- Backend-driven Dashboard, Learning Path, Lesson, Challenge, and Profile data.
- Lesson answer submission flow.
- Challenge `Run` and `Submit` integration.
- Monaco Editor for code editing.
- xterm.js or equivalent terminal surface.
- Optional React Query and Zustand only when a concrete slice needs them.

## Backend Boundary

The backend is a Go REST API using the standard library `net/http` router.

Current endpoints:

```text
GET  /healthz
GET  /api/v1/tracks
GET  /api/v1/levels
GET  /api/v1/lessons/{id}
GET  /api/v1/progress
GET  /api/v1/profile
POST /api/v1/challenges/{id}/run
POST /api/v1/challenges/{id}/submit
```

Current API characteristics:

- Learning endpoints are read-only.
- Progress and profile temporarily use `demo-user` until auth is added.
- Challenge endpoints execute a local prototype runner and return immediate results.
- Unknown lesson ids return `404` with `{"error":"lesson not found"}`.
- Unknown challenge ids return `404` with `{"error":"challenge not found"}`.

Future backend boundaries:

- Auth endpoints: register, login, logout, and `GET /api/v1/me`.
- Cookie-based session middleware.
- Attempts persistence.
- XP, streak, daily goal, checkpoint, and unlock domain rules.
- Stable API error and validation conventions.
- Async or queued sandbox execution when Redis-backed sandbox jobs are introduced.

## Learning Layer

The `backend/internal/learning` package owns the current learning read model.

Current boundary:

- `Store` interface provides `Tracks`, `Levels`, `Lesson`, `Progress`, and `Profile`.
- `MemoryStore` is the default store used by `httpserver.NewRouter()`.
- `PostgresStore` is available for read-only PostgreSQL-backed data.
- `cmd/api` selects the store at startup:
  - empty `GOPATH_DATABASE_URL` -> `MemoryStore`;
  - non-empty `GOPATH_DATABASE_URL` -> open PostgreSQL connection and use `PostgresStore`.

Current limitation:

- `PostgresStore` still uses temporary derived/static defaults for fields not yet represented in the schema.
- This is acceptable until the persistence/content roadmap phases remove those defaults.

## Database Boundary

The backend has a minimal database connection layer in `backend/internal/database`.

Current PostgreSQL assets:

- `backend/migrations/000001_init_schema.sql` - initial schema.
- `backend/migrations/000002_seed_learning_data.sql` - demo learning seed data.
- `scripts/migrate-up.sh` - applies migrations in lexicographic order through the local PostgreSQL service.

Current schema covers:

- users
- tracks
- levels
- lessons
- lesson options
- challenges
- challenge tests
- attempts
- user progress
- skill progress

Current limitation:

- API write flows do not yet use the database.
- Attempts, XP, streak, daily goal, and checkpoint progression are not yet persisted through product endpoints.
- Migration tracking is not yet implemented.

## Challenge Runner Boundary

The `backend/internal/challenges` package owns the current challenge execution prototype.

Current behavior:

- Supports challenge `retry-context`.
- `Run` executes public tests.
- `Submit` executes public and hidden tests.
- Creates a temporary directory.
- Writes `go.mod`, `solution.go`, and `solution_test.go`.
- Runs `gofmt` and `go test` through `exec.CommandContext`.
- Uses a 2 second timeout.
- Removes the temporary directory after execution.

Current statuses:

```text
passed
failed
compile_error
timeout
internal_error
```

Current limitation:

- This is not a production sandbox.
- It does not provide Docker isolation.
- It does not use Redis jobs.
- It does not enforce network, syscall, CPU, or memory isolation beyond the current process timeout behavior.
- It does not persist attempts or update XP/progress.

Future sandbox work:

- Separate prototype runner from sandbox runner.
- Add Redis-backed job state when needed.
- Add Docker-based isolated execution.
- Disable network for sandbox containers.
- Enforce CPU, memory, and time limits.
- Persist run and submit attempts.
- Add rate limits and cleanup checks.

## Local Infrastructure

Local infrastructure is defined in `infra/compose.yml`.

Current services:

- `postgres` using `postgres:17-alpine`.
- `redis` using `redis:7-alpine`.

Current local defaults:

- Environment example: `infra/.env.example`.
- PostgreSQL data is stored in a named Docker volume.
- Redis data is stored in a named Docker volume.
- Healthchecks are configured for both services.

Current limitation:

- The backend only connects to PostgreSQL when `GOPATH_DATABASE_URL` is set.
- Redis is present in local infrastructure but is not wired into backend runtime yet.
- Docker stack is local development infrastructure, not production deployment.

## Sources of Truth

- Product requirements: `docs/gopath-technical-spec.md`.
- Delivery order: `docs/delivery-roadmap.md`.
- Technical architecture: this document.
- Bootstrap history: `docs/bootstrap-plan.md`.

## Future Boundaries

The following areas are intentionally future work and should be implemented as small roadmap-driven commits:

- GitHub Actions CI.
- Root developer commands.
- Local development guide.
- Auth and sessions.
- Frontend API integration.
- Attempts persistence.
- Progression domain rules.
- Content source and validation pipeline.
- Redis-backed sandbox jobs and rate limits.
- Docker-based sandbox hardening.
- Observability and security hardening.
- Localhost final acceptance.
- Deployment after localhost completion.

## Development Rules

- Keep one small approved step per commit.
- Do not mix docs, frontend, backend, infrastructure, and CI changes unless a step explicitly requires it.
- Prefer existing project patterns and standard library until a dependency is justified by a concrete roadmap step.
- Run the narrowest meaningful checks for the changed layer before committing.
- Do not start deploy work until localhost final acceptance is complete.
