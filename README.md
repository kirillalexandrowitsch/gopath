# GoPath

[![Backend](https://github.com/kirillalexandrowitsch/gopath/actions/workflows/backend.yml/badge.svg?branch=main)](https://github.com/kirillalexandrowitsch/gopath/actions/workflows/backend.yml)
[![Frontend](https://github.com/kirillalexandrowitsch/gopath/actions/workflows/frontend.yml/badge.svg?branch=main)](https://github.com/kirillalexandrowitsch/gopath/actions/workflows/frontend.yml)
[![Infrastructure](https://github.com/kirillalexandrowitsch/gopath/actions/workflows/infra.yml/badge.svg?branch=main)](https://github.com/kirillalexandrowitsch/gopath/actions/workflows/infra.yml)
[![Migrations](https://github.com/kirillalexandrowitsch/gopath/actions/workflows/migrations.yml/badge.svg?branch=main)](https://github.com/kirillalexandrowitsch/gopath/actions/workflows/migrations.yml)

GoPath — веб-приложение для русскоязычного обучения Go backend-разработке через короткие уроки, карьерную карту, XP, streak, checkpoints, задания с выбором ответа и code challenges.

Проект развивается маленькими проверяемыми commit по localhost-first roadmap: сначала доводим приложение до production-ready состояния на `localhost`, и только после этого переходим к deploy.

## MVP scope

- Русскоязычный интерфейс и учебный контент.
- Карьерная карта backend-разработчика: `Стажер Backend`, `Junior Backend`, `Middle Backend`, `Senior Backend`.
- 5 основных экранов: `Dashboard`, `Learning Path`, `Lesson`, `Code Challenge`, `Profile / Progress`.
- Progression mechanics: XP, streak, daily goal, checkpoints, active/completed/locked states.
- Go code challenges с public/hidden tests.
- REST API для auth, learning, attempts, challenges и progress.

## Stack

Frontend:

- React
- TypeScript
- Vite
- Tailwind CSS
- React Query
- Zustand
- Monaco Editor
- xterm.js

Backend:

- Go
- REST API
- PostgreSQL
- Redis
- Docker Compose

## Sources of Truth

- [Technical specification](docs/gopath-technical-spec.md)
- [Delivery roadmap](docs/delivery-roadmap.md)
- [Architecture notes](docs/architecture.md)
- [API contract notes](docs/api-contract.md)
- [Content authoring guide](docs/content-authoring.md)
- [Bootstrap plan](docs/bootstrap-plan.md) — historical bootstrap plan

## Local Development

Локальный запуск frontend, backend, PostgreSQL, Redis, migrations и smoke checks описан в [Local Development](docs/local-development.md).

## CI

GitHub Actions baseline:

- `Backend` запускает Go tests в `backend/`.
- `Frontend` запускает `npm ci`, `npm run lint` и `npm run build` в `frontend/`.
- `Infrastructure` проверяет Docker Compose config и shell syntax для scripts.
- `Migrations` запускает lightweight validation для SQL migration files.

## Development Status

Уже есть:

- `frontend/` на Vite, React, TypeScript, Tailwind CSS и React Router.
- Static MVP shell и 5 mock screens: `Dashboard`, `Learning Path`, `Lesson`, `Code Challenge`, `Profile`.
- `backend/` Go REST API на стандартном `net/http`.
- `GET /healthz`, read-only learning API endpoints и challenge runner endpoints.
- Local PostgreSQL/Redis Docker Compose config в `infra/`.
- SQL migrations, demo seed data и migration runner script.
- Optional PostgreSQL-backed learning store через `GOPATH_DATABASE_URL`.
- Local Go challenge runner prototype для `Run` и `Submit`.
- GitHub Actions CI baseline для backend, frontend, infrastructure и migrations.

Пока не готово:

- Frontend API integration.
- Auth, sessions и `GET /api/v1/me`.
- Attempts persistence.
- Real XP, streak, daily goal и progression writes.
- Production sandbox isolation.
- Deploy.

Следующий порядок работ описан в [Delivery roadmap](docs/delivery-roadmap.md).
