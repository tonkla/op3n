const scraper = require('setscraper')
const mongoose = require('mongoose')
mongoose.connect('mongodb://localhost/op3n', { useMongoClient: true })
mongoose.Promise = global.Promise

const feedSchema = mongoose.Schema({
  symbol: { type: String, index: true },
  date: { type: Number, index: true },
  open: Number,
  high: Number,
  low: Number,
  close: Number,
  volume: Number
})
const Feed = mongoose.model('Feed', feedSchema)

function getSummary (req, res) {
  scraper.get().then(response => {
    res.json(response)
  })
}

function getBySymbol (req, res) {
  const symbol = req.params.symbol
  const q = req.query
  if (q.from) {
    let dateCond = { '$gte': parseInt(q.from) }
    if (q.to) Object.assign(dateCond, { '$lte': parseInt(q.to) })
    Feed.find({ symbol: symbol.toUpperCase(), date: dateCond })
        .select('-_id symbol date open high low close volume')
        .sort({ date: -1 })
        .exec((err, result) => {
          if (err) throw err
          res.json(result)
        })
  } else {
    scraper.get(symbol).then(response => {
      res.json(response)
    })
  }
}

module.exports = { getSummary, getBySymbol }
