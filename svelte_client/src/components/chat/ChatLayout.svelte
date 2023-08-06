<script>

	import { CurrentChat } from "$lib";
	import ChatList from "./ChatList.svelte";
	import ChatPlace from "./ChatPlace.svelte";

</script>

<div class={`
    flex-1 flex flex-col
    bg-[whitesmoke]
    max-md:absolute max-md:left-0 max-md:top-0 max-md:h-full max-md:w-full
    ${$CurrentChat ? "max-md:-translate-x-full" : ""}
    transition-all ease-linear
`}>
    <!-- chat_app/chats -->
    <!-- Search bar -->
    <div class="w-full h-[3.5rem] grid place-items-center">
        Search
    </div>
    <!-- Chats list -->
    <ChatList chats={[
        {
            name:"Saved messages",
            last_message:"hello"
        }
    ]}/>
</div>
<div class=" flex-[3] flex">
    <!-- chat_app/dialog -->
    <ChatPlace />
</div>