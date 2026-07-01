import {
  Navigate,
  NavLink,
  Route,
  Routes,
  useParams,
  type NavLinkRenderProps,
} from 'react-router-dom'
import { shellSummary, type LearningBlockStatus } from './data/shell'

const primaryNav = [
  { to: '/app', label: 'Главная' },
  { to: '/learn', label: 'Карта обучения' },
  { to: `/lesson/${shellSummary.currentLesson.id}`, label: 'Уроки' },
  { to: '/challenge/retry-context', label: 'Практика' },
  { to: '/profile', label: 'Профиль' },
]

const shellMetrics = [
  { label: 'Уровень', value: shellSummary.level },
  { label: 'XP', value: `${shellSummary.xp.toLocaleString('ru-RU')} XP` },
  { label: 'Стрик', value: `${shellSummary.streakDays} дней подряд` },
]

function App() {
  return (
    <div className="min-h-screen bg-black text-white">
      <div className="border-b border-white/10">
        <header className="mx-auto flex min-h-16 max-w-7xl items-center justify-between gap-6 px-4 py-3 sm:px-6 lg:px-8">
          <NavLink to="/app" className="flex items-center gap-3">
            <span className="flex size-10 items-center justify-center border border-white/25 bg-white text-sm font-semibold text-black">
              GP
            </span>
            <span className="text-lg font-semibold">GoPath</span>
          </NavLink>
          <nav className="hidden items-center gap-1 lg:flex" aria-label="Основная навигация">
            {primaryNav.map((item) => (
              <ShellNavLink key={item.to} to={item.to}>
                {item.label}
              </ShellNavLink>
            ))}
          </nav>
          <div className="flex items-center gap-4 text-sm text-white/70">
            <span>{shellSummary.streakDays} дней подряд</span>
            <span>{shellSummary.xp.toLocaleString('ru-RU')} XP</span>
          </div>
        </header>
      </div>

      <div className="mx-auto grid max-w-7xl lg:grid-cols-[260px_1fr]">
        <aside className="border-b border-white/10 px-4 py-4 sm:px-6 lg:min-h-[calc(100vh-65px)] lg:border-b-0 lg:border-r lg:px-5">
          <nav className="grid gap-2" aria-label="Разделы приложения">
            {primaryNav.map((item) => (
              <ShellNavLink key={item.to} to={item.to} variant="side">
                {item.label}
              </ShellNavLink>
            ))}
          </nav>

          <section className="mt-8 hidden border-t border-white/10 pt-6 lg:block">
            <p className="text-xs uppercase tracking-[0.2em] text-white/40">Прогресс</p>
            <div className="mt-4 grid gap-3">
              {shellMetrics.map((metric) => (
                <div key={metric.label} className="border border-white/10 p-3">
                  <p className="text-xs text-white/45">{metric.label}</p>
                  <p className="mt-1 text-sm font-medium">{metric.value}</p>
                </div>
              ))}
            </div>
          </section>
        </aside>

        <main className="min-h-[calc(100vh-65px)] px-4 py-6 sm:px-6 lg:px-8">
          <Routes>
            <Route path="/" element={<Navigate to="/app" replace />} />
            <Route path="/app" element={<DashboardView />} />
            <Route path="/learn" element={<LearningPathView />} />
            <Route path="/lesson/:lessonId" element={<LessonView />} />
            <Route path="/challenge/:challengeId" element={<ChallengeView />} />
            <Route path="/profile" element={<ProfileView />} />
            <Route path="*" element={<Navigate to="/app" replace />} />
          </Routes>
        </main>
      </div>
    </div>
  )
}

function ShellNavLink({
  children,
  to,
  variant = 'top',
}: {
  children: string
  to: string
  variant?: 'top' | 'side'
}) {
  const baseClass =
    variant === 'side'
      ? 'block border px-4 py-3 text-sm transition'
      : 'border px-3 py-2 text-sm transition'

  return (
    <NavLink to={to} className={(state) => navClassName(state, baseClass)}>
      {children}
    </NavLink>
  )
}

function navClassName({ isActive }: NavLinkRenderProps, baseClass: string) {
  return [
    baseClass,
    isActive
      ? 'border-white bg-white text-black'
      : 'border-white/10 text-white/65 hover:border-white/35 hover:text-white',
  ].join(' ')
}

function DashboardView() {
  const dailyGoalPercent = Math.round(
    (shellSummary.dailyGoal.earnedXp / shellSummary.dailyGoal.targetXp) * 100,
  )
  const levelProgressPercent = Math.round((shellSummary.xp / shellSummary.nextLevelXp) * 100)

  return (
    <section className="grid gap-6">
      <div className="flex flex-col justify-between gap-4 border-b border-white/10 pb-6 lg:flex-row lg:items-end">
        <div>
          <p className="text-sm uppercase tracking-[0.22em] text-white/40">Dashboard</p>
          <h1 className="mt-3 text-3xl font-semibold tracking-normal sm:text-5xl">Главная</h1>
          <p className="mt-3 max-w-2xl text-base leading-7 text-white/60">
            Продолжайте путь к уровню Middle Backend через короткие уроки, практику и
            checkpoints.
          </p>
        </div>
        <NavLink
          to={`/lesson/${shellSummary.currentLesson.id}`}
          className="inline-flex w-fit border border-white bg-white px-5 py-3 text-sm font-semibold text-black"
        >
          Продолжить обучение
        </NavLink>
      </div>

      <div className="grid gap-4 xl:grid-cols-[1fr_340px]">
        <div className="grid gap-4">
          <section className="grid gap-4 md:grid-cols-3">
            <DashboardMetric
              label="Ваш уровень"
              value={shellSummary.level}
              detail={`${shellSummary.xp.toLocaleString('ru-RU')} / ${shellSummary.nextLevelXp.toLocaleString('ru-RU')} XP`}
              progress={levelProgressPercent}
            />
            <DashboardMetric
              label="Стрик"
              value={`${shellSummary.streakDays} дней подряд`}
              detail="Текущая серия занятий"
            />
            <DashboardMetric
              label="Сегодняшняя цель"
              value={`${shellSummary.dailyGoal.earnedXp} / ${shellSummary.dailyGoal.targetXp} XP`}
              detail={`${shellSummary.dailyGoal.completedTasks} из ${shellSummary.dailyGoal.targetTasks} задач выполнено`}
              progress={dailyGoalPercent}
            />
          </section>

          <section className="border border-white/10 p-5">
            <div className="flex flex-col justify-between gap-4 lg:flex-row lg:items-center">
              <div>
                <p className="text-sm text-white/45">Текущий урок</p>
                <h2 className="mt-2 text-2xl font-semibold">{shellSummary.currentLesson.title}</h2>
                <p className="mt-2 text-sm text-white/55">{shellSummary.currentLesson.step}</p>
              </div>
              <div className="border border-white/15 px-4 py-3 text-sm text-white/70">
                +{shellSummary.currentLesson.xpReward} XP
              </div>
            </div>
            <ProgressBar value={shellSummary.currentLesson.progressPercent} />
            <div className="mt-5 flex flex-wrap gap-3">
              <NavLink
                to={`/lesson/${shellSummary.currentLesson.id}`}
                className="border border-white bg-white px-4 py-2 text-sm font-medium text-black"
              >
                Открыть урок
              </NavLink>
              <NavLink
                to="/learn"
                className="border border-white/15 px-4 py-2 text-sm font-medium text-white/75"
              >
                Карта обучения
              </NavLink>
            </div>
          </section>

          <section className="grid gap-4 lg:grid-cols-[1fr_1fr]">
            <div className="border border-white/10 p-5">
              <div className="flex items-center justify-between gap-4">
                <div>
                  <p className="text-sm text-white/45">Слабые темы</p>
                  <h2 className="mt-2 text-xl font-semibold">Приоритет для практики</h2>
                </div>
                <NavLink to="/challenge/retry-context" className="text-sm text-white/60">
                  Практика
                </NavLink>
              </div>
              <div className="mt-5 grid gap-3">
                {shellSummary.weakTopics.map((topic) => (
                  <div key={topic.title} className="grid gap-2 border border-white/10 p-3">
                    <div className="flex items-center justify-between gap-4">
                      <span className="font-medium">{topic.title}</span>
                      <span className="text-sm text-white/55">{topic.action}</span>
                    </div>
                    <ProgressBar value={topic.confidencePercent} compact />
                  </div>
                ))}
              </div>
            </div>

            <div className="border border-white/10 p-5">
              <p className="text-sm text-white/45">Следующий checkpoint</p>
              <h2 className="mt-2 text-xl font-semibold">{shellSummary.checkpoint.title}</h2>
              <p className="mt-3 text-sm leading-6 text-white/60">
                {shellSummary.checkpoint.description}
              </p>
              <dl className="mt-5 grid grid-cols-3 gap-2 text-sm">
                <CheckpointStat label="Задач" value={String(shellSummary.checkpoint.tasksTotal)} />
                <CheckpointStat label="Награда" value={`+${shellSummary.checkpoint.rewardXp} XP`} />
                <CheckpointStat label="Сложность" value={shellSummary.checkpoint.difficulty} />
              </dl>
              <NavLink
                to="/learn"
                className="mt-5 inline-flex border border-white/15 px-4 py-2 text-sm font-medium text-white/75"
              >
                Подробнее
              </NavLink>
            </div>
          </section>
        </div>

        <aside className="grid gap-4">
          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Активность за неделю</p>
            <div className="mt-5 flex h-36 items-end gap-2 border-b border-white/15 pb-2">
              {shellSummary.weeklyActivity.map((item) => (
                <div key={item.day} className="flex flex-1 flex-col items-center gap-2">
                  <div
                    className="w-full bg-white"
                    style={{ height: `${Math.max(18, item.xp / 2)}px` }}
                    title={`${item.day}: ${item.xp} XP`}
                  />
                  <span className="text-xs text-white/45">{item.day}</span>
                </div>
              ))}
            </div>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Последняя активность</p>
            <div className="mt-5 grid gap-3">
              {shellSummary.recentActivity.map((activity) => (
                <div key={`${activity.timestamp}-${activity.title}`} className="border-b border-white/10 pb-3 last:border-b-0 last:pb-0">
                  <div className="flex items-start justify-between gap-4">
                    <div>
                      <p className="text-sm font-medium leading-6">{activity.title}</p>
                      <p className="mt-1 text-xs text-white/45">{activity.timestamp}</p>
                    </div>
                    <span className="shrink-0 text-sm text-white/70">+{activity.xp} XP</span>
                  </div>
                </div>
              ))}
            </div>
          </section>
        </aside>
      </div>
    </section>
  )
}

function LearningPathView() {
  const activeBlock = shellSummary.learningPath.activeBlock

  return (
    <section className="grid gap-6">
      <div className="flex flex-col justify-between gap-4 border-b border-white/10 pb-6 lg:flex-row lg:items-end">
        <div>
          <p className="text-sm uppercase tracking-[0.22em] text-white/40">Learning Path</p>
          <h1 className="mt-3 text-3xl font-semibold tracking-normal sm:text-5xl">
            Карта обучения
          </h1>
          <p className="mt-3 max-w-2xl text-base leading-7 text-white/60">
            Карьерная карта backend-разработчика: от стажера до Senior через уроки,
            практику и checkpoints.
          </p>
        </div>
        <div className="border border-white/10 p-4 lg:w-72">
          <p className="text-sm text-white/45">Готовность к Middle</p>
          <p className="mt-2 text-2xl font-semibold">
            {shellSummary.learningPath.nextLevelProgressPercent}%
          </p>
          <ProgressBar value={shellSummary.learningPath.nextLevelProgressPercent} compact />
        </div>
      </div>

      <div className="grid gap-4 xl:grid-cols-[1fr_360px]">
        <section className="grid gap-6">
          {shellSummary.learningPath.stages.map((stage, stageIndex) => (
            <div key={stage.title} className="grid gap-4 lg:grid-cols-[190px_1fr]">
              <div className="border-t border-white/10 pt-4">
                <p className="text-sm font-semibold">{stage.title}</p>
                <p className="mt-1 text-sm text-white/45">{stage.subtitle}</p>
                <p className="mt-3 text-xs uppercase tracking-[0.18em] text-white/35">
                  {stage.completedBlocks} / {stage.totalBlocks} блоков
                </p>
              </div>
              <div className="relative grid gap-3 border-l border-white/10 pl-5">
                <div className="absolute -left-[5px] top-5 size-2.5 border border-white/40 bg-black" />
                {stage.blocks.map((block) => (
                  <LearningBlockCard
                    key={block.id}
                    block={block}
                    isFirstStage={stageIndex === 0}
                    isActive={block.id === activeBlock.id}
                  />
                ))}
              </div>
            </div>
          ))}
        </section>

        <aside className="grid gap-4 self-start xl:sticky xl:top-6">
          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Активный блок</p>
            <h2 className="mt-2 text-2xl font-semibold">{activeBlock.title}</h2>
            <p className="mt-3 text-sm leading-6 text-white/60">{activeBlock.description}</p>
            <ProgressBar value={activeBlock.progressPercent} />

            <div className="mt-5">
              <p className="text-sm text-white/45">Навыки уровня</p>
              <div className="mt-3 flex flex-wrap gap-2">
                {activeBlock.skills.map((skill) => (
                  <LearningTag key={skill}>{skill}</LearningTag>
                ))}
              </div>
            </div>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Требования</p>
            <ul className="mt-3 grid gap-3 text-sm text-white/70">
              {activeBlock.requirements.map((requirement) => (
                <li key={requirement} className="flex items-center gap-3">
                  <span className="size-2 border border-white/40 bg-white" />
                  {requirement}
                </li>
              ))}
            </ul>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Карьерные результаты</p>
            <ul className="mt-3 grid gap-3 text-sm text-white/70">
              {activeBlock.careerOutcomes.map((outcome) => (
                <li key={outcome} className="border border-white/10 px-3 py-2">
                  {outcome}
                </li>
              ))}
            </ul>
          </section>

          <NavLink
            to={`/lesson/${shellSummary.currentLesson.id}`}
            className="border border-white bg-white px-4 py-3 text-center text-sm font-semibold text-black"
          >
            Перейти к урокам
          </NavLink>
        </aside>
      </div>
    </section>
  )
}

function LearningBlockCard({
  block,
  isActive,
  isFirstStage,
}: {
  block: (typeof shellSummary.learningPath.stages)[number]['blocks'][number]
  isActive: boolean
  isFirstStage: boolean
}) {
  return (
    <div className={learningBlockClassName(block.status, isActive)}>
      <div className="flex flex-col justify-between gap-4 sm:flex-row sm:items-center">
        <div>
          <div className="flex flex-wrap items-center gap-2">
            <h2 className="text-xl font-semibold">{block.title}</h2>
            {isFirstStage ? <span className="text-xs text-white/35">база</span> : null}
          </div>
          <div className="mt-2 flex flex-wrap gap-2">
            {block.tags.map((tag) => (
              <LearningTag key={tag}>{tag}</LearningTag>
            ))}
          </div>
        </div>
        <div className="flex items-center gap-3">
          <span className="border border-white/15 px-3 py-1 text-sm text-white/70">
            {learningStatusLabel(block.status)}
          </span>
          {block.status === 'active' ? (
            <NavLink
              to={`/lesson/${shellSummary.currentLesson.id}`}
              className="border border-white bg-white px-3 py-1 text-sm font-medium text-black"
            >
              Продолжить
            </NavLink>
          ) : null}
        </div>
      </div>
      {typeof block.progressPercent === 'number' ? (
        <ProgressBar value={block.progressPercent} compact />
      ) : null}
      {block.status === 'checkpoint' ? (
        <p className="mt-3 text-sm text-white/55">
          {block.tasksCompleted ?? 0} / {block.tasksTotal ?? 0} заданий
        </p>
      ) : null}
    </div>
  )
}

function LearningTag({ children }: { children: string }) {
  return <span className="border border-white/15 px-2 py-1 text-xs text-white/60">{children}</span>
}

function learningStatusLabel(status: LearningBlockStatus) {
  const labels: Record<LearningBlockStatus, string> = {
    active: 'Активно',
    checkpoint: 'Checkpoint',
    completed: 'Завершено',
    locked: 'Заблокировано',
  }

  return labels[status]
}

function learningBlockClassName(status: LearningBlockStatus, isActive: boolean) {
  const baseClass = 'border p-5 transition'

  if (isActive) {
    return `${baseClass} border-white bg-white/[0.03]`
  }

  if (status === 'locked') {
    return `${baseClass} border-white/10 opacity-60`
  }

  if (status === 'checkpoint') {
    return `${baseClass} border-dashed border-white/30`
  }

  return `${baseClass} border-white/10`
}

function lessonOptionClassName(isSelected: boolean) {
  const baseClass = 'flex flex-col gap-4 border p-4 transition sm:flex-row sm:items-center'

  return isSelected ? `${baseClass} border-white bg-white/[0.03]` : `${baseClass} border-white/10`
}

function LessonView() {
  const { lessonId } = useParams()
  const lesson = shellSummary.currentLesson

  return (
    <section className="grid gap-6">
      <div className="flex flex-col justify-between gap-4 border-b border-white/10 pb-6 xl:flex-row xl:items-end">
        <div>
          <p className="text-sm uppercase tracking-[0.22em] text-white/40">Lesson</p>
          <h1 className="mt-3 text-3xl font-semibold tracking-normal sm:text-5xl">
            {lesson.title}
          </h1>
          <p className="mt-3 text-sm text-white/55">
            {lesson.step} · Lesson ID: {lessonId ?? lesson.id}
          </p>
        </div>
        <div className="grid gap-3 sm:grid-cols-2 xl:w-72">
          <div className="border border-white/10 p-4">
            <p className="text-sm text-white/45">Награда</p>
            <p className="mt-2 text-2xl font-semibold">+{lesson.xpReward} XP</p>
          </div>
          <div className="border border-white/10 p-4">
            <p className="text-sm text-white/45">Попытки</p>
            <p className="mt-2 text-2xl font-semibold">
              {lesson.attemptsRemaining} / {lesson.totalAttempts}
            </p>
          </div>
        </div>
      </div>

      <div className="grid gap-4 xl:grid-cols-[1fr_360px]">
        <section className="grid gap-4">
          <div className="border border-white/10 p-5">
            <div className="flex flex-col justify-between gap-3 sm:flex-row sm:items-center">
              <div>
                <p className="text-sm text-white/45">Прогресс урока</p>
                <p className="mt-2 text-xl font-semibold">{lesson.progressPercent}%</p>
              </div>
              <NavLink
                to="/learn"
                className="w-fit border border-white/15 px-4 py-2 text-sm text-white/70"
              >
                Карта обучения
              </NavLink>
            </div>
            <ProgressBar value={lesson.progressPercent} />
          </div>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Вопрос</p>
            <h2 className="mt-3 text-2xl font-semibold leading-tight">{lesson.question.prompt}</h2>

            <div className="mt-6 grid gap-3">
              {lesson.question.options.map((option) => (
                <div key={option.id} className={lessonOptionClassName(option.isSelected)}>
                  <div className="flex size-9 shrink-0 items-center justify-center border border-white/25 text-sm font-semibold">
                    {option.label}
                  </div>
                  <p className="text-base leading-6 sm:flex-1">{option.text}</p>
                  {option.isCorrect ? (
                    <span className="w-fit shrink-0 border border-white bg-white px-3 py-1 text-sm font-semibold text-black sm:ml-auto">
                      Верно
                    </span>
                  ) : null}
                </div>
              ))}
            </div>

            <div className="mt-6 flex flex-wrap items-center justify-between gap-3">
              <button
                type="button"
                className="border border-white bg-white px-5 py-3 text-sm font-semibold text-black"
              >
                Проверить
              </button>
              <button
                type="button"
                className="border border-white/15 px-5 py-3 text-sm text-white/70"
              >
                Отложить урок
              </button>
            </div>
          </section>

          <section className="grid gap-4 lg:grid-cols-[260px_1fr]">
            <div className="border border-white/10 p-5">
              <div className="flex size-16 items-center justify-center border border-white text-3xl">
                ✓
              </div>
              <h2 className="mt-5 text-2xl font-semibold">{lesson.question.feedbackTitle}</h2>
              <p className="mt-3 text-sm leading-6 text-white/60">{lesson.question.feedbackText}</p>
            </div>

            <div className="border border-white/10 p-5">
              <p className="text-sm text-white/45">Объяснение</p>
              <p className="mt-3 text-sm leading-7 text-white/65">{lesson.question.explanation}</p>
              <div className="mt-5 flex justify-end">
                <NavLink
                  to="/challenge/retry-context"
                  className="border border-white/15 px-4 py-2 text-sm font-medium text-white/75"
                >
                  Перейти к практике
                </NavLink>
              </div>
            </div>
          </section>
        </section>

        <aside className="grid gap-4 self-start xl:sticky xl:top-6">
          <section className="border border-white/10 p-5">
            <p className="text-sm uppercase tracking-[0.18em] text-white/35">
              {lesson.sidebar.conceptTitle}
            </p>
            <div className="mt-4 grid gap-3 text-sm leading-6 text-white/65">
              {lesson.sidebar.conceptText.map((paragraph) => (
                <p key={paragraph}>{paragraph}</p>
              ))}
            </div>
            <pre className="mt-5 overflow-x-auto border border-white/10 bg-white/[0.03] p-4 text-sm leading-6 text-white/75">
              <code>{lesson.sidebar.codeSnippet}</code>
            </pre>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Практика в backend</p>
            <ul className="mt-4 grid gap-3 text-sm text-white/65">
              {lesson.sidebar.practiceNotes.map((note) => (
                <li key={note} className="flex gap-3">
                  <span className="mt-1.5 size-2 shrink-0 border border-white/50 bg-white" />
                  <span>{note}</span>
                </li>
              ))}
            </ul>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Полезно знать</p>
            <p className="mt-3 text-sm leading-6 text-white/65">{lesson.sidebar.usefulNote}</p>
            <div className="mt-5 flex flex-wrap gap-2">
              {lesson.sidebar.relatedTopics.map((topic) => (
                <LearningTag key={topic}>{topic}</LearningTag>
              ))}
            </div>
          </section>
        </aside>
      </div>
    </section>
  )
}

function ChallengeView() {
  const { challengeId } = useParams()

  return (
    <ShellPage
      eyebrow="Code Challenge"
      title="Практика"
      description="Каркас code challenge с будущим редактором, терминалом, public и hidden tests."
    >
      <PlaceholderPanel
        label="Challenge ID"
        value={challengeId ?? 'retry-context'}
        description="Monaco Editor, xterm.js и sandbox integration будут добавлены отдельными commit."
      />
    </ShellPage>
  )
}

function ProfileView() {
  return (
    <ShellPage
      eyebrow="Profile"
      title="Профиль"
      description="Каркас профиля с будущими графиками навыков, календарем streak и рекомендациями."
    >
      <div className="grid gap-4 md:grid-cols-2">
        <SummaryCard label="Пользователь" value={shellSummary.userName} />
        <SummaryCard label="Готовность к Middle" value="67%" />
      </div>
    </ShellPage>
  )
}

function ShellPage({
  children,
  description,
  eyebrow,
  title,
}: {
  children: React.ReactNode
  description: string
  eyebrow: string
  title: string
}) {
  return (
    <section className="grid gap-6">
      <div className="max-w-3xl">
        <p className="text-sm uppercase tracking-[0.22em] text-white/40">{eyebrow}</p>
        <h1 className="mt-3 text-3xl font-semibold tracking-normal sm:text-5xl">{title}</h1>
        <p className="mt-4 text-base leading-7 text-white/60">{description}</p>
      </div>
      {children}
    </section>
  )
}

function SummaryCard({ label, value }: { label: string; value: string }) {
  return (
    <div className="border border-white/10 p-5">
      <p className="text-sm text-white/45">{label}</p>
      <p className="mt-3 text-2xl font-semibold">{value}</p>
    </div>
  )
}

function DashboardMetric({
  detail,
  label,
  progress,
  value,
}: {
  detail: string
  label: string
  progress?: number
  value: string
}) {
  return (
    <div className="border border-white/10 p-5">
      <p className="text-sm text-white/45">{label}</p>
      <p className="mt-3 text-2xl font-semibold leading-tight">{value}</p>
      <p className="mt-2 text-sm text-white/55">{detail}</p>
      {typeof progress === 'number' ? <ProgressBar value={progress} compact /> : null}
    </div>
  )
}

function CheckpointStat({ label, value }: { label: string; value: string }) {
  return (
    <div className="border border-white/10 p-3">
      <dt className="text-xs text-white/45">{label}</dt>
      <dd className="mt-1 font-medium">{value}</dd>
    </div>
  )
}

function PlaceholderPanel({
  description,
  label,
  value,
}: {
  description: string
  label: string
  value: string
}) {
  return (
    <div className="border border-dashed border-white/20 p-5">
      <p className="text-sm text-white/45">{label}</p>
      <p className="mt-2 font-mono text-lg text-white">{value}</p>
      <p className="mt-4 max-w-2xl text-sm leading-6 text-white/60">{description}</p>
    </div>
  )
}

function ProgressBar({ compact = false, value }: { compact?: boolean; value: number }) {
  return (
    <div className={compact ? 'mt-3' : 'mt-4'}>
      <div className="h-2 border border-white/10 bg-white/10">
        <div className="h-full bg-white" style={{ width: `${value}%` }} />
      </div>
      {!compact ? <p className="mt-2 text-sm text-white/50">{value}%</p> : null}
    </div>
  )
}

export default App
