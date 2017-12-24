;(async function () {
  const scraper = require('setscraper')
  const mongoose = require('mongoose')
  mongoose.connect('mongodb://localhost/op3n', { useMongoClient: true })
  mongoose.Promise = global.Promise

  // MongoDB CLI: db.stocks.drop()

  const stockSchema = mongoose.Schema({
    symbol: { type: String, index: true },
    market: { type: String, index: true },
    name: String,
    nameTH: String,
    indexes: { type: [String], index: true },
    isActive: { type: Boolean, index: true },
    updated: { type: Date, default: Date.now }
  })
  const Stock = mongoose.model('Stock', stockSchema)

  const indexSchema = mongoose.Schema({
    symbol: { type: String, index: true },
    indexes: { type: [String], index: true },
    updated: { type: Date, default: Date.now }
  })
  const IndexStock = mongoose.model('IndexStock', indexSchema)

  async function updateStocks () {
    const stocks = await scraper.getStocks({ lang: 'th' })
    for (let s of stocks) {
      const stock = new Stock({ symbol: s.symbol, market: s.market, name: s.name, nameTH: s.nameTH })
      stock.save(err => {
        if (err) console.error(err)
        console.log(`Created: ${s.symbol}`)
      })
    }
  }

  async function updateIndexStocks () {
    const results = []
    const indexes = ['SET50', 'SET100', 'SETHD', 'sSET']
    for (let index of indexes) {
      results.push(scraper.get(index))
    }
    return saveIndexes(await Promise.all(results))
  }

  function saveIndexes (results) {
    for (let data of results) {
      for (let s of data.stocks) {
        Stock.update({ symbol: s.symbol }, { $push: { indexes: data.index } }, (err, raw) => {
          if (err) console.error(err)
          console.log(`Updated: ${s.symbol}`)
        })

        IndexStock.findOneAndUpdate({ symbol: s.symbol },
          { $push: { indexes: data.index } },
          { new: true,
            upsert: true,
            setDefaultsOnInsert: true
          }, (err, raw) => {
            if (err) console.error(err)
            console.log(`Index Updated: ${s.symbol}`)
          })
      }
    }
  }

  await updateStocks()
  await updateIndexStocks()
})()
