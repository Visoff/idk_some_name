<script lang="ts">
	import { browser } from "$app/environment";
	import { answer, call } from "$lib/rtc";
    import { onMount } from "svelte";

    // @ts-ignore
    let video_tag:HTMLVideoElement;
    // @ts-ignore
    let video_tag_v2:HTMLVideoElement;
    function Call() {
        navigator.mediaDevices.getUserMedia({video:true, audio:true}).then(stream => {
            call(stream, video_tag, video_tag_v2)
        })
    }

    function Answer() {
        const id = prompt()||""
        answer(id, video_tag, video_tag_v2)
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
    <!-- svelte-ignore a11y-media-has-caption -->
    <video
        autoplay bind:this={video_tag_v2}
        class="
            flex-1
        "
    ></video>
    <button on:click={Call}>call</button>
    <button on:click={Answer}>answer</button>
</div>