const express = require('express')
const app = express()

app.get('/',(req,res)=> {
console.log('received request')
res.send('home')
})

app.listen('8080',()=> {
console.log('listening on port')})
