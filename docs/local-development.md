# Local Development

Эта инструкция описывает локальный запуск GoPath на `localhost`. Она покрывает текущий dev-flow: frontend, backend, PostgreSQL, Redis, migrations, проверки и базовые smoke checks.

GoPath сначала доводится до production-ready состояния локально. Deploy начинается только после localhost acceptance gate из `docs/delivery-roadmap.md`.

## Prerequisites

Нужны:

- Go `1.26`.
- Node.js и `npm`.
- Docker Compose.
- `curl` для smoke checks.

Проверьте версии:

```sh
go version
node --version
npm --version
docker compose version
```

## Environment Files

Tracked defaults находятся в:

```text
infra/.env.example
```

Для локальных изменений скопируйте файл:

```sh
cp infra/.env.example infra/.env
```

`infra/.env` ignored by git. Не коммитьте локальные credentials, ports или secrets.

Текущие defaults:

```text
GOPATH_POSTGRES_DB=gopath
GOPATH_POSTGRES_USER=gopath
GOPATH_POSTGRES_PASSWORD=gopath_dev_password
GOPATH_POSTGRES_PORT=5432
GOPATH_DATABASE_URL=postgres://gopath:gopath_dev_password@localhost:5432/gopath?sslmode=disable
GOPATH_REDIS_PORT=6379
GOPATH_REDIS_ADDR=localhost:6379
```

## Root Command Workflow

Основной local workflow теперь идет через root `Makefile`.

Посмотреть доступные команды:

```sh
make help
```

Первичный setup frontend dependencies:

```sh
cd frontend
npm install
cd ..
```

Проверить local Docker Compose config без запуска stack:

```sh
make stack-config
```

Если используете локальный `infra/.env`, передайте override:

```sh
COMPOSE_ENV=infra/.env make stack-config
```

Запустить PostgreSQL и Redis:

```sh
make stack-up
```

Для локального `infra/.env`:

```sh
COMPOSE_ENV=infra/.env make stack-up
```

Применить migrations после запуска PostgreSQL:

```sh
./scripts/migrate-up.sh
```

Запустить backend без PostgreSQL dependency, в `MemoryStore` mode:

```sh
cd backend
go run ./cmd/api
```

Запустить backend в PostgreSQL-backed mode:

```sh
cd backend
GOPATH_DATABASE_URL='postgres://gopath:gopath_dev_password@localhost:5432/gopath?sslmode=disable' go run ./cmd/api
```

Запустить frontend:

```sh
cd frontend
npm run dev
```

Запустить проверки:

```sh
make backend-verify
make frontend-verify
```

Запустить API smoke checks, когда backend уже поднят:

```sh
make smoke-api
```

Если backend слушает другой address:

```sh
API_BASE_URL=http://localhost:8081 make smoke-api
```

Остановить local stack без удаления volumes:

```sh
make stack-down
```

Для локального `infra/.env`:

```sh
COMPOSE_ENV=infra/.env make stack-down
```

Не используйте `--volumes`, если явно не хотите удалить локальные PostgreSQL/Redis данные.

## Validate Compose Config

Проверить Docker Compose config без запуска stack:

```sh
make stack-config
```

Если используете локальный `infra/.env`:

```sh
COMPOSE_ENV=infra/.env make stack-config
```

## Start PostgreSQL and Redis

Запустить локальные PostgreSQL и Redis:

```sh
COMPOSE_ENV=infra/.env make stack-up
```

Если `infra/.env` не создан, можно использовать tracked example:

```sh
make stack-up
```

Посмотреть состояние stack:

```sh
docker compose --env-file infra/.env -f infra/compose.yml ps
```

## Apply Migrations

После запуска PostgreSQL примените migrations:

```sh
./scripts/migrate-up.sh
```

Script выбирает `infra/.env`, если файл существует, иначе использует `infra/.env.example`.

Важно:

- script применяет все `backend/migrations/*.sql` в lexicographic order;
- script выполняет SQL против локального PostgreSQL service;
- script не удаляет volumes и не делает destructive reset;
- текущий migration runner пока не ведет отдельную таблицу applied migrations.

## Run Backend

Backend находится в `backend/`.

### MemoryStore Mode

Если `GOPATH_DATABASE_URL` не задан, API стартует с in-memory learning data:

```sh
cd backend
go run ./cmd/api
```

По умолчанию API слушает:

```text
http://localhost:8080
```

Можно изменить address через:

```sh
GOPATH_API_ADDR=:8081 go run ./cmd/api
```

### PostgreSQL-backed Mode

Чтобы API читал learning data из PostgreSQL, задайте `GOPATH_DATABASE_URL` из `infra/.env.example` или `infra/.env`:

```sh
cd backend
GOPATH_DATABASE_URL='postgres://gopath:gopath_dev_password@localhost:5432/gopath?sslmode=disable' go run ./cmd/api
```

В этом режиме PostgreSQL stack должен быть запущен, а migrations должны быть применены.

## Run Frontend

Frontend находится в `frontend/`.

Установить dependencies:

```sh
cd frontend
npm install
```

Запустить Vite dev server:

```sh
npm run dev
```

Frontend сейчас использует local mock data. Backend API integration будет отдельной roadmap phase.

## Checks

Backend checks:

```sh
make backend-verify
```

Frontend checks:

```sh
make frontend-verify
```

Repository whitespace check:

```sh
make diff-check
```

## Smoke Checks

Smoke checks предполагают, что backend запущен на `localhost:8080`.

```sh
make smoke-api
```

Для backend на другом address:

```sh
API_BASE_URL=http://localhost:8081 make smoke-api
```

Smoke script проверяет `GET /healthz`, tracks, levels, current lesson, progress и profile.

Challenge `run` и `submit` требуют JSON body с Go code. Полная frontend integration и persisted attempts будут добавлены позже.

## Stop Docker Stack

Остановить local stack без удаления volumes:

```sh
COMPOSE_ENV=infra/.env make stack-down
```

Если запускали stack через example env:

```sh
make stack-down
```

Не используйте `--volumes`, если явно не хотите удалить локальные PostgreSQL/Redis данные.

## Current Limitations

- Frontend пока использует local typed mock data.
- Auth, sessions и `GET /api/v1/me` еще не реализованы.
- Attempts persistence еще не подключен к API flows.
- XP, streak, daily goal и progression writes пока не работают как real domain logic.
- Redis есть в local stack, но backend runtime пока не использует Redis.
- Challenge runner пока local prototype через `go test`, а не production sandbox.
- GitHub Actions CI пока отсутствует.
- Deploy намеренно отложен до закрытия localhost acceptance gate.
