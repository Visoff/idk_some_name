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

connect("cbd06730-e267-462c-a431-ec0aa26b1fe4", "7e0c054f-613d-4e86-8d39-467d1ecad046")