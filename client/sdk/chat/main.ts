
function Connect() {
    const ws = new WebSocket("ws://localhost:8080")
    ws.onopen = (e) => {
      ws.send(JSON.stringify({
        "type":"auth",
        "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiaWx5YSJ9.iM3IqVsTImhXxNybC1oUzGaEkVSNb0KkM2gkxApTb5w"
      }))
      // send("hello")
    }
  
    function send(message:string) {
      ws.send(JSON.stringify({
        "type":"message",
        "content":message
      }))
    }
  
    var events = {}
  
    ws.onmessage = (e) => {
      console.log(e.data)
      var message = ""
      try {
        message = JSON.parse(e.data)
        events[(message as any).type](message)
      } catch {}
    }
  
    function addEventListener(event:string, handler:(e:any) => void) {
      events[event] = handler
    }
  
    return {
      addEventListener, send
    }
}