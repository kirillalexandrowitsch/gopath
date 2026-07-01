BEGIN;

INSERT INTO users (
    id,
    email,
    handle,
    display_name,
    password_hash,
    role
) VALUES (
    'demo-user',
    'alex.kim@example.com',
    '@alexkim',
    'Alex Kim',
    'demo-password-hash',
    'learner'
) ON CONFLICT (id) DO UPDATE SET
    email = EXCLUDED.email,
    handle = EXCLUDED.handle,
    display_name = EXCLUDED.display_name,
    role = EXCLUDED.role,
    updated_at = now();

INSERT INTO tracks (
    id,
    title,
    description
) VALUES (
    'go-backend',
    'Go Backend',
    'Карьерный путь Go backend-разработчика от основ до production reliability.'
) ON CONFLICT (id) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    updated_at = now();

INSERT INTO levels (
    id,
    track_id,
    title,
    career_stage,
    status,
    required_xp,
    required_skills,
    unlock_rules,
    order_index
) VALUES
    (
        'trainee-backend',
        'go-backend',
        'Стажер Backend',
        'trainee',
        'completed',
        600,
        ARRAY['Go Basics', 'HTTP Basics', 'SQL Basics'],
        '{"completedBlocks":2,"totalBlocks":2}'::jsonb,
        0
    ),
    (
        'junior-backend',
        'go-backend',
        'Junior Backend',
        'junior',
        'active',
        2000,
        ARRAY['Go', 'Testing', 'Errors', 'REST API'],
        '{"completedBlocks":2,"totalBlocks":6}'::jsonb,
        1
    ),
    (
        'middle-backend',
        'go-backend',
        'Middle Backend',
        'middle',
        'locked',
        5000,
        ARRAY['Concurrency', 'PostgreSQL', 'Kafka', 'Observability'],
        '{"completedBlocks":0,"totalBlocks":6}'::jsonb,
        2
    ),
    (
        'senior-backend',
        'go-backend',
        'Senior Backend',
        'senior',
        'locked',
        9000,
        ARRAY['System Design', 'Reliability', 'Architecture'],
        '{"completedBlocks":0,"totalBlocks":6}'::jsonb,
        3
    )
ON CONFLICT (id) DO UPDATE SET
    track_id = EXCLUDED.track_id,
    title = EXCLUDED.title,
    career_stage = EXCLUDED.career_stage,
    status = EXCLUDED.status,
    required_xp = EXCLUDED.required_xp,
    required_skills = EXCLUDED.required_skills,
    unlock_rules = EXCLUDED.unlock_rules,
    order_index = EXCLUDED.order_index,
    updated_at = now();

INSERT INTO lessons (
    id,
    level_id,
    title,
    lesson_type,
    theory,
    question,
    explanation,
    xp_reward,
    tags,
    order_index
) VALUES (
    'go-errors-tests',
    'junior-backend',
    'Go: ошибки и тесты',
    'multiple_choice',
    'В Go функции, которые могут завершиться ошибкой, возвращают значение типа error. Такой подход держит обработку сбоя рядом с местом, где есть достаточно контекста.',
    'Почему в Go ошибку обычно возвращают явно?',
    'В Go ошибки обрабатываются явно, чтобы вызывающая сторона понимала контекст и могла выбрать стратегию обработки сбоя.',
    40,
    ARRAY['Go', 'Testing', 'Errors'],
    4
) ON CONFLICT (id) DO UPDATE SET
    level_id = EXCLUDED.level_id,
    title = EXCLUDED.title,
    lesson_type = EXCLUDED.lesson_type,
    theory = EXCLUDED.theory,
    question = EXCLUDED.question,
    explanation = EXCLUDED.explanation,
    xp_reward = EXCLUDED.xp_reward,
    tags = EXCLUDED.tags,
    order_index = EXCLUDED.order_index,
    updated_at = now();

INSERT INTO lesson_options (
    id,
    lesson_id,
    label,
    body,
    is_correct,
    order_index
) VALUES
    (
        'caller-controls-error',
        'go-errors-tests',
        'A',
        'Чтобы caller сам решил, как обработать сбой',
        true,
        0
    ),
    (
        'hide-problem',
        'go-errors-tests',
        'B',
        'Чтобы скрыть проблему от пользователя',
        false,
        1
    ),
    (
        'replace-tests',
        'go-errors-tests',
        'C',
        'Чтобы заменить unit tests',
        false,
        2
    ),
    (
        'speed-runtime',
        'go-errors-tests',
        'D',
        'Чтобы ускорить runtime',
        false,
        3
    )
ON CONFLICT (id) DO UPDATE SET
    lesson_id = EXCLUDED.lesson_id,
    label = EXCLUDED.label,
    body = EXCLUDED.body,
    is_correct = EXCLUDED.is_correct,
    order_index = EXCLUDED.order_index,
    updated_at = now();

INSERT INTO challenges (
    id,
    level_id,
    title,
    prompt,
    starter_code,
    hints,
    difficulty,
    required_skills,
    xp_reward,
    time_limit_ms,
    memory_limit_mb,
    order_index
) VALUES (
    'retry-context',
    'junior-backend',
    'Напиши безопасный retry',
    'Реализуй функцию FetchUser с механизмом retry для временно неудачных запросов. Повтор должен выполняться с экспоненциальной задержкой и уважать context.Context.',
    $starter$
package main

import (
    "context"
    "errors"
    "io"
    "net/http"
    "time"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type Doer interface {
    Do(req *http.Request) (*http.Response, error)
}

var client Doer

func FetchUser(ctx context.Context, id int) (User, error) {
    // TODO: реализуй функцию согласно требованиям
    return User{}, errors.New("not implemented")
}

func closeBody(resp *http.Response) {
    if resp != nil && resp.Body != nil {
        _, _ = io.Copy(io.Discard, resp.Body)
        _ = resp.Body.Close()
    }
}
$starter$,
    ARRAY['Используй time.NewTimer и обязательно останавливай timer, чтобы избежать утечек goroutine.'],
    'junior',
    ARRAY['context.Context', 'Retry', 'HTTP Client', 'Error Handling'],
    120,
    2000,
    256,
    0
) ON CONFLICT (id) DO UPDATE SET
    level_id = EXCLUDED.level_id,
    title = EXCLUDED.title,
    prompt = EXCLUDED.prompt,
    starter_code = EXCLUDED.starter_code,
    hints = EXCLUDED.hints,
    difficulty = EXCLUDED.difficulty,
    required_skills = EXCLUDED.required_skills,
    xp_reward = EXCLUDED.xp_reward,
    time_limit_ms = EXCLUDED.time_limit_ms,
    memory_limit_mb = EXCLUDED.memory_limit_mb,
    order_index = EXCLUDED.order_index,
    updated_at = now();

UPDATE levels
SET checkpoint_challenge_id = 'retry-context',
    updated_at = now()
WHERE id = 'junior-backend';

INSERT INTO challenge_tests (
    id,
    challenge_id,
    visibility,
    name,
    test_code,
    order_index
) VALUES
    (
        'retry-context-success',
        'retry-context',
        'public',
        'TestFetchUser_Success',
        $test$func TestFetchUser_Success(t *testing.T) {
    // Verifies successful response decoding.
}$test$,
        0
    ),
    (
        'retry-context-temporary-retry',
        'retry-context',
        'public',
        'TestFetchUser_RetryOnTemporary',
        $test$func TestFetchUser_RetryOnTemporary(t *testing.T) {
    // Verifies retry behavior for temporary errors.
}$test$,
        1
    ),
    (
        'retry-context-cancel',
        'retry-context',
        'public',
        'TestFetchUser_ContextCancel',
        $test$func TestFetchUser_ContextCancel(t *testing.T) {
    // Verifies context cancellation is respected.
}$test$,
        2
    ),
    (
        'retry-context-non-retryable',
        'retry-context',
        'hidden',
        'TestFetchUser_NonRetryable',
        $test$func TestFetchUser_NonRetryable(t *testing.T) {
    // Verifies non-temporary errors are not retried.
}$test$,
        0
    ),
    (
        'retry-context-max-retries',
        'retry-context',
        'hidden',
        'TestFetchUser_MaxRetries',
        $test$func TestFetchUser_MaxRetries(t *testing.T) {
    // Verifies retry count does not exceed the requirement.
}$test$,
        1
    )
ON CONFLICT (id) DO UPDATE SET
    challenge_id = EXCLUDED.challenge_id,
    visibility = EXCLUDED.visibility,
    name = EXCLUDED.name,
    test_code = EXCLUDED.test_code,
    order_index = EXCLUDED.order_index,
    updated_at = now();

INSERT INTO user_progress (
    user_id,
    current_level_id,
    current_lesson_id,
    current_checkpoint_id,
    xp,
    streak_days,
    completed_lessons,
    completed_checkpoints,
    last_activity_at
) VALUES (
    'demo-user',
    'junior-backend',
    'go-errors-tests',
    'retry-context',
    1420,
    7,
    24,
    2,
    now()
) ON CONFLICT (user_id) DO UPDATE SET
    current_level_id = EXCLUDED.current_level_id,
    current_lesson_id = EXCLUDED.current_lesson_id,
    current_checkpoint_id = EXCLUDED.current_checkpoint_id,
    xp = EXCLUDED.xp,
    streak_days = EXCLUDED.streak_days,
    completed_lessons = EXCLUDED.completed_lessons,
    completed_checkpoints = EXCLUDED.completed_checkpoints,
    last_activity_at = EXCLUDED.last_activity_at,
    updated_at = now();

INSERT INTO skill_progress (
    user_id,
    skill,
    score
) VALUES
    ('demo-user', 'Go', 84),
    ('demo-user', 'REST API', 78),
    ('demo-user', 'PostgreSQL', 72),
    ('demo-user', 'Redis', 68),
    ('demo-user', 'Docker', 75),
    ('demo-user', 'Kafka', 60),
    ('demo-user', 'Kubernetes', 58),
    ('demo-user', 'Prometheus', 50)
ON CONFLICT (user_id, skill) DO UPDATE SET
    score = EXCLUDED.score,
    updated_at = now();

COMMIT;
