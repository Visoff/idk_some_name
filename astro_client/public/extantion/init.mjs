var extantions = JSON.parse(localStorage.getItem("extantions"))

if (extantions == null) {
    localStorage.setItem("extantions", JSON.stringify([]))
}

extantions.forEach(async (el, i) => {
    console.log(el)
    const url = el.url
    const ext = await import(url)
    extantions[i] = {...ext, url}
});

async function add_extantion(url) {
    const ext = await import(url)
    localStorage.setItem("extantions", JSON.stringify([...extantions, {config:ext.config, url}]))
    extantions = [...extantions, {...ext, url}]
}

function list_extantions() {
    extantions.forEach(el => {console.log(el)})
}

function update_extantions() {
    const save = [...extantions]
    extantions = []
    save.forEach(el => {
        add_extantion(el.url)
    })
}

window.add_extantion = add_extantion
window.list_extantions = list_extantions
window.update_extantions = update_extantions