<script>
    import "./page.scss";
</script>

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
        width: fit-content;
        height: fit-content;

        padding: 20px;
    }

    h1 {
        min-width: 200px;
    }

    img {
        pointer-events: none;
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
