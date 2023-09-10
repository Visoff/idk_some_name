<script lang="ts">
    import rtc from 'network-class-rtc-sdk'
    import { onMount } from 'svelte';

    let streams:MediaStream[] = [];
    let local_stream:MediaStream|undefined;
    let videoContainer:HTMLElement;

    onMount(() => {
        navigator.mediaDevices.getUserMedia({video:true, audio:true}).then(stream => {
            local_stream = stream
            rtc.connect({url:"ws://localhost:8080" ,stream:stream, room_id:"room1"})
            rtc.addEventListener("streamchange", (new_streams:any) => {
                streams = Object.keys(new_streams).map(key => new_streams[key])
            })
        })
    })
    
    function update(local_stream:MediaStream|undefined, streams:MediaStream[]) {
        videoContainer.innerHTML = ""
        const vid = document.createElement("video")
        vid.srcObject = local_stream||new MediaStream()
        vid.autoplay = true
        vid.muted = true
        videoContainer.append(vid)
        streams.forEach(stream => {
            const vid = document.createElement("video")
            vid.srcObject = stream
            vid.autoplay = true
            videoContainer.append(vid)
        })
    }
    $: videoContainer && update(local_stream, streams)
</script>

<div>
    <div class="videoplace" bind:this={videoContainer}></div>
</div>