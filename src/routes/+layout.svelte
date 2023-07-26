<script>
    import "./page.scss"
</script>

<div>
    <img src="/Dragon.svg" alt="Logo" width="50">
    <div>
        <h1>DnD Sheet</h1>

        <div>
            <a href="/">Home</a>
            <a href="/characters">Characters</a>
            <a href="/settings">Settings</a>
        </div>
    </div>
</div>

<style lang="scss">


    div {
        padding: 0;
        margin: 0;

        width: 100%;
        height: 5vh;

        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 20px;

        border-bottom: 2px solid black;

        & div{
            padding: 0;
            margin: 0;

            width: fit-content;
            height: 100%;

            align-items: last baseline;

            border: none;
        }
    }

    a {
        margin: 0;
        padding: 0;

        text-decoration: none;
        color: black;

        font-size: 1.25rem;

        &:hover {
            color: darkblue;
        }
    }
</style>

<slot />