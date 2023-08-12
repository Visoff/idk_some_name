<script lang="ts">
	import { browser } from "$app/environment";
	import { host } from "$lib/env";
	import Chat from "./Chat.svelte";

    let chats:any[] = []
    if (browser) {
        fetch(host+"/user/chats", {
            method:"GET",
            headers:{
                Authorization:`Bearer ${localStorage.getItem("user_token")}`
            }
        }).then(response => response.json())
        .then(c => {
            console.log(c)
            if (!c || c.length == 0) {
                chats = []
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
        fetch(host+"/chat/", {
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

    if (browser) {
        let longpoll = new Date().toISOString()
        setInterval(() => {
            fetch(host+"/user/chats", {
                method:"GET",
                headers:{
                    Authorization:`Bearer ${localStorage.getItem("user_token")}`,
                    LongPoll:longpoll
                }
            }).then(response => response.json())
            .then(c => {
                if (!c || c.length == 0) {
                    return
                }
                longpoll = c[0].last_update
                chats = [...chats, ...c]
            }).catch(console.error)
        }, 500)
    }
</script>

<div class="flex-1 flex flex-col gap-2">
    {#each chats as chat}
        <Chat id={chat.id} name={chat.name} description={chat.lastMessage||chat.description} />
    {/each}
    <button
        class="mx-4 py-1 rounded-2xl border-blue-500 border hover:bg-blue-200"
        on:click={AddChat}
    >
        +
    </button>
</div>