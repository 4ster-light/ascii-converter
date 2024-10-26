<script lang="ts">
    import FileUploader from "./components/FileUploader.svelte";
</script>

<main class="max-w-2xl mx-auto">
    <!-- Feature notices -->
    <div class="bg-orange-100 dark:bg-gray-800 rounded-lg p-4 mb-6">
        <h2
            class="text-lg font-semibold text-orange-700 dark:text-orange-400 mb-2"
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

    <div class="bg-orange-100 dark:bg-gray-800 rounded-lg p-4 mb-6">
        <h2
            class="text-lg font-semibold text-orange-700 dark:text-orange-400 mb-2"
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
    <div class="bg-white dark:bg-gray-800 rounded-lg p-4 mb-6 shadow-lg">
        <p class="text-gray-700 dark:text-gray-300 mb-4">
            This is a website to convert any image to ASCII art. It allows you
            to upload an image and convert it to an ASCII text format in a
            heartbeat.
        </p>
        <p class="text-gray-700 dark:text-gray-300 mb-6">
            At the moment, it supports common image formats like PNG, JPEG, JPG,
            and WEBP.
        </p>
    </div>

    <div class="bg-white dark:bg-gray-800 rounded-lg p-6 shadow-lg">
        <FileUploader />
    </div>
</main>
