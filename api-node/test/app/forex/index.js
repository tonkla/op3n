/* global describe it */

const chai = require('chai')
const chaiHttp = require('chai-http')
const server = require('../../../server')
const should = chai.should()

chai.use(chaiHttp)

describe('Forex', () => {
  it('GET all symbols', (done) => {
    chai.request(server)
      .get('/forex')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(10)
        done()
      })
  })

  it('GET by symbol', (done) => {
    chai.request(server)
      .get('/forex/eurusd')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(1)
        done()
      })
  })

  it('GET by symbols', (done) => {
    chai.request(server)
      .get('/forex/eurusd,usdjpy')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(2)
        done()
      })
  })
})
