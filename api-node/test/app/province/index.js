/* global describe it */

const chai = require('chai')
const chaiHttp = require('chai-http')
const server = require('../../../server')
const should = chai.should()

chai.use(chaiHttp)

describe('Thailand Provinces', () => {
  it('GET all provinces', (done) => {
    chai.request(server)
      .get('/thailand/provinces')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(77)
        done()
      })
  })

  it('GET by province', (done) => {
    chai.request(server)
      .get('/thailand/provinces/46')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(1)
        done()
      })
  })

  it('GET all districts', (done) => {
    chai.request(server)
      .get('/thailand/provinces/46/districts')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(7)
        done()
      })
  })

  it('GET by district', (done) => {
    chai.request(server)
      .get('/thailand/provinces/46/districts/628')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(1)
        done()
      })
  })

  it('GET all subdistricts', (done) => {
    chai.request(server)
      .get('/thailand/provinces/46/districts/628/subdistricts')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(7)
        done()
      })
  })

  it('GET by subdistrict', (done) => {
    chai.request(server)
      .get('/thailand/provinces/46/districts/628/subdistricts/4987')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(1)
        done()
      })
  })

  it('GET by postcode', (done) => {
    chai.request(server)
      .get('/thailand/postcodes/58110')
      .end((err, res) => {
        res.should.have.status(200)
        res.body.should.be.a('array')
        res.body.length.should.be.eql(13)
        done()
      })
  })
})
