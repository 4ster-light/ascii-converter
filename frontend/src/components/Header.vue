<template>
    <header class="flex justify-between items-center mb-12">
        <h1 class="text-5xl font-bold text-orange-500">
            ✰ArtSCII✰
        </h1>
        <button @click="toggleDarkMode"
            class="p-2 rounded-full bg-orange-400 dark:bg-gray-800 text-black dark:text-orange-400">
            <Sun v-if="isDarkMode" :size="24" />
            <Moon v-else :size="24" />
        </button>
    </header>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { Sun, Moon } from 'lucide-vue-next';

const isDarkMode = ref(false);

onMounted(() => {
    // Check system preference on mount
    isDarkMode.value = window.matchMedia('(prefers-color-scheme: dark)').matches;
    applyDarkMode();

    // Listen for changes in system preference
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
        isDarkMode.value = e.matches;
        applyDarkMode();
    });
});

const toggleDarkMode = () => {
    isDarkMode.value = !isDarkMode.value;
    applyDarkMode();
};

const applyDarkMode = () => {
    if (isDarkMode.value) {
        document.documentElement.classList.add('dark');
    } else {
        document.documentElement.classList.remove('dark');
    }
};
</script>
