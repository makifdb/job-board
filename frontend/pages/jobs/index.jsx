import Search from '../../components/Search'

export default function Home() {
    return (
        <div className="mx-auto my-6">
            <main className="flex flex-col justify-between">
                <h1 className="text-2xl md:text-4xl font-bold text-left mb-3">Search 🔎</h1>
                <Search />
            </main>
        </div>
    )
}
