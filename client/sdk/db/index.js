
/** @type {{url?:URL, apikey?:string}} */
var clientdata = {
    url:undefined, apikey:undefined
}

/** @type {boolean} */
var listening_for_events = false

/**
 * 
 * @param {{url?:string, apikey?:string}} param0 
 */
function connect({url, apikey}) {
    if (url == undefined || apikey == undefined) {
        console.error("Please provide url and apikey to connect to faketime db!!!")
        return
    }
    const Url = new URL(url)
    clientdata = {url:Url, apikey}

    return {
        select,
        insert,
        update,
        delete:db_delete,
        addEventListener
    }
}

/**
 * 
 * @param {{table:string, values:{}}}} param0 
 */
function insert({table, values}) {
    const url = new URL(clientdata.url.toString())
    if (table == "" || values == "" || url.origin == "") {
        console.error("Incorrect data error")
        return
    }
    if (url.pathname == "/") {
        url.pathname = `/${table}`
    } else {
        url.pathname=[...url.pathname.split("/"), table].join("/")
    }
    fetch(url, {
        method:"POST",
        body:JSON.stringify(values),
        headers:{
            "Content-Type":"application/json"
        }
    }).catch(async resp => {console.error(resp.text())})
}

/**
 * 
 * @param {string[]|"*"} rows
 */
function select(rows) {
    const url = new URL(clientdata.url.toString())
    if (url == undefined) {
        console.error("Please connet to db!")
        return
    }
    if (rows != "*") {
        url.searchParams.append("select", rows.map(el => `"${el}"`).join(", "))
    }

    /**
     * @param {string} table 
     */
    function from(table) {
        if (url.pathname == "/") {
            url.pathname = `/${table}`
        } else {
            url.pathname=[...url.pathname.split("/"), table].join("/")
        }

        return {
            where,
            // join,
            query
        }
    }

    /**
     * 
     * @param {?string} condition 
     */
    function where(condition) {
        if (condition == undefined) {
            return
        }
        url.searchParams.append("where", condition)

        return {
            query
        }
    }

    async function query() {
        return fetch(url).then(async resp => {
            if (resp.status != 200) {
                throw new Error(await resp.text())
            } else {
                return resp.json()
            }
        })
    }

    return {
        from
    }
}

/**
 * 
 * @param {{query:string, values:{[any:string]:any}}}} param0 
 */
function update({table, query, values}) {
    const url = new URL(clientdata.url.toString())
    if (url.pathname == "/") {
        url.pathname = `/${table}`
    } else {
        url.pathname=[...url.pathname.split("/"), table].join("/")
    }
    url.searchParams.append("where", query)
    fetch(url, {
        method:"PATCH",
        body:JSON.stringify(values),
        headers:{
            "Content-Type":"application/json"
        }
    })
}

/**
 * @param {Object} param0 
 * @param {string} param0.table 
 * @param {string} param0.query 
 */
function db_delete({table, query}) {
    const url = new URL(clientdata.url.toString())
    if (url.pathname == "/") {
        url.pathname = `/${table}`
    } else {
        url.pathname=[...url.pathname.split("/"), table].join("/")
    }
    url.searchParams.append("where", query)
    fetch(url, {
        method:"DELETE",
        headers:{
            "Content-Type":"application/json"
        }
    })
}

/**
 * @param {"insert"|"update"|"delete"} event
 * @param {*} f
 */
function addEventListener(table, event, f) {
    const url = new URL(clientdata.url.toString())
    if (url.pathname == "/") {
        url.pathname = `/${table}/subscribe`
    } else {
        url.pathname=[...url.pathname.split("/"), table, "subscribe"].join("/")
    }
    const sse = new EventSource(url)
    sse.addEventListener(event, (e) => {
        f(e.data)
    })
}

module.exports = {connect}