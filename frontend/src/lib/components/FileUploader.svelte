<script lang="ts">
    import { Upload } from "lucide-svelte";

    let { fileName, preserveColors, output, file } = $state({
        fileName: "",
        preserveColors: false,
        output: "",
        file: null as File | null
    });

    function decodeUnicode(input: string) {
        return input
            .replace(/\\u003c/g, "<")
            .replace(/\\u003e/g, ">")
            .replace(/\\u0026/g, "&")
            .replace(/\\"/g, '"');
    }

    function handleFileUpload(event: Event) {
        const uploadedFile = (event.target as HTMLInputElement).files?.[0];
        if (uploadedFile) {
            file = uploadedFile;
            fileName = uploadedFile.name;
        }
    }

    function clearImage() {
        fileName = "";
        file = null;
        output = "";
    }

    async function generateAscii() {
        if (!file) {
            alert("Please select an image first");
            return;
        }

        const formData = new FormData();
        formData.append("image", file);
        formData.append("preserve_color", preserveColors ? "on" : "off");

        try {
            const response = await fetch("/convert-to-ascii", {
                method: "POST",
                body: formData,
            });

            if (!response.ok) {
                throw new Error("Network response was not ok");
            }

            output = decodeUnicode(await response.text());
        } catch (error) {
            for (let [key, value] of formData.entries()) {
                console.log(key, value);
            }
            console.error("Error:", error);
            alert("An error occurred while generating ASCII art");
        }
    }
</script>

<div class="mb-6">
    <label class="block font-bold mb-2 text-orange-500" for="file-upload">
        Choose file
    </label>
    <div class="flex items-center">
        <label
            for="file-upload"
            class="cursor-pointer font-bold py-2 px-4 rounded inline-flex items-center bg-orange-500 dark:bg-orange-500 text-white dark:text-black hover:bg-orange-400 dark:hover:bg-orange-600"
        >
            <Upload class="mr-2" />
            <span>Select Image</span>
        </label>
        <input
            id="file-upload"
            type="file"
            accept="image/*"
            class="hidden"
            onchange={handleFileUpload}
        />
        <span class="ml-3 font-bold">{fileName || "No file chosen"}</span>
    </div>
</div>

<div class="flex items-center mb-6">
    <input
        type="checkbox"
        id="preserve-colors"
        bind:checked={preserveColors}
        class="mr-2"
    />
    <label for="preserve-colors" class="text-orange-500 font-bold">
        Preserve original colors
    </label>
</div>

<div class="flex space-x-4 mb-8">
    <button
        onclick={clearImage}
        class="font-bold py-2 px-4 rounded bg-orange-300 dark:bg-gray-800 text-gray-800 dark:text-orange-400 hover:bg-orange-400 dark:hover:bg-gray-700"
    >
        Clear Image
    </button>
    <button
        onclick={generateAscii}
        class="font-bold py-2 px-4 rounded bg-orange-500 dark:bg-orange-500 text-white dark:text-black hover:bg-orange-400 dark:hover:bg-orange-600"
    >
        Generate ASCII
    </button>
</div>

<div
    class="rounded bg-white dark:bg-black border border-orange-300 dark:border-orange-700 p-4 flex items-center justify-center overflow-scroll sm:overflow-auto"
>
    <pre
        class="w-auto whitespace-pre-wrap text-[1px] md:text-[2px] bg-black text-white font-bold tracking-widest">
        {@html output}
    </pre>
</div>
