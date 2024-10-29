const express = require('express')
const app = express();
const pg = require('pg');
require('dotenv').config();

const client  = new pg.Client({
    user: process.env.POSTGRES_USER,
    password: process.env.POSTGRES_ROOT_PASSWORD,
    host: 'pg-service',
    port: process.env.POSTGRESDB_LOCAL_PORT,
    database: process.env.POSTGRES_DATABASE,
    statement_timeout: 10000,
    query_timeout: 10000,
});

client.connect().then(()=> console.log('connected to postgres db'));

app.get('/',(req,res)=> {
console.log('received request')
res.send('home')
})

app.listen('8080',()=> {
console.log('listening on port')})
