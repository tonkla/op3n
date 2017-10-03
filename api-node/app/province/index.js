const repo = require('./repository')

function getProvinces (req, res) {
  repo.getProvinces().then(response => {
    res.json(response)
  }).catch(error => {
    console.err(error)
  })
}

function getByProvince (req, res) {
  const province = req.params.province
  repo.getByProvince(province).then(response => {
    res.json(response)
  }).catch(error => {
    console.err(error)
  })
}

module.exports = { getProvinces, getByProvince }
