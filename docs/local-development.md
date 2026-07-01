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

## Validate Compose Config

Проверить Docker Compose config без запуска stack:

```sh
docker compose --env-file infra/.env.example -f infra/compose.yml config
```

Если используете локальный `infra/.env`:

```sh
docker compose --env-file infra/.env -f infra/compose.yml config
```

## Start PostgreSQL and Redis

Запустить локальные PostgreSQL и Redis:

```sh
docker compose --env-file infra/.env -f infra/compose.yml up -d
```

Если `infra/.env` не создан, можно использовать tracked example:

```sh
docker compose --env-file infra/.env.example -f infra/compose.yml up -d
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
cd backend
GOCACHE=/private/tmp/gopath-gocache go test ./...
```

Frontend checks:

```sh
cd frontend
npm run lint
npm run build
```

Repository whitespace check:

```sh
git diff --check
```

## Smoke Checks

Smoke checks предполагают, что backend запущен на `localhost:8080`.

Health:

```sh
curl -sS http://localhost:8080/healthz
```

Tracks:

```sh
curl -sS http://localhost:8080/api/v1/tracks
```

Current lesson:

```sh
curl -sS http://localhost:8080/api/v1/lessons/go-errors-tests
```

Challenge `run` и `submit` требуют JSON body с Go code. Полная frontend integration и persisted attempts будут добавлены позже.

## Stop Docker Stack

Остановить local stack без удаления volumes:

```sh
docker compose --env-file infra/.env -f infra/compose.yml down
```

Если запускали stack через example env:

```sh
docker compose --env-file infra/.env.example -f infra/compose.yml down
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
