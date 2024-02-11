<script>
    import { page } from "$app/stores";

    let ruleset = $page.url.href.split("/").at(3);
</script>

<div class="page">
    <div class="content">
        <p class="ruleset" />
        <div class="paperwrap">
            <div class="paper">
                <slot />
            </div>
        </div>
    </div>

    <div class="sidebar" />
</div>

<style lang="scss">
    .page {
        width: 100%;
        height: 100%;

        padding-bottom: 10px;

        display: grid;
        grid-template-columns: 8fr minmax(200px, 2fr);
        gap: 10px;
    }

    .sidebar,
    .content {
        width: 100%;
        height: 100%;

        border-radius: 10px;

        background-color: var(--main-bg);

        box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.25);
    }
</style>
