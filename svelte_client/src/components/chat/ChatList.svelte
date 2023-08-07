<script lang="ts">
	import { browser } from "$app/environment";
	import Chat from "./Chat.svelte";

    let chats:{id:string, name:string, description:string}[] = []
    if (browser) {
        fetch("http://localhost:8080/api/user/chats", {
            method:"GET",
            headers:{
                Authorization:`Bearer ${localStorage.getItem("user_token")}`
            }
        }).then(response => response.json())
        .then(c => {
            if (!c || c.length == 0) {
                chats = [
                    {
                        id:"1",
                        name:"hello",
                        description:"world"
                    }
                ]
                return
            }
            chats = c
        }).catch(console.error)
    }

    function AddChat() {
        const name = prompt("Chat name")
        if (name == undefined) return;
        const description = prompt("Chat description")
        if (description == undefined) return;
        fetch("http://localhost:8080/api/chat/", {
            method:"POST",
            headers:{
                Authorization:`Bearer ${localStorage.getItem("user_token")}`,
                "Content-Type":"application/json"
            },
            body:JSON.stringify({
                name, description
            })
        })
    }
</script>

<div class="flex-1 flex flex-col gap-2">
    {#each chats as chat}
        <Chat id={chat.id} name={chat.name} description={chat.description} />
    {/each}
    <button
        class="mx-4 py-1 rounded-2xl border-blue-500 border hover:bg-blue-200"
        on:click={AddChat}
    >
        +
    </button>
</div>