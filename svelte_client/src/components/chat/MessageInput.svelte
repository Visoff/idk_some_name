<script lang="ts">
    import {CurrentChat} from '$lib'
	import { host } from '$lib/env';
    let typing:Boolean = false
    function CheckTyping(e:Event&{currentTarget: HTMLInputElement}) {
        typing = (e.target as HTMLInputElement).value != ""
    }

    function sendMessage(e:SubmitEvent) {
        e.preventDefault()
        const content = (e.target as any)[0].value
        if (content == "") return;
        fetch(host+"/message", {
            method:"POST",
            headers:{
                Authorization:`Bearer ${localStorage.getItem("user_token")}`,
                Chat:($CurrentChat||"")
            },
            body:content
        }).then(response => response.json())
        .then(inserted => {
            console.log(inserted)
        }).catch(console.error);
        (e.target as HTMLFormElement).reset()
    }
</script>

<div class="flex-1 flex flex-row rounded-lg bg-gray-200">
    <div class="m-1 aspect-square rounded-full bg-red-500"></div>
    <form on:submit={sendMessage} class="flex flex-1">
        <input
            type="text" placeholder="Message"
            class="flex-1 bg-transparent focus:outline-none"
            on:input={CheckTyping}
        >
    </form>
    <div class="m-1 aspect-square rounded-full bg-red-500">
    </div>
</div>
<div class="h-full aspect-square rounded-full bg-purple-600 p-[.4rem] flex">
        <img
            src={typing ? "/icons/social/send.svg" : "/icons/social/microphone.svg"}
            alt={typing ? "record" : "send"}
            class="flex-1 duration-200 transition-all" />
</div>