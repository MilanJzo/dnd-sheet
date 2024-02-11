<script>
    import "../app.css";
    import { page } from "$app/stores";

    let active_footer_item = $page.url.href.split("/").at(3);
</script>

<svelte:head>
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="true" />
    <link href="https://fonts.googleapis.com/css2?family=Rubik&display=swap" rel="stylesheet" />
</svelte:head>

<div class="page">
    <div class="slot">
        <slot />
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
    @import url("https://fonts.googleapis.com/css2?family=Lato:ital,wght@0,100;0,300;0,400;0,700;0,900;1,100;1,300;1,400;1,700;1,900&family=Preahvihear&display=swap");

    .slot {
        width: 100%;
        height: 100%;
    }

    .page {
        font-family: "Lato", sans-serif;

        width: 100%;
        height: 100vh;

        padding: 10px;

        display: grid;
        grid-template-rows: 1fr 60px;

        background-color: var(--accent-bg);
        color: var(--main-text);
    }

    .footer {
        width: 100%;
        height: 100%;

        padding: 10px;

        display: flex;
        align-items: center;
        gap: 10px;

        border-radius: 10px;

        background-color: var(--main-bg);

        box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.25);
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
