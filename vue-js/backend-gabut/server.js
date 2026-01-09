const express = require('express')
const mysql = require('mysql2')
const cors = require('cors')
const app = express()

app.use(cors())
app.use(express.json())
const db = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: '',
    database: 'gabut_db'
})

app.get('/items', (req, res) => {
    db.query('SELECT * FROM items', (err, result) => {
        res.send(result)
    })
})
app.post('/items', (req, res) => {
    const { name, description} = req.body
    db.query('INSERT INTO items (name, description) VALUES (?, ?)', [name, description], (err, result) => {
        res.send(result)
    })
})

app.listen(3000, () => console.log('server dh jalan, aman wok'))