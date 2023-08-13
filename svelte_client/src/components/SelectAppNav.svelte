<script>
	import { auth0 } from "$lib/auth0";
	import App from "./SelectAppNav/App.svelte";
    
    function login() {
        if (localStorage.getItem("user_token") == undefined) {
            auth0.authorize({
                responseType:"token",
                redirectUri:location.href,
                domain:"dev-vrsblas6ves78ir5.us.auth0.com",
            })
        } else {
            auth0.client.userInfo(localStorage.getItem("user_token")||"", console.log)
        }
    }
</script>
<div class="
    w-[5rem] p-2 gap-2 bg-blue-100
    max-md:absolute max-md:left-0 max-md:top-0 max-md:-translate-x-full max-md:h-full max-md:w-[16vw]
    transition-all ease-linear
">
    <div class="
        aspect-square w-full p-2 flex
    "
    >
        <a href="/" class="
            flex-1 border-blue-700 border rounded-full flex
        "><button on:click={login} class="flex-1 rounded-full"></a>
    </div>
    <div class="
        w-full h-px bg-black
    "></div>
    <App src="/" popover="chat" />
    <App src="/confirence" popover="video" />
</div>