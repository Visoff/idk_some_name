<script>
  import Appicon from './appicon.svelte';

    let opened = false;
</script>

<nav
    class="
        sm:w-16 sm:h-full
        w-full h-20
        p-2
        flex flex-row sm:flex-col gap-2
        max-sm:justify-evenly
        bg-gray-200
        sm:static
        z-0
    "
>
    <div class="
        max-sm:h-full sm:w-full
        aspect-square rounded-full
        bg-gray-400
        mb-2
        max-sm:hidden
    "></div>
    <Appicon url="/app/messanger" popover="Messanger" />
    <Appicon url="/app/meets" popover="Sync meets" />
    <Appicon />
    <Appicon />
    <Appicon />
</nav>