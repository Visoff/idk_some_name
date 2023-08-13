<script lang="ts">
	import { browser } from "$app/environment";
	import SelectAppNav from "./SelectAppNav.svelte";
    /*
    import { Auth0Lock } from 'auth0-lock'
    if (browser) {
        const lock = new Auth0Lock(
            "8BzcKDUjJFpMAGe8KPgkUor7pHRmkI2x",
            "dev-vrsblas6ves78ir5.us.auth0.com"
        )
        // lock.show()
    }*/
    import {auth0} from '$lib/auth0'
	import { host } from "$lib/env";

    if (browser) {
        if (location.hash != "" && localStorage.getItem("user_token") == undefined) {
            auth0.parseHash((err, res) => {
                if (err != null) {console.error(err); return}
                const access_token = res?.accessToken||""
                fetch(host+"/user/sign", {
                    method:"POST",
                    body:access_token
                }).then(response => response.text())
                .then(token => {
                    localStorage.setItem("user_token", token)
                    window.location.replace("/")
                })
                .catch(console.error)
            })
        }
    }
</script>
<div class="absolute top-0 left-0 w-screen h-screen flex overflow-hidden">
    <SelectAppNav />
    <slot />
</div>