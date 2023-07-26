<script>
    import "./page.scss"
</script>

<div>
    <img src="/Dragon.svg" alt="Logo" width="50">
    <h1>DnD Sheet</h1>
</div>

<style>
    div {
        padding: 0;
        margin: 0;

        width: 100%;
        height: 5vh;

        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 10px;

        /* background-color: red; */

        border-bottom: 2px solid black;
    }
</style>

<slot />