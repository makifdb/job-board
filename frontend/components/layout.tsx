import Navbar from '../components/navbar'

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex flex-col min-h-screen max-w-4xl mx-auto py-3 px-10">
        <Navbar />
        <main>{children}</main>
    </div>
  )
}