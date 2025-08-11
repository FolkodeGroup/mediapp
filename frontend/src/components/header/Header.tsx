import DarkModeToggle from './DarkModeToggle'

const Header = () => {
  return (
    <header className="p-4 shadow-md bg-gray-100 dark:bg-gray-800 flex justify-between items-center">
      <h1 className="text-xl font-bold">Mi Aplicación</h1>
      <DarkModeToggle />
    </header>
  )
}

export default Header
