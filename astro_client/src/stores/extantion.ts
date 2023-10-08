var extantions:any = JSON.parse(localStorage.getItem("extantions")||"{}")

async function add_extantion(url:string) {
    const ext = await import(url)
    if (ext.config == undefined || ext.config.name == undefined) {
        console.error("incorrect plugin")
    }
    localStorage.setItem("extantions", JSON.stringify({...extantions, [ext.config.name]:{config:ext.config, url}}))
    extantions = {...extantions, [ext.config.name]:{...ext, url}}
}

export function list_extantions() {
    Object.keys(extantions).forEach((key:string) => {console.log(extantions[key])})
}

export function update_extantions() {
    const save = {...extantions}
    extantions = {}
    Object.keys(save).forEach(key => {
        add_extantion(save[key].url)
    })
}

export function run_command(extantion:string, command:string) {
    if (!(extantion in extantions)) {
        return new Error("Incorrect extantion name")
    }
    if (!extantions[extantion].config.commands.includes(command)) {
        return new Error("Specified plugin does not have specified command")
    }
    extantions[extantion].command_handler(command)
}

export function list_commands() {
    var commands:{[name:string]:any} = {}
    Object.keys(extantions).forEach(key => {
        commands[key] = extantions[key].config.commands
    })
    return commands
}

export function init() {
    Object.keys(extantions).forEach(async (key) => {
        const url = extantions[key].url
        const ext = await import(url)
        extantions[key] = {...ext, url}
    });
}
