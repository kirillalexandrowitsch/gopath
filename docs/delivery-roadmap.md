# GoPath Delivery Roadmap

Этот документ является главным delivery source of truth для разработки GoPath. Product source of truth остается в `docs/gopath-technical-spec.md`.

Roadmap описывает путь от текущего состояния проекта до полностью завершенного production-ready приложения. Главный принцип: сначала доводим GoPath до полного production-ready состояния на `localhost`, и только после этого начинаем deploy.

## 1. Рабочие правила

### 1.1 Один шаг - один commit

- Один согласованный маленький шаг выполняется одним commit.
- Нельзя закрывать большую тему одним commit.
- Нельзя объединять несколько независимых этапов в один commit.
- Нельзя смешивать unrelated changes.
- Если commit меняет frontend, backend, infra и docs одновременно, это должно быть явно необходимо для одного маленького шага.
- После каждого commit работа останавливается до отдельной команды пользователя.

### 1.2 Процесс каждого commit

Каждый будущий commit проходит один и тот же цикл:

1. Выбрать следующий пункт из этого roadmap.
2. Исследовать текущий код и документы, относящиеся только к этому пункту.
3. Сформировать точный план commit.
4. Получить подтверждение пользователя.
5. Реализовать только подтвержденный scope.
6. Запустить релевантные проверки.
7. Проверить `git diff --check`.
8. Проверить `git status`.
9. Создать один commit с английским commit message.
10. Push в `main`.
11. Проверить GitHub CI или зафиксировать, что workflows отсутствуют.
12. Если запускался Docker stack, остановить его без удаления volumes.
13. Остановиться.

### 1.3 Язык и стиль

- Product UI, уроки, подсказки, feedback и документация для пользователя пишутся на русском.
- Code identifiers, API paths, branch names, commit messages, commands и technical package names пишутся на английском.
- Интерфейс остается черно-белым, плотным, современным, без копирования Duolingo brand, mascot, colors или visual identity.
- GoPath использует только общую механику progression-based learning: lessons, XP, streak, checkpoints, career map.

### 1.4 Localhost first

До deploy запрещено начинать production hosting work, если не закрыт `Localhost Final Acceptance`.

Под production-ready на localhost понимается:

- приложение полностью запускается локально;
- все основные пользовательские сценарии работают end-to-end;
- данные сохраняются в PostgreSQL;
- Redis используется там, где он нужен для очередей, rate limits или cache;
- auth и sessions работают;
- lesson flow работает;
- challenge flow работает;
- XP, streak, daily goal, progress и checkpoint gating работают;
- sandbox безопаснее простого prototype runner;
- тесты и CI закрывают основные риски;
- документация описывает локальный запуск, проверки и troubleshooting.

Deploy начинается только после отдельного docs/decision step, который явно подтверждает закрытие localhost gate.

### 1.5 Изменение roadmap

- Если порядок работ меняется, сначала делается отдельный docs commit.
- Если новая тема появляется в процессе, она добавляется в roadmap отдельным маленьким docs commit.
- Если пункт roadmap оказался слишком большим, он дробится отдельным docs commit перед реализацией.
- Этот документ не является замороженным планом навсегда, но любые изменения должны быть видимыми в git history.

## 2. Текущее состояние проекта

### 2.1 Уже закрытые commits

1. `docs: add project bootstrap specification`
   - Добавлены начальные project docs, bootstrap plan и technical specification.

2. `build: add frontend skeleton`
   - Создан `frontend/` на Vite, React, TypeScript и Tailwind CSS.

3. `build: add backend skeleton`
   - Создан `backend/` Go module.
   - Добавлен HTTP server и `GET /healthz`.

4. `docs: add architecture notes`
   - Добавлен `docs/architecture.md` с начальными архитектурными решениями.

5. `ui: add static mvp shell`
   - Добавлен SPA shell.
   - Добавлены routes `/app`, `/learn`, `/lesson/:lessonId`, `/challenge/:challengeId`, `/profile`.

6. `ui: add dashboard mock screen`
   - Добавлен статичный Dashboard mock screen.

7. `ui: add learning path mock screen`
   - Добавлен статичный Learning Path mock screen.

8. `ui: add lesson mock screen`
   - Добавлен статичный Lesson mock screen.

9. `ui: add challenge mock screen`
   - Добавлен статичный Code Challenge mock screen.

10. `ui: add profile mock screen`
    - Добавлен статичный Profile mock screen.

11. `api: add core learning mock endpoints`
    - Добавлены read-only endpoints `/api/v1/tracks`, `/api/v1/levels`, `/api/v1/lessons/{id}`, `/api/v1/progress`, `/api/v1/profile`.

12. `infra: add local postgres and redis`
    - Добавлен локальный Docker Compose stack для PostgreSQL и Redis.

13. `api: add initial database schema`
    - Добавлена первая SQL schema migration.

14. `infra: add migration runner script`
    - Добавлен `scripts/migrate-up.sh`.
    - Добавлена краткая infra инструкция.

15. `api: add database connection layer`
    - Добавлен `internal/database`.
    - Добавлен PostgreSQL driver через `pgx` stdlib.

16. `api: add learning service boundary`
    - Добавлен `internal/learning`.
    - `MemoryStore` вынесен из `httpserver`.

17. `api: add learning seed migration`
    - Добавлены demo seed data.
    - Migration script научен применять все SQL migrations по порядку.

18. `api: add postgres learning store`
    - Добавлен read-only `PostgresStore` для `learning.Store`.

19. `api: wire optional postgres learning store`
    - Backend выбирает `MemoryStore` по умолчанию.
    - Если задан `GOPATH_DATABASE_URL`, используется PostgreSQL store.

20. `api: add local challenge runner prototype`
    - Добавлены `POST /api/v1/challenges/{id}/run` и `/submit`.
    - Добавлен local Go test runner prototype.

### 2.2 Что уже есть

- Monorepo layout: `frontend/`, `backend/`, `infra/`, `docs/`, `scripts/`.
- Frontend SPA shell и 5 основных mock screens.
- Backend HTTP API на Go.
- Health endpoint.
- Read-only learning API.
- PostgreSQL schema и demo seed migration.
- Optional PostgreSQL learning store.
- Local `MemoryStore` default.
- Local Go challenge runner prototype через `go test` в temp directory.
- Docker Compose config для PostgreSQL и Redis.

### 2.3 Что пока mock/prototype

- Frontend screens используют local mock data.
- Frontend не делает backend API calls.
- Learning content частично hardcoded.
- `PostgresStore` использует temporary static defaults для полей, которых еще нет в schema.
- Challenge runner является prototype и не является production sandbox.
- Attempts не сохраняются через API.
- XP, streak, daily goal и unlock logic не обновляются реальными domain rules.

### 2.4 Чего пока нет

- GitHub Actions workflows.
- Root developer commands.
- Полная local startup инструкция.
- Auth/register/login/logout.
- Cookie sessions.
- `GET /api/v1/me`.
- Attempts persistence.
- Frontend API integration.
- Real lesson submission flow.
- Real challenge persistence flow.
- Redis job queue.
- Docker-based sandbox isolation.
- Content authoring system.
- Admin/content tooling.
- Observability.
- Security hardening.
- Deploy.

### 2.5 Устаревшие источники

- `README.md` частично описывает старый bootstrap state.
- `docs/architecture.md` частично описывает старую backend/frontend границу.
- `docs/bootstrap-plan.md` остается историческим bootstrap plan и не должен использоваться как текущий delivery roadmap.

## 3. Block 1 - Complete Production-Ready Localhost App

Этот блок является главным рабочим блоком. Он должен быть закрыт полностью до начала deploy.

## Phase A - Project Control and Documentation

### Goal

Сделать документацию актуальной системой управления проектом, чтобы дальнейшие commits шли по понятному плану.

### Why

Сейчас проект уже ушел дальше bootstrap, но часть документов осталась в старом состоянии. Без актуального roadmap легко начать делать технические commits без общей картины.

### Commits

1. `docs: add full delivery roadmap`
   - Добавить этот roadmap.
   - Зафиксировать localhost-first подход.
   - Добавить roadmap в `README.md` Sources of Truth.

2. `docs: refresh architecture status`
   - Обновить `docs/architecture.md` под фактическое состояние.
   - Зафиксировать текущие frontend/backend/infra boundaries.
   - Отметить `MemoryStore`, optional PostgreSQL store и challenge runner prototype.

3. `docs: refresh root readme status`
   - Обновить root `README.md`.
   - Убрать устаревшую фразу про repository containing only bootstrap docs.
   - Добавить краткую current status секцию.

4. `docs: add local development guide`
   - Добавить инструкцию локального запуска frontend/backend/infra.
   - Описать env files.
   - Описать migrate-up.
   - Описать как остановить Docker stack без удаления volumes.

5. `docs: add api contract notes`
   - Описать текущие `/api/v1` endpoints.
   - Описать JSON error format.
   - Отметить temporary/demo fields.

6. `docs: add content authoring guide`
   - Описать будущую структуру learning content.
   - Зафиксировать language rules: Russian UI/content, English code/tech terms.
   - Описать требования к lesson, options, challenge tests и hints.

### Includes

- Только документация.
- Обновление источников истины.
- Уточнение текущего состояния и порядка работ.

### Excludes

- Runtime code changes.
- API changes.
- Schema changes.
- Docker runtime changes.
- CI workflows.

### Checks

- `git diff --check`.
- `git status`.
- Для docs-only commits `go test` и `npm run build` не обязательны, если code не менялся.

### Closure Criteria

- Все основные документы не противоречат текущему состоянию проекта.
- Новый roadmap явно является delivery source of truth.
- Пользователь может понять следующий commit без чтения всей истории чата.

## Phase B - Local Developer Experience

### Goal

Сделать локальную разработку повторяемой: один набор команд для проверки, запуска и smoke validation.

### Why

Пока команды разбросаны между frontend/backend/infra. Перед активной интеграцией нужен предсказуемый local workflow.

### Commits

1. `build: add root developer commands`
   - Добавить root `Makefile` или scripts wrapper.
   - Команды: backend test, frontend lint, frontend build, compose config, diff check.
   - Не запускать Docker stack автоматически.

2. `build: add backend verify command`
   - Добавить backend-focused verify command.
   - Использовать `GOCACHE=/private/tmp/gopath-gocache` в документации/Makefile, если нужно.
   - Не менять backend logic.

3. `build: add frontend verify command`
   - Добавить frontend-focused verify command.
   - Запускать `npm run lint` и `npm run build`.
   - Не менять UI.

4. `infra: add local stack helper commands`
   - Добавить команды для `docker compose up -d`, `down`, `config`.
   - `down` без `--volumes`.
   - Не запускать stack в commit автоматически.

5. `test: add api smoke script`
   - Добавить script, который проверяет `GET /healthz` и read-only learning endpoints.
   - Script предполагает запущенный backend.
   - Не добавлять auth/challenge smoke пока flow не стабилизирован.

6. `docs: document local command workflow`
   - Описать последовательность: install deps, compose config, stack up, migrate, backend start, frontend start, smoke.

### Includes

- Root commands/scripts.
- Документация команд.
- Non-destructive local workflow.

### Excludes

- CI.
- Production Dockerfiles.
- Deploy.
- Изменения product behavior.

### Checks

- `sh -n` для shell scripts.
- `make` target dry checks, если Makefile добавлен.
- `docker compose --env-file infra/.env.example -f infra/compose.yml config`.
- `go test ./...` внутри `backend/`.
- `npm run lint` и `npm run build` внутри `frontend/`.
- `git diff --check`.

### Closure Criteria

- Новый разработчик может запустить и проверить проект локально по документации.
- Есть единые команды для частых проверок.
- Docker volumes не удаляются командами по умолчанию.

## Phase C - CI Foundation

### Goal

Добавить автоматические проверки для каждого push.

### Why

Сейчас `.github/workflows` отсутствует, поэтому ошибки можно заметить только локально. Перед дальнейшей интеграцией нужен CI baseline.

### Commits

1. `ci: add backend test workflow`
   - GitHub Actions для Go tests.
   - Использовать Go toolchain version из `backend/go.mod`.
   - Не запускать PostgreSQL integration tests, если они еще не добавлены.

2. `ci: add frontend build workflow`
   - GitHub Actions для `npm ci`, `npm run lint`, `npm run build` в `frontend/`.
   - Использовать Node version, совместимую с Vite/package lock.

3. `ci: add infrastructure validation workflow`
   - Проверять Docker Compose config.
   - Проверять shell syntax для scripts.
   - Не поднимать долгоживущий Docker stack.

4. `ci: add migration validation check`
   - Добавить lightweight SQL/script validation.
   - Не применять migrations к production database.
   - Если нужен PostgreSQL service container, добавить отдельным commit.

5. `docs: add ci status to readme`
   - Добавить badges/описание только после появления workflows.

### Includes

- `.github/workflows`.
- CI checks для уже существующих слоев.

### Excludes

- Deploy workflows.
- Secrets.
- Release automation.
- Long-running e2e.

### Checks

- Локально проверить workflow YAML syntax доступными средствами.
- После push проверить GitHub Actions result.
- `git diff --check`.

### Closure Criteria

- Каждый push проверяет backend, frontend и infra baseline.
- CI failures видны сразу.
- CI не требует production secrets.

## Phase D - Backend API Contract

### Goal

Стабилизировать API conventions перед frontend integration и write flows.

### Why

Read-only endpoints уже есть, но error format, validation и response contracts пока минимальны. Перед подключением frontend лучше зафиксировать правила.

### Commits

1. `api: standardize error responses`
   - Ввести единый JSON error shape.
   - Сохранить существующие paths.
   - Обновить tests.

2. `api: add request validation helpers`
   - Helper для JSON decode, empty body, unknown fields при необходимости.
   - Применить только к challenge endpoints или одному маленькому endpoint.

3. `api: document learning response contracts`
   - Зафиксировать fields для tracks/levels/lesson/progress/profile.
   - Docs-only или tests-only, без расширения behavior.

4. `api: stabilize challenge response contract`
   - Зафиксировать statuses: `passed`, `failed`, `compile_error`, `timeout`, `internal_error`.
   - Подготовить место для `queued` и `running`, но не включать async behavior.

5. `api: add not found contract tests`
   - Покрыть unknown lesson/challenge cases.
   - Не менять frontend.

6. `api: prepare auth contract notes`
   - Docs-only contract для register/login/logout/me.
   - Реализация auth идет позже.

### Includes

- API conventions.
- Handler tests.
- Contract documentation.

### Excludes

- Auth implementation.
- Database writes.
- Frontend integration.
- Async sandbox.

### Checks

- `go fmt ./...`.
- `go test ./...`.
- `git diff --check`.

### Closure Criteria

- API errors and DTOs достаточно стабильны для frontend API client.
- Backend tests фиксируют основные contract cases.

## Phase E - Database and Persistence

### Goal

Перевести core product state с mock/prototype подхода на проверяемую PostgreSQL persistence model.

### Why

Без persistence нельзя сделать auth, attempts, XP, streak и реальные progress updates.

### Commits

1. `api: add migration tracking table`
   - Добавить schema migration для tracking applied migrations, если выбираем custom runner.
   - Или подготовить переход на migration tool отдельным решением.
   - Не менять runtime API.

2. `infra: make migrate script idempotent`
   - Обновить `scripts/migrate-up.sh`, чтобы повторный запуск был безопасным.
   - Не делать destructive reset.

3. `api: add database test helpers`
   - Добавить test utilities для integration tests.
   - Без обязательного запуска Docker stack в обычных unit tests.

4. `api: add users repository`
   - Repository для users.
   - Tests на create/find by email/find by id.
   - Не добавлять auth handlers.

5. `api: add sessions schema`
   - Migration для sessions.
   - TTL/expiration fields.
   - Не добавлять middleware.

6. `api: add sessions repository`
   - Create/find/delete session.
   - Tests.
   - Не добавлять HTTP auth.

7. `api: add attempts repository`
   - Repository для attempts table.
   - Сохранение lesson/challenge attempts.
   - Не подключать handlers.

8. `api: add progress repository`
   - Repository для user_progress и skill_progress.
   - Read/update XP, streak, completed lessons/checkpoints.
   - Не менять domain rules пока.

9. `api: remove postgres learning static defaults`
   - Добавить missing schema/content fields маленькими migrations.
   - Убрать fallback на `mockProfile`/`mockProgress` там, где данные уже есть в DB.

10. `api: add database integration tests`
    - Tests for migrations and repositories.
    - Использовать отдельный integration tag, если запуск требует PostgreSQL.

### Includes

- Migrations.
- Repository boundaries.
- Integration test foundation.

### Excludes

- Auth HTTP flows.
- Frontend integration.
- Domain XP/streak rules.
- Deploy database.

### Checks

- `go fmt ./...`.
- `go test ./...`.
- Integration tests with local PostgreSQL only when stack is intentionally started.
- `docker compose ... config`.
- `git diff --check`.

### Closure Criteria

- Core data can be stored and read through repositories.
- Migrations can be applied repeatedly in local development without data loss.
- Temporary DB fallbacks are reduced or documented.

## Phase F - Auth and Sessions

### Goal

Добавить user registration, login, logout и cookie-based session auth для web app.

### Why

Progress, attempts, XP и streak должны принадлежать конкретному пользователю. Demo `demo-user` не подходит для полноценного продукта.

### Commits

1. `api: add password hashing`
   - Добавить password hash/verify helpers.
   - Tests for valid/invalid password.
   - Не добавлять HTTP handlers.

2. `api: add auth service boundary`
   - Service для register/login/logout/me logic.
   - Использовать users/sessions repositories.
   - Не менять routes в этом commit.

3. `api: add register endpoint`
   - `POST /api/v1/auth/register`.
   - Validate email, handle, password.
   - Create user and session.

4. `api: add login endpoint`
   - `POST /api/v1/auth/login`.
   - Verify credentials.
   - Create session cookie.

5. `api: add logout endpoint`
   - `POST /api/v1/auth/logout`.
   - Delete current session.
   - Clear cookie.

6. `api: add me endpoint`
   - `GET /api/v1/me`.
   - Return authenticated user summary.

7. `api: add auth middleware`
   - Require session for user-specific endpoints.
   - Keep health endpoint public.

8. `api: protect progress and profile endpoints`
   - Replace hardcoded `demo-user`.
   - Use authenticated user id.

9. `docs: document local auth flow`
   - Document register/login/logout/me for local development.

### Includes

- Password hashing.
- Cookie sessions.
- Auth endpoints.
- Auth middleware.
- User-specific progress/profile.

### Excludes

- OAuth.
- Email verification.
- Password reset.
- Admin UI.
- Production secrets manager.

### Checks

- `go fmt ./...`.
- `go test ./...`.
- Auth handler tests with cookies.
- `git diff --check`.

### Closure Criteria

- A local user can register, login, call `GET /api/v1/me`, view own profile/progress, and logout.
- User-specific endpoints no longer depend on hardcoded `demo-user`.

## Phase G - Frontend API Integration

### Goal

Подключить frontend screens к backend API вместо primary mock data.

### Why

Сейчас UI выглядит как продукт, но не работает с backend. Нужно перевести mock screens в real localhost app.

### Commits

1. `ui: add frontend api config`
   - Add env-based API base URL.
   - Document local Vite env.
   - No backend calls yet.

2. `ui: add api client foundation`
   - Add typed fetch wrapper.
   - Handle JSON errors.
   - No React Query yet unless needed in this step.

3. `ui: add learning api types`
   - Add TypeScript DTOs for tracks, levels, lesson, progress, profile.
   - Keep mapping separate from UI components.

4. `ui: fetch dashboard summary`
   - Dashboard reads progress/profile/lesson data from API.
   - Keep fallback minimal for loading/error only.

5. `ui: fetch learning path`
   - Learning Path reads tracks/levels from API.
   - Preserve current visual layout.

6. `ui: fetch lesson data`
   - Lesson page reads `/api/v1/lessons/:id`.
   - No answer submission yet.

7. `ui: fetch profile data`
   - Profile page reads `/api/v1/profile`.
   - Preserve skill graph mock rendering if API provides skill values.

8. `ui: add app loading states`
   - Shared loading state components.
   - Avoid layout jumps.

9. `ui: add app error states`
   - Shared API error displays.
   - Russian user-facing messages.

10. `ui: remove primary shell mock dependency`
    - Keep static mock only for development examples if still useful.
    - Main routes use API as source.

### Includes

- Frontend API client.
- Typed DTOs.
- Loading/error states.
- Read-only integration.

### Excludes

- Lesson answer submission.
- Challenge run/submit integration.
- Auth UI, unless auth phase requires it first.
- Zustand/React Query unless introduced by explicit commit.

### Checks

- `npm run lint`.
- `npm run build`.
- Backend tests if API contract is touched.
- Manual local smoke after backend is running.
- `git diff --check`.

### Closure Criteria

- Dashboard, Learning Path, Lesson and Profile read real backend data on localhost.
- Mock data is no longer the main data source for read-only screens.

## Phase H - Lesson Flow

### Goal

Сделать обычный lesson flow end-to-end: выбрать ответ, проверить, получить feedback, сохранить attempt, обновить progress.

### Why

MVP считается готовым только если пользователь может пройти обычный урок, а XP/streak/progress реально меняются.

### Commits

1. `api: add lesson attempt request contract`
   - Define request/response for lesson answer submission.
   - Include selected option id and lesson id.
   - No persistence yet.

2. `api: add lesson answer evaluation`
   - Service checks correct option.
   - Returns correctness, feedback, explanation and XP preview.

3. `api: persist lesson attempts`
   - Store attempt in `attempts`.
   - Use authenticated user id.

4. `api: add lesson xp awarding`
   - Award XP for first correct completion according to simple rule.
   - Prevent duplicate XP for repeated correct answer.

5. `api: update lesson progress`
   - Increment completed lessons when appropriate.
   - Update current lesson id.

6. `api: update streak for lesson activity`
   - Basic daily activity rule.
   - Tests for same day and next day.

7. `api: update daily goal for lesson activity`
   - Track earned XP and completed tasks for current day.
   - Add schema if needed.

8. `ui: submit lesson answer`
   - Lesson page calls backend.
   - Button states: idle, checking, done.

9. `ui: show lesson feedback from api`
   - Render correct/incorrect feedback.
   - Use Russian copy from API or mapped UI copy.

10. `ui: add next lesson transition`
    - Continue button moves to next lesson or learning path.

11. `test: add lesson flow smoke test`
    - Cover API or e2e path for answer submission and progress update.

### Includes

- Lesson submission endpoint.
- Attempt persistence.
- XP/progress/streak updates.
- Frontend interaction.

### Excludes

- Multiple answer types beyond current multiple choice.
- Full content authoring system.
- Challenge submissions.
- Complex anti-cheat.

### Checks

- `go test ./...`.
- `npm run lint`.
- `npm run build`.
- Local smoke with running backend.
- `git diff --check`.

### Closure Criteria

- A logged-in local user can complete a lesson and see progress change.
- Duplicate submissions do not incorrectly farm XP.
- Tests cover correct and incorrect answers.

## Phase I - Code Challenge Flow

### Goal

Сделать challenge flow end-to-end на localhost: run public tests, submit public+hidden tests, сохранить attempts, обновить XP/progress.

### Why

Code challenges являются ключевой частью GoPath. Текущий runner работает только как API prototype без persistence и frontend integration.

### Commits

1. `api: add challenge catalog store boundary`
   - Separate challenge metadata/tests from runner catalog.
   - Prepare DB-backed challenge source.

2. `api: load challenge tests from database`
   - Runner receives tests from store instead of hardcoded catalog.
   - Keep `retry-context` behavior.

3. `api: persist challenge run attempts`
   - Save `run` results without XP award.
   - Store status/stdout/stderr/tests/duration.

4. `api: persist challenge submit attempts`
   - Save submit results.
   - Include public+hidden test count.

5. `api: award xp for successful challenge submit`
   - Award XP once per completed challenge.
   - Prevent repeated XP farming.

6. `api: update progress after challenge submit`
   - Update practice tasks, checkpoint completion if applicable.

7. `api: add challenge attempt history endpoint`
   - Return recent attempts for current user and challenge.

8. `ui: connect challenge run action`
   - `Run` button calls `/run`.
   - Show running state and result.

9. `ui: connect challenge submit action`
   - `Submit` button calls `/submit`.
   - Show hidden tests summary only after submit.

10. `ui: render terminal output from api`
    - Display stdout/stderr in terminal panel.
    - Preserve developer-tool styling.

11. `ui: render challenge attempt history`
    - Show persisted attempts.
    - No local mock history as primary source.

12. `test: add challenge flow smoke test`
    - Cover run, failed submit, passed submit, XP update.

### Includes

- DB-backed challenge metadata/tests.
- Attempt persistence.
- XP/progress update.
- Frontend run/submit integration.

### Excludes

- Production-grade sandbox isolation.
- Redis async queue.
- Monaco/xterm if not yet in UX phase.
- Multi-language challenges.

### Checks

- `go test ./...`.
- `npm run lint`.
- `npm run build`.
- Local challenge smoke.
- `git diff --check`.

### Closure Criteria

- A logged-in local user can solve `retry-context` end-to-end.
- Run does not award XP.
- Successful submit awards XP once.
- Attempt history is persisted and visible.

## Phase J - Sandbox Hardening for Localhost

### Goal

Заменить unsafe prototype runner на более изолированный local sandbox flow.

### Why

`exec.CommandContext` в temp directory полезен для bootstrap, но не подходит как production-ready sandbox even on localhost.

### Commits

1. `api: split challenge runner interfaces`
   - Separate synchronous prototype runner and sandbox job runner interfaces.
   - Keep existing endpoints behavior.

2. `api: add sandbox job types`
   - Define job id, status, request, result.
   - Include `queued`, `running`, terminal statuses.

3. `api: add redis connection layer`
   - Config from env.
   - Ping on startup only when enabled.
   - Tests without real Redis for config behavior.

4. `api: add sandbox job queue`
   - Enqueue/dequeue job metadata in Redis or minimal queue abstraction.
   - No Docker runner yet.

5. `api: add async challenge run contract`
   - Decide whether endpoints remain sync for MVP or return job ids.
   - Update docs/tests.

6. `sandbox: add docker runner prototype`
   - Run challenge in short-lived Docker container.
   - Use local image or Go image.
   - No production deploy.

7. `sandbox: disable network`
   - Ensure challenge container has no network.
   - Add verification test or smoke script.

8. `sandbox: add cpu memory timeout limits`
   - Enforce CPU, memory and wall-clock limits.
   - Document defaults.

9. `sandbox: add workspace cleanup`
   - Ensure temporary files/containers are removed.
   - Do not delete Docker volumes.

10. `api: add sandbox rate limits`
    - Per-user basic rate limits for run/submit.
    - Use Redis if available.

11. `test: add sandbox safety smoke checks`
    - Timeout case.
    - Network blocked case.
    - Cleanup case.

12. `docs: document local sandbox limitations`
    - Be explicit about what is safe and what remains future production hardening.

### Includes

- Redis-backed or Redis-ready sandbox jobs.
- Docker-based local isolation.
- Limits and cleanup.
- Rate limiting.

### Excludes

- Kubernetes sandbox.
- Remote execution cluster.
- Multi-tenant production hardening beyond local MVP.
- Paid anti-abuse systems.

### Checks

- `go test ./...`.
- Docker Compose config.
- Sandbox smoke scripts only when Docker is intentionally used.
- Stop Docker stack/containers without deleting volumes.
- `git diff --check`.

### Closure Criteria

- Challenge code no longer runs directly in the API process environment.
- Network is blocked for sandbox runs.
- CPU/memory/time limits are enforced locally.
- Failed/timeout runs are cleaned up.

## Phase K - Learning Content System

### Goal

Сделать learning content управляемым через repository files and import/validation pipeline, а не через hardcoded mocks.

### Why

GoPath ценен только при расширяемом качественном контенте. Hardcoded data мешает добавлять уроки и checkpoints.

### Commits

1. `content: add content directory structure`
   - Add `content/` layout for tracks, levels, lessons, challenges.
   - Add README for format.
   - No importer yet.

2. `content: add content schema draft`
   - Define JSON/YAML/Markdown format decision.
   - Document required fields.

3. `content: add content validation script`
   - Validate ids, required fields, language rules and references.
   - No database import.

4. `content: move demo lesson content`
   - Move `go-errors-tests` content into content source.
   - Keep API behavior unchanged.

5. `content: move retry challenge content`
   - Move `retry-context` prompt, starter code, tests and hints.
   - Keep runner behavior.

6. `content: add seed generation script`
   - Generate SQL seed or import data from content files.
   - No runtime admin UI.

7. `api: load seed content from generated source`
   - Make migrations/import use generated content source.
   - Avoid manual duplication.

8. `content: add trainee backend module`
   - Add initial lessons for Go basics, HTTP basics, SQL basics.
   - Keep each lesson group as small commits if large.

9. `content: add junior backend module`
   - Add lessons for errors, testing, context, REST API, PostgreSQL basics.
   - Split into multiple commits.

10. `content: add first checkpoint mini project`
    - Add REST API checkpoint.
    - Include requirements, tests and rubric.

11. `docs: add content review checklist`
    - Technical accuracy.
    - Russian wording.
    - No Duolingo copy.
    - Test coverage.

### Includes

- File-based content source.
- Validation.
- Seed/import pipeline.
- Initial real lesson/challenge content.

### Excludes

- Admin UI.
- User-generated content.
- Full Senior/Middle catalog at once.
- External CMS.

### Checks

- Content validation script.
- SQL generation/import checks if added.
- `go test ./...` if backend import code changes.
- `npm run build` only if frontend content rendering changes.
- `git diff --check`.

### Closure Criteria

- Content can be added through documented files.
- IDs are validated.
- Seed data is generated or imported from content source.
- Hardcoded content is no longer the primary source.

## Phase L - Progression Mechanics

### Goal

Formalize and implement XP, streak, daily goal, unlocks, checkpoints and skill graph updates.

### Why

Progression mechanics are central to GoPath. They must be deterministic, tested and hard to exploit.

### Commits

1. `api: document progression rules`
   - Define XP rules for lessons and challenges.
   - Define streak and daily goal rules.
   - Docs-only.

2. `api: add xp domain service`
   - Pure Go logic for XP award.
   - Unit tests.
   - No HTTP wiring.

3. `api: add streak domain service`
   - Same-day, next-day and missed-day behavior.
   - Unit tests.

4. `api: add daily goal domain service`
   - Earned XP and completed tasks per day.
   - Add schema if needed in separate migration commit.

5. `api: add lesson completion rules`
   - First correct completion logic.
   - Prevent XP farming.

6. `api: add challenge completion rules`
   - Successful submit logic.
   - Prevent XP farming.

7. `api: add unlock rules`
   - Active/completed/locked state based on completed lessons/checkpoints/XP.
   - Tests.

8. `api: add checkpoint gating`
   - Block next level until checkpoint passed.
   - Tests.

9. `api: add skill progress update rules`
   - Map lessons/challenges to skills.
   - Update skill graph.

10. `ui: show real progression states`
    - Learning Path uses backend locked/active/completed state.
    - Dashboard/Profile show backend progress.

### Includes

- Domain services.
- DB support where needed.
- UI rendering of real progression state.

### Excludes

- Complex gamification economy.
- Leaderboards.
- Paid streak freeze mechanics.
- Social mechanics.

### Checks

- `go test ./...`.
- `npm run lint`.
- `npm run build`.
- End-to-end local progression smoke.
- `git diff --check`.

### Closure Criteria

- XP, streak, daily goal and unlock state are backend-driven.
- Checkpoints gate level progression.
- UI reflects real progression.

## Phase M - Product UX Polish

### Goal

Довести UI/UX до состояния завершенного developer education product на localhost.

### Why

Mock screens уже выглядят как продукт, но реальное приложение должно быть responsive, accessible, интерактивным и устойчивым к loading/error states.

### Commits

1. `ui: audit responsive shell layout`
   - Fix layout issues on desktop/tablet/mobile.
   - No new features.

2. `ui: add shared ui primitives`
   - Buttons, panels, progress bars, status labels if useful.
   - Avoid over-abstracting.

3. `ui: improve navigation states`
   - Active, focus, disabled, keyboard navigation.

4. `ui: polish dashboard real data view`
   - Tighten layout after API integration.
   - No new backend fields unless planned.

5. `ui: polish learning path real data view`
   - Ensure locked/active/completed states are clear without color-only cues.

6. `ui: polish lesson interaction`
   - Answer selection states.
   - Feedback panel.
   - Next action.

7. `ui: add monaco editor`
   - Install Monaco dependency.
   - Replace static code editor mock.
   - Keep one route scope: challenge page.

8. `ui: add terminal component`
   - Install xterm.js or equivalent if still chosen.
   - Render run/submit output.

9. `ui: polish challenge workspace`
   - Resizable or stable three-column layout if needed.
   - Run/Submit/result states.

10. `ui: polish profile progress view`
    - Skill graph, readiness, achievements, activity calendar.

11. `ui: add empty states`
    - No attempts.
    - No activity.
    - Locked content.

12. `ui: add error recovery actions`
    - Retry buttons.
    - Clear Russian messages.

13. `test: add visual smoke screenshots`
    - Use Playwright or equivalent.
    - Check key routes render.

### Includes

- UI polish.
- Responsive/accessibility improvements.
- Monaco/xterm integration in separate commits.

### Excludes

- New product modules.
- Marketing landing page.
- Duolingo-like mascot/colors.
- Admin UI.

### Checks

- `npm run lint`.
- `npm run build`.
- Frontend e2e/visual smoke when added.
- `git diff --check`.

### Closure Criteria

- All 5 main screens work with real data.
- UI remains black-and-white developer-tool style.
- Core flows are usable on common screen sizes.
- No obvious text overlap or broken layout.

## Phase N - Testing Strategy

### Goal

Сделать тестирование системным, а не точечным.

### Why

Перед production-ready localhost gate нужны unit, integration и e2e checks для основных рисков.

### Commits

1. `test: document testing strategy`
   - Define which tests run locally and in CI.
   - Define integration test prerequisites.

2. `test: add backend integration test setup`
   - PostgreSQL test setup.
   - Migrations apply in test DB.
   - No destructive local user DB reset.

3. `test: cover auth integration flow`
   - Register/login/me/logout.

4. `test: cover lesson integration flow`
   - Submit answer, attempt saved, XP updated.

5. `test: cover challenge integration flow`
   - Run/submit, attempt saved, XP updated.

6. `test: cover progression integration flow`
   - Streak, daily goal, unlock rules.

7. `test: add frontend route smoke tests`
   - Main routes render.
   - Auth state handling.

8. `test: add e2e login flow`
   - Browser test for login.

9. `test: add e2e lesson flow`
   - Browser test for lesson completion.

10. `test: add e2e challenge flow`
    - Browser test for challenge run/submit.

11. `ci: run integration tests`
    - Add service containers if needed.
    - Keep runtime reasonable.

12. `ci: run e2e smoke tests`
    - Start backend/frontend in CI.
    - Run minimal smoke suite.

### Includes

- Backend unit/integration tests.
- Frontend smoke/e2e tests.
- CI execution for important flows.

### Excludes

- Full load testing.
- Browser matrix beyond what is needed for MVP.
- External production monitoring.

### Checks

- `go test ./...`.
- Integration tests with PostgreSQL.
- `npm run lint`.
- `npm run build`.
- E2E smoke.
- CI green.
- `git diff --check`.

### Closure Criteria

- Key user journeys are covered by automated tests.
- CI runs enough checks to catch major regressions.
- Local test docs are accurate.

## Phase O - Admin and Content Operations

### Goal

Добавить минимальные инструменты для управления контентом и подготовки будущей админки.

### Why

Обучающий продукт требует постоянного обновления content. Даже если full admin UI не нужен для MVP, должны быть clear operations.

### Commits

1. `api: add admin role checks`
   - Use existing user role.
   - Middleware/helper only.

2. `content: add content preview command`
   - Render/print lesson/challenge summary from content source.

3. `content: add content import dry run`
   - Validate and show changes without applying.

4. `content: add content import command`
   - Apply content to local DB.
   - Non-destructive by default.

5. `api: add admin content read endpoints`
   - Read-only admin endpoints for content overview.
   - Protected by admin role.

6. `ui: add minimal admin content placeholder`
   - Only if needed for local operations.
   - Keep separate from learner UX.

7. `docs: add content operations guide`
   - How to add lesson.
   - How to validate.
   - How to import.
   - How to rollback manually.

### Includes

- Admin role foundation.
- Content validation/import operations.
- Documentation.

### Excludes

- Full CMS.
- Public teacher accounts.
- Rich text editor.
- Multi-author workflow.

### Checks

- Content validation.
- `go test ./...`.
- `npm run build` if UI changes.
- `git diff --check`.

### Closure Criteria

- Content can be validated and imported locally.
- Admin-only operations are protected.
- Content changes do not require manual DB edits.

## Phase P - Local Observability and Reliability

### Goal

Добавить локальную observability и reliability baseline.

### Why

Production-ready localhost app должен быть диагностируемым: logs, request IDs, health/readiness, metrics and clear failure modes.

### Commits

1. `api: add structured logging`
   - Use standard library or minimal dependency decision.
   - Include method/path/status/duration.

2. `api: add request id middleware`
   - Generate or pass request id.
   - Include in logs and responses if appropriate.

3. `api: add readiness endpoint`
   - Check dependencies when configured.
   - Keep `/healthz` simple.

4. `api: add graceful shutdown`
   - Handle OS signals.
   - Shutdown HTTP server with timeout.

5. `api: add basic metrics endpoint`
   - Decide on Prometheus dependency or simple local metrics.
   - Keep scope minimal.

6. `api: add sandbox metrics`
   - Runs, failures, timeouts, duration.

7. `infra: add optional local monitoring notes`
   - Docs first or optional compose profile later.
   - Do not make monitoring required for basic local startup.

8. `docs: add troubleshooting guide`
   - Database connection issues.
   - Redis issues.
   - Sandbox Docker issues.
   - Frontend API connection issues.

### Includes

- Logs.
- Request IDs.
- Health/readiness.
- Graceful shutdown.
- Local metrics/troubleshooting.

### Excludes

- Production alerting.
- External SaaS monitoring.
- Full Grafana dashboards unless explicitly added later.

### Checks

- `go test ./...`.
- Local smoke for health/readiness.
- `docker compose ... config` if infra changes.
- `git diff --check`.

### Closure Criteria

- Local failures can be diagnosed from logs and endpoints.
- App shuts down cleanly.
- Dependency health is visible.

## Phase Q - Local Security Hardening

### Goal

Закрыть базовые security risks до localhost final acceptance.

### Why

Даже до deploy приложение должно проектироваться как production-ready: auth, input limits, CSRF, rate limits, sandbox abuse controls.

### Commits

1. `docs: add security baseline`
   - Define local/prod security expectations.
   - Document known limitations.

2. `api: add request size limits`
   - Limit JSON body size.
   - Especially challenge code submissions.

3. `api: add auth rate limits`
   - Limit login/register attempts.
   - Redis-backed if available.

4. `api: add challenge rate limits`
   - Limit run/submit frequency.
   - Per user/session.

5. `api: add csrf protection strategy`
   - Choose CSRF approach for cookie sessions.
   - Implement in small follow-up commit.

6. `api: harden session cookies`
   - HttpOnly, SameSite, Secure behavior depending on local/prod mode.

7. `api: add security headers`
   - Reasonable headers for local app.
   - Avoid breaking Vite dev flow.

8. `sandbox: add forbidden pattern guardrails`
   - Basic pre-run checks only if useful.
   - Not a replacement for sandbox isolation.

9. `build: add dependency audit command`
   - `npm audit` or equivalent decision.
   - Go vuln check if available.

10. `docs: add secrets handling guide`
    - Env vars.
    - `.env` rules.
    - What never gets committed.

### Includes

- Input limits.
- Rate limits.
- CSRF/session hardening.
- Security docs.

### Excludes

- Enterprise SSO.
- Full threat modeling workshop.
- Production secrets manager integration.

### Checks

- `go test ./...`.
- `npm run build` if frontend auth/security changes.
- Security-related handler tests.
- `git diff --check`.

### Closure Criteria

- Obvious abuse paths are limited.
- Cookie auth has CSRF/session protections.
- Secrets and env handling are documented.

## Phase R - Localhost Final Acceptance

### Goal

Prove that GoPath is complete and production-ready on localhost before deploy work starts.

### Why

This is the gate that prevents premature deploy. Deploy should not begin while core product flows are unfinished locally.

### Commits

1. `docs: add localhost acceptance checklist`
   - Define exact checklist.
   - Include commands and expected results.

2. `test: add full local smoke script`
   - Automated or semi-automated full journey check.
   - Does not delete volumes.

3. `test: verify auth local journey`
   - Register/login/me/logout.

4. `test: verify learning local journey`
   - Dashboard -> Lesson -> answer -> progress.

5. `test: verify challenge local journey`
   - Challenge run -> submit -> attempts -> XP.

6. `test: verify checkpoint gating`
   - Locked/active/completed states.

7. `test: verify content import journey`
   - Validate/import content locally.

8. `test: verify local sandbox constraints`
   - Timeout, network block, cleanup.

9. `docs: close localhost readiness report`
   - Summarize what is complete.
   - List known limitations.
   - Explicitly approve moving to deploy block only after user review.

### Includes

- Acceptance checklist.
- Full local smoke validation.
- Final localhost readiness report.

### Excludes

- Actual deploy.
- Production infrastructure.
- Domain/TLS.

### Checks

- Full local verify command.
- Backend tests.
- Frontend tests/build.
- Integration/e2e tests.
- CI green.
- Docker stack stopped without deleting volumes after checks.
- `git diff --check`.

### Closure Criteria

- User can run GoPath locally end-to-end.
- All MVP criteria from `docs/gopath-technical-spec.md` are satisfied locally.
- Tests and CI pass.
- User explicitly approves moving to deploy block.

## 4. Block 2 - Deploy After Localhost Is Complete

Deploy starts only after Phase R is closed. Until then, deploy work is out of scope.

## Phase S - Deployment Decision

### Goal

Выбрать deploy target and production architecture.

### Why

Deploy choices affect Dockerfiles, env config, database/Redis provisioning, migration strategy and monitoring.

### Commits

1. `docs: add deployment options analysis`
   - Compare realistic hosting targets.
   - Include cost, complexity and fit for GoPath.

2. `docs: choose deployment target`
   - Record decision.
   - Define staging and production environments.

3. `docs: define production data services`
   - Managed PostgreSQL strategy.
   - Managed Redis strategy.
   - Backup expectations.

4. `docs: define deployment release process`
   - How code reaches staging and production.
   - Who runs migrations.
   - Rollback expectations.

### Includes

- Decision docs.
- Environment strategy.
- Release process.

### Excludes

- Dockerfile implementation.
- Actual cloud resources.
- Secrets creation.

### Checks

- `git diff --check`.
- Docs review.

### Closure Criteria

- Deployment target is chosen.
- Staging/production strategy is documented.
- Implementation can proceed without guessing.

## Phase T - Containerization and Release Packaging

### Goal

Package frontend and backend for production deployment.

### Why

Local dev server setup is not production packaging. We need repeatable artifacts.

### Commits

1. `build: add backend dockerfile`
   - Multi-stage build if appropriate.
   - Non-root runtime if possible.
   - Healthcheck or docs.

2. `build: add frontend production build packaging`
   - Static build strategy.
   - Decide whether backend serves static files or separate hosting.

3. `build: add production compose example`
   - Example only, if useful.
   - No secrets committed.

4. `build: add production env example`
   - Required env vars.
   - Safe placeholders.

5. `infra: add migration deploy command`
   - Production-safe migration invocation.
   - No destructive reset.

6. `docs: add release packaging guide`
   - Build, tag, run, rollback basics.

### Includes

- Dockerfiles/build packaging.
- Production env examples.
- Migration deploy command.

### Excludes

- Actual cloud deployment.
- Domain/TLS.
- Monitoring setup.

### Checks

- Docker build.
- Backend tests.
- Frontend build.
- Migration command dry/syntax checks.
- `git diff --check`.

### Closure Criteria

- App can be built as production artifacts.
- Required env vars are documented.
- Migration process is safe and repeatable.

## Phase U - Staging Deploy

### Goal

Deploy GoPath to staging before production.

### Why

Staging catches environment-specific problems before public production.

### Commits

1. `infra: add staging deployment config`
   - Platform-specific config.
   - No production secrets.

2. `ci: add staging deployment workflow`
   - Manual trigger or protected branch strategy.
   - No automatic production deploy.

3. `infra: add staging migration workflow`
   - Apply migrations safely to staging.

4. `test: add staging smoke check`
   - Health, auth, lesson, challenge minimal checks.

5. `docs: add staging runbook`
   - How to deploy.
   - How to inspect logs.
   - How to rollback.

### Includes

- Staging environment.
- Staging deploy workflow.
- Staging smoke checks.

### Excludes

- Production deploy.
- Public launch.
- Load testing beyond smoke.

### Checks

- CI staging workflow result.
- Staging smoke test.
- `git diff --check`.

### Closure Criteria

- Staging URL works.
- Critical flows pass on staging.
- Rollback path is documented.

## Phase V - Production Deploy

### Goal

Deploy GoPath to production.

### Why

After localhost and staging are complete, production deploy makes the product available outside local development.

### Commits

1. `infra: add production deployment config`
   - Platform-specific production config.
   - No secrets committed.

2. `ci: add production deployment workflow`
   - Manual approval.
   - Protected environment.

3. `infra: add production migration workflow`
   - Safe migration process.
   - Backup/rollback note.

4. `infra: configure domain and tls`
   - Domain/TLS config or docs depending on platform.

5. `test: add production smoke check`
   - Health, auth, read-only app, core flow minimal checks.

6. `docs: add production runbook`
   - Deploy.
   - Rollback.
   - Logs.
   - Migrations.
   - Incident response.

### Includes

- Production deployment config.
- Production workflow.
- Domain/TLS.
- Production smoke.

### Excludes

- New product features.
- Major architecture changes.
- Unplanned database rewrites.

### Checks

- Production deploy workflow result.
- Production smoke test.
- Monitoring health.
- `git diff --check`.

### Closure Criteria

- Public production URL works.
- Critical flows pass.
- Rollback process is documented.
- User confirms production deploy success.

## Phase W - Production Observability

### Goal

Make production operation visible and actionable.

### Why

Production is not complete if failures cannot be detected or diagnosed.

### Commits

1. `infra: add production log access docs`
   - Where logs are.
   - How to filter by request id.

2. `infra: add production metrics setup`
   - Platform-specific metrics.
   - API/sandbox/database indicators.

3. `infra: add production alerts`
   - App down.
   - High error rate.
   - Sandbox failures/timeouts.

4. `docs: add incident checklist`
   - Triage.
   - Rollback.
   - Communication.

5. `docs: add operational review checklist`
   - Weekly checks.
   - Dependency updates.
   - Backup restore checks.

### Includes

- Production logs.
- Metrics.
- Alerts.
- Incident docs.

### Excludes

- Complex SRE platform.
- On-call scheduling system.
- Enterprise compliance.

### Checks

- Alert test if possible.
- Smoke test after observability changes.
- `git diff --check`.

### Closure Criteria

- Production health is visible.
- Critical failures alert.
- There is a clear incident response path.

## Phase X - Production Acceptance

### Goal

Close the roadmap by confirming GoPath is complete as a production MVP.

### Why

The project should end with a clear acceptance point, not an endless list of improvements.

### Commits

1. `docs: add production acceptance checklist`
   - Define final checklist.
   - Map back to technical specification.

2. `test: run production acceptance smoke`
   - Record commands and results.
   - Do not include secrets.

3. `docs: add mvp completion report`
   - What is complete.
   - Known limitations.
   - Future post-MVP ideas.

4. `docs: mark delivery roadmap complete`
   - Mark roadmap closed for MVP.
   - Add next-phase note only if user wants post-MVP roadmap.

### Includes

- Final verification.
- Completion report.
- Roadmap closure.

### Excludes

- New features.
- Scope expansion.
- Post-MVP roadmap unless explicitly requested.

### Checks

- Production smoke.
- CI green.
- Docs review.
- `git diff --check`.

### Closure Criteria

- Technical specification MVP criteria are satisfied.
- Production deploy is healthy.
- User confirms project is complete for MVP.

## 5. Global Future Topics

These topics are not separate phases unless the roadmap is updated. They should be inserted as small commits only when the current phase needs them.

- React Query adoption.
- Zustand adoption.
- Monaco Editor.
- xterm.js.
- Redis cache beyond sandbox/rate limits.
- Advanced PostgreSQL indexing and query tuning.
- More learning tracks.
- Middle Backend content.
- Senior Backend content.
- Community features.
- Leaderboards.
- Payments or subscriptions.
- Mobile-specific UI.
- Internationalization beyond Russian.

## 6. Final Definition of Done

GoPath is considered complete for this roadmap only when all of these are true:

- `Localhost Final Acceptance` is closed.
- Deploy phases are closed after localhost completion.
- Production URL is available.
- Main user journey works:
  - register;
  - login;
  - open dashboard;
  - open learning path;
  - complete a lesson;
  - run and submit a challenge;
  - see XP/streak/progress changes;
  - see profile progress;
  - logout.
- PostgreSQL stores users, sessions, content, attempts and progress.
- Redis is used for the chosen sandbox/rate-limit/cache responsibilities.
- Sandbox has local isolation, limits and cleanup.
- CI is green.
- Documentation explains local and production operation.
- User explicitly confirms roadmap completion.
