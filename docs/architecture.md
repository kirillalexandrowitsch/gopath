# GoPath Architecture Notes

This document records current technical decisions for the GoPath MVP bootstrap. The product requirements remain in `docs/gopath-technical-spec.md`.

## Repository Layout

Current directories:

- `docs/` — product specification, bootstrap plan, and architecture notes.
- `frontend/` — React, TypeScript, Vite, and Tailwind CSS application.
- `backend/` — Go REST API module.

Planned directories:

- `infra/` — local Docker Compose and infrastructure configuration.
- `content/` — seed learning content, mock tracks, lessons, and challenges.
- `.github/workflows/` — CI checks.

## Frontend Boundary

The frontend is a single-page application built with React, TypeScript, Vite, and Tailwind CSS.

Planned application routes:

```text
/app
/learn
/lesson/:lessonId
/challenge/:challengeId
/profile
```

Routing, React Query, Zustand, Monaco Editor, and xterm.js will be added in separate commits when the related product slice needs them. The current frontend skeleton must stay minimal until the static MVP shell begins.

## Backend Boundary

The backend is a Go REST API. The current server uses the standard library `net/http` and exposes only:

```text
GET /healthz
```

The MVP API surface will live under:

```text
/api/v1
```

The first product endpoints will follow the API list from `docs/gopath-technical-spec.md`: auth, learning, attempts, challenges, and progress. Database access, Redis, migrations, and product handlers are intentionally future work.

## Defaults for Future Work

- Auth: cookie-based sessions for the web app by default.
- PostgreSQL: users, lessons, attempts, and progress storage.
- Migrations: add a migration tool and schema in a dedicated future commit.
- Redis: short-lived sandbox jobs, rate limits, and cache.
- Sandbox: hybrid runner; start with a simple prototype, then add stronger isolation.
- Docker Compose: local PostgreSQL, Redis, and app services in a future infrastructure step.
- Kafka, RabbitMQ, Kubernetes, Prometheus, Grafana, ClickHouse, and similar systems start as learning content, not required infrastructure for the MVP application.

## Development Rules

- Keep one small approved step per commit.
- Do not mix docs, frontend, backend, infrastructure, and CI changes unless a step explicitly requires it.
- Prefer standard library and existing project patterns until additional dependencies are justified.
- Run the narrowest meaningful checks for the changed layer before committing.
