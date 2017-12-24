/* global describe it */

const chai = require('chai')
const chaiHttp = require('chai-http')
const server = require('../../../server')
const should = chai.should()

chai.use(chaiHttp)

describe('Stock', () => {
  it('GET summary', (done) => {
    chai.request(server)
      .get('/stock')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('object')
        res.body.should.have.property('updatedAt')
        res.body.should.have.property('indexes')
        done()
      })
  })

  it('GET real-time price of the stock', (done) => {
    chai.request(server)
      .get('/stock/ADVANC')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('object')
        res.body.should.have.property('updatedAt')
        res.body.should.have.property('symbol')
        done()
      })
  })

  it('GET historical prices of the stock', (done) => {
    chai.request(server)
      .get('/stock/ADVANC?from=20171017&to=20171019')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(3)
        res.body[0].should.have.property('date')
        res.body[0].date.should.equal(20171019)
        res.body[2].date.should.equal(20171017)
        done()
      })
  })

  it('GET highlights of the stock', (done) => {
    chai.request(server)
      .get('/stock/ADVANC/highlights')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('object')
        res.body.should.have.property('symbol')
        res.body.should.have.property('highlights')
        res.body.highlights.should.be.a('array')
        res.body.highlights.length.should.be.above(1)
        done()
      })
  })
})
