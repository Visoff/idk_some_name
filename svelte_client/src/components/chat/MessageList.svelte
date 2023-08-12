<script lang="ts">
    import { CurrentChat } from '$lib'
	import { browser } from "$app/environment";
	import Message from "./Message.svelte";

    function onChatChange(chat_id:string|false) {
        if (!browser) {
            return [
                {
                    id:"",
                    content:"hiiii"
                },
                {
                    id:"",
                    content:"second",
                    self:true
                }
            ]
        }
        fetch("http://localhost:8080/api/chat/messages/", {
            method:"GET",
            headers:{
                Authorization:`Bearer ${localStorage.getItem("user_token")}`,
                Chat:chat_id||""
            }
        }).then(response => response.json())
        .then((m?:any[]) => {
            if (m == undefined) {
                messages = []
                return
            }
            messages = m
            console.log(m)
        }).catch(console.error)
    }
    let messages:any[] = []
    CurrentChat.subscribe(onChatChange)
    if (browser) {
        let last_time = new Date().toISOString()
        setInterval(() => {
            fetch("http://localhost:8080/api/chat/messages/", {
                method:"GET",
                headers:{
                    Authorization:`Bearer ${localStorage.getItem("user_token")}`,
                    Chat:$CurrentChat||"",
                    LongPoll:last_time
                }
            }).then(response => response.json())
            .then((m?:any[]) => {
                if (m == undefined || m.length == 0) {
                    return
                }
                messages = [...messages, ...m]
                last_time = m.at(-1).created_at
                console.log(m)
            }).catch(console.error)
        }, 500)
    }
</script>

<div class=" flex flex-col gap-2 px-2 pt-2">
    {#each messages as message}
        <Message message={message}/>
    {/each}
</div>