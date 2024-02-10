<script>
    import "../app.css";
</script>

<div class="page">
    <div class="header">
        <h1>DnDSheet</h1>
        <a class="settings" href="/settings">
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="white"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                class="lucide lucide-cog"
            >
                <path d="M12 20a8 8 0 1 0 0-16 8 8 0 0 0 0 16Z" />
                <path d="M12 14a2 2 0 1 0 0-4 2 2 0 0 0 0 4Z" />
                <path d="M12 2v2" />
                <path d="M12 22v-2" />
                <path d="m17 20.66-1-1.73" />
                <path d="M11 10.27 7 3.34" />
                <path d="m20.66 17-1.73-1" />
                <path d="m3.34 7 1.73 1" />
                <path d="M14 12h8" />
                <path d="M2 12h2" />
                <path d="m20.66 7-1.73 1" />
                <path d="m3.34 17 1.73-1" />
                <path d="m17 3.34-1 1.73" />
                <path d="m11 13.73-4 6.93" />
            </svg>
        </a>
    </div>

    <div class="welcome">
        <p>
            Welcome to DnDSheet! This is a project that I am working on to help me learn Svelte and
            to help me keep track of my DnD characters. I hope you enjoy it!
        </p>
    </div>
</div>

<style lang="scss">
    .welcome {
        width: 600px;
        padding: 25px 75px;

        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 25px;

        border-radius: 10px;

        font-size: 1.1rem;

        background-color: var(--main-bg);
        box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.25);
    }

    .page {
        height: 100%;
        width: 100%;

        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 150px;

        background-color: var(--accent-bg);
        color: var(--main-text);
    }

    .header {
        width: 100%;

        padding: 10px 75px;

        display: flex;
        justify-content: space-between;
        align-items: center;

        border-radius: 10px;

        background-color: var(--main-bg);
        box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.25);
    }

    .header h1 {
        height: fit-content;

        font-size: 2.5rem;
        font-weight: bold;
    }

    .settings {
        padding-inline: 20px;
        padding-block: 10px;

        display: flex;
        align-items: center;

        border: none;
        border-radius: 10px;

        background-color: var(--main-bg);
        color: var(--main-text);

        font-weight: bold;
        font-size: var(--font-size);

        text-decoration: none;

        transition: background-color 0.25s ease-in-out;

        &:hover,
        &:focus {
            background-color: var(--accent-bg);
        }
    }
</style>
