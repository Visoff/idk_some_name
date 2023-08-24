<script lang="ts">
    export let url:string = "/app"
    export let icon:string = "#"
    export let popover:string = "Some text"
</script>

<div class="group w-full h-fit flex flex-col relative">
    <a href={url}>
        <img src={icon} alt="" class="w-full aspect-square rounded-full border border-gray-400">
    </a>
    <div class="
        absolute -right-2 top-1/2 translate-x-full -translate-y-1/2
        w-fit h-fit
        transition-transform origin-left
        scale-x-0 ease-out
        group-hover:scale-x-100 group-hover:ease-in
        ">
        <span class="
            px-2 py-3 rounded-r-full
            text-sm whitespace-nowrap
            bg-gray-400
        ">
            {popover}
        </span>
    </div>
</div>