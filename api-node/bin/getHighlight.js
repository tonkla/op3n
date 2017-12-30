const scraper = require('setscraper')
const mongoose = require('mongoose')

const { stockSchema, highlightSchema } = require('../app/stock/schema')

mongoose.connect('mongodb://localhost/op3n', { useMongoClient: true })
mongoose.Promise = global.Promise

const Stock = mongoose.model('Stock', stockSchema)
const Highlight = mongoose.model('Highlight', highlightSchema)

Stock.find({ symbol: { $in: [/^A+/] } }).then(stocks => {
  const results = []
  for (let stock of stocks) {
    results.push(scraper.getHighlights(stock.symbol))
  }
  Promise.all(results).then(stocks => {
    for (let s of stocks) {
      for (let h of s.highlights) {
        const _h = new Highlight({
          symbol: s.symbol,
          date: h.date,
          asset: h.asset,
          liability: h.liability,
          equity: h.equity,
          revenue: h.revenue,
          profit: h.profit,
          eps: h.eps,
          roa: h.roa,
          roe: h.roe,
          npm: h.npm,
          price: h.price,
          mktCap: h.mktCap,
          pe: h.pe,
          pbv: h.pbv,
          bvs: h.bvs,
          yield: h.yield
        })
        _h.save(err => {
          if (err) console.error(err)
          console.log(`Highlight Created: ${s.symbol} - ${h.date}`)
        })
      }
    }
    // mongoose.disconnect()
  }).catch(err => {
    console.log(`Exception: ${err}`)
  })
})
