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
                <h1 className="text-4xl text-left font-bold">Create a Job</h1>

                <form
                    className="flex flex-col items-center justify-center w-full mt-12"
                    onSubmit={handleSubmit}
                >
                    <div className="flex flex-col items-center justify-center w-full">
                        <label htmlFor="title" className="text-2xl">
                            Title
                        </label>
                        <input
                            type="text"
                            name="title"
                            id="title"
                            className="w-full p-2 mt-2 border dark:bg-black border-gray-300 rounded-md"
                            value={formData.title}
                            onChange={handleChange}
                        />
                    </div>

                    <div className="flex flex-col items-center justify-center w-full mt-4">
                        <label htmlFor="location" className="text-2xl">
                            Location
                        </label>
                        <input
                            type="text"
                            name="location"
                            id="location"
                            className="w-full p-2 mt-2 border dark:bg-black border-gray-300 rounded-md"
                            value={formData.location}
                            onChange={handleChange}
                        />
                    </div>

                    <div className="flex flex-col items-center justify-center w-full mt-4">
                        <label htmlFor="company_name" className="text-2xl">
                            Company Name
                        </label>
                        <input
                            type="text"
                            name="company_name"
                            id="company_name"
                            className="w-full p-2 mt-2 border dark:bg-black border-gray-300 rounded-md"
                            value={formData.company_name}
                            onChange={handleChange}
                        />
                    </div>

                    <div className="flex flex-col items-center justify-center w-full mt-4">
                        <label htmlFor="company_logo" className="text-2xl">
                            Company Logo
                        </label>
                        <input
                            type="text"
                            name="company_logo"
                            id="company_logo"
                            className="w-full p-2 mt-2 border dark:bg-black border-gray-300 rounded-md"
                            value={formData.company_logo}
                            onChange={handleChange}
                        />
                    </div>

                    <div className="flex flex-col items-center justify-center w-full mt-4">
                        <label htmlFor="description" className="text-2xl">
                            Description
                        </label>
                        <textarea
                            name="description"
                            id="description"
                            className="w-full p-2 mt-2 border dark:bg-black border-gray-300 rounded-md"
                            rows={8}
                            value={formData.description}
                            onChange={handleChange}
                        />
                    </div>

                    <button
                        type="submit"
                        className="w-full p-2 mt-4 text-white bg-gray-500 rounded-md"
                    >
                        Create
                    </button>
                </form>
            </main>
        </div>
    );
}