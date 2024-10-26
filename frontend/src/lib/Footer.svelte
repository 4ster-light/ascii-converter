<script lang="ts">
    import { Github, Globe, Twitter } from "lucide-svelte";
</script>

<footer class="mt-8 text-center text-gray-600 dark:text-gray-400">
    <p>Made by David Vivar Bogónez, aka: Aster</p>
    <div class="flex justify-center gap-4 mt-2">
        <a href="https://x.com/4ster_light" class="hover:text-orange-500">
            <span class="sr-only">Twitter</span>
            <Twitter class="w-5 h-5" />
        </a>
        <a href="https://github.com/4ster-light" class="hover:text-orange-500">
            <span class="sr-only">GitHub</span>
            <Github class="w-5 h-5" />
        </a>
        <a href="/" class="hover:text-orange-500">
            <span class="sr-only">Website</span>
            <Globe class="w-5 h-5" />
        </a>
    </div>
</footer>
