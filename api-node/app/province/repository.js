const sqlite3 = require('sqlite3')

function getProvinces () {
  return new Promise((resolve, reject) => {
    const db = new sqlite3.Database('../data/provinces.db', sqlite3.OPEN_READONLY)
    db.serialize(() => {
      db.all('SELECT * FROM provinces', (err, rows) => {
        if (err) reject(err)
        resolve(rows)
      })
    })
    db.close()
  })
}

function getByProvince (province) {
  return new Promise((resolve, reject) => {
    const db = new sqlite3.Database('../data/provinces.db', sqlite3.OPEN_READONLY)
    db.serialize(() => {
      db.all('SELECT * FROM provinces WHERE id = ?', [province], (err, rows) => {
        if (err) reject(err)
        resolve(rows[0])
      })
    })
    db.close()
  })
}

module.exports = { getProvinces, getByProvince }
