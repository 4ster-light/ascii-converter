<script lang="ts">
    import { Moon, Sun } from "lucide-svelte";
    import { onMount } from "svelte";

    let darkMode = $state(false);

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
    class="p-2 rounded-full bg-orange-400 dark:bg-gray-800 text-black dark:text-orange-400"
    onclick={toggleTheme}
    aria-label="Toggle theme"
>
    {#if darkMode}
        <Sun class="w-6 h-6" />
    {:else}
        <Moon class="w-6 h-6" />
    {/if}
</button>
