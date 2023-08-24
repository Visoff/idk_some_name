<script lang="ts">
    import req from '@lib/fetch'
  import { file } from 'dist/server/chunks/prerender.1097080f.mjs';
    
    let path:string = "/"
    let files:any[] = []

    $: req("/bucket/dir"+path, {
        headers:{
            Authorization:"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoidXNlcjEyMyJ9.dC_hrWMOlI2Pg9M-Chgg4DIwZjJA7RzA0BeKhBpYlgE"
        }
    }).then(response => response.json())
    .then(response_files => {
        files = response_files
    })
    .catch(console.error)
    
    function readFile(name:string) {
        req("/bucket/file"+"/"+[...(path.split("/")), name].filter(el => el != "").join("/"), {
            headers:{
                Authorization:"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoidXNlcjEyMyJ9.dC_hrWMOlI2Pg9M-Chgg4DIwZjJA7RzA0BeKhBpYlgE"
            }
        }).then(response => response.text())
        .then(file_content => {
            alert(file_content)
        })
        .catch(console.error)
    }
</script>

<div>
    {path}
    {#if path != "/"}
        <button
            on:click={(e) => {
                const a = path.split("/").filter(el => el != "")
                a.pop()
                path="/"+a.join("/")
            }}
        >...</button>
    {/if}
    {#each files as file}
        {#if file.IsDir}
            <button
                class="bg-blue-400"
                on:click={(e) => {
                    path="/"+[...(path.split("/")), file.Name].filter(el => el != "").join("/")
                }}
            >
                {file.Name}
            </button>
        {:else}
            <button
                class="bg-blue-200"
                on:click={() => {
                    readFile(file.Name)
                }}
            >
                {file.Name}
            </button>
        {/if}
    {/each}
</div>