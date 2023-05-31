import Image from 'next/image';
import Link from 'next/link';

export default function Home() {
  return (
    <div className="max-w-5xl mx-auto my-6">
      <main className="flex flex-col items-left justify-between">
        {/* center the secreen */}
        <div className="grid grid-cols-8 mt-16 lg:mt-32">
          <div className="flex flex-col items-left col-span-3">
            <h1 className="text-3xl md:text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r dark:from-gray-200 dark:to-gray-800 from-gray-800 to-gray-200 animate-gradient-x mb-3">
              Welcome to
            </h1>
            <h1 className="text-6xl md:text-8xl font-bold text-transparent bg-clip-text bg-gradient-to-r dark:from-gray-200 dark:to-gray-800 from-gray-800 to-gray-200 animate-gradient-x mb-3">
              Job Board
            </h1>
            <p className="animate-text mt-8 bg-gradient-to-r from-gray-900 via-gray-600 to-gray-400 dark:from-gray-200 dark:via-gray-400 dark:to-gray-600 bg-clip-text text-transparent text-2xl">
              Find your dream job or your dream employee
            </p>
          </div>
          <Image
            className='rounded-full col-span-5'
            src="/cover2.png"
            width={800}
            height={800}
            alt="Picture of the author"
          />
        </div>

        <div className="flex flex-col items-left justify-between mt-32 space-y-16 lg:space-y-32">

          <div>
            <Link href="/jobs" className='animate-text bg-gradient-to-r from-gray-900 via-gray-600 to-gray-400 dark:from-gray-200 dark:via-gray-400 dark:to-gray-600 bg-clip-text text-transparent text-4xl hover:animate-pulse'>
              Find a job and start your career →
            </Link>
          </div>

          <div>
            <Link href="/jobs/create" className='animate-text bg-gradient-to-r from-gray-900 via-gray-600 to-gray-400 dark:from-gray-200 dark:via-gray-400 dark:to-gray-600 bg-clip-text text-transparent text-4xl hover:animate-pulse'>
              Find your next employee →
            </Link>
          </div>

        </div>
      </main>
    </div>
  )
}
