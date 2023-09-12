<script lang="ts">
    export let url:string = "/app"
    export let icon:string = "#"
    export let popover:string = "Some text"
</script>

<div class="
    group
    h-pcalc(100%-1rem)] w-fit
    sm:h-fit sm:w-full
    aspect-square rounded-full
    border border-gray-400
    flex flex-col
    relative
    max-sm:mb-4
">
    <a href={url} class="h-full w-full">
        <img src={icon} alt="" class="h-full w-full border-none">
    </a>
    <div class="
        absolute
        w-fit h-fit

        sm:-right-2 sm:top-1/2 sm:translate-x-full sm:-translate-y-1/2
        max-sm:-bottom-5 max-sm:left-1/2 -translate-x-1/2

        transition-transform origin-left
        sm:scale-x-0 ease-out
        group-hover:scale-x-100 group-hover:ease-in
        ">
        <span class="
            px-2 py-3 sm:rounded-r-full
            text-sm whitespace-nowrap
            sm:bg-gray-400
        ">
            {popover}
        </span>
    </div>
</div>