<script lang="ts">
    import { enhance } from "$app/forms";
    import type { PageData } from "./$types";
    import { theme, secondary, primary, text } from "../../lib/stores";

    export let data: PageData;
</script>

<div class="entry">
    <h1>Profile</h1>
    <div class="bullet-point">
        <h2>Username</h2>
        <div>
            <p>{data.username}</p>
        </div>
    </div>
    <div class="bullet-point">
        <h2>Account ID</h2>
        <div>
            <p>{data.userId}</p>
        </div>
    </div>
    <form method="post" action="?/logout" use:enhance>
        <button type="submit">Sign out</button>
    </form>
</div>
<div class="entry">
    <h1>Theme</h1>
    <div class="bullet-point">
        <h2>Preset</h2>
        <div>
            <select bind:value={$theme}>
                <option value="LIGHT">Light</option>
                <option value="DARK">Dark</option>
                <!-- <option value="CUSTOM">Custom</option> -->
            </select>
        </div>
    </div>
     <!-- {#if $theme == "CUSTOM"}
        <div class="bullet-point">
            <h2>Background</h2>
            <div>
                <input type="text" bind:value={$primary} class="colorIN" />
            </div>
        </div>
        <div class="bullet-point">
            <h2>Accent</h2>
            <div>
                <input type="text" bind:value={$secondary} class="colorIN" />
            </div>
        </div>
        <div class="bullet-point">
            <h2>Text</h2>
            <div>
                <input type="text" bind:value={$text} class="colorIN" />
            </div>
        </div>
    {/if}  -->
</div>

<style lang="scss">
    .colorIN {
        width: fit-content;
        height: 100%;

        border: none;
        border-bottom: 1px solid var(--text);
    }

    .entry {
        // padding-top: 10px;
        // padding-inline: 10px;

        padding-bottom: 10px;

        display: flex;
        flex-direction: column;
        gap: 10px;

        border-bottom: 1px solid var(--text);
    }

    .bullet-point {
        width: 300px;

        display: flex;
        flex-direction: row;
        justify-content: space-between;

        & div {
            width: 150px;
        }
    }

    form {
        padding: 0;
        padding-top: 5px;
    }

    h1 {
        font-size: 1.5rem;
    }
</style>
