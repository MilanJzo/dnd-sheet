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
    <input type="checkbox" bind:checked={skill.classSkill} />
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
        width: 100%;
        height: fit-content;

        padding: 5px;

        background: red;
    }
</style>
