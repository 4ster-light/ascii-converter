<script setup lang="ts">
import { ref } from 'vue';
import { Upload } from 'lucide-vue-next';

const fileName = ref('');
const preserveColors = ref(false);
const output = ref('');
const file = ref<File | null>(null);

const handleFileUpload = (event: Event) => {
    const uploadedFile = (event.target as HTMLInputElement).files![0];
    file.value = uploadedFile;
    fileName.value = uploadedFile ? uploadedFile.name : '';
};

const clearImage = () => {
    fileName.value = '';
    file.value = null;
    output.value = '';
};

const generateAscii = async () => {
    if (!file.value) {
        alert('Please select an image first');
        return;
    }

    const formData = new FormData();
    formData.append('image', file.value);
    formData.append('preserve_color', preserveColors.value ? 'on' : 'off');

    try {
        const response = await fetch('/convert-to-image', {
            method: 'POST',
            body: formData,
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        output.value = await response.text();
    } catch (error) {
        for (let [key, value] of formData.entries()) {
            console.log(key, value);
        }
        console.error('Error:', error);
        alert('An error occurred while generating ASCII art');
    }
};
</script>

<template>
    <div class="mb-6">
        <label class="block font-bold mb-2 text-orange-500" for="file-upload">
            Choose file
        </label>
        <div class="flex items-center">
            <label :for="'file-upload'"
                class="cursor-pointer font-bold py-2 px-4 rounded inline-flex items-center bg-orange-500 dark:bg-orange-500 text-white dark:text-black hover:bg-orange-400 dark:hover:bg-orange-600">
                <Upload class="mr-2" />
                <span>Select Image</span>
            </label>
            <input id="file-upload" type="file" accept="image/*" class="hidden" @change="handleFileUpload" />
            <span class="ml-3 font-bold">{{ fileName || 'No file chosen' }}</span>
        </div>
    </div>

    <div class="flex items-center mb-6">
        <input type="checkbox" id="preserve-colors" v-model="preserveColors" class="mr-2" />
        <label for="preserve-colors" class="text-orange-500 font-bold">Preserve original colors</label>
    </div>

    <div class="flex space-x-4 mb-8">
        <button @click="clearImage"
            class="font-bold py-2 px-4 rounded bg-orange-300 dark:bg-gray-800 text-gray-800 dark:text-orange-400 hover:bg-orange-400 dark:hover:bg-gray-700">
            Clear Image
        </button>
        <button @click="generateAscii"
            class="font-bold py-2 px-4 rounded bg-orange-500 dark:bg-orange-500 text-white dark:text-black hover:bg-orange-400 dark:hover:bg-orange-600">
            Generate ASCII
        </button>
    </div>

    <div class="rounded bg-white dark:bg-black border border-orange-300 dark:border-orange-700 p-4">
        <pre class="whitespace-pre-wrap font-mono text-sm" v-html="output"></pre>
    </div>
</template>
