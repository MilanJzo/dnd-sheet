<script lang="ts">
    import { enhance } from "$app/forms";
    import type { ActionData } from "./$types";

    export let form: ActionData;
</script>

<h1>Sign in</h1>
<form method="post" use:enhance>
    <input name="username" id="username" placeholder="Username" /><br />
    <input type="password" name="password" id="password" placeholder="Password" /><br />
    <div>
        <a href="/signup">Create an account</a>
        <button type="submit">Sign in</button>
    </div>
</form>

{#if form?.message}
    <p>{form.message}</p>
{/if}

<style>
    div {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
    }
</style>
