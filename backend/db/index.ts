import express from 'express'
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

app.use(express.json())

const subscribers = new Map()

function sendEvent(event:{type:"insert", table:string, row:string}|{type:"update", table:string, rows:string[]}|{type:"delete", table:string}) {
    if (!subscribers.has(event.table)) {return}
    subscribers.get(event.table).forEach(res => {
        res.write(`event:${event.type}\ndata:${JSON.stringify(event)}\n\n`)
        res.flushHeaders()
    });
}

app.get("/:table/subscribe", (req, res) => {
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
            "table":table
        })
    }).catch(err => {
        res.status(500).send(err)
    })
})

app.listen(8080)