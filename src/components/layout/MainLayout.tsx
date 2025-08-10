import Header from '../header/Header'

type Props = {
  children: React.ReactNode
}

const MainLayout = ({ children }: Props) => {
  return (
    <div className="min-h-screen bg-white dark:bg-gray-900 text-gray-900 dark:text-white transition-colors">
      <Header />
      <main className="p-4">{children}</main>
    </div>
  )
}

export default MainLayout
