<script lang="ts">
    import FileUploader from "./components/FileUploader.svelte";
</script>

<main
    class="rounded-lg shadow-lg p-8 bg-gradient-to-br from-orange-200 to-orange-300 dark:from-gray-900 dark:to-gray-950"
>
    <!-- Feature notices -->
        <div class="rounded p-4 mb-6 bg-orange-100 dark:bg-black border border-orange-400 dark:border-orange-700">
            <h2
                class="text-lg font-semibold text-orange-500 dark:text-orange-400 mb-2"
            >
                SOME FEATURES ARE MISSING
            </h2>
            <p class="text-gray-700 dark:text-gray-300 mb-2">
                But they'll be coming back soon:
            </p>
            <ul class="list-disc list-inside text-gray-700 dark:text-gray-300">
                <li>Downloading images.</li>
            </ul>
        </div>

        <div class="rounded p-4 mb-6 bg-orange-100 dark:bg-black border border-orange-400 dark:border-orange-700">
            <h2
                class="text-lg font-semibold text-orange-500 dark:text-orange-400 mb-2"
            >
                THERE'S MORE COMING:
            </h2>
            <p class="text-gray-700 dark:text-gray-300">
                I'm working on more features, like:
            </p>
            <ul class="list-disc list-inside text-gray-700 dark:text-gray-300">
                <li>A progress bar.</li>
                <li>Better performance.</li>
                <li>Custom characters.</li>
            </ul>
        </div>
        <!-- Main converter -->
        <p class="text-gray-700 dark:text-gray-300 mb-4">
            This is a website to convert any image to ASCII art. It allows you
            to upload an image and convert it to an ASCII text format in a
            heartbeat.
        </p>
        <p class="text-gray-700 dark:text-gray-300 mb-6">
            At the moment, it supports common image formats like PNG, JPEG, JPG,
            and WEBP.
        </p>

        <FileUploader />
</main>
