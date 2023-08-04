<script lang="ts">
	import Chat from "./Chat.svelte";

    export let chats: {name:string, last_message:string}[] = []
</script>

{#each chats as chat}
    <Chat {...chat} />
{/each}