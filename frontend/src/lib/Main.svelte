<script lang="ts">
    import FileUploader from "./components/FileUploader.svelte";

    type CalloutProps = {
        title: string;
        head: string;
        content: string[];
    };
</script>

<main
    class="rounded-lg shadow-lg p-8 bg-gradient-to-br from-orange-200 to-orange-300 dark:from-gray-900 dark:to-gray-950"
>
    {@render Callout({
        title: "SOME FEATURES ARE MISSING",
        head: "But they'll be coming back soon:",
        content: [
            "Downloading images",
        ],
    })}

    {@render Callout({
        title: "THERE'S MORE COMING",
        head: "I'm working on more features, like:",
        content: [
            "A progress bar.",
            "Better performance.",
            "Custom characters.",
        ],
    })}

    <p class="text-gray-700 dark:text-gray-300 mb-4">
        This is a website to convert any image to ASCII art. It allows you to upload an image and convert it to an ASCII text format in a heartbeat.
    </p>
    <p class="text-gray-700 dark:text-gray-300 mb-6">
        At the moment, it supports common image formats like PNG, JPEG, JPG, and WEBP.
    </p>

    <FileUploader />
</main>

{#snippet Callout({ title, head, content }: CalloutProps)}
    <div class="rounded p-4 mb-6 bg-orange-100 dark:bg-black border border-orange-400 dark:border-orange-700">
        <h2 class="text-lg font-semibold text-orange-500 mb-2">{title}</h2>
        <p class="mb-2">{head}</p>
        <ul class="list-disc list-inside">
            {#each content as item}
                <li>{item}</li>
            {/each}
        </ul>
    </div>
{/snippet}
