<script lang="ts">
    export let message:any
</script>

<div class="
    relative flex
    w-fit max-w-[50vw] min-w-[4rem] px-3 pt-2 pb-6 rounded-2xl {message.self ? "justify-end self-end rounded-br-none" : "rounded-bl-none"}
    bg-blue-500 text-white
">
    {message.content}
    <div class="
        absolute {message.self ? "right-1" : "right-1"} bottom-1
        text-sm
    ">
        {new Date(message.created_at).toLocaleTimeString()}
    </div>
</div>