<script>
    let modtype = `STR`;
    let modValue = 0;
    let value = 0;
</script>

<main>
    <input type="checkbox" />
    <div class="wrap">
        <label for="" class="skillLa">{value}</label>
        <p>AbilityName</p>
    </div>
    <p>=</p>
    <div class="wrap">
        <label for="" class="skillLa">{modValue}</label>
        <select bind:value={modtype} class="select">
            <option value="STR">STR</option>
            <option value="DEX">DEX</option>
            <option value="CON">CON</option>
            <option value="INT">INT</option>
            <option value="WIS">WIS</option>
            <option value="CHA">CHA</option>
        </select>
    </div>
</main>

<style lang="scss">
    main {
        width: auto;
        height: fit-content;

        padding: 5px;

        display: flex;
        gap: 10px;

        border: 1px solid var(--text);
        border-radius: 5px;

        & p {
            padding: 0;
            margin: 0;

            font-size: 0.8rem;
        }
    }

    .wrap {
        width: 100px;
        height: 80%;

        display: flex;
        flex-direction: column;
        justify-content: center;
        gap: 5px;
    }

    .select {
        width: auto;
        height: 80%;

        padding: 0;
        padding-inline: 5px;

        font-size: 0.7rem;

        border: none;

        background-color: var(--primary);
    }

    .skillLa {
        width: auto;
        height: fit-content;

        padding: 0;
        padding-inline: 5px;

        border: 1px solid var(--text);
        border-radius: 5px;

        background-color: var(--primary);
    }
</style>
