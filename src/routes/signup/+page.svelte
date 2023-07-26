<script lang="ts">
    import { enhance } from "$app/forms";
    import type { ActionData } from "./$types";

    export let form: ActionData;
</script>

<h1>Sign up</h1>
<form method="post" use:enhance>
	<input name="username" id="username" placeholder="Username" /><br />
	<input type="password" name="password" id="password" placeholder="Password" /><br />
    <button type="submit">Sign up</button>
</form>
<a href="/login">Sign in</a>

{#if form?.message}
    <p>{form.message}</p>
{/if}
