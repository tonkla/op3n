const express = require('express')
const bodyParser = require('body-parser')

const app = express()
app.use(bodyParser.urlencoded({ extended: true }))
app.use(bodyParser.text())
app.use(bodyParser.json())

app.use((req, res, next) => {
  const ip = req.headers['x-forwarded-for'] || req.connection.remoteAddress
  console.log(`${new Date(Date.now()).toLocaleString()} | ${ip} | ${req.originalUrl}`)
  next()
})

app.get('/', (req, res) => res.json({ 'message': 'Please visit https://github.com/tonkla/op3n for documentation.' }))

const forex = require('./app/forex')
app.get('/forex', forex.getAll)
app.get('/forex/:symbol', forex.getBySymbol)

const province = require('./app/province')
app.get('/thailand', (req, res) => { res.redirect(307, '/thailand/provinces') })
app.get('/thailand/provinces', province.getProvinces)
app.get('/thailand/provinces/:province', province.getByProvince)
app.get('/thailand/provinces/:province/districts', province.getDistricts)
app.get('/thailand/provinces/:province/districts/:district', province.getByDistrict)
app.get('/thailand/provinces/:province/districts/:district/subdistricts', province.getSubdistricts)
app.get('/thailand/provinces/:province/districts/:district/subdistricts/:subdistrict', province.getBySubdistrict)
app.get('/thailand/postcodes', (req, res) => { res.status(404).send([]) })
app.get('/thailand/postcodes/:postcode', province.getByPostcode)

const stock = require('./app/stock')
app.get('/stock', stock.getSummary)
app.get('/stock/:symbol', stock.getBySymbol)

app.listen(3000)
console.log('Server started on port 3000')

// for testing that is called by chai-http
module.exports = app
