<script setup lang="ts">
import Callout from '../components/Callout.vue';
import FileForm from '../components/FileForm.vue';
</script>

<template>
    <p class="mb-4">
        This is a simple website to convert images to ASCII art. It allows you to upload an image and
        convert it to an ASCII text format.
    </p>
    <p class="mb-4">
        At the moment, it supports common image formats like PNG, JPEG, JPG, and WEBP.
    </p>
    <p class="mb-6">
        The output is black and white text at the moment. Now you can download as a PNG image too!!!
    </p>

    <Callout title="EXPERIMENTAL FEATURE: Preserve Colors">
        <p class="mb-2">You can preserve the colour of the original image but:</p>
        <ul class="list-disc list-inside">
            <li>It may stress some computers hardware.</li>
            <li>It will take longer to generate the ASCII art.</li>
            <li>When downloading as image it will still be black and white.</li>
        </ul>
    </Callout>

    <FileForm />

    <div class="h-32 rounded bg-white dark:bg-black border border-orange-300 dark:border-orange-700"></div>
</template>
