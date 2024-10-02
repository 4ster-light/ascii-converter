<template>
    <footer class="mt-8 text-center">
        <p class="mb-4">Made by David Vivar Bogónez, aka: Aster</p>
        <div class="flex justify-center space-x-4">
            <a href="https://x.com/4ster_light" class="text-orange-500 hover:text-orange-600">
                <Twitter />
            </a>
            <a href="#" class="text-orange-500 hover:text-orange-600">
                <Github />
            </a>
            <a href="#" class="text-orange-500 hover:text-orange-600">
                <Globe />
            </a>
        </div>
    </footer>
</template>

<script setup lang="ts">
import { Twitter, Github, Globe } from 'lucide-vue-next';
</script>
