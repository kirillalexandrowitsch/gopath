# GoPath Bootstrap Plan

Этот документ фиксирует порядок первых маленьких шагов разработки. Каждый шаг выполняется отдельным commit после отдельного подтверждения.

## Principles

- Один согласованный шаг — один commit.
- Не смешивать документацию, frontend, backend, infrastructure и CI в одном изменении.
- После каждого шага запускать проверки, которые уже доступны на этом этапе.
- Если запускался Docker stack, останавливать его без удаления volumes.
- Не менять unrelated files и не переписывать git history.

## Proposed Commit Sequence

1. `docs: add project bootstrap specification`
   - Добавить `README.md`.
   - Добавить `.gitignore`.
   - Добавить `docs/gopath-technical-spec.md`.
   - Добавить `docs/bootstrap-plan.md`.

2. `build: add frontend skeleton`
   - Создать Vite + React + TypeScript приложение в `frontend/`.
   - Подключить Tailwind CSS.
   - Добавить базовые scripts для lint/build.
   - Без реализации продуктовых экранов.

3. `build: add backend skeleton`
   - Создать Go module в `backend/`.
   - Добавить минимальный HTTP server.
   - Добавить health endpoint.
   - Добавить базовый test.

4. `docs: add architecture notes`
   - Описать структуру monorepo.
   - Зафиксировать API boundaries.
   - Зафиксировать первые решения по auth, migrations и sandbox.

5. `ui: add static mvp shell`
   - Добавить общий layout.
   - Добавить routing для `/app`, `/learn`, `/lesson/:lessonId`, `/challenge/:challengeId`, `/profile`.
   - Использовать mock data.

6. `ui: add static mvp screens`
   - Реализовать 5 high-fidelity экранов на mock data.
   - Сохранить русский UI и черно-белый developer-tool стиль.

7. `api: add core learning endpoints`
   - Добавить первые REST endpoints для tracks, levels, lessons и progress.
   - Пока без PostgreSQL.

8. `infra: add local postgres and redis`
   - Добавить Docker Compose для PostgreSQL и Redis.
   - Добавить базовые env examples.

9. `api: add persistence model`
   - Добавить migrations.
   - Добавить базовые repository/service слои.

10. `api: add challenge runner prototype`
    - Добавить простой runner для Go code challenges.
    - Вернуть статусы `passed`, `failed`, `compile_error`, `timeout`.

## Initial Repository Layout

```text
docs/
frontend/
backend/
infra/
content/
.github/workflows/
```

Эта структура будет создаваться постепенно. На первом commit создается только `docs/` и базовые root-документы.
