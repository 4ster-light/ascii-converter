<script setup lang="ts">
import Callout from '../components/Callout.vue';
import FileForm from '../components/FileForm.vue';
</script>

<template>
    <Callout title="SOME FEATURES ARE MISSING">
        <p class="mb-2">But they'll be coming back soon:</p>
        <ul class="list-disc list-inside">
            <li>Downloading images.</li>
            <li>Preserving the colours.</li>
        </ul>
    </Callout>

    <p class="mb-4">
        This is a website to convert any image to ASCII art. It allows you to upload an image and
        convert it to an ASCII text format in a heartbeat.
    </p>
    <p class="mb-6">
        At the moment, it supports common image formats like PNG, JPEG, JPG, and WEBP.
    </p>

    <FileForm />
</template>
