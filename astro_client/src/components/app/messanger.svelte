<script lang="ts">
    let open:boolean = false
</script>

<div class="w-full h-full relative flex flex-row overflow-hidden bg-gray-100">
    <nav class="flex-1 bg-gray-300">
        <button on:click={() => {open = true}}>open</button>
    </nav>
    <main class="
        sm:flex-[6]
        max-sm:absolute max-sm:right-0 {open ? "" : "max-sm:translate-x-full"}
        max-sm:w-screen max-sm:h-screen
        max-sm:transition-transform
        bg-gray-400
    ">

    </main>
</div>