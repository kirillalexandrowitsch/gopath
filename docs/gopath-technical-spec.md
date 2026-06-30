# GoPath: техническое задание MVP v1

## Инструкция для нового чата Codex

Это ТЗ для нового проекта `GoPath`.

Нужно разработать веб-приложение для русскоязычных пользователей, которое помогает учиться Go backend-разработке через короткие уроки, карьерную карту, XP, streak, тесты и code challenges.

Важные правила:

- Основной язык продукта: русский.
- Код, команды и названия технологий остаются на английском.
- Интерфейс: черно-белый, современный, плотный, в стиле серьезного developer tool.
- Не копировать бренд, маскота, ассеты или визуальную айдентику Duolingo.
- Можно использовать только механику progression-based learning: уроки, карта, XP, streak, checkpoints.
- Сначала делать MVP, не пытаться сразу строить огромную production-платформу.

---

# 1. Концепция продукта

`GoPath` — веб-приложение для обучения Go backend-разработке в формате коротких интерактивных уроков, практических задач, code challenges, XP, streak и карьерной карты развития.

Продукт вдохновляется механиками Duolingo:

- регулярная практика;
- короткие уроки;
- карта прогресса;
- уровни;
- награды;
- ежедневная цель;
- быстрый feedback.

При этом GoPath не должен копировать:

- бренд Duolingo;
- маскота;
- визуальный стиль;
- цвета;
- ассеты;
- тексты;
- точную композицию экранов.

GoPath должен ощущаться как серьезный современный developer education tool для backend-разработчиков.

---

# 2. Язык продукта

Основной язык приложения — русский.

На русском должны быть:

- интерфейс;
- onboarding;
- уроки;
- задания;
- объяснения;
- подсказки;
- ошибки;
- feedback;
- прогресс;
- названия разделов;
- описания уровней;
- рекомендации.

На английском остаются:

- код;
- названия технологий: `Go`, `PostgreSQL`, `Redis`, `Kafka`, `RabbitMQ`, `Docker`, `Kubernetes`;
- команды: `go test`, `go fmt`, `docker compose up`, `kubectl apply`;
- API-термины: `REST API`, `gRPC`, `endpoint`, `request`, `response`;
- SQL;
- Docker/Kubernetes сущности;
- технические термины, которые в реальной разработке обычно не переводятся.

Пример правильного стиля:

> Напиши handler, который возвращает `201 Created`, если пользователь успешно создан.

Пример неправильного стиля:

> Напиши обработчик, который возвращает “201 Создано”.

---

# 3. Целевая аудитория

GoPath рассчитан на русскоязычных разработчиков, которые хотят расти в Go backend:

- начинающие, которые готовятся к стажировке;
- стажеры;
- Junior Backend;
- Middle Backend;
- Senior Backend;
- разработчики из других языков, которые переходят в Go.

Продукт не должен быть “для совсем нулевых”. Даже уровень “Стажер” должен подразумевать практический backend-контекст.

---

# 4. Основная механика

Пользователь проходит карьерную карту backend-разработчика.

Основной цикл:

1. Пользователь открывает приложение.
2. Видит текущий уровень, XP, streak и ежедневную цель.
3. Нажимает “Продолжить обучение”.
4. Проходит короткий урок.
5. Решает вопрос или практическое задание.
6. Получает instant feedback.
7. Получает XP.
8. Поддерживает streak.
9. Открывает следующие уроки.
10. Проходит checkpoint.
11. Переходит на следующий карьерный уровень.

Типы заданий:

- выбор одного ответа;
- выбор нескольких ответов;
- сопоставление;
- заполнение пропуска в коде;
- исправление бага;
- написание функции;
- SQL-запрос;
- анализ логов;
- анализ ошибки;
- code challenge;
- mini-project checkpoint.

---

# 5. Карьерная градация

Градация должна быть похожа на реальную карьерную лестницу backend-разработчика, а не на легкий учебный список тем.

## 5.1 Стажер Backend

Цель уровня: пользователь уже не совсем с нуля и способен выполнять простые задачи под контролем.

Должен знать:

- базовый синтаксис Go;
- переменные, функции, условия, циклы;
- структуры;
- интерфейсы на базовом уровне;
- slices;
- maps;
- простую обработку ошибок через `error`;
- базовый Git: `commit`, `branch`, `merge`, pull request;
- Linux basics: файлы, процессы, environment variables, permissions, `curl`;
- базовый SQL: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `JOIN`;
- HTTP basics: methods, status codes, headers, JSON;
- простые REST endpoints;
- базовые unit tests;
- Docker Compose на уровне локального запуска сервисов.

Типовые задания:

- написать простую Go-функцию;
- исправить ошибку компиляции;
- написать простой SQL-запрос;
- выбрать правильный HTTP status code;
- реализовать простой HTTP handler;
- объяснить, почему функция возвращает ошибку;
- запустить локальный сервис через `docker compose`.

Checkpoint уровня:

- написать небольшой REST endpoint;
- добавить простой SQL-запрос;
- обработать ошибку;
- покрыть базовым тестом.

## 5.2 Junior Backend

Цель уровня: пользователь может самостоятельно выполнять небольшие backend-задачи, но архитектурные решения еще согласует.

Должен знать:

- Go idioms;
- package structure;
- receiver methods;
- interface design на базовом уровне;
- нормальную обработку ошибок;
- `errors.Is`, `errors.As`, error wrapping;
- `context.Context` на базовом уровне;
- table-driven tests;
- REST API design;
- request validation;
- auth basics;
- PostgreSQL migrations;
- indexes;
- constraints;
- transactions;
- Redis cache, TTL, rate limiting basics;
- Dockerfile;
- containers и images;
- CI/CD basics;
- базовое логирование;
- базовые метрики;
- основы безопасности API.

Типовые задания:

- реализовать CRUD-фичу;
- добавить миграцию;
- покрыть код тестами;
- добавить request validation;
- добавить кеширование;
- обработать ошибку API;
- разобраться с `500 Internal Server Error`;
- написать Dockerfile;
- поправить failing pipeline.

Checkpoint уровня:

- реализовать небольшую backend-фичу end-to-end:
  - REST endpoint;
  - PostgreSQL migration;
  - validation;
  - tests;
  - error handling;
  - Docker-ready запуск.

## 5.3 Middle Backend

Цель уровня: пользователь владеет сервисом или крупным модулем и принимает технические решения внутри задачи.

Должен знать:

- goroutines;
- channels;
- mutex;
- race conditions;
- `context` cancellation;
- timeout;
- deadline propagation;
- profiling;
- CPU profile;
- memory profile;
- goroutine leaks;
- gRPC;
- proto contracts;
- interceptors;
- Kafka;
- RabbitMQ;
- producer/consumer;
- ack;
- retry;
- DLQ;
- idempotency;
- PostgreSQL query plans;
- indexes глубже;
- locks;
- isolation levels;
- slow query analysis;
- Redis cache invalidation;
- MongoDB;
- ClickHouse;
- когда выбирать SQL/NoSQL/OLAP;
- Kubernetes basics;
- deployment;
- service;
- configmap;
- secret;
- probes;
- Prometheus metrics;
- Grafana dashboards;
- structured logs;
- tracing basics;
- production debugging.

Типовые задания:

- спроектировать и реализовать сервисную фичу;
- найти slow query;
- исправить race condition;
- добавить retry/idempotency;
- настроить DLQ;
- добавить метрики;
- собрать Grafana dashboard;
- подготовить сервис к production deploy;
- разобраться с деградацией latency;
- найти причину memory leak.

Checkpoint уровня:

- реализовать production-like фичу:
  - API;
  - background worker;
  - очередь;
  - PostgreSQL;
  - Redis;
  - observability;
  - retries;
  - tests;
  - documented trade-offs.

## 5.4 Senior Backend

Цель уровня: пользователь отвечает не только за код, а за архитектуру, надежность, масштабирование и техническое качество команды.

Должен знать:

- distributed systems;
- consistency;
- availability;
- backpressure;
- retries;
- timeouts;
- service boundaries;
- API contracts;
- versioning;
- backward compatibility;
- zero-downtime migrations;
- performance;
- scalability;
- reliability engineering;
- SLO;
- error budget;
- incident response;
- advanced PostgreSQL;
- partitioning;
- replication basics;
- locks;
- vacuum;
- query tuning;
- production Kafka/RabbitMQ patterns;
- Kubernetes production patterns;
- security basics;
- threat modeling basics;
- secrets management;
- auth flows;
- supply chain risks;
- observability на уровне пользовательских симптомов;
- code review standards;
- mentoring;
- technical leadership.

Типовые задания:

- выбрать архитектуру под нагрузку;
- провести design review;
- расследовать production incident;
- снизить latency;
- снизить error rate;
- спроектировать миграцию без downtime;
- построить SLO;
- улучшить observability;
- провести performance review;
- наставлять Junior/Middle;
- улучшить инженерные стандарты команды.

Checkpoint уровня:

- решить архитектурный кейс:
  - есть нагрузка;
  - есть ограничения;
  - есть несколько вариантов архитектуры;
  - нужно выбрать решение;
  - объяснить trade-offs;
  - описать rollout;
  - описать observability;
  - описать failure modes.

---

# 6. Основные экраны MVP

## 6.1 Dashboard

Назначение: главный экран после входа.

Должен показывать:

- текущий карьерный уровень;
- XP;
- streak;
- ежедневную цель;
- текущий урок;
- кнопку “Продолжить обучение”;
- слабые темы;
- ближайший checkpoint;
- краткую активность за неделю.

Примеры текста:

- “Продолжить обучение”
- “Junior Backend”
- “7 дней подряд”
- “Сегодняшняя цель”
- “Слабые темы”
- “Следующий checkpoint”

## 6.2 Learning Path

Назначение: карта развития backend-разработчика.

Должна показывать:

- вертикальную карту уровней;
- `Стажер Backend`;
- `Junior Backend`;
- `Middle Backend`;
- `Senior Backend`;
- завершенные уроки;
- активный урок;
- заблокированные уроки;
- checkpoints;
- требования уровня;
- навыки уровня;
- прогресс до следующего уровня.

Состояния уроков:

- `Завершено`;
- `Активно`;
- `Заблокировано`;
- `Checkpoint`.

## 6.3 Lesson

Назначение: прохождение обычного урока.

Должен показывать:

- название урока;
- номер шага;
- прогресс;
- вопрос;
- варианты ответа;
- кнопку “Проверить”;
- feedback;
- объяснение;
- XP за правильный ответ;
- переход дальше.

Пример задания:

> Почему в Go ошибку обычно возвращают явно?

Варианты:

- Чтобы caller сам решил, как обработать сбой.
- Чтобы скрыть проблему от пользователя.
- Чтобы заменить unit tests.
- Чтобы ускорить runtime.

## 6.4 Code Challenge

Назначение: практическая задача с написанием кода.

Должен показывать:

- условие задачи на русском;
- Go code editor;
- терминал;
- public tests;
- hidden tests;
- кнопку `Run`;
- кнопку `Submit`;
- sandbox status;
- лимиты sandbox;
- подсказку;
- результат проверки.

Пример задания:

> Напиши функцию `FetchUser(ctx context.Context, id int) (User, error)`, которая корректно уважает отмену `context.Context`.

## 6.5 Profile / Progress

Назначение: экран личного прогресса.

Должен показывать:

- текущий уровень;
- XP;
- streak calendar;
- skill graph;
- прогресс по технологиям;
- завершенные уроки;
- готовность к следующему уровню;
- рекомендации;
- историю активности.

Технологии в skill graph:

- Go;
- REST API;
- PostgreSQL;
- Redis;
- Docker;
- Kafka;
- RabbitMQ;
- Kubernetes;
- Prometheus;
- Grafana.

---

# 7. Стек приложения

## 7.1 Frontend

Рекомендуемый стек:

- React;
- TypeScript;
- Vite;
- Tailwind CSS;
- React Query;
- Zustand;
- Monaco Editor;
- xterm.js.

Основные маршруты:

```text
/app
/learn
/lesson/:lessonId
/challenge/:challengeId
/profile
```

## 7.2 Backend

Рекомендуемый стек:

- Go;
- REST API;
- PostgreSQL;
- Redis;
- Docker Compose.

## 7.3 Инфраструктура MVP

MVP должен использовать:

- PostgreSQL для пользователей, уроков, попыток и прогресса;
- Redis для короткоживущих sandbox jobs, rate limits и кеша;
- Docker Compose для локального запуска;
- backend runner для code challenges.

Позже можно добавить:

- gRPC;
- Kafka;
- RabbitMQ;
- Kubernetes;
- Prometheus;
- Grafana;
- ClickHouse.

В MVP эти технологии могут присутствовать сначала как учебный контент, а не как обязательная инфраструктура самого приложения.

---

# 8. API MVP

Auth:

```text
POST /api/v1/auth/register
POST /api/v1/auth/login
GET  /api/v1/me
```

Learning:

```text
GET /api/v1/tracks
GET /api/v1/levels
GET /api/v1/lessons/:id
```

Attempts:

```text
POST /api/v1/attempts
```

Code challenges:

```text
POST /api/v1/challenges/:id/run
POST /api/v1/challenges/:id/submit
```

Progress:

```text
GET /api/v1/progress
GET /api/v1/profile
```

---

# 9. Code Sandbox

MVP использует hybrid sandbox.

Пользователь:

- пишет Go-код в браузере;
- видит терминал;
- нажимает `Run`;
- получает результат public tests;
- нажимает `Submit`;
- получает результат public + hidden tests.

Backend:

- принимает код;
- запускает изолированный runner;
- выполняет `gofmt`;
- выполняет `go test`;
- применяет public tests;
- применяет hidden tests;
- возвращает результат.

Ограничения sandbox:

- сеть запрещена;
- лимит памяти;
- лимит CPU;
- лимит времени;
- запрет опасных системных вызовов;
- изоляция файловой системы;
- одноразовое окружение на запуск.

Пример ответа:

```json
{
  "status": "passed",
  "stdout": "go test ./... ok",
  "stderr": "",
  "testsPassed": 5,
  "testsTotal": 5,
  "durationMs": 148,
  "hint": ""
}
```

Возможные статусы:

- `queued`;
- `running`;
- `passed`;
- `failed`;
- `timeout`;
- `compile_error`;
- `blocked`;
- `internal_error`.

---

# 10. Контентная модель

Track:

- id;
- title;
- description;
- levels.

Level:

- id;
- title;
- career stage;
- required skills;
- lessons;
- checkpoint;
- unlock rules.

Lesson:

- id;
- level id;
- title;
- type;
- theory;
- question;
- options;
- explanation;
- XP reward.

Challenge:

- id;
- title;
- prompt;
- starter code;
- public tests;
- hidden tests;
- hints;
- difficulty;
- required skills.

Attempt:

- user id;
- lesson id;
- answer;
- correctness;
- XP awarded;
- timestamp.

Progress:

- user id;
- completed lessons;
- current level;
- current checkpoint;
- XP;
- streak;
- skill graph.

---

# 11. Визуальный стиль

Интерфейс должен быть:

- черно-белым;
- современным;
- плотным;
- техническим;
- серьезным;
- красивым;
- похожим на startup SaaS / developer tool 2026 года.

Запрещено:

- яркая игровая палитра;
- зеленый маскот;
- любые Duolingo-like персонажи;
- декоративные gradient blobs;
- marketing hero первым экраном;
- карточки внутри карточек;
- перегруженный cartoon-style.

Разрешено:

- контрастная типографика;
- аккуратные borders;
- hatching;
- монохромные иконки;
- progress rings;
- тонкие линии;
- terminal/code surfaces;
- плотные dashboard layouts.

Успех/ошибка должны показываться не цветом, а:

- иконками;
- толщиной border;
- паттерном;
- контрастом;
- текстовым статусом.

---

# 12. Макеты MVP

Нужно подготовить 5 high-fidelity экранов:

1. Dashboard.
2. Learning Path.
3. Lesson.
4. Code Challenge.
5. Profile / Progress.

Общие требования к макетам:

- desktop 1440px;
- русский UI;
- код и технологии на английском;
- черно-белый стиль;
- без брендинга Duolingo;
- без mascot;
- без landing page;
- без декоративных градиентов;
- без наложения текста;
- интерфейс должен выглядеть как готовый продукт.

---

# 13. Критерии готовности MVP

MVP считается готовым, если:

- есть 5 основных экранов;
- весь интерфейс на русском;
- уроки и задания на русском;
- код и технологии на английском;
- есть реалистичная карьерная карта;
- можно пройти обычный урок;
- можно пройти code challenge;
- XP обновляется;
- streak обновляется;
- Learning Path показывает active/completed/locked состояния;
- checkpoints блокируют переход между уровнями;
- профиль показывает прогресс и готовность к следующему уровню;
- backend API покрывает базовые сценарии;
- sandbox возвращает понятный результат;
- UI выглядит как современный developer education product.

---

# 14. Рекомендуемый первый этап разработки

Первый этап не должен пытаться реализовать все сразу.

Рекомендуемый порядок:

1. Создать репозиторий.
2. Создать frontend skeleton.
3. Создать backend skeleton.
4. Описать базовую модель данных.
5. Сделать статичный русский UI для 5 экранов.
6. Подключить mock data.
7. Реализовать REST API.
8. Подключить frontend к API.
9. Реализовать простой sandbox runner.
10. Добавить PostgreSQL/Redis.
11. Добавить тесты.
12. Улучшить UX и visual polish.

Главный принцип: двигаться маленькими проверяемыми шагами.
