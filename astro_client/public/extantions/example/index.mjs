const config = {
    name:"messanger",
    icons:["#"],
    commands:["hello world", "open window"]
}

export function command_handler(command) {
    switch (command) {
        case "hello world":
            console.log("hiii")
            break
        case "open window":
            console.log("working on that")
            break
    }
}

export {config}
export function init() {
    console.log(123)
}