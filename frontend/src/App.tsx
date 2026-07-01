import {
  Navigate,
  NavLink,
  Route,
  Routes,
  useParams,
  type NavLinkRenderProps,
} from 'react-router-dom'
import {
  shellSummary,
  type ChallengeTest,
  type ChallengeTestStatus,
  type LearningBlockStatus,
  type ProfileActivityDay,
  type ProfileCareerStep,
  type ProfileReadinessItem,
  type ProfileSkill,
} from './data/shell'

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
  const challenge = shellSummary.currentChallenge
  const displayChallengeId = challengeId ?? challenge.id

  return (
    <section className="grid gap-4">
      <div className="flex flex-col justify-between gap-4 border-b border-white/10 pb-5 xl:flex-row xl:items-end">
        <div>
          <p className="text-sm uppercase tracking-[0.22em] text-white/40">
            Практика · Code Challenge
          </p>
          <h1 className="mt-3 text-3xl font-semibold tracking-normal sm:text-4xl">
            {challenge.title}
          </h1>
          <p className="mt-3 text-sm text-white/55">
            Challenge ID: {displayChallengeId} · {challenge.difficulty}
          </p>
        </div>
        <div className="grid gap-3 sm:grid-cols-3 xl:w-[520px]">
          <ChallengeMetaCard label="Награда" value={`+${challenge.rewardXp} XP`} />
          <ChallengeMetaCard label="Runtime" value={challenge.sandbox.goVersion} />
          <ChallengeMetaCard label="Sandbox" value={challenge.sandbox.status} />
        </div>
      </div>

      <div className="grid gap-4 2xl:grid-cols-[330px_minmax(0,1fr)_320px]">
        <aside className="grid gap-4 self-start">
          <section className="border border-white/10 p-5">
            <LearningTag>{challenge.difficulty}</LearningTag>
            <h2 className="mt-4 text-2xl font-semibold leading-tight">{challenge.title}</h2>
            <p className="mt-4 text-sm leading-7 text-white/65">{challenge.description}</p>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Требования</p>
            <ul className="mt-4 grid gap-3 text-sm leading-6 text-white/70">
              {challenge.requirements.map((requirement) => (
                <li key={requirement} className="flex gap-3">
                  <span className="mt-2 size-1.5 shrink-0 bg-white" />
                  <span>{requirement}</span>
                </li>
              ))}
            </ul>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Сигнатура</p>
            <pre className="mt-3 overflow-x-auto border border-white/10 bg-white/[0.03] p-3 text-sm text-white/75">
              <code>{challenge.signature}</code>
            </pre>
            <p className="mt-5 text-sm text-white/45">Типы</p>
            <pre className="mt-3 overflow-x-auto border border-white/10 bg-white/[0.03] p-3 text-sm leading-6 text-white/75">
              <code>{challenge.types}</code>
            </pre>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Подсказка</p>
            <p className="mt-3 text-sm leading-6 text-white/65">{challenge.hint}</p>
          </section>
        </aside>

        <section className="grid gap-4">
          <section className="border border-white/10">
            <div className="flex flex-wrap items-center justify-between gap-3 border-b border-white/10 px-4 py-3">
              <div className="flex items-center gap-2">
                <span className="border border-white/15 px-3 py-1 text-sm font-medium">
                  solution.go
                </span>
                <span className="text-sm text-white/45">{challenge.sandbox.goVersion}</span>
              </div>
              <div className="flex items-center gap-2">
                <button type="button" className="border border-white/15 px-4 py-2 text-sm">
                  Run
                </button>
                <button
                  type="button"
                  className="border border-white bg-white px-4 py-2 text-sm font-semibold text-black"
                >
                  Submit
                </button>
              </div>
            </div>
            <pre className="min-h-[520px] overflow-x-auto p-5 text-sm leading-6 text-white/75">
              <code>{challenge.starterCode}</code>
            </pre>
          </section>

          <section className="border border-white/10">
            <div className="flex items-center justify-between border-b border-white/10 px-4 py-3">
              <p className="text-sm font-medium">Терминал</p>
              <span className="text-sm text-white/45">go test ./...</span>
            </div>
            <pre className="overflow-x-auto p-5 text-sm leading-6 text-white/70">
              <code>{challenge.terminalOutput.join('\n')}</code>
            </pre>
          </section>
        </section>

        <aside className="grid gap-4 self-start">
          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Тесты</p>
            <ProgressBar
              value={(challenge.result.passedTests / challenge.result.totalTests) * 100}
              compact
            />
            <ChallengeTestPanel title="public tests" tests={challenge.publicTests} />
            <ChallengeTestPanel title="hidden tests" tests={challenge.hiddenTests} />
          </section>

          <section className="border border-white/10 p-5">
            <div className="flex items-start gap-4">
              <div className="flex size-14 shrink-0 items-center justify-center border border-white text-3xl">
                ✓
              </div>
              <div>
                <p className="text-xl font-semibold">{challenge.result.title}</p>
                <p className="mt-2 text-sm leading-6 text-white/60">
                  {challenge.result.description}
                </p>
              </div>
            </div>
            <dl className="mt-5 grid grid-cols-2 gap-3 text-sm">
              <CheckpointStat
                label="Тесты"
                value={`${challenge.result.passedTests} / ${challenge.result.totalTests}`}
              />
              <CheckpointStat label="Награда" value={`+${challenge.rewardXp} XP`} />
            </dl>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Sandbox limits</p>
            <div className="mt-4 grid gap-2 text-sm text-white/65">
              <ChallengeLimit label="Network" value={challenge.sandbox.network} />
              <ChallengeLimit label="Memory" value={challenge.sandbox.memoryLimit} />
              <ChallengeLimit label="Timeout" value={challenge.sandbox.timeout} />
            </div>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">История попыток</p>
            <div className="mt-4 grid gap-3">
              {challenge.attempts.map((attempt) => (
                <div
                  key={`${attempt.timestamp}-${attempt.summary}`}
                  className="flex items-start justify-between gap-4 border-b border-white/10 pb-3 last:border-b-0 last:pb-0"
                >
                  <div>
                    <p className="text-sm font-medium">{attempt.timestamp}</p>
                    <p className="mt-1 text-sm text-white/55">{attempt.summary}</p>
                  </div>
                  <span className={challengeAttemptClassName(attempt.status)}>
                    +{attempt.xp} XP
                  </span>
                </div>
              ))}
            </div>
          </section>
        </aside>
      </div>
    </section>
  )
}

function ChallengeMetaCard({ label, value }: { label: string; value: string }) {
  return (
    <div className="border border-white/10 p-3">
      <p className="text-xs text-white/45">{label}</p>
      <p className="mt-2 text-sm font-semibold">{value}</p>
    </div>
  )
}

function ChallengeTestPanel({ tests, title }: { tests: ChallengeTest[]; title: string }) {
  const passedTests = tests.filter((test) => test.status === 'passed').length

  return (
    <div className="mt-5">
      <div className="flex items-center justify-between gap-4">
        <p className="text-sm font-medium">{title}</p>
        <span className="text-sm text-white/45">
          {passedTests} / {tests.length}
        </span>
      </div>
      <div className="mt-3 grid gap-2">
        {tests.map((test) => (
          <div key={test.name} className={challengeTestClassName(test.status)}>
            <span className="flex size-6 shrink-0 items-center justify-center border border-white/30 text-xs">
              {challengeTestStatusLabel(test.status)}
            </span>
            <span className="min-w-0 flex-1 truncate text-sm">{test.name}</span>
            <span className="text-xs text-white/45">
              {typeof test.durationMs === 'number' ? `${test.durationMs}ms` : '-'}
            </span>
          </div>
        ))}
      </div>
    </div>
  )
}

function ChallengeLimit({ label, value }: { label: string; value: string }) {
  return (
    <div className="flex items-center justify-between gap-4 border border-white/10 px-3 py-2">
      <span>{label}</span>
      <span className="text-white/45">{value}</span>
    </div>
  )
}

function challengeTestStatusLabel(status: ChallengeTestStatus) {
  const labels: Record<ChallengeTestStatus, string> = {
    failed: '!',
    locked: '-',
    passed: '✓',
  }

  return labels[status]
}

function challengeTestClassName(status: ChallengeTestStatus) {
  const baseClass = 'flex items-center gap-3 border px-3 py-2'

  if (status === 'passed') {
    return `${baseClass} border-white/15 text-white/75`
  }

  if (status === 'locked') {
    return `${baseClass} border-white/10 text-white/40`
  }

  return `${baseClass} border-white/25 text-white`
}

function challengeAttemptClassName(status: 'passed' | 'failed') {
  const baseClass = 'shrink-0 border px-2 py-1 text-xs'

  return status === 'passed'
    ? `${baseClass} border-white bg-white text-black`
    : `${baseClass} border-white/15 text-white/50`
}

function ProfileView() {
  const profile = shellSummary.profile
  const totalWeekXp = profile.weeklyXp.reduce((total, item) => total + item.xp, 0)

  return (
    <section className="grid gap-6">
      <div className="flex flex-col justify-between gap-4 border-b border-white/10 pb-6 xl:flex-row xl:items-end">
        <div>
          <p className="text-sm uppercase tracking-[0.22em] text-white/40">Profile</p>
          <h1 className="mt-3 text-3xl font-semibold tracking-normal sm:text-4xl">Профиль</h1>
          <p className="mt-3 max-w-2xl text-sm leading-6 text-white/60">
            Личный прогресс, streak, навыки и готовность к следующему карьерному уровню.
          </p>
        </div>
        <div className="grid gap-3 sm:grid-cols-3 xl:w-[520px]">
          <ChallengeMetaCard label="Уровень" value={profile.user.level} />
          <ChallengeMetaCard label="XP" value={`${shellSummary.xp.toLocaleString('ru-RU')} XP`} />
          <ChallengeMetaCard label="Стрик" value={`${shellSummary.streakDays} дней`} />
        </div>
      </div>

      <div className="grid gap-4 2xl:grid-cols-[300px_minmax(0,1fr)_360px]">
        <aside className="grid gap-4 self-start">
          <section className="border border-white/10 p-5">
            <div className="flex size-24 items-center justify-center rounded-full border border-white text-3xl font-semibold">
              {profile.user.initials}
            </div>
            <h2 className="mt-5 text-2xl font-semibold">{profile.user.name}</h2>
            <p className="mt-1 text-sm text-white/45">{profile.user.handle}</p>
            <div className="mt-4 w-fit border border-white px-3 py-2 text-sm font-semibold">
              {profile.user.level}
            </div>
            <div className="mt-5">
              <p className="text-3xl font-semibold">{shellSummary.xp.toLocaleString('ru-RU')} XP</p>
              <p className="mt-1 text-sm text-white/45">
                из {shellSummary.nextLevelXp.toLocaleString('ru-RU')} до следующего уровня
              </p>
              <ProgressBar
                value={Math.round((shellSummary.xp / shellSummary.nextLevelXp) * 100)}
                compact
              />
            </div>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Текущий streak</p>
            <p className="mt-3 text-2xl font-semibold">{shellSummary.streakDays} дней подряд</p>
            <div className="mt-4 grid grid-cols-7 gap-2">
              {['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс'].map((day, index) => (
                <div key={day} className="grid gap-2 text-center">
                  <span className="text-xs text-white/45">{day}</span>
                  <span className={profileStreakDayClassName(index < 6)}>
                    {index < 6 ? '✓' : ''}
                  </span>
                </div>
              ))}
            </div>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Итоги</p>
            <div className="mt-4 grid gap-3">
              <ProfileSummaryStat
                label="Завершенные уроки"
                value={profile.totals.completedLessons}
              />
              <ProfileSummaryStat
                label="Практические задания"
                value={profile.totals.practiceTasks}
              />
              <ProfileSummaryStat label="Проекты" value={profile.totals.projectsCompleted} />
              <ProfileSummaryStat
                label="Checkpoint"
                value={`${profile.totals.checkpointsCompleted} / ${profile.totals.checkpointsTotal}`}
              />
            </div>
            <div className="mt-5 flex items-center justify-between border-t border-white/10 pt-4 text-sm">
              <span className="text-white/45">Вступил</span>
              <span>{profile.user.joinedAt}</span>
            </div>
          </section>
        </aside>

        <section className="grid gap-4">
          <section className="grid gap-4 xl:grid-cols-[360px_minmax(0,1fr)]">
            <div className="border border-white/10 p-5">
              <div className="flex items-center justify-between gap-4">
                <div>
                  <p className="text-sm text-white/45">Skill graph</p>
                  <h2 className="mt-2 text-xl font-semibold">Навыки backend</h2>
                </div>
                <LearningTag>30D</LearningTag>
              </div>
              <ProfileSkillGraph skills={profile.skills} />
            </div>

            <div className="border border-white/10 p-5">
              <p className="text-sm text-white/45">Прогресс по технологиям</p>
              <div className="mt-5 grid gap-3">
                {profile.skills.map((skill) => (
                  <ProfileSkillBar key={skill.name} skill={skill} />
                ))}
              </div>
            </div>
          </section>

          <section className="border border-white/10 p-5">
            <div className="flex flex-col justify-between gap-2 sm:flex-row sm:items-start">
              <div>
                <p className="text-sm text-white/45">XP за неделю</p>
                <h2 className="mt-2 text-xl font-semibold">Недельный темп</h2>
              </div>
              <div className="text-left sm:text-right">
                <p className="text-2xl font-semibold">{totalWeekXp.toLocaleString('ru-RU')} XP</p>
                <p className="mt-1 text-sm text-white/45">на этой неделе</p>
              </div>
            </div>
            <ProfileWeeklyChart items={profile.weeklyXp} />
          </section>

          <section className="grid gap-4 xl:grid-cols-[280px_minmax(0,1fr)]">
            <div className="border border-white/10 p-5">
              <p className="text-sm text-white/45">Карьерный путь</p>
              <div className="mt-5 grid gap-4">
                {profile.careerPath.map((step) => (
                  <ProfileCareerPathItem key={step.title} step={step} />
                ))}
              </div>
            </div>

            <div className="border border-white/10 p-5">
              <div className="flex items-start justify-between gap-4">
                <div>
                  <p className="text-sm text-white/45">
                    Готовность к {profile.readiness.targetLevel}
                  </p>
                  <h2 className="mt-2 text-xl font-semibold">
                    {profile.readiness.progressPercent}%
                  </h2>
                </div>
                <LearningTag>{profile.user.level}</LearningTag>
              </div>
              <ProgressBar value={profile.readiness.progressPercent} />
              <div className="mt-5 grid gap-3">
                {profile.readiness.items.map((item) => (
                  <ProfileReadinessRow key={item.title} item={item} />
                ))}
              </div>
            </div>
          </section>
        </section>

        <aside className="grid gap-4 self-start">
          <section className="border border-white/10 p-5">
            <div className="flex items-center justify-between gap-4">
              <p className="text-sm text-white/45">Календарь активности</p>
              <span className="text-sm text-white/60">Май 2024</span>
            </div>
            <ProfileActivityCalendar days={profile.activityCalendar} />
            <div className="mt-4 flex items-center justify-between border-t border-white/10 pt-4 text-sm">
              <span>{shellSummary.streakDays} дней подряд</span>
              <span className="text-white/45">Лучший streak: 11 дней</span>
            </div>
          </section>

          <section className="border border-white/10 p-5">
            <div className="flex items-center justify-between gap-4">
              <p className="text-sm text-white/45">Недавние достижения</p>
              <span className="text-sm text-white/45">Смотреть все</span>
            </div>
            <div className="mt-5 grid gap-4">
              {profile.achievements.map((achievement) => (
                <ProfileAchievementItem key={achievement.title} achievement={achievement} />
              ))}
            </div>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">Следующая цель</p>
            <div className="mt-4 border border-white/10 p-4">
              <h2 className="text-lg font-semibold">{profile.nextGoal.title}</h2>
              <p className="mt-2 text-sm text-white/55">{profile.nextGoal.description}</p>
              <ProgressBar value={profile.nextGoal.progressPercent} compact />
              <p className="mt-3 text-sm text-white/55">
                {profile.nextGoal.currentXp.toLocaleString('ru-RU')} /{' '}
                {profile.nextGoal.targetXp.toLocaleString('ru-RU')} XP
              </p>
            </div>
          </section>

          <section className="border border-white/10 p-5">
            <p className="text-sm text-white/45">{profile.recommendation.title}</p>
            <p className="mt-3 text-sm leading-6 text-white/65">
              {profile.recommendation.description}
            </p>
            <NavLink
              to={profile.recommendation.to}
              className="mt-5 inline-flex border border-white/15 px-4 py-2 text-sm text-white/75"
            >
              Открыть практику
            </NavLink>
          </section>
        </aside>
      </div>
    </section>
  )
}

function profileStreakDayClassName(isCompleted: boolean) {
  const baseClass = 'border py-1 text-xs'

  return isCompleted
    ? `${baseClass} border-white bg-white text-black`
    : `${baseClass} border-dashed border-white/35 text-white/45`
}

function ProfileSummaryStat({ label, value }: { label: string; value: number | string }) {
  return (
    <div className="flex items-center justify-between gap-4 border border-white/10 px-3 py-2 text-sm">
      <span className="text-white/55">{label}</span>
      <span className="font-semibold">{value}</span>
    </div>
  )
}

function ProfileSkillGraph({ skills }: { skills: ProfileSkill[] }) {
  const polygonPoints = skills
    .map((skill, index) => {
      const point = profileGraphPoint(index, skills.length, skill.score)
      return `${point.x},${point.y}`
    })
    .join(' ')

  return (
    <svg viewBox="0 0 260 260" role="img" aria-label="Skill graph" className="mt-5 w-full">
      {[25, 50, 75, 100].map((value) => (
        <polygon
          key={value}
          points={skills
            .map((_, index) => {
              const point = profileGraphPoint(index, skills.length, value)
              return `${point.x},${point.y}`
            })
            .join(' ')}
          fill="none"
          stroke="rgba(255,255,255,0.12)"
          strokeWidth="1"
        />
      ))}
      {skills.map((skill, index) => {
        const outerPoint = profileGraphPoint(index, skills.length, 100)
        const valuePoint = profileGraphPoint(index, skills.length, skill.score)
        const labelPoint = profileGraphPoint(index, skills.length, 118)

        return (
          <g key={skill.name}>
            <line
              x1="130"
              y1="130"
              x2={outerPoint.x}
              y2={outerPoint.y}
              stroke="rgba(255,255,255,0.14)"
            />
            <circle cx={valuePoint.x} cy={valuePoint.y} r="3.5" fill="white" />
            <text
              x={labelPoint.x}
              y={labelPoint.y}
              fill="rgba(255,255,255,0.7)"
              fontSize="9"
              textAnchor="middle"
              dominantBaseline="middle"
            >
              {skill.name}
            </text>
          </g>
        )
      })}
      <polygon
        points={polygonPoints}
        fill="rgba(255,255,255,0.2)"
        stroke="white"
        strokeWidth="2"
      />
    </svg>
  )
}

function profileGraphPoint(index: number, total: number, value: number) {
  const angle = -Math.PI / 2 + (index / total) * Math.PI * 2
  const radius = 82 * (value / 100)

  return {
    x: Number((130 + Math.cos(angle) * radius).toFixed(2)),
    y: Number((130 + Math.sin(angle) * radius).toFixed(2)),
  }
}

function ProfileSkillBar({ skill }: { skill: ProfileSkill }) {
  return (
    <div className="grid gap-2">
      <div className="flex items-center justify-between gap-4 text-sm">
        <span>{skill.name}</span>
        <span className="text-white/55">{skill.score} / 100</span>
      </div>
      <ProgressBar value={skill.score} compact />
    </div>
  )
}

function ProfileWeeklyChart({ items }: { items: { day: string; xp: number }[] }) {
  const maxXp = Math.max(...items.map((item) => item.xp))

  return (
    <div className="mt-6 flex h-44 items-end gap-3 border-b border-white/15 pb-2">
      {items.map((item) => (
        <div key={item.day} className="flex flex-1 flex-col items-center gap-2">
          <span className="text-xs text-white/55">{item.xp}</span>
          <div
            className="w-full bg-white"
            style={{ height: `${Math.max(18, (item.xp / maxXp) * 120)}px` }}
            title={`${item.day}: ${item.xp} XP`}
          />
          <span className="text-xs text-white/45">{item.day}</span>
        </div>
      ))}
    </div>
  )
}

function ProfileCareerPathItem({ step }: { step: ProfileCareerStep }) {
  return (
    <div className="flex gap-3">
      <span className={profileCareerDotClassName(step.status)} />
      <div>
        <p className="text-sm font-semibold">{step.title}</p>
        <p className="mt-1 text-sm text-white/45">{step.detail}</p>
      </div>
    </div>
  )
}

function profileCareerDotClassName(status: ProfileCareerStep['status']) {
  const baseClass = 'mt-1 size-4 shrink-0 border'

  if (status === 'completed') {
    return `${baseClass} border-white bg-white`
  }

  if (status === 'active') {
    return `${baseClass} border-white bg-black`
  }

  return `${baseClass} border-white/25 bg-black`
}

function ProfileReadinessRow({ item }: { item: ProfileReadinessItem }) {
  return (
    <div className="grid gap-2 border-b border-white/10 pb-3 last:border-b-0 last:pb-0">
      <div className="flex items-center justify-between gap-4 text-sm">
        <span>{item.title}</span>
        <span className="text-white/55">
          {item.current.toLocaleString('ru-RU')} / {item.target.toLocaleString('ru-RU')}
        </span>
      </div>
      <div className="flex items-center gap-3">
        <span className={item.isComplete ? 'size-3 bg-white' : 'size-3 border border-white/40'} />
        <div className="h-2 flex-1 border border-white/10 bg-white/10">
          <div
            className="h-full bg-white"
            style={{ width: `${Math.min(100, Math.round((item.current / item.target) * 100))}%` }}
          />
        </div>
      </div>
    </div>
  )
}

function ProfileActivityCalendar({ days }: { days: ProfileActivityDay[] }) {
  return (
    <div className="mt-5 grid grid-cols-7 gap-2">
      {days.slice(0, 7).map((day) => (
        <span
          key={`label-${day.label}-${day.value}`}
          className="text-center text-xs text-white/45"
        >
          {day.label}
        </span>
      ))}
      {days.map((day, index) => (
        <span
          key={`${day.value}-${index}`}
          className={profileActivityDayClassName(day.status)}
          title={`${day.label} ${day.value}`}
        >
          {day.status === 'completed' ? '✓' : day.value}
        </span>
      ))}
    </div>
  )
}

function profileActivityDayClassName(status: ProfileActivityDay['status']) {
  const baseClass = 'flex aspect-square items-center justify-center border text-xs'

  if (status === 'completed') {
    return `${baseClass} border-white bg-white text-black`
  }

  if (status === 'active') {
    return `${baseClass} border-2 border-white text-white`
  }

  if (status === 'planned') {
    return `${baseClass} border-dashed border-white/50 text-white`
  }

  return `${baseClass} border-white/10 text-white/45`
}

function ProfileAchievementItem({
  achievement,
}: {
  achievement: (typeof shellSummary.profile.achievements)[number]
}) {
  return (
    <div className="flex items-start justify-between gap-4 border-b border-white/10 pb-4 last:border-b-0 last:pb-0">
      <div>
        <p className="text-sm font-semibold leading-6">{achievement.title}</p>
        <p className="mt-1 text-sm leading-5 text-white/50">{achievement.description}</p>
      </div>
      <span className="shrink-0 text-xs text-white/45">{achievement.date}</span>
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
