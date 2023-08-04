<script>
    import '../app.css'
	import ChatList from '../components/chat/ChatList.svelte';
	import ChatPlace from '../components/chat/ChatPlace.svelte';
</script>

<div class="absolute top-0 left-0 w-screen h-screen flex overflow-hidden">
    <div class="
        flex-[.25] bg-blue-100
        max-md:absolute max-md:left-0 max-md:top-0 max-md:-translate-x-full
    ">
        <!-- apps -->
    </div>
    <div class="
        flex-1 flex flex-col
        bg-[whitesmoke]
        max-md:absolute max-md:left-0 max-md:top-0 max-md:-translate-x-full">
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
</div>