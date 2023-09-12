<script lang="ts">
    let chat_id:string = "123";
    const chats = [
        {
            id:"12345",
            title:"Title"
        },
        {
            id:"456",
            title:"Another"
        }
    ]
</script>

<div class="w-full h-full relative flex flex-row overflow-x-hidden bg-gray-100">
    <nav class="flex-1 flex flex-col gap-1 bg-gray-300">
        {#each chats as chat}
            <button class="" on:click={() => {chat_id = chat.id}}>{chat.title}</button>
        {/each}
    </nav>
    <main class="
        sm:flex-[6]
        max-sm:fixed max-sm:right-0 {chat_id != "" ? "" : "max-sm:translate-x-full"}
        max-sm:w-screen max-sm:h-screen
        max-sm:transition-transform
        bg-gray-400
        flex flex-col
        items-center
    ">
        <div class="w-full h-16 bg-gray-600">
            <button class="
                h-full aspect-square
                bg-gray-300
                sm:hidden
            " on:click={() => {chat_id = ""}}></button>
        </div>
        <div class="flex-1 w-full max-w-3xl xl:max-w-[43.5rem] xl:-translate-x-[2.25rem]">

        </div>
        <form class="w-full max-w-3xl h-16 flex gap-2 px-2 py-1" on:submit={(e) => {e.preventDefault()}}>
            <div class="h-full flex gap-1 flex-1 rounded-xl px-2 py-1 bg-blue-200">
                <button class="aspect-square rounded-full border border-black"></button>
                <input type="text" class="flex-1 bg-transparent outline-none">
                <button class="aspect-square rounded-full border border-black"></button>
            </div>
            <button class="aspect-square rounded-full bg-blue-400"></button>
        </form>
    </main>
</div>