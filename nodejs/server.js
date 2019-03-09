const http = require('http')
const { readFileSync } = require('fs')
const { join } = require('path')
const port = process.env.PORT || 3000
const data = readFileSync(join(__dirname, '../data.json'))
const requestHandler = (_, response) => {
    response.setHeader('Content-Type', 'application/json')
    response.end(data)
}
const server = http.createServer(requestHandler)
server.listen(port, (err) => {
    if (err) {
        return console.log('something bad happened', err)
    }
    console.log(`server is listening on ${port}`)
})
