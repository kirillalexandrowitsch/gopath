BEGIN;

CREATE TABLE users (
    id TEXT PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    handle TEXT NOT NULL UNIQUE,
    display_name TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'learner' CHECK (role IN ('learner', 'admin')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE tracks (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE levels (
    id TEXT PRIMARY KEY,
    track_id TEXT NOT NULL REFERENCES tracks(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    career_stage TEXT NOT NULL CHECK (career_stage IN ('trainee', 'junior', 'middle', 'senior')),
    status TEXT NOT NULL DEFAULT 'locked' CHECK (status IN ('completed', 'active', 'locked')),
    required_xp INTEGER NOT NULL DEFAULT 0 CHECK (required_xp >= 0),
    required_skills TEXT[] NOT NULL DEFAULT '{}',
    unlock_rules JSONB NOT NULL DEFAULT '{}'::jsonb,
    checkpoint_challenge_id TEXT,
    order_index INTEGER NOT NULL CHECK (order_index >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (track_id, order_index)
);

CREATE TABLE lessons (
    id TEXT PRIMARY KEY,
    level_id TEXT NOT NULL REFERENCES levels(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    lesson_type TEXT NOT NULL CHECK (lesson_type IN ('theory', 'multiple_choice', 'checkpoint')),
    theory TEXT NOT NULL DEFAULT '',
    question TEXT NOT NULL DEFAULT '',
    explanation TEXT NOT NULL DEFAULT '',
    xp_reward INTEGER NOT NULL DEFAULT 0 CHECK (xp_reward >= 0),
    tags TEXT[] NOT NULL DEFAULT '{}',
    order_index INTEGER NOT NULL CHECK (order_index >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (level_id, order_index)
);

CREATE TABLE lesson_options (
    id TEXT PRIMARY KEY,
    lesson_id TEXT NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    label TEXT NOT NULL,
    body TEXT NOT NULL,
    is_correct BOOLEAN NOT NULL DEFAULT false,
    order_index INTEGER NOT NULL CHECK (order_index >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (lesson_id, label),
    UNIQUE (lesson_id, order_index)
);

CREATE TABLE challenges (
    id TEXT PRIMARY KEY,
    level_id TEXT NOT NULL REFERENCES levels(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    prompt TEXT NOT NULL,
    starter_code TEXT NOT NULL,
    hints TEXT[] NOT NULL DEFAULT '{}',
    difficulty TEXT NOT NULL CHECK (difficulty IN ('trainee', 'junior', 'middle', 'senior')),
    required_skills TEXT[] NOT NULL DEFAULT '{}',
    xp_reward INTEGER NOT NULL DEFAULT 0 CHECK (xp_reward >= 0),
    time_limit_ms INTEGER NOT NULL DEFAULT 2000 CHECK (time_limit_ms > 0),
    memory_limit_mb INTEGER NOT NULL DEFAULT 256 CHECK (memory_limit_mb > 0),
    order_index INTEGER NOT NULL CHECK (order_index >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (level_id, order_index)
);

CREATE TABLE challenge_tests (
    id TEXT PRIMARY KEY,
    challenge_id TEXT NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    visibility TEXT NOT NULL CHECK (visibility IN ('public', 'hidden')),
    name TEXT NOT NULL,
    test_code TEXT NOT NULL,
    order_index INTEGER NOT NULL CHECK (order_index >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (challenge_id, name),
    UNIQUE (challenge_id, visibility, order_index)
);

ALTER TABLE levels
    ADD CONSTRAINT levels_checkpoint_challenge_id_fkey
    FOREIGN KEY (checkpoint_challenge_id) REFERENCES challenges(id) ON DELETE SET NULL;

CREATE TABLE attempts (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lesson_id TEXT REFERENCES lessons(id) ON DELETE CASCADE,
    challenge_id TEXT REFERENCES challenges(id) ON DELETE CASCADE,
    answer TEXT NOT NULL DEFAULT '',
    is_correct BOOLEAN NOT NULL DEFAULT false,
    xp_awarded INTEGER NOT NULL DEFAULT 0 CHECK (xp_awarded >= 0),
    status TEXT NOT NULL CHECK (
        status IN (
            'queued',
            'running',
            'passed',
            'failed',
            'timeout',
            'compile_error',
            'blocked',
            'internal_error'
        )
    ),
    stdout TEXT NOT NULL DEFAULT '',
    stderr TEXT NOT NULL DEFAULT '',
    tests_passed INTEGER NOT NULL DEFAULT 0 CHECK (tests_passed >= 0),
    tests_total INTEGER NOT NULL DEFAULT 0 CHECK (tests_total >= 0),
    duration_ms INTEGER NOT NULL DEFAULT 0 CHECK (duration_ms >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (((lesson_id IS NOT NULL)::int + (challenge_id IS NOT NULL)::int) = 1),
    CHECK (tests_passed <= tests_total)
);

CREATE TABLE user_progress (
    user_id TEXT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    current_level_id TEXT REFERENCES levels(id) ON DELETE SET NULL,
    current_lesson_id TEXT REFERENCES lessons(id) ON DELETE SET NULL,
    current_checkpoint_id TEXT REFERENCES challenges(id) ON DELETE SET NULL,
    xp INTEGER NOT NULL DEFAULT 0 CHECK (xp >= 0),
    streak_days INTEGER NOT NULL DEFAULT 0 CHECK (streak_days >= 0),
    completed_lessons INTEGER NOT NULL DEFAULT 0 CHECK (completed_lessons >= 0),
    completed_checkpoints INTEGER NOT NULL DEFAULT 0 CHECK (completed_checkpoints >= 0),
    last_activity_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE skill_progress (
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    skill TEXT NOT NULL,
    score INTEGER NOT NULL DEFAULT 0 CHECK (score BETWEEN 0 AND 100),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, skill)
);

CREATE INDEX levels_track_id_idx ON levels(track_id);
CREATE INDEX lessons_level_id_idx ON lessons(level_id);
CREATE INDEX lesson_options_lesson_id_idx ON lesson_options(lesson_id);
CREATE INDEX challenges_level_id_idx ON challenges(level_id);
CREATE INDEX challenge_tests_challenge_id_idx ON challenge_tests(challenge_id);
CREATE INDEX attempts_user_id_created_at_idx ON attempts(user_id, created_at DESC);
CREATE INDEX attempts_lesson_id_idx ON attempts(lesson_id);
CREATE INDEX attempts_challenge_id_idx ON attempts(challenge_id);

COMMIT;
