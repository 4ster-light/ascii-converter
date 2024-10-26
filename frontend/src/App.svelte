<script lang="ts">
    import Footer from "./lib/Footer.svelte";
    import Header from "./lib/Header.svelte";
    import Main from "./lib/Main.svelte";
</script>

<div
    class="min-h-screen transition-colors duration-200 bg-[#FFDAB9] dark:bg-gray-900 text-gray-900 dark:text-gray-100"
>
    <div class="container mx-auto px-4 py-8">
        <Header />
        <Main />
        <Footer />
    </div>
</div>
