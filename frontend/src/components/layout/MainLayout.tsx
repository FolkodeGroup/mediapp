import Header from '../header/Header'

type Props = {
  children: React.ReactNode
}

const MainLayout = ({ children }: Props) => {
  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-white transition-all duration-300 main-layout">
      <Header />
      <main>{children}</main>
    </div>
  )
}

export default MainLayout
