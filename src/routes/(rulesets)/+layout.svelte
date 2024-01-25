<script>
    import "../../app.css";
    import { page } from "$app/stores";

    let active_footer_item = $page.url.href.split("/").at(3);
</script>

<svelte:head>
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="true" />
    <link href="https://fonts.googleapis.com/css2?family=Rubik&display=swap" rel="stylesheet" />
</svelte:head>

<div class="page">
    <div class="content">
        <slot name="page" />
    </div>

    <div class="sidebar">
        <slot name="sidebar" />
    </div>

    <div class="footer">
        <a
            class={active_footer_item === "" ? "active_footer_item" : "footer_item"}
            href="/"
            on:click={() => {
                active_footer_item = "";
            }}>Home</a
        >
        <a
            class={active_footer_item === "dnd3_5" ? "active_footer_item" : "footer_item"}
            href="/dnd3_5"
            on:click={() => {
                active_footer_item = "dnd3_5";
            }}>DnD 3.5</a
        >
        <a
            class={active_footer_item === "dnd5e" ? "active_footer_item" : "footer_item"}
            href="/dnd5e"
            on:click={() => {
                active_footer_item = "dnd5e";
            }}>DnD 5e</a
        >
        <a
            class={active_footer_item === "pathfinder1" ? "active_footer_item" : "footer_item"}
            href="/pathfinder1"
            on:click={() => {
                active_footer_item = "pathfinder1";
            }}>Pathfinder 1</a
        >
    </div>
</div>

<style lang="scss">
    .page {
        width: 100%;
        height: 100vh;

        padding: 10px;

        display: grid;
        grid-template-columns: 8fr 2fr;
        grid-template-rows: 1fr 60px;
        gap: 10px;

        background-color: var(--accent-bg);
        color: var(--main-text);
    }

    .content,
    .sidebar,
    .footer {
        width: 100%;
        height: 100%;

        border-radius: 10px;

        background-color: var(--main-bg);
    }

    .footer {
        grid-column: span 2;

        padding: 10px;

        display: flex;
        align-items: center;
        gap: 10px;

        overflow-x: scroll;
        overflow-y: none;
    }

    .footer_item,
    .active_footer_item {
        width: fit-content;
        height: fit-content;

        padding: 10px;

        border-radius: 10px;

        background-color: var(--main-bg);
        color: var(--main-text);

        font-weight: bold;
        font-size: calc(var(--font-size) * 0.9);

        text-decoration: none;

        transition: background-color 0.25s ease-in-out;

        &:hover,
        &:focus {
            background-color: var(--accent-bg);
        }
    }

    .active_footer_item {
        background-color: var(--accent-bg);
    }
</style>
