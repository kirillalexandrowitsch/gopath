import {
  Navigate,
  NavLink,
  Route,
  Routes,
  useParams,
  type NavLinkRenderProps,
} from 'react-router-dom'
import { shellSummary } from './data/shell'

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
  return (
    <ShellPage
      eyebrow="Learning Path"
      title="Карта обучения"
      description="Здесь будет вертикальная карьерная карта с active/completed/locked состояниями и checkpoints."
    >
      <div className="grid gap-3">
        {shellSummary.levels.map((level, index) => (
          <div key={level.title} className="flex items-center justify-between border border-white/10 p-4">
            <div>
              <p className="text-xs text-white/45">Уровень {index + 1}</p>
              <h2 className="mt-1 text-xl font-semibold">{level.title}</h2>
            </div>
            <span className="border border-white/15 px-3 py-1 text-sm text-white/65">
              {level.status === 'completed'
                ? 'Завершено'
                : level.status === 'active'
                  ? 'Активно'
                  : 'Заблокировано'}
            </span>
          </div>
        ))}
      </div>
    </ShellPage>
  )
}

function LessonView() {
  const { lessonId } = useParams()

  return (
    <ShellPage
      eyebrow="Lesson"
      title="Урок"
      description="Каркас обычного урока с будущими вопросами, вариантами ответа, feedback и XP."
    >
      <PlaceholderPanel
        label="Lesson ID"
        value={lessonId ?? shellSummary.currentLesson.id}
        description="На этом маршруте позже появится интерактивный урок."
      />
    </ShellPage>
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
