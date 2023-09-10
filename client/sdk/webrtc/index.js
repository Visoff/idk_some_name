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

var room = {}
var datachannelmessagehandler = () => {}

/**
 * 
 * @param {{url:string, room_id:string, stream:MediaStream}} 
 */
function connect({url, room_id, stream}) {
    const url_obj = new URL(url)
    url_obj.pathname += encodeURI(room_id)
    const ws = new WebSocket(url_obj)

    var auth = ""

    ws.addEventListener("message", async (e) => {
        const data = JSON.parse(e.data)
        console.log(data)
        if (data.type == "auth") {
            auth = data.auth
        } else if (data.type == "history") {
            data.users.forEach(async user => {
                if (user == auth) {return}
                const pc = new RTCPeerConnection(servers)
                const datachannel = pc.createDataChannel("datachannel")

                datachannel.onopen = console.log
                datachannel.onmessage = datachannelmessagehandler

                stream.getTracks().forEach(track => {
                    pc.addTrack(track, stream)
                })
                pc.onicecandidate = (e) => {
                    if (e.candidate == undefined) {return}
                    ws.send(JSON.stringify({
                        type:"ice",
                        candidate:e.candidate.toJSON(),
                        to:user
                    }))
                }
                const offer = await pc.createOffer()
                pc.setLocalDescription(offer)
                ws.send(JSON.stringify({
                    type:"offer",
                    to:user,
                    offer:{sdp:offer.sdp, type:offer.type}
                }))
                room[user] = {pc, datachannel}
            })
        } else if (data.type == "ice") {
            room[data.from].pc.addIceCandidate(new RTCIceCandidate(data.candidate))
        } else if (data.type == "offer") {
            const pc = new RTCPeerConnection(servers)
            pc.ondatachannel = (e) => {
                room[data.from].datachannel = e.channel
                e.channel.onmessage = datachannelmessagehandler
            }
            
            pc.onicecandidate = (e) => {
                if (e.candidate == undefined) {return}
                ws.send(JSON.stringify({
                    type:"ice",
                    candidate:e.candidate.toJSON(),
                    to:data.from
                }))
            }
            
            stream.getTracks().forEach(track => {
                pc.addTrack(track, stream)
            })

            const remote_stream = new MediaStream()

            pc.ontrack = (e) => {
                if (e.track == undefined) {return}
                remote_stream.addTrack(e.track)
            }

            addstream(remote_stream, data.from)

            pc.setRemoteDescription(new RTCSessionDescription(data.offer))
            const answer = await pc.createAnswer()
            pc.setLocalDescription(answer)
            ws.send(JSON.stringify({
                type:"answer",
                to:data.from,
                offer:{sdp:answer.sdp, type:answer.type}
            }))
            room[data.from] = {pc}
        } else if (data.type == "answer") {
            if (room[data.from] == undefined) {
                return
            }
            const stream = new MediaStream()
            
            room[data.from].pc.ontrack = (e) => {
                if (e.track == undefined) {return}
                stream.addTrack(e.track)
            }
            window.remote_stream = stream

            room[data.from].pc.setRemoteDescription(data.offer)

            addstream(stream, data.from)
        } else if (data.type == "disconnect") {
            delete(room[data.from])
            removestream(data.from)
        }
    })
}

/**
 * 
 * @param {"streamchange"|"datachannel"} event 
 * @param {(event: any) => void} func 
 */
function addEventListener(event, func) {
    if (event == "datachannel") {
            datachannelmessagehandler = func
        return
    }
    events[event] = func
}

/**
 * 
 * @param {string} message 
 */
function sendData(message) {
    Object.keys(room).forEach(key => {
        /** @type {RTCDataChannel} */
        const datachannel = room[key].datachannel
        datachannel.send(message)
    })
}

module.exports = {
    addEventListener,
    connect,
    sendData
}