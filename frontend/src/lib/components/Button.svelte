<script lang="ts">
    export let variant: "primary" | "secondary" = "primary";
</script>

<button
    class="px-4 py-2 rounded-lg font-medium transition-colors
           {variant === 'primary'
        ? 'bg-orange-500 hover:bg-orange-600 text-white'
        : 'bg-gray-200 dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-600'}"
    on:click
>
    <slot />
</button>
