<script lang="ts">
    let typing:Boolean = false
    function CheckTyping(e:Event&{currentTarget: HTMLInputElement}) {
        typing = (e.target as HTMLInputElement).value != ""
    }
</script>

<div class="flex-1 flex flex-row rounded-lg bg-gray-200">
    <div class="m-1 aspect-square rounded-full bg-red-500"></div>
    <form on:submit={e => e.preventDefault()} class="flex flex-1">
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