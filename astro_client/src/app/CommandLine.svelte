<script lang="ts">
    import { onMount } from "svelte";
    import {list_commands, run_command} from '../stores/extantion'
    import {main_space_url} from '../stores/windows'

    let opened:Boolean = false
    let input:HTMLInputElement

    let commands:any = {}
    onMount(() => {
        // console.log(123)
        window.onkeydown = (event) => {
            if (event.key == "F1") {
                input.value = ""
                event.preventDefault()
                opened = !opened
                commands = list_commands()
                if (opened) {
                    setTimeout(() => {input.focus()})
                }
            }
            if (opened && event.key == "Escape") {
                event.preventDefault()
                opened = false
            }
        }
        input.oninput = (event) => {
            const value = (event.target as HTMLInputElement).value
            commands = list_commands()
            Object.keys(commands).forEach(key => {
                if (key.includes(value)) {return}
                commands[key] = commands[key].filter((command:string) => {
                    return command.includes(value)
                });
            })
        }
    })

    function run(event:any) {
        event.preventDefault()
        opened = false
        $main_space_url = "http://localhost:3000/"
        run_command(Object.keys(commands)[0], commands[Object.keys(commands)[0]][0])
    }
</script>

<form class="
    {opened ? "" : "hidden"}
    fixed top-2 left-1/2 -translate-x-1/2
    w-screen h-fit flex flex-col items-center
" on:submit={run}>
    <input bind:this={input} type="text" class="flex-1">
    <ul>
        {#each Object.keys(commands) as extantion}
            {#each commands[extantion] as command}
                <li>{extantion}:{command}</li>
            {/each}
        {/each}
    </ul>
</form>