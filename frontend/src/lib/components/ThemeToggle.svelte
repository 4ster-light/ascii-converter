<script lang="ts">
    import { Moon, Sun } from "lucide-svelte";
    import { onMount } from "svelte";

    export let darkMode = false;

    onMount(() => {
        darkMode = window.matchMedia("(prefers-color-scheme: dark)").matches;
        document.documentElement.classList.toggle("dark", darkMode);
    });

    function toggleTheme() {
        darkMode = !darkMode;
        document.documentElement.classList.toggle("dark", darkMode);
    }
</script>

<button
    class="p-2 rounded-lg hover:bg-orange-200 dark:hover:bg-gray-700 transition-colors"
    on:click={toggleTheme}
    aria-label="Toggle theme"
>
    {#if darkMode}
        <Sun class="w-6 h-6" />
    {:else}
        <Moon class="w-6 h-6" />
    {/if}
</button>
