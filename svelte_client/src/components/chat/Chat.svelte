<script lang="ts">
    import { CurrentChat } from "$lib";
    export let name:string = "Saved messages"
    export let description:string = "hello"
    export let id:string;
</script>

<div class="relative h-[4.5rem] w-full flex flex-row p-2 gap-2 hover:bg-gray-200">
    <button class="opacity-0 absolute top-0 left-0 w-full h-full" on:click={() => {$CurrentChat = id}}></button>
    <div class="bg-pink-400 aspect-square rounded-full"></div>
    <div class="flex flex-1 flex-col gap-1 overflow-x-hidden whitespace-nowrap">
        <div class="flex-1 leading-5 flex overflow-hidden items-end font-bold">{name}</div>
        <div class="flex-1 leading-5 flex overflow-hidden items-start">{description}</div>
    </div>
</div>