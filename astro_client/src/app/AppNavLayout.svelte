<script lang="ts">
    let navOpened:boolean = false
</script>

<main class="
    absolute top-0 left-0 w-screen h-screen
    flex
">
    <nav class="
        h-screen w-screen md:w-auto
        md:flex-1
        max-md:fixed top-0 left-0 transition-transform {navOpened ? "" : "max-md:-translate-x-full"}
        bg-blue-200
    ">
        <nav class="
            md:hidden
            w-full h-12
            p-2
        ">
            <button class="
                h-full aspect-square
            "
            on:click={() => {navOpened = !navOpened}}
            >=</button>
        </nav>
        <ul>
            <li>1</li>
            <li>2</li>
            <li>3</li>
            <li>4</li>
            <li>5</li>
        </ul>
    </nav>
    <div class="
        flex-[6]
        flex flex-col
        bg-blue-200
    ">
        <nav class="
            md:hidden
            w-full h-12
            bg-red-200
            p-2
        ">
            <button class="
                h-full aspect-square
            "
            on:click={() => {navOpened = !navOpened}}
            >=</button>
        </nav>
        <slot/>
    </div>
</main>