const events = {}

const remote_streams = {}

/**
 * 
 * @param {MediaStream} stream 
 * @param {string} auth 
 */
function addstream(stream, auth) {
    remote_streams[auth] = stream
    events["streamchange"] && events["streamchange"](remote_streams)
}

/**
 * 
 * @param {string} auth 
 */
function removestream(auth) {
    delete(remote_streams[auth])
    events["streamchange"] && events["streamchange"](remote_streams)
}

const servers = {
    iceServers:[
        {
            urls:["stun:stun1.l.google.com:19302", "stun:stun2.l.google.com:19302"]
        }
    ],
    iceCandidatesPoolSize:10
}

const room = {}

/**
 * 
 * @param {MediaStream} stream 
 * @param {string} room_id 
 */
async function connect(stream, room_id) {
    const url = new URL("ws://localhost:8080/"+room_id)
    const ws = new WebSocket(url)

    var auth = ""

    ws.addEventListener("message", async (e) => {
        const data = JSON.parse(e.data)
        console.log(data)
        if (data.type == "auth") {
            auth = data.auth
        } else if (data.type == "history") {
            data.users.forEach(async user => {
                if (user == auth) {return}
                room[user] = new RTCPeerConnection(servers)
                stream.getTracks().forEach(track => {
                    room[user].addTrack(track, stream)
                })
                room[user].onicecandidate = (e) => {
                    if (e.candidate == undefined) {return}
                    ws.send(JSON.stringify({
                        type:"ice",
                        candidate:e.candidate.toJSON(),
                        to:user
                    }))
                }
                const offer = await room[user].createOffer()
                room[user].setLocalDescription(offer)
                ws.send(JSON.stringify({
                    type:"offer",
                    to:user,
                    offer:{sdp:offer.sdp, type:offer.type}
                }))
            })
        } else if (data.type == "ice") {
            room[data.from].addIceCandidate(new RTCIceCandidate(data.candidate))
        } else if (data.type == "offer") {
            room[data.from] = new RTCPeerConnection(servers)
            
            room[data.from].onicecandidate = (e) => {
                if (e.candidate == undefined) {return}
                ws.send(JSON.stringify({
                    type:"ice",
                    candidate:e.candidate.toJSON(),
                    to:data.from
                }))
            }
            
            stream.getTracks().forEach(track => {
                room[data.from].addTrack(track, stream)
            })

            const remote_stream = new MediaStream()

            room[data.from].ontrack = (e) => {
                if (e.track == undefined) {return}
                remote_stream.addTrack(e.track)
            }

            addstream(remote_stream, data.from)

            room[data.from].setRemoteDescription(new RTCSessionDescription(data.offer))
            const answer = await room[data.from].createAnswer()
            room[data.from].setLocalDescription(answer)
            ws.send(JSON.stringify({
                type:"answer",
                to:data.from,
                offer:{sdp:answer.sdp, type:answer.type}
            }))
        } else if (data.type == "answer") {
            if (room[data.from] == undefined) {
                return
            }
            const stream = new MediaStream()
            
            room[data.from].ontrack = (e) => {
                if (e.track == undefined) {return}
                stream.addTrack(e.track)
            }
            window.remote_stream = stream

            room[data.from].setRemoteDescription(data.offer)

            addstream(stream, data.from)
        } else if (data.type == "disconnect") {
            room[data.from] = undefined
            removestream(data.from)
        }
    })
}

/**
 * 
 * @param {string} event 
 * @param {(event: any) => void} func 
 */
function addEventListener(event, func) {
    events[event] = func
}

module.exports = {
    addEventListener,
    connect
}