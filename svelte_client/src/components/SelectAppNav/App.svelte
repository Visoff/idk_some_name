<script lang="ts">
    export let src:string = "/"
    export let img:string = "#"
    export let popover:string = ""
</script>

<div class="
    group relative
    aspect-square w-full p-2 flex
">
    <a href={src} class="
        flex-1 border-blue-700 border rounded-[50%] transition-all ease-linear duration-75 hover:rounded-2xl
    ">
        <img src={img} alt="">
    </a>
    <div class="
        hidden group-hover:block
        absolute top-1/2 -translate-y-1/2 -right-2 px-2 py-1 translate-x-full
        bg-blue-100 rounded-r-full
        z-10
    ">
        {popover}
    </div>
</div>