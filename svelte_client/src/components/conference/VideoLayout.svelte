<script>

	import { CurrentChat } from "$lib";
	import { call } from "$lib/rtc";
	import ChatList from "../chat/ChatList.svelte";
	import VideoPlace from "./VideoPlace.svelte";

</script>

<div class={`
    flex-1 flex flex-col
    bg-[whitesmoke]
    max-md:absolute max-md:left-0 max-md:top-0 max-md:h-full max-md:w-full
    ${$CurrentChat ? "max-md:-translate-x-full" : ""}
    transition-all ease-linear
`}>
    <!-- video_app/planned -->
    <button>call</button>
</div>
<div class=" flex-[3] flex">
    <!-- video_app/dialog -->
    <VideoPlace />
</div>