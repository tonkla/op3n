const mongoose = require('mongoose')

const feedSchema = mongoose.Schema({
  symbol: { type: String, index: true },
  date: { type: Number, index: true },
  open: Number,
  high: Number,
  low: Number,
  close: Number,
  volume: Number
})

const stockSchema = mongoose.Schema({
  symbol: { type: String, index: true },
  market: { type: String, index: true },
  name: String,
  nameTH: String,
  indexes: { type: [String], index: true },
  isActive: { type: Boolean, index: true },
  updated: { type: Date, default: Date.now }
})

const indexSchema = mongoose.Schema({
  symbol: { type: String, index: true },
  indexes: { type: [String], index: true },
  updated: { type: Date, default: Date.now }
})

const highlightSchema = mongoose.Schema({
  symbol: { type: String, index: true },
  date: { type: String, index: true },
  asset: Number,
  liability: Number,
  equity: Number,
  revenue: Number,
  profit: Number,
  eps: Number,
  roa: Number,
  roe: Number,
  npm: Number,
  price: Number,
  mktCap: Number,
  pe: Number,
  pbv: Number,
  bvs: Number,
  yield: Number
})

module.exports = { feedSchema, stockSchema, indexSchema, highlightSchema }
