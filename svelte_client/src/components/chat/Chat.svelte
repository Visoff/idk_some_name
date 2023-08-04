<script lang="ts">
    export let name:string = "Saved messages"
    export let last_message:string = "hello"
</script>

<div class="h-[4.5rem] w-full flex flex-row p-2 gap-2 hover:bg-gray-200">
    <div class="bg-pink-400 aspect-square rounded-full"></div>
    <div class="flex flex-1 flex-col gap-1">
        <div class="flex-1 leading-5 flex overflow-hidden whitespace-nowrap text-ellipsis items-end font-bold">{name}</div>
        <div class="flex-1 leading-5 flex overflow-hidden whitespace-nowrap text-ellipsis items-start">{last_message}</div>
    </div>
</div>