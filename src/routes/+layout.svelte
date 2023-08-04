<script>
    import "./page.scss";
    import { theme, primary, secondary, text } from "../lib/stores";

    // $:{
    //     if(typeof window !== "undefined"){
    //         document.body.style.setProperty("--primary", $primary);
    //         document.body.style.setProperty("--secondary", $secondary);
    //         document.body.style.setProperty("--text", $text);
    //     }
    // }
</script>

<svelte:head>
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="true" />
    <link href="https://fonts.googleapis.com/css2?family=Rubik&display=swap" rel="stylesheet" />

    {#if $theme == "DARK"}
        <style>
            :root {
                --invert: 1;
                --secondary: darkgrey;
                --primary: #2c2922;
                --text: #ffffff;
            }
        </style>
        <!-- {:else if $theme == "CUSTOM"}
        <style>
            :root {
                --secondary: {$secondary};
                --primary: {$primary};
                --text: {$text};
            }
        </style> -->
    {:else}
        <style>
            :root {
                --invert: 0;
                --secondary: #6c6c6c;
                --primary: #fff3cf;
                --text: #000000;
            }
        </style>
    {/if}
</svelte:head>

<div class="navbar">
    <img src="/Dragon.svg" alt="Logo" width="50" />
    <div class="title">
        <h1>DnD Sheet</h1>

        <div class="options">
            <a href="/">Home</a>
            <a href="/characters">Characters</a>
            <a href="/settings">Settings</a>
        </div>
    </div>
</div>

<div class="content">
    <slot />
</div>

<style lang="scss">
    .content {
        width: auto;
        min-height: 400px;

        padding: 20px;

        display: flex;
        flex-direction: column;
        // justify-content: center;
        // align-items: center;
    }

    h1 {
        min-width: 200px;
    }

    img {
        pointer-events: none;
        -webkit-filter: invert(var(--invert)); /* safari 6.0 - 9.0 */
        filter: invert(var(--invert));
    }

    .navbar {
        padding: 0;
        margin: 0;

        width: 100%;
        height: fit-content;

        display: flex;
        flex-direction: row;
        gap: 20px;

        border-bottom: 2px solid var(--text);
    }

    .title {
        padding: 0;
        margin: 0;

        width: 100%;
        height: auto;

        display: flex;
        flex-direction: row;
        align-items: last baseline;
        gap: 20px;

        border: none;
    }

    .options {
        padding: 0;
        margin: 0;

        width: 100%;
        height: 100%;

        display: flex;
        flex-direction: row;
        align-items: last baseline;
        gap: 20px;

        border: none;
    }

    a {
        font-size: 1.25rem;
    }

    @media (max-width: 800px) {
        .title {
            padding: 0;
            flex-direction: column;
            justify-content: last baseline;
            align-items: flex-start;
            gap: 0px;
        }

        .options {
            padding-left: 3px;
        }
    }
</style>
