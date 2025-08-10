import { useEffect, useState } from 'react'

const DarkModeToggle = () => {
  const [dark, setDark] = useState(() => {
    return localStorage.getItem('theme') === 'dark'
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
      className="px-2 py-1 border rounded text-sm"
    >
      {dark ? '🌞 Claro' : '🌙 Oscuro'}
    </button>
  )
}

export default DarkModeToggle
