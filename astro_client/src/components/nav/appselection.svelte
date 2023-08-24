<script>
  import Appicon from './appicon.svelte';
</script>

<nav
    class="
        w-16 h-full p-2
        flex flex-col gap-2
        bg-gray-200
    "
>
    <div class="w-full aspect-square rounded-full bg-gray-400"></div>
    <hr class="border-gray-400">
    <Appicon url="/app/messanger" popover="Messanger" />
    <Appicon />
    <Appicon />
    <Appicon />
    <Appicon />
</nav>