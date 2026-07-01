# Content Authoring Guide

Этот документ фиксирует первые правила для учебного контента GoPath. Он нужен, чтобы будущие lessons, challenges и checkpoints писались в одном стиле и не расходились с product specification.

Это authoring guide, а не content schema implementation. Директория `content/`, validation script и import pipeline будут добавлены отдельными future commits по `docs/delivery-roadmap.md`.

## Sources of Truth

- Product requirements: `docs/gopath-technical-spec.md`.
- Delivery order: `docs/delivery-roadmap.md`.
- Current API snapshot: `docs/api-contract.md`.
- Current database shape: `backend/migrations/000001_init_schema.sql`.
- Current demo seed examples: `backend/migrations/000002_seed_learning_data.sql`.

## Current Content Status

Сейчас учебный контент временно живет в нескольких местах:

- SQL seed data в `backend/migrations/000002_seed_learning_data.sql`.
- Backend in-memory mocks в `backend/internal/learning`.
- Challenge runner catalog в `backend/internal/challenges`.
- Frontend typed mock data в `frontend/src/data/shell.ts`.

Это временное состояние bootstrap. В будущем `content/` должен стать основным source для tracks, levels, lessons, challenges, tests, hints и seed/import pipeline.

## Content Principles

- Основной язык учебного контента - русский.
- Code, identifiers, API paths, commands и названия технологий остаются на английском.
- Контент должен быть backend-oriented, а не академическим набором синтаксиса.
- Даже `Стажер Backend` должен решать задачи в практическом backend context.
- Уроки должны быть короткими, но не поверхностными.
- Каждая единица контента должна давать progression value: знание, навык, проверку или подготовку к checkpoint.
- Feedback должен объяснять причину, а не просто говорить `Верно` или `Неверно`.
- Hints должны помогать думать, но не раскрывать полное решение.

## Brand and Style Rules

Запрещено:

- копировать Duolingo brand;
- использовать mascot;
- копировать Duolingo colors, texts, assets или visual identity;
- писать игровую или cartoon-style подачу;
- превращать GoPath в продукт для совсем нулевых.

Разрешено:

- использовать progression-based learning mechanics;
- использовать XP, streak, career map, checkpoints и short lessons;
- писать строгий, современный, практичный developer education content.

## Language Rules

На русском:

- interface copy;
- lesson titles;
- theory;
- task prompts;
- feedback;
- explanations;
- hints;
- recommendations;
- checkpoint descriptions.

На английском:

- `Go`, `PostgreSQL`, `Redis`, `Kafka`, `RabbitMQ`, `Docker`, `Kubernetes`;
- `REST API`, `gRPC`, `endpoint`, `request`, `response`;
- `handler`, `middleware`, `context.Context`, `error`;
- commands like `go test`, `go fmt`, `docker compose up`;
- code, SQL, JSON and HTTP status names like `201 Created`.

Правильно:

```text
Напиши handler, который возвращает `201 Created`, если пользователь успешно создан.
```

Неправильно:

```text
Напиши обработчик, который возвращает "201 Создано".
```

## IDs and Naming

IDs должны быть stable, lowercase и URL-safe.

Recommended format:

- Track: `go-backend`.
- Level: `trainee-backend`, `junior-backend`, `middle-backend`, `senior-backend`.
- Lesson: `go-errors-tests`, `http-handlers`, `postgres-indexes`.
- Challenge: `retry-context`, `rest-api-checkpoint`.
- Test: `retry-context-success`, `retry-context-max-retries`.

Rules:

- Использовать kebab-case.
- Не использовать spaces.
- Не использовать русские символы в ids.
- Не переиспользовать id для другого смысла.
- Если title меняется, id должен оставаться стабильным, пока смысл content unit тот же.

## Track Requirements

Track описывает карьерный путь.

Required fields conceptually:

- `id`;
- `title`;
- `description`;
- ordered levels.

Current example:

```text
id: go-backend
title: Go Backend
description: Карьерный путь Go backend-разработчика от основ до production reliability.
```

Track должен отвечать на вопрос: какой путь проходит пользователь и к какой роли он движется.

## Level Requirements

Level описывает карьерную ступень.

Current career stages:

- `trainee`;
- `junior`;
- `middle`;
- `senior`.

Level должен включать:

- stable `id`;
- user-facing `title`;
- `career_stage`;
- required skills;
- ordered lessons/challenges;
- unlock rules;
- checkpoint challenge.

Level content должен соответствовать реальной ответственности роли:

- `Стажер Backend` - простые задачи под контролем;
- `Junior Backend` - небольшие backend задачи самостоятельно;
- `Middle Backend` - ownership модуля или сервиса;
- `Senior Backend` - architecture, reliability, trade-offs, mentoring.

## Lesson Requirements

Lesson должна быть короткой и проверяемой.

Current schema fields:

- `id`;
- `level_id`;
- `title`;
- `lesson_type`;
- `theory`;
- `question`;
- `explanation`;
- `xp_reward`;
- `tags`;
- `order_index`.

Current lesson types:

- `theory`;
- `multiple_choice`;
- `checkpoint`.

Lesson должна содержать:

- одну ясную learning objective;
- короткую theory section;
- вопрос или практическую проверку;
- объяснение правильного ответа;
- XP reward;
- tags для skill/progress mapping;
- order внутри level.

Lesson не должна:

- проверять несколько unrelated concepts одновременно;
- использовать слишком абстрактные вопросы без backend context;
- давать ambiguous answer options;
- быть завязана на UI mock data как источник истины.

## Lesson Options

Lesson option используется для multiple-choice задач.

Current schema fields:

- `id`;
- `lesson_id`;
- `label`;
- `body`;
- `is_correct`;
- `order_index`.

Rules:

- Для single-choice question должен быть ровно один correct option.
- Labels должны быть короткими: `A`, `B`, `C`, `D`.
- Wrong options должны быть правдоподобными, но явно неверными после объяснения.
- Option text пишется на русском, code/terms остаются на английском.
- `is_correct` не должен быть exposed в production assessment flow напрямую.

Current bootstrap API exposes `correctId` для lesson UI. Это временное решение и должно быть пересмотрено перед real assessment flow.

## Challenge Requirements

Challenge проверяет практический навык через код.

Current schema fields:

- `id`;
- `level_id`;
- `title`;
- `prompt`;
- `starter_code`;
- `hints`;
- `difficulty`;
- `required_skills`;
- `xp_reward`;
- `time_limit_ms`;
- `memory_limit_mb`;
- `order_index`.

Current difficulty values:

- `trainee`;
- `junior`;
- `middle`;
- `senior`.

Challenge должен содержать:

- clear backend scenario;
- exact task requirements;
- starter code that compiles or intentionally points to TODO;
- function signature or API contract;
- deterministic public tests;
- deterministic hidden tests;
- hints;
- realistic time and memory limits;
- required skills.

Challenge не должен:

- требовать external network;
- зависеть от current time без control;
- зависеть от flaky sleeps;
- требовать secret values;
- проверять undefined behavior;
- иметь hidden tests, которые противоречат prompt.

## Challenge Tests

Current test visibility values:

- `public`;
- `hidden`.

Public tests:

- показывают пользователю базовый expected behavior;
- должны быть понятными по name;
- не должны покрывать все edge cases.

Hidden tests:

- проверяют edge cases;
- не должны менять требования задачи;
- не должны быть неожиданной новой задачей.

Test names:

- Go test names должны начинаться с `Test`;
- names должны объяснять behavior: `TestFetchUser_ContextCancel`;
- seed ids должны быть stable: `retry-context-cancel`.

Test code:

- должен быть deterministic;
- должен запускаться через `go test`;
- должен избегать real network;
- должен контролировать timeouts и context;
- должен давать понятную failure message.

## Hints

Hint должен помогать, но не решать задачу за пользователя.

Good hint:

```text
Используй `time.NewTimer` и обязательно останавливай timer, чтобы избежать утечек goroutine.
```

Bad hint:

```text
Скопируй готовый retry loop с тремя попытками и delay 100ms.
```

Rules:

- Hint должен указывать направление.
- Hint может назвать relevant API.
- Hint не должен раскрывать весь алгоритм.
- Hint должен быть связан с конкретной ошибкой или concept.

## Checkpoint Requirements

Checkpoint проверяет готовность перейти дальше по career path.

Checkpoint должен:

- объединять несколько skills уровня;
- иметь practical backend scenario;
- быть сложнее обычного lesson;
- иметь clear acceptance criteria;
- иметь public and hidden tests или rubric;
- влиять на unlock/progression rules.

Examples:

- `Стажер Backend`: простой REST endpoint с SQL query и базовым test.
- `Junior Backend`: REST API feature with validation, migrations, tests and error handling.
- `Middle Backend`: production-like feature with worker, queue, PostgreSQL, Redis, retries and observability.
- `Senior Backend`: architecture case with constraints, trade-offs, rollout and failure modes.

## XP and Progression Guidance

Until progression rules are implemented, XP values are content metadata.

Guidelines:

- Простые lessons: small XP reward.
- Practice tasks: medium XP reward.
- Checkpoints/challenges: larger XP reward.
- XP should reflect difficulty and effort.
- XP must not be farmable after persistence rules are implemented.

Content authors should not invent custom XP behavior per lesson. XP rules will be centralized in future progression domain services.

## Quality Checklist

Before content is accepted, verify:

- Russian text is clear and natural.
- English technical terms are preserved where appropriate.
- The content matches the target career level.
- The backend scenario is realistic.
- The lesson has one primary objective.
- The correct answer is unambiguous.
- Wrong options are plausible but not misleading.
- Explanation teaches the concept.
- Challenge starter code matches prompt.
- Public tests match visible requirements.
- Hidden tests do not introduce new requirements.
- Hints do not reveal full solutions.
- IDs are stable and URL-safe.
- Tags/skills are useful for progress tracking.
- No Duolingo brand, mascot, colors, text or visual identity is copied.

## Current Limitations

- There is no `content/` directory yet.
- There is no content validation script yet.
- There is no seed generation pipeline yet.
- There is no admin/content tooling yet.
- SQL seed data and Go/frontend mocks still duplicate content.
- `correctId` is exposed by the current lesson API for bootstrap and must be revisited before production assessment.

## Future Direction

Future content work should happen in small commits:

1. Add `content/` structure.
2. Draft content file format.
3. Add content validation.
4. Move demo lesson content into content source.
5. Move `retry-context` challenge content into content source.
6. Add seed generation or import pipeline.
7. Expand `Стажер Backend` content.
8. Expand `Junior Backend` content.
9. Add first checkpoint mini-projects.
