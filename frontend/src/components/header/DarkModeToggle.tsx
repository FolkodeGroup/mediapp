import { useEffect, useState } from 'react'

const DarkModeToggle = () => {
  const [dark, setDark] = useState(() => {
    // Verificar localStorage primero, luego sistema
    const saved = localStorage.getItem('theme')
    if (saved) {
      return saved === 'dark'
    }
    // Si no hay preferencia guardada, usar la del sistema
    return window.matchMedia('(prefers-color-scheme: dark)').matches
  })

  useEffect(() => {
    const root = document.documentElement
    if (dark) {
      root.classList.add('dark')
      localStorage.setItem('theme', 'dark')
    } else {
      root.classList.remove('dark')
      localStorage.setItem('theme', 'light')
    }
  }, [dark])

  return (
    <button
      onClick={() => setDark(!dark)}
      className="relative inline-flex items-center justify-center p-2 rounded-md text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white hover:bg-gray-100 dark:hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-blue-500 transition-colors duration-200"
      aria-label={dark ? 'Cambiar a modo claro' : 'Cambiar a modo oscuro'}
    >
      <span className="text-lg">
        {dark ? '🌞' : '🌙'}
      </span>
      <span className="ml-1 text-sm font-medium">
        {dark ? 'Claro' : 'Oscuro'}
      </span>
    </button>
  )
}

export default DarkModeToggle
