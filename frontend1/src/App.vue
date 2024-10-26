<script setup lang="ts">
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';
</script>

<template>
    <div class="min-h-screen p-8 bg-orange-100 dark:bg-black text-gray-800 dark:text-orange-200">
        <div class="max-w-4xl mx-auto">
            <Header />
            <main
                class="rounded-lg shadow-lg p-8 bg-gradient-to-br from-orange-200 to-orange-300 dark:from-gray-900 dark:to-gray-950">
                <router-view></router-view>
            </main>
            <Footer />
        </div>
    </div>
</template>
