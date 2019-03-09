const uWS = require('uWebSockets.js')
const { readFileSync } = require('fs')
const { join } = require('path')
const port = process.env.PORT || 9001
const data = readFileSync(join(__dirname, '../data.json'))
const app = uWS.App().get('/', (res) => {
  res.end(data)
}).listen(port, (token) => {
  if (token) {
    console.log('Listening to port ' + port)
  } else {
    console.log('Failed to listen to port ' + port)
  }
})