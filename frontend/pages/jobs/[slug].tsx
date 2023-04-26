import { useRouter } from 'next/router'
import { useState, useEffect } from 'react'

export default function Job() {
    const router = useRouter()
    const { slug } = router.query

    let [job, setJob] = useState({})
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        const getJob = async () => {
            if (!slug) return
            const res = await fetch(`http://localhost:3000/api/v1/jobs/${slug}`)
            const data = await res.json()
            setJob(data)
            setLoading(false)
        }
        getJob()
    }, [slug])

    if (loading) {
        return (
            <div className="flex flex-col items-center justify-center w-full py-2 mt-40">
                <svg className="animate-spin h-32 w-32 mr-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                    <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                    <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z"></path>
                </svg>
            </div>
        )
    }

    return (
        <div className="flex flex-col items-center justify-center w-full py-2">
            <main className="flex flex-col justify-center w-full flex-1 px-20">
                <h1 className="text-4xl text-left font-bold my-6">
                    {job.data.title}
                </h1>
                <div className="flex flex-row items-center space-x-4 bg-gray-300 dark:bg-gray-600 rounded-md p-4">
                    <img className="rounded-full w-32"
                        src={job.data.company_logo}
                        alt={job.data.company_name}
                    />
                    <div>
                        <h2 className="text-4xl font-bold">
                            {job.data.company_name}
                        </h2>
                        <h3 className="text-2xl font-bold">
                            {job.data.location}
                        </h3>

                    </div>

                </div>

                <p className="text-xl mt-10">
                    {job.data.description}
                </p>
            </main>
        </div>
    )
}