<script>
    import "../app.css";
    import { page } from "$app/stores";
</script>

<svelte:head>
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="true" />
    <link href="https://fonts.googleapis.com/css2?family=Rubik&display=swap" rel="stylesheet" />
</svelte:head>

<div class="page">
    <div class="content">
        <slot />
    </div>

    <div class="sidebar">
        <slot name="sidebar" />
    </div>

    <div class="footer">
        <a class="footer_item" href="/">Home</a>
        <a class="footer_item" href="/dnd3_5">DnD 3.5</a>
        <a class="footer_item" href="/dnd5e">DnD 5e</a>
        <a class="footer_item" href="/pathfinder1">Pathfinder 1</a>
    </div>
</div>

<style lang="scss">
    .page {
        width: 100%;
        height: 100vh;

        padding: 10px;

        display: grid;
        grid-template-columns: 8fr 2fr;
        grid-template-rows: 1fr 40px;
        gap: 10px;

        background-color: var(--main-bg);
    }

    .content,
    .sidebar,
    .footer {
        width: 100%;
        height: 100%;
    }
</style>
