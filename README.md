# The Open Data

This project aims to translate an open data in generic forms into a JSON format.

## Forex (currency price feed)

The API serves a real-time price feed of many currency pairs that have been provided publicly by [TrueFX](https://www.truefx.com). (npm: [truefx](https://github.com/tonkla/truefx.js)).

* [https://api.op3n.ga/forex](https://api.op3n.ga/forex)
* [https://api.op3n.ga/forex/eurusd](https://api.op3n.ga/forex/eurusd)
* [https://api.op3n.ga/forex/eurusd,usdjpy](https://api.op3n.ga/forex/eurusd,usdjpy)

## SET (stock price feed)

The API serves a real-time price feed and historical prices of the stocks that have been provided publicly by [The Stock Exchange of Thailand](https://marketdata.set.or.th/mkt/marketsummary.do). (npm: [setscraper](https://github.com/tonkla/setscraper)).

* [https://api.op3n.ga/stock](https://api.op3n.ga/stock)
* [https://api.op3n.ga/stock/advanc](https://api.op3n.ga/stock/advanc)
* [https://api.op3n.ga/stock/advanc?from=20171001&to=20171010](https://api.op3n.ga/stock/advanc?from=20171001&to=20171010)

## Administrative divisions of Thailand

The API serves the data of the [Thailand administrative divisions](https://en.wikipedia.org/wiki/Provinces_of_Thailand), that consists of provinces (Changwat), districts (Amphoe), and subdistricts (Tambon).

The current version is accessible by only its database ID,

* [https://api.op3n.ga/thailand/provinces](https://api.op3n.ga/thailand/provinces)
* [https://api.op3n.ga/thailand/provinces/1](https://api.op3n.ga/thailand/provinces/1)
* [https://api.op3n.ga/thailand/provinces/1/districts](https://api.op3n.ga/thailand/provinces/1/districts)
* [https://api.op3n.ga/thailand/provinces/1/districts/1](https://api.op3n.ga/thailand/provinces/1/districts/1)
* [https://api.op3n.ga/thailand/provinces/1/districts/1/subdistricts](https://api.op3n.ga/thailand/provinces/1/districts/1/subdistricts)
* [https://api.op3n.ga/thailand/provinces/1/districts/1/subdistricts/1](https://api.op3n.ga/thailand/provinces/1/districts/1/subdistricts/1)

List the subdistricts by its postal code (one district can have many postal codes, because of transportation),

* [https://api.op3n.ga/thailand/postcodes/10200](https://api.op3n.ga/thailand/postcodes/10200)

**TODO:**

* Add access by the standard code such as [ISO 3166-2:TH](https://en.wikipedia.org/wiki/ISO_3166-2:TH) or [Thai Industry Standard 1099-2548 (TIS 1099)](https://en.wikipedia.org/wiki/Thai_Industrial_Standard_1099-2548)
* Add geolocations (latitude, longitude)
* Add villages (Muban)
