import Link from 'next/link'
import { useState, useEffect } from 'react'
import ThemeSwitch from './theme-switch'

export default function Header() {
    const [mounted, setMounted] = useState(false)

    useEffect(() => setMounted(true), [])

    return (
        <header className="flex flex-row items-center justify-between w-full py-8 mx-auto">
            <Link href="/">
                <h1 className="flex flex-row items-center text-3xl font-bold space-x-2">
                    Job Board
                </h1>
            </Link>

            <div className="flex flex-row items-center space-x-4">
                {/* <Link href="/jobs/create" className="bg-gray-200 hover:bg-gray-400 dark:bg-gray-600 dark:hover:bg-gray-800 p-1.5 rounded-md">
                    <p className="flex flex-row items-center text-xl space-x-2">
                        Post a Job
                    </p>
                </Link> */}
                <ThemeSwitch />
            </div>
        </header>
    )
}