const truefx = require('truefx')

function getAll (req, res) {
  truefx.get().then(response => {
    res.json(response)
  })
}

function getBySymbol (req, res) {
  const symbol = req.params.symbol
  truefx.get(symbol).then(response => {
    res.json(response)
  })
}

module.exports = { getAll, getBySymbol }
