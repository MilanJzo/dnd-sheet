<script>
    $: skill = {
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

    $: deactivate = !skill.classSkill && skill.hasToBeClassSkill;
    //$: halfedEffect = !skill.classSkill && !skill.hasToBeClassSkill;
    $: modValue = 0;
    $: currentCalculatedValue =
        modValue +
        (!skill.classSkill && !skill.hasToBeClassSkill ? skill.ranks / 2 : skill.ranks) +
        skill.miscellaneous;

    //TODO update logic
</script>

<main>
    <div class={skill.classSkill ? "checked" : "nocheck"} />
    <input class="check" type="checkbox" bind:checked={skill.classSkill} />
    <div class="wrap">
        <p>{skill.name}:</p>
        <label for="" class="label">{currentCalculatedValue}</label>
    </div>
    <p>=</p>
    <div class="wrap">
        <div class="modSelect">
            <select bind:value={skill.abilityModifier} class="select">
                <option value="STR">STR</option>
                <option value="DEX">DEX</option>
                <option value="CON">CON</option>
                <option value="INT">INT</option>
                <option value="WIS">WIS</option>
                <option value="CHA">CHA</option>
            </select>
            <p>:</p>
        </div>
        <label for="" class="label">{modValue}</label>
    </div>
    <div class="wrap">
        <p>Ranks:</p>
        <input type="number" class="input" bind:value={skill.ranks} disabled={deactivate} />
    </div>
    <div class="wrap">
        <p>Miscellaneous:</p>
        <input type="number" class="input" bind:value={skill.miscellaneous} />
    </div>
    <button class="delbtn">delete</button>
</main>

<style lang="scss">
    :root {
        --hover: white;
    }

    main {
        width: fit-content;
        min-width: fit-content;
        height: fit-content;

        padding: 5px;

        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 15px;

        border: 2px solid black;
        border-radius: 10px;

        position: relative;
    }

    p {
        user-select: none;
    }

    .checked,
    .nocheck {
        width: 20px;
        height: 20px;

        border: 2px solid black;
        border-radius: 5px;

        position: relative;
    }

    .checked::before {
        content: "";
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);

        width: 16px;
        height: 16px;

        background-color: black;
        border-radius: 5px;
    }

    .check {
        position: absolute;
        top: 5px;
        left: 5px;

        width: 24px;
        height: 24px;

        border: 2px solid black;

        cursor: pointer;

        opacity: 0;

        &:hover .checked,
        &:hover .nocheck {
            background: var(--hover);
        }
    }

    .wrap {
        height: 19px;

        display: flex;
        flex-direction: row;
        gap: 5px;

        border-bottom: 1px solid black;
    }

    .modSelect {
        display: flex;
        flex-direction: row;
        align-items: center;
    }

    .label {
        width: auto;
        height: 18px;

        border: none;

        display: flex;
        // justify-content: center;
        align-items: center;

        user-select: none;
    }

    .input {
        width: 50px;
        height: auto;

        border: none;
    }

    .select {
        width: auto;
        height: 19px;

        padding: 0;
        margin: 0;

        border: none;

        cursor: pointer;

        &:hover {
            background: var(--hover);
        }
    }

    .delbtn {
        width: fit-content;
        height: 24px;

        border: 2px solid black;
        border-radius: 5px;

        cursor: pointer;

        &:hover {
            background: var(--hover);
        }
    }
</style>
