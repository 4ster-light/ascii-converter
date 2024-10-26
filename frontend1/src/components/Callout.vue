<script setup lang="ts">
defineProps<{
    title: string;
}>();
</script>

<template>
    <div class="rounded p-4 mb-6 bg-orange-100 dark:bg-black border border-orange-400 dark:border-orange-700">
        <h2 class="text-lg font-semibold text-orange-500 mb-2">{{ title }}</h2>
        <slot />
    </div>
</template>
