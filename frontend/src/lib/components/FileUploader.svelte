<script lang="ts">
    import Button from "./Button.svelte";
    import { createEventDispatcher } from "svelte";
    import { Upload } from "lucide-svelte";

    let preserveColors = false;
    let selectedFile: File | null = null;
    let asciiResult = "";

    const dispatch = createEventDispatcher<{ fileSelect: File }>();
    let fileInput: HTMLInputElement;

    function handleFileChange(event: Event) {
        const target = event.target as HTMLInputElement;
        if (target.files && target.files[0]) {
            dispatch("fileSelect", target.files[0]);
        }
    }

    async function handleFileSelect(event: CustomEvent<File>) {
        selectedFile = event.detail;
    }

    function clearImage() {
        selectedFile = null;
        asciiResult = "";
    }

    async function generateAscii() {
        if (!selectedFile) return;

        const formData = new FormData();
        formData.append("file", selectedFile);
        formData.append("preserveColors", preserveColors.toString());

        try {
            const response = await fetch("/api/convert", {
                method: "POST",
                body: formData,
            });

            if (!response.ok) throw new Error("Conversion failed");

            const result = await response.json();
            asciiResult = result.result;
        } catch (error) {
            console.error("Error:", error);
        }
    }
</script>

<div class="relative">
    <input
        type="file"
        bind:this={fileInput}
        on:change={handleFileChange}
        accept="image/*"
        class="hidden"
    />
    <button
        class="w-full px-4 py-2 border-2 border-dashed border-orange-300 dark:border-gray-600 rounded-lg
               hover:border-orange-400 dark:hover:border-gray-500 transition-colors flex items-center justify-center gap-2"
        on:click={() => fileInput.click()}
    >
        <Upload class="w-5 h-5" />
        <span>Select Image</span>
    </button>

    <div class="flex items-center gap-2 my-4">
        <input
            type="checkbox"
            id="preserveColors"
            bind:checked={preserveColors}
            class="w-4 h-4 text-orange-600"
        />
        <label
            for="preserveColors"
            class="text-gray-700 dark:text-gray-300"
        >
            Preserve original colors
        </label>
    </div>

    <div class="flex gap-4 mt-4">
        <Button variant="secondary" on:click={clearImage}
            >Clear Image</Button
        >
        <Button variant="primary" on:click={generateAscii}
            >Generate ASCII</Button
        >
    </div>

    {#if asciiResult}
        <pre
            class="mt-6 p-4 bg-gray-100 dark:bg-gray-700 rounded overflow-x-auto">
            {asciiResult}
        </pre>
    {/if}
</div>
