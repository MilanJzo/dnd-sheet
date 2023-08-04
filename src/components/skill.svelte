<script>
    let skill = {
        id: 0,
        classSkill: false,
        hasToBeClassSkill: false,
        isAffectedByArmorPenalty: false,

        name: `SkillName`,
        total: 0,
        abilityModifier: `STR`,
        ranks: 0,
        miscellaneous: 0
    };

    let deactivate = !skill.classSkill && skill.hasToBeClassSkill;
    let halfedEffect = !skill.classSkill && !skill.hasToBeClassSkill;
    let modValue = 0;
    let currentCalculatedValue =
        modValue + (halfedEffect ? skill.ranks / 2 : skill.ranks) + skill.miscellaneous;

    //TODO update logic
</script>

<main>
    <input type="checkbox" bind:value={skill.classSkill} />
    <div class="wrap">
        <label for="" class="skillLa">{currentCalculatedValue}</label>
        <p>{skill.name}</p>
    </div>
    <p>=</p>
    <div class="wrap">
        <label for="" class="skillLa">{modValue}</label>
        <select bind:value={skill.abilityModifier} class="select">
            <option value="STR">STR</option>
            <option value="DEX">DEX</option>
            <option value="CON">CON</option>
            <option value="INT">INT</option>
            <option value="WIS">WIS</option>
            <option value="CHA">CHA</option>
        </select>
    </div>
    <div class="wrap">
        <input type="number" class="input" bind:value={skill.ranks} disabled={deactivate} />
        <p>Ranks</p>
    </div>
    <div class="wrap">
        <input type="number" class="input" bind:value={skill.miscellaneous} />
        <p>Miscellaneous</p>
    </div>
</main>

<style lang="scss">
    main {
        width: fit-content;
        height: fit-content;

        margin: 0;
        padding: 5px;
        padding-inline: 10px;

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
        height: 15px;

        padding: 0;
        padding-inline: 5px;

        font-size: 0.7rem;

        border: none;

        background-color: var(--primary);
    }

    .skillLa {
        width: auto;
        height: 20px;

        padding: 0;
        padding-inline: 3px;

        border: 1px solid var(--text);
        border-radius: 5px;

        background-color: var(--primary);

        font-size: var(--text-size);
    }

    .input {
        width: 100%;
        height: 21px;

        padding: 0;
    }
</style>
