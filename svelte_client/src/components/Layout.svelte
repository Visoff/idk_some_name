<script>
	import { browser } from "$app/environment";
	import SelectAppNav from "./SelectAppNav.svelte";
    import {clerk as clerk_import, clerk_ready} from "$lib"

    let signOut = () => {}
    
    const clerk = $clerk_import
    if (browser) {
        (async () => {
            await clerk.load({
                "afterSignInUrl":"http://localhost:5173",
                "afterSignUpUrl":"http://localhost:5173",
            })
            if (clerk.user == undefined) {
                clerk.redirectToSignIn()
                return
            }
            if (localStorage.getItem("user_token") == undefined) {
                fetch("http://localhost:8080/api/user/sign/", {
                    method:"POST",
                    body:clerk.user.id
                }).then(response => response.text())
                .then(token => {
                    localStorage.setItem("user_token", token)
                })
                .catch(console.error)
            }
            // @ts-ignore
            window.signOut = clerk.signOut
            $clerk_ready = true
        })()
    }
</script>
<div class="absolute top-0 left-0 w-screen h-screen flex overflow-hidden">
    <SelectAppNav />
    <slot />
</div>