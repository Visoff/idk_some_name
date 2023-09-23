const key_commands = {
    "f1":() => console.log("cmd")
}

window.addEventListener("keydown", event => {
    const key = event.key.toLocaleLowerCase()
    if (key in key_commands) {
        event.preventDefault()
        key_commands[key]()
    }
})