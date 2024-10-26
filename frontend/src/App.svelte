<script lang="ts">
    import Footer from "./lib/Footer.svelte";
    import Header from "./lib/Header.svelte";
    import Main from "./lib/Main.svelte";
</script>

<div
    class="min-h-screen p-8 bg-orange-100 dark:bg-black text-gray-800 dark:text-orange-200"
>
    <div class="max-w-4xl mx-auto">
        <Header />
        <Main />
        <Footer />
    </div>
</div>
