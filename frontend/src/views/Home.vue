<template>
    <div class="min-h-screen p-8 bg-orange-50 dark:bg-black text-gray-800 dark:text-orange-200">
        <div class="max-w-4xl mx-auto">
            <Header />

            <main
                class="rounded-lg shadow-lg p-8 bg-gradient-to-br from-orange-100 to-orange-200 dark:from-gray-900 dark:to-black">
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

                <div
                    class="rounded p-4 mb-6 bg-orange-100 dark:bg-black border border-orange-300 dark:border-orange-700">
                    <h2 class="text-lg font-semibold text-orange-500 mb-2">EXPERIMENTAL FEATURE:</h2>
                    <p class="mb-2">You can preserve the colour of the original image but:</p>
                    <ul class="list-disc list-inside">
                        <li>It may stress some computers hardware.</li>
                        <li>It will take longer to generate the ASCII art.</li>
                        <li>When downloading as image it will still be black and white.</li>
                    </ul>
                </div>

                <div class="mb-6">
                    <label class="block text-sm font-medium mb-2 text-orange-500" for="file-upload">
                        Choose file
                    </label>
                    <div class="flex items-center">
                        <label :for="'file-upload'"
                            class="cursor-pointer font-bold py-2 px-4 rounded inline-flex items-center bg-orange-400 dark:bg-orange-500 text-white dark:text-black hover:bg-orange-500 dark:hover:bg-orange-600">
                            <Upload class="mr-2" />
                            <span>Select Image</span>
                        </label>
                        <input id="file-upload" type="file" class="hidden" @change="handleFileUpload" />
                        <span class="ml-3 text-orange-300">{{ fileName || 'No file chosen' }}</span>
                    </div>
                </div>

                <div class="flex items-center mb-6">
                    <input type="checkbox" id="preserve-colors" v-model="preserveColors" class="mr-2" />
                    <label for="preserve-colors" class="text-orange-500">Preserve original colors</label>
                </div>

                <div class="flex space-x-4 mb-8">
                    <button
                        class="font-bold py-2 px-4 rounded bg-orange-200 dark:bg-gray-800 text-gray-800 dark:text-orange-400 hover:bg-orange-300 dark:hover:bg-gray-700">
                        Clear Image
                    </button>
                    <button
                        class="font-bold py-2 px-4 rounded bg-orange-400 dark:bg-orange-500 text-white dark:text-black hover:bg-orange-500 dark:hover:bg-orange-600">
                        Generate ASCII
                    </button>
                </div>

                <div class="h-32 rounded bg-white dark:bg-black border border-orange-300 dark:border-orange-700"></div>
            </main>

            <Footer />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Upload } from 'lucide-vue-next';
import Header from '../components/Header.vue';
import Footer from '../components/Footer.vue'

const fileName = ref('');
const preserveColors = ref(false);

const handleFileUpload = (event: Event) => {
    const file = (event.target as HTMLInputElement).files![0];
    fileName.value = file ? file.name : '';
};
</script>
