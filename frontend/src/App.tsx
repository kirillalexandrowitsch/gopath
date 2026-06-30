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
  return (
    <ShellPage
      eyebrow="Dashboard"
      title="Главная"
      description="Каркас главного экрана. Полноценные виджеты, состояние уроков и прогресс будут добавлены в отдельном UI commit."
    >
      <div className="grid gap-4 md:grid-cols-3">
        {shellMetrics.map((metric) => (
          <SummaryCard key={metric.label} label={metric.label} value={metric.value} />
        ))}
      </div>
      <section className="grid gap-4 lg:grid-cols-[2fr_1fr]">
        <div className="border border-white/10 p-5">
          <p className="text-sm text-white/45">Текущий урок</p>
          <h2 className="mt-3 text-2xl font-semibold">{shellSummary.currentLesson.title}</h2>
          <p className="mt-2 text-sm text-white/55">{shellSummary.currentLesson.step}</p>
          <ProgressBar value={shellSummary.currentLesson.progressPercent} />
          <NavLink
            to={`/lesson/${shellSummary.currentLesson.id}`}
            className="mt-5 inline-flex border border-white bg-white px-4 py-2 text-sm font-medium text-black"
          >
            Продолжить обучение
          </NavLink>
        </div>
        <div className="border border-white/10 p-5">
          <p className="text-sm text-white/45">Слабые темы</p>
          <div className="mt-4 flex flex-wrap gap-2">
            {shellSummary.weakTopics.map((topic) => (
              <span key={topic} className="border border-white/15 px-3 py-2 text-sm text-white/75">
                {topic}
              </span>
            ))}
          </div>
        </div>
      </section>
    </ShellPage>
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

function ProgressBar({ value }: { value: number }) {
  return (
    <div className="mt-4">
      <div className="h-2 border border-white/10 bg-white/10">
        <div className="h-full bg-white" style={{ width: `${value}%` }} />
      </div>
      <p className="mt-2 text-sm text-white/50">{value}%</p>
    </div>
  )
}

export default App
