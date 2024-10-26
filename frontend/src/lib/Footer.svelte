<script lang="ts">
    import { Github, Globe, Twitter } from "lucide-svelte";
</script>

<footer class="mt-8 text-center">
    <p class="mb-4">Made by David Vivar Bogónez, aka: Aster</p>
    <div class="flex justify-center space-x-4">
        <a href="https://x.com/4ster_light" class="text-orange-500 hover:text-orange-600">
            <span class="sr-only">Twitter</span>
            <Twitter />
        </a>
        <a href="https://github.com/4ster-light" class="text-orange-500 hover:text-orange-600">
            <span class="sr-only">GitHub</span>
            <Github />
        </a>
        <a href="/" class="text-orange-500 hover:text-orange-600">
            <span class="sr-only">Website</span>
            <Globe />
        </a>
    </div>
</footer>
