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
    <div class="options">
        <a class="link" href="/dnd3_5"> DnD 3.5 </a>
        <a class="link" href="/dnd5e"> DnD 5e </a>
        <a class="link" href="/pathfinder1"> Pathfinder 1 </a>
    </div>
</div>

<style lang="scss">
    .page {
        min-height: 100vh;

        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 50px;

        background-color: var(--accent-bg);
        color: var(--main-text);
    }

    .options {
        max-width: 800px;
        height: auto;

        padding: 20px;

        display: flex;
        align-items: center;
        justify-content: center;
        flex-wrap: wrap;
        gap: 10px;

        border-radius: 10px;

        background-color: var(--main-bg);
    }

    .header {
        width: 100%;

        padding: 20px 50px;

        display: flex;
        justify-content: space-between;
        align-items: center;

        background-color: var(--main-bg);
    }

    .link,
    .settings {
        padding-inline: 30px;
        padding-block: 20px;

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

    .settings {
        padding-inline: 20px;
        padding-block: 10px;
    }
</style>
