const repo = require('./repository')

function getProvinces (req, res) {
  repo.getProvinces().then(response => {
    res.json(response)
  }).catch(error => {
    console.err(error)
  })
}

function getByProvince (req, res) {
  const p = req.params.province
  repo.getByProvince(p).then(response => {
    res.json(response)
  }).catch(error => {
    console.err(error)
  })
}

function getDistricts (req, res) {
  const p = req.params.province
  repo.getDistricts(p).then(response => {
    res.json(response)
  }).catch(error => {
    console.err(error)
  })
}

function getByDistrict (req, res) {
  const p = req.params.province
  const d = req.params.district
  repo.getByDistrict(p, d).then(response => {
    res.json(response)
  }).catch(error => {
    console.err(error)
  })
}

function getSubdistricts (req, res) {
  const p = req.params.province
  const d = req.params.district
  repo.getSubdistricts(p, d).then(response => {
    res.json(response)
  }).catch(error => {
    console.err(error)
  })
}

function getBySubdistrict (req, res) {
  const p = req.params.province
  const d = req.params.district
  const s = req.params.subdistrict
  repo.getBySubdistrict(p, d, s).then(response => {
    res.json(response)
  }).catch(error => {
    console.err(error)
  })
}

function getByPostcode (req, res) {
  const p = req.params.postcode
  repo.getByPostcode(p).then(response => {
    res.json(response)
  }).catch(error => {
    console.err(error)
  })
}

module.exports = {
  getProvinces,
  getByProvince,
  getDistricts,
  getByDistrict,
  getSubdistricts,
  getBySubdistrict,
  getByPostcode
}
