<script lang="ts">
	import { browser } from "$app/environment";
    import { onMount } from "svelte";

    // @ts-ignore
    let video_tag:HTMLVideoElement;
    if (browser) {
        let pc = new RTCPeerConnection({});
        onMount(() => {
          if ("getUserMedia" in navigator) {
            // @ts-ignore
            navigator.getUserMedia({video:true}, async (stream:MediaStream) => {
                video_tag.srcObject = stream
                stream.getTracks().forEach(track => {
                    pc.addTrack(track, stream)
                })
                const descr = await pc.createOffer()
                console.log(descr)
                pc.setLocalDescription(descr)
            }, console.error)
          }  
        })
    }
</script>
<div class="flex-1 flex flex-row items-center gap-1">
    <!-- svelte-ignore a11y-media-has-caption -->
    <video
        autoplay bind:this={video_tag}
        class="
            flex-1
        "
    ></video>
</div>