const events = {}

function connect(chat_id, user_id) {
    ws = new WebSocket("ws:localhost:8080/"+chat_id)
    ws.addEventListener("error", console.error)
    ws.addEventListener("message", (e) => {
        // console.log(e.data)
        const data = JSON.parse(e.data)
        console.log(data)
        switch (data.type) {
            case "message":
                events["message"] && events["message"](data)
                break
        }
    })
    ws.addEventListener("open", () => {
        ws.send(JSON.stringify({auth:user_id}))
    })
}

console.log("hello")