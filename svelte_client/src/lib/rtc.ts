import { host } from "./env"

/*const servers = {
    iceServers:[
        {
            "urls":[
                "stun:stun.l.google.com:19302",
                "stun:stun1.l.google.com:19302"
            ]
        }
    ],
    iceCandidatePoolSize:10
}
*/
const servers = {}

function WaitForAnswer(pc:RTCPeerConnection, call:string) {
    fetch(host+"/conference/answer", {
        method:"GET",
        headers:{
            Call:call
        }
    }).then(response => response.json())
    .then(answer => {
        if (answer == undefined) {
            // console.log("waiting for answer...")
            return setTimeout(() => WaitForAnswer(pc, call), 500)
        }
        // console.log("got answer")
        pc.setRemoteDescription(
            new RTCSessionDescription(
                answer
            )
        )
    })
}

function listenForIce(call:string, callback:(candidate:RTCIceCandidate) => void) {
    let date = new Date().toISOString()
    setInterval(() => {
        fetch(host+"/conference/icepoll", {
            method:"GET",
            headers:{
                Call:call,
                LongPoll:date
            }
        }).then(response => response.json())
        .then(ice => {
            if (!ice || ice.length == 0) {
                return
            }
            date = new Date().toISOString()
            ice.forEach((el:any) => {
                callback(new RTCIceCandidate(JSON.parse(el.Candidate)))
            })
        })
    }, 1000)
}

function AddIceCandidate(ice:RTCIceCandidate, call:string) {
    fetch(host+"/conference/ice", {
        method:"POST",
        headers:{
            Call:call
        },
        body:JSON.stringify(ice.toJSON())
    })
}

export async function call(stream:MediaStream, local_video_tag:HTMLVideoElement, remote_video_tag:HTMLVideoElement) {
    const pc = new RTCPeerConnection(servers)
    // @ts-ignore
    window.pc = pc
    stream.getTracks().forEach(track => {
        pc.addTrack(track, stream)
    });
    const offer = await pc.createOffer()
    const call = await fetch(host+"/conference/call", {
        method:"POST",
        body:JSON.stringify({
            sdp:offer.sdp,
            type:offer.type
        })
    }).then(resposne => resposne.text())
    pc.onicecandidate = (e) => {
        if (e.candidate == undefined) {return}
        AddIceCandidate(e.candidate, call)
    }
    console.log(call)
    pc.setLocalDescription(offer)
    WaitForAnswer(pc, call)
    setTimeout(() => {
        listenForIce(call, (ice) => {
            if (pc.remoteDescription == null) {
                setTimeout(() => {
                    pc.addIceCandidate(ice)
                }, 1000)
                return
            }
            pc.addIceCandidate(ice)
        })
    }, 1000)
    const RemoteStream = new MediaStream()
    pc.ontrack = (e) => {
        if (e.streams.length == 0) {return}
        e.streams[0].getTracks().forEach(track => {
            RemoteStream.addTrack(track)
        })
    }
    remote_video_tag.srcObject = RemoteStream
    local_video_tag.srcObject = stream
}

export async function answer(id:string, local_video_tag:HTMLVideoElement, remote_video_tag:HTMLVideoElement) {
    const pc = new RTCPeerConnection(servers)
    // @ts-ignore
    window.pc = pc
    const stream = new MediaStream()
    const localStream = await navigator.mediaDevices.getUserMedia({"video":true})
    localStream.getTracks().forEach(track => {
        pc.addTrack(track, localStream)
    })
    const starterice = await fetch(host+"/conference/ice", {
        method:"GET",
        headers:{
            Call:id
        }
    }).then(response => response.json())
    pc.onicecandidate = (e) => {
        if (e.candidate == undefined) {return}
        console.log(e.candidate)
        AddIceCandidate(e.candidate, id)
    }
    pc.ontrack = (e) => {
        if (e.streams.length == 0) {return}
        console.log(e.track)
        e.streams[0].getTracks().forEach(track => {
            stream.addTrack(track)
        })
    }
    const offer = await (await fetch(host+"/conference/call", {
        method:"GET",
        headers:{
            "Call":id
        }
    })).json()
    pc.setRemoteDescription(new RTCSessionDescription(offer))
    starterice.forEach((ice:any) => {
        pc.addIceCandidate(new RTCIceCandidate(JSON.parse(ice.Candidate)))
    })
    const answer = await pc.createAnswer()
    fetch(host+"/conference/answer", {
        method:"POST",
        headers:{
            "Call":id
        },
        body:JSON.stringify({
            sdp:answer.sdp,
            type:answer.type
        })
    }).then(response => response.text())
    .then(console.log)
    .catch(console.error)
    listenForIce(id, (ice) => {
        pc.addIceCandidate(ice)
    })
    pc.setLocalDescription(answer)
    remote_video_tag.srcObject = stream
    local_video_tag.srcObject = localStream
}