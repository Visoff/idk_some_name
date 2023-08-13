<script>
	import { browser } from "$app/environment";
	import SelectAppNav from "./SelectAppNav.svelte";
    if (browser) {
    }
</script>
<div class="absolute top-0 left-0 w-screen h-screen flex overflow-hidden">
    <SelectAppNav />
    <slot />
</div>