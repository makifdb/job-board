import React from 'react';
import Link from 'next/link'
import { InstantSearch, Highlight, Configure, SearchBox } from 'react-instantsearch-hooks-web';
import { instantMeiliSearch } from '@meilisearch/instant-meilisearch';
import { InfiniteHits } from './infinite-hits';

const searchClient = instantMeiliSearch(
    'http://localhost:7700',
    'masterKey',
);

export default function Search() {
    return (
        <InstantSearch
            indexName="jobs"
            searchClient={searchClient}
        >
            <SearchBox
                classNames={{
                    root: 'min-w-full items-center rounded-md w-full justify-center',
                    input: 'text-2xl w-full dark:bg-black outline-gray-900 border-2 rounded-md p-2'
                }}
                submitIconComponent={({ }) => (null)}
                resetIconComponent={({ }) => (null)}
            />
            <InfiniteHits hitComponent={Hit} />
            <Configure hitsPerPage={8} />

        </InstantSearch>
    )
}

const Hit = ({ hit }) => (
    <div className="flex flex-row mt-4 p-4 w-full min-w-full rounded-lg items-center space-x-4 bg-gray-700 hover:bg-gray-800">
        <Link href={`/jobs/${hit.id}`} key={hit.id} className="flex flex-row w-full items-center space-x-4">
            <img src={hit.company_logo} alt={hit.company_name} width={100} height={100} className='rounded-full bg-white' />
            <div className="flex flex-col items-start justify-center">
                <h1 className="text-xl text-white">
                    <Highlight attribute="title" hit={hit} />
                </h1>
                <h2 className="text-lg text-gray-300">
                    <Highlight attribute="company_name" hit={hit} />
                </h2>
                <span className="text-gray-400">
                    {hit.source}
                </span>
            </div>
        </Link>
    </div>
);
