<script setup lang="ts">
import Callout from '../components/Callout.vue';
import FileForm from '../components/FileForm.vue';
</script>

<template>
    <Callout title="EXPERIMENTAL FEATURE: Preserve Colors">
        <p class="mb-2">You can preserve the colour of the original image but:</p>
        <ul class="list-disc list-inside">
            <li>It may stress some computers hardware.</li>
            <li>It will take longer to generate the ASCII art.</li>
            <li>When downloading as image it will still be black and white.</li>
        </ul>
    </Callout>

    <p class="mb-4">
        This is a website to convert images to ASCII art. It allows you to upload an image and
        convert it to an ASCII text format. You can download as a PNG image too!!!
    </p>
    <p class="mb-6">
        At the moment, it supports common image formats like PNG, JPEG, JPG, and WEBP.
    </p>

    <FileForm />
</template>
