<script>
	import Message from "./Message.svelte";
    import { CurrentChat } from "$lib";
</script>
<div class="flex-1 flex flex-col items-center gap-1">
    <div class="
        h-[3.5rem] w-full flex p-2 gap-2
    ">
        <button class="
            aspect-square h-full
            bg-blue-200
            md:hidden
        " on:click={() => {$CurrentChat = false}}></button>
        <div class="
            aspect-square rounded-full h-full
            bg-blue-200
        "></div>
        <div class="
            flex-1 h-full
            bg-blue-200
        "></div>
        <div class="
            aspect-square h-full
            bg-blue-200
        "></div>
    </div>
    <div class="flex-1 w-full lg:w-96 bg-blue-100">
        <Message />
    </div>
    <div class="h-[3.375rem] flex p-3 pt-0 w-full lg:w-96 justify-center">
        <div class="flex-1 flex w-full gap-2">
            <div class="flex-1 flex flex-row rounded-lg bg-gray-200">
                <div class="m-1 aspect-square rounded-full bg-red-500"></div>
                <form on:submit={e => e.preventDefault()} class="flex flex-1">
                    <input type="text" placeholder="Message" class="flex-1 bg-transparent focus:outline-none">
                </form>
                <div class="m-1 aspect-square rounded-full bg-red-500"></div>
            </div>
            <div class="h-full aspect-square rounded-full bg-blue-400"></div>
        </div>
    </div>
</div>