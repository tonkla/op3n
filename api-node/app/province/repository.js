const sqlite3 = require('sqlite3')

function get (sql, values = []) {
  return new Promise((resolve, reject) => {
    const db = new sqlite3.Database('../data/provinces.db', sqlite3.OPEN_READONLY)
    db.serialize(() => {
      db.all(sql, values, (err, rows) => {
        if (err) reject(err)
        resolve(rows)
      })
    })
    db.close()
  })
}

function getProvinces () {
  return get('SELECT * FROM provinces')
}

function getByProvince (p) {
  return get('SELECT * FROM provinces WHERE id = ?', [p])
}

function getDistricts (p) {
  return get('SELECT * FROM districts WHERE province_id = ?', [p])
}

function getByDistrict (p, d) {
  return get('SELECT * FROM districts WHERE province_id = ? AND id = ?', [p, d])
}

function getSubdistricts (p, d) {
  return get('SELECT * FROM subdistricts WHERE province_id = ? AND district_id = ?', [p, d])
}

function getBySubdistrict (p, d, s) {
  return get('SELECT * FROM subdistricts WHERE province_id = ? AND district_id = ? AND id = ?', [p, d, s])
}

function getByPostcode (p) {
  return get('SELECT * FROM subdistricts WHERE postcode = ?', [p])
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
