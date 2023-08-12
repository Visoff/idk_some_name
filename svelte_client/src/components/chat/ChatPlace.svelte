<script lang="ts">
	import MessageList from "./MessageList.svelte";
    import { CurrentChat } from "$lib";
	import MessageInput from "./MessageInput.svelte";
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
        ">{$CurrentChat}</div>
        <div class="
            aspect-square h-full
            bg-blue-200
        "></div>
    </div>
    <div class="flex-1 overflow-y-auto w-full lg:w-96 bg-blue-100">
        <MessageList />
    </div>
    <div class="h-[3.375rem] flex p-3 pt-0 w-full lg:w-96 justify-center">
        <div class="flex-1 flex w-full gap-2">
            <MessageInput />
        </div>
    </div>
</div>