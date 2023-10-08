import express from 'express'
import fetch from 'node-fetch'
import pg from 'pg'

const db = new pg.Client({
    "host":process.env.POSTGRES_HOST||"localhost",
    "port":parseInt(process.env.POSTGRES_PORT||"0")||5432,
    "ssl":eval(process.env.POSTGRES_SSL||"false")||false,
    "user":process.env.POSTGRES_USER||"admin",
    "password":process.env.POSTGRES_PASSWORD||"31415926",
    "database":process.env.POSTGRES_DATABASE||"dev"
})
db.connect().then(() => {console.log("db connected")})
const app = express();

const cors = {
    origin:process.env.CORS_ORIGIN||"*",
    methods:process.env.CORS_METHODS||"*",
    headers:process.env.CORS_HEADERS||"*",
}

const apikey = process.env.APIKEY||""

app.use(express.json())
app.use((req, res, next) => {
    if (req.method != "OPTIONS" && apikey && apikey != "" && req.headers.apikey != apikey) {
        res.status(401).send(`Please provide correct "apikey" header`)
        return
    }
    res.setHeader("Access-Control-Allow-Origin", cors.origin)
    res.setHeader("Access-Control-Allow-Methods", cors.methods)
    res.setHeader("Access-Control-Allow-Headers", cors.headers)
    next()
})

const subscribers = new Map()
const webhook_subscribers = new Map()

function sendEvent(event:{type:"insert", table:string, row:string}|{type:"update", table:string, rows:string[]}|{type:"delete", table:string, row:string}) {
    if (!(subscribers.has(event.table) || webhook_subscribers.has(event.table))) {return}
    subscribers.get(event.table)?.forEach(res => {
        res.write(`event:${event.type}\ndata:${JSON.stringify(event)}\n\n`)
        res.flushHeaders()
    });
    
    webhook_subscribers.get(event.table)?.forEach(hook => {
        fetch(hook.url, {
            method:"POST",
            headers:JSON.parse(hook.headers),
            body:{...hook.body, ...event}
        })
    })
}

app.route("/:table/subscribe")
.get((req, res) => {
    const {table} = req.params

    res.setHeader("Content-Type", "text/event-stream")
    res.setHeader("Connection", "keep-alive")
    res.setHeader("Cache-Control", "no-cache")

    res.flushHeaders()

    if (!subscribers.has(table)) {
        subscribers.set(table, [])
    }

    subscribers.get(table).push(res)

    req.on("close", () => {
        subscribers.set(
            table, 
            subscribers.get(table).filter(el => res !== el)
        )
    })
})
.post((req, res) => {
    const {table} = req.params
    const {body} = req
    if (!webhook_subscribers.has(table)) {
        webhook_subscribers.set(table, [])
    }
    if (body.type == "webhook") {
        try {
            body.headers = JSON.parse(body.headers)
        } catch {
            body.headers = undefined
        }
        webhook_subscribers.get(table).push({
            body:body.body,
            headers:body.headers,
            url:body.url
        })
        res.status(200).send(`webhook was placed for url ${body.url}`)
    }
    res.status(404).send("Sorry, I dont know how to subscribe to that data")
})

app.route("/:table")
.get((req, res) => {
    const {table} = req.params
    var req_query = req.query
    const rows = req_query.select != undefined ? req_query.select : "*"
    delete(req.query.select)
    const query = Object.keys(req_query).map(key => `${key} ${req_query[key]}`).join(" ")

    db.query(`select ${rows} from "${table}" ${query};`).then(result => {
        res.status(200).send(result.rows)
    }).catch(err => {
        res.status(500).send(err)
    })
})
.post((req, res) => {
    const {table} = req.params
    const {body} = req
    const columns = Object.keys(body)

    db.query(`insert into "${table}"(${columns.map(el => `"${el}"`).join(", ")}) values(${columns.map(key => `'${body[key]}'`).join(", ")}) returning *`).then(result => {
        sendEvent({
            "row":result.rows[0],
            "table":table,
            "type":"insert"
        })
        res.status(200).send(result.rows)
    }).catch(err => {
        res.status(500).send(err)
    })
})
.patch((req, res) => {
    const {table} = req.params
    var req_query = req.query
    delete(req.query.select)
    const query = Object.keys(req_query).map(key => `${key} ${req_query[key]}`).join(" ")
    const {body} = req
    const columns = Object.keys(body)

    db.query(`update "${table}" set ${columns.map(key => `"${key}" = '${body[
        key]}'`).join(", ")} ${query} returning *`).then(result => {
            sendEvent({
                "rows":result.rows,
                "table":table,
                "type":"update"
            })
        res.status(200).send(result.rows)
    }).catch(err => {
        res.status(500).send(err)
    })
})
.delete((req, res) => {
    const {table} = req.params
    var req_query = req.query
    delete(req.query.select)
    const query = Object.keys(req_query).map(key => `${key} ${req_query[key]}`).join(" ")

    db.query(`delete from "${table}" ${query} returning *`).then(result => {
        res.status(200).send(result.rows)
        sendEvent({
            "type":"delete",
            "table":table,
            "row":result.rows[0]
        })
    }).catch(err => {
        res.status(500).send(err)
    })
})

app.listen(8080)