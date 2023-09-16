# Examples
```ts
    import {connect} from 'faketime-db-sdk'
    const db = connect({url:"your url", apikey:"your key"})

    // select everything
    db.select("*").from("User").query().then(rows => {
        // ...
    })

    // "advanced" select
    db.select("row1", "row2").from("Message").where(`"row3" = 'value'`).query().then(rows => {
        // ...
    })

    // insert
    db.insert({
        table:"User",
        values:{
            username:"data"
        }
    })

    // place event listener
    db.addEventListener("User", "insert", (e) => {
        console.log(e.type)// insert
        // ...
    })
```