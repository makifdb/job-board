import { useState } from "react";
import Head from "next/head";
import { useRouter } from "next/router";

export default function JobCreate() {
    const router = useRouter();

    const form = {
        title: "",
        description: "",
        location: "",
        company_name: "",
        company_logo: "",
        tags: [],
    };

    const [formData, setFormData] = useState(form);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prevState) => ({
            ...prevState,
            [name]: value,
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const res = await fetch("http://localhost:3000/api/v1/jobs", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(formData),
        });
        const data = await res.json();
        router.push(`/jobs/${data.data.id}`);
    };

    return (
        <div className="flex flex-col items-center justify-center w-full py-2 mt-4">
            <Head>
                <title>Job Board</title>
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <main className="flex flex-col justify-center w-full flex-1 text-center">
                <form
                    className="flex flex-col items-center justify-center w-full space-y-6"
                    onSubmit={handleSubmit}
                >
                    <div className="flex flex-row justify-between items-center w-full px-3 py-4 border border-black dark:border-white rounded-md mb-4">
                        <h1 className="text-4xl text-left font-bold">Post a Job</h1>
                        <button
                            type="submit"
                            className="p-2 text-gray-200 dark:text-gray-900 bg-black dark:bg-gray-200 rounded-md hover:opacity-90 dark:hover:bg-gray-300"
                        >
                            CREATE
                        </button>

                    </div>
                    <div className="flex flex-col justify-center w-full">
                        <div className="flex flex-row justify-between items-center w-full border border-black dark:border-white rounded-md p-2">
                            <label htmlFor="title" className="text-2xl text-left font-semibold tracking-wider ">
                                Title
                            </label>
                            <div className="flex flex-row justify-between items-center space-x-3">
                                <p className="text-sm text-gray-700 dark:text-gray-200">
                                    Job title
                                </p>
                                <p className="text-sm text-gray-500 dark:text-gray-400">
                                    {formData.title.length} / 60
                                </p>
                            </div>
                        </div>
                        <input
                            type="text"
                            name="title"
                            id="title"
                            className="w-full p-2 mt-2 border dark:bg-black border-gray-300 rounded-md"
                            value={formData.title}
                            onChange={handleChange}
                        />
                    </div>

                    <div className="flex flex-col justify-center w-full mt-4">
                        <div className="flex flex-row justify-between items-center w-full border border-black dark:border-white rounded-md p-2">
                            <label htmlFor="title" className="text-2xl text-left font-semibold tracking-wider ">
                                Location
                            </label>
                            <div className="flex flex-row justify-between items-center space-x-3">
                                <p className="text-sm text-gray-700 dark:text-gray-200">
                                    Location of the job (e.g. Remote)
                                </p>
                                <p className="text-sm text-gray-500 dark:text-gray-400">
                                    {formData.location.length} / 60
                                </p>
                            </div>
                        </div>
                        <input
                            type="text"
                            name="location"
                            id="location"
                            className="w-full p-2 mt-2 border dark:bg-black border-gray-300 rounded-md"
                            value={formData.location}
                            onChange={handleChange}
                        />
                    </div>

                    <div className="flex flex-col justify-center w-full mt-4">
                        <div className="flex flex-row justify-between items-center w-full border border-black dark:border-white rounded-md p-2">
                            <label htmlFor="title" className="text-2xl text-left font-semibold tracking-wider ">
                                Company Name
                            </label>
                            <div className="flex flex-row justify-between items-center space-x-3">
                                <p className="text-sm text-gray-700 dark:text-gray-200">
                                    Name of the company hiring
                                </p>
                                <p className="text-sm text-gray-500 dark:text-gray-400">
                                    {formData.company_name.length} / 60
                                </p>
                            </div>
                        </div>
                        <input
                            type="text"
                            name="company_name"
                            id="company_name"
                            className="w-full p-2 mt-2 border dark:bg-black border-gray-300 rounded-md"
                            value={formData.company_name}
                            onChange={handleChange}
                        />
                    </div>

                    <div className="flex flex-col justify-center w-full mt-4">
                        <div className="flex flex-row justify-between items-center w-full border border-black dark:border-white rounded-md p-2">
                            <label htmlFor="title" className="text-2xl text-left font-semibold tracking-wider ">
                                Company Logo
                            </label>
                            <div className="flex flex-row justify-between items-center space-x-3">
                                <p className="text-sm text-gray-700 dark:text-gray-200">
                                    You can use a link to an image
                                </p>
                                <p className="text-sm text-gray-500 dark:text-gray-400">
                                    {formData.company_logo.length} / 60
                                </p>
                            </div>
                        </div>
                        <input
                            type="text"
                            name="company_logo"
                            id="company_logo"
                            className="w-full p-2 mt-2 border dark:bg-black border-gray-300 rounded-md"
                            value={formData.company_logo}
                            onChange={handleChange}
                        />
                    </div>

                    <div className="flex flex-col justify-center w-full mt-4">
                        <div className="flex flex-row justify-between items-center w-full border border-black dark:border-white rounded-md p-2">
                            <label htmlFor="title" className="text-2xl text-left font-semibold tracking-wider ">
                                Description
                            </label>
                            <div className="flex flex-row justify-between items-center space-x-3">
                                <p className="text-sm text-gray-700 dark:text-gray-200">
                                    Don't forget to include how to apply!
                                </p>
                                <p className="text-sm text-gray-500 dark:text-gray-400">
                                    {formData.description.length} / 5000
                                </p>
                            </div>
                        </div>
                        <textarea
                            name="description"
                            id="description"
                            className="w-full p-2 mt-2 border dark:bg-black border-gray-300 rounded-md"
                            rows={20}
                            value={formData.description}
                            onChange={handleChange}
                        />
                    </div>
                </form>
            </main>
        </div>
    );
}