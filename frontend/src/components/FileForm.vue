<script setup lang="ts">
import { ref } from 'vue';
import { Upload } from 'lucide-vue-next';

const fileName = ref('');
const preserveColors = ref(false);

const handleFileUpload = (event: Event) => {
    const file = (event.target as HTMLInputElement).files![0];
    fileName.value = file ? file.name : '';
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
            <input id="file-upload" type="file" class="hidden" @change="handleFileUpload" />
            <span class="ml-3 font-bold">{{ fileName || 'No file chosen' }}</span>
        </div>
    </div>

    <div class="flex items-center mb-6">
        <input type="checkbox" id="preserve-colors" v-model="preserveColors" class="mr-2" />
        <label for="preserve-colors" class="text-orange-500 font-bold">Preserve original colors</label>
    </div>

    <div class="flex space-x-4 mb-8">
        <button
            class="font-bold py-2 px-4 rounded bg-orange-300 dark:bg-gray-800 text-gray-800 dark:text-orange-400 hover:bg-orange-400 dark:hover:bg-gray-700">
            Clear Image
        </button>
        <button
            class="font-bold py-2 px-4 rounded bg-orange-500 dark:bg-orange-500 text-white dark:text-black hover:bg-orange-400 dark:hover:bg-orange-600">
            Generate ASCII
        </button>
    </div>
</template>
