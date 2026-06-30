# GoPath

GoPath — веб-приложение для русскоязычного обучения Go backend-разработке через короткие уроки, карьерную карту, XP, streak, checkpoints, задания с выбором ответа и code challenges.

Проект находится на этапе bootstrap. Сначала фиксируются источники истины и структура разработки, затем будут добавляться frontend и backend skeleton небольшими отдельными commit.

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
- [Bootstrap plan](docs/bootstrap-plan.md)

## Development Status

Текущий репозиторий содержит только bootstrap-документацию. Frontend, backend, infrastructure и CI будут добавляться отдельными маленькими commit после согласования каждого шага.
