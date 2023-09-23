<script lang="ts">
    let navOpened:boolean = false
</script>

<main class="
    absolute top-0 left-0 w-screen h-screen
    flex
">
    <script type="module" src="/extantion/init.mjs"></script>
    <script type="module" src="/keyboard/init.mjs"></script>
    <nav class="
        h-screen w-screen
        md:w-16
        max-md:fixed top-0 left-0 transition-transform {navOpened ? "" : "max-md:-translate-x-full"}
        bg-blue-100
    ">
        <nav class="
            md:hidden
            w-full h-12
            p-2
        ">
        </nav>
        <ul>
            <li>1</li>
            <li>2</li>
            <li>3</li>
            <li>4</li>
            <li>5</li>
            <input type="text">
        </ul>
    </nav>
    <div class="
        flex-1
        flex flex-col
    ">
        <slot/>
    </div>
</main>