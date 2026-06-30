function App() {
  return (
    <main className="min-h-screen bg-black text-white">
      <section className="mx-auto flex min-h-screen w-full max-w-5xl flex-col justify-between px-6 py-8 sm:px-10">
        <header className="flex items-center justify-between border-b border-white/15 pb-6">
          <div className="flex items-center gap-3">
            <div className="flex size-10 items-center justify-center border border-white/25 bg-white text-sm font-semibold text-black">
              GP
            </div>
            <span className="text-lg font-semibold tracking-normal">GoPath</span>
          </div>
          <span className="text-sm text-white/60">frontend skeleton</span>
        </header>

        <div className="grid gap-8 py-16">
          <p className="text-sm uppercase tracking-[0.24em] text-white/50">
            Go backend learning platform
          </p>
          <div className="max-w-3xl space-y-5">
            <h1 className="text-4xl font-semibold leading-tight tracking-normal sm:text-6xl">
              Скелет frontend готов
            </h1>
            <p className="max-w-2xl text-base leading-7 text-white/70 sm:text-lg">
              Минимальное приложение на React, TypeScript, Vite и Tailwind CSS.
              Продуктовые экраны, маршруты и mock data будут добавлены отдельными
              commit.
            </p>
          </div>
        </div>

        <footer className="grid gap-3 border-t border-white/15 pt-6 text-sm text-white/60 sm:grid-cols-3">
          <span>React + TypeScript</span>
          <span>Vite build</span>
          <span>Tailwind CSS</span>
        </footer>
      </section>
    </main>
  )
}

export default App
