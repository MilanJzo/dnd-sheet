<script lang="ts">
    import { enhance } from "$app/forms";
    import { redirect } from "@sveltejs/kit";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h1>Profile</h1>

<p>User id: {data.userId}</p>
<p>Username: {data.username}</p>
<form method="post" action="?/logout" use:enhance>
    <button type="submit">Sign out</button>
</form>
