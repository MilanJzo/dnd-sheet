<script lang="ts">
    import { enhance } from "$app/forms";
    import { redirect } from "@sveltejs/kit";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h1>Profile</h1>

<div class="profile">
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

<style lang="scss">
    .profile {
        padding-top: 10px;
        padding-inline: 10px;

        display: flex;
        flex-direction: column;
        gap: 10px;
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

    form{
        padding: 0;
        padding-top: 5px;
    }

    h1 {
        font-size: 1.5rem;
    }
</style>
