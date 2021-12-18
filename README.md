### Tripwire

A Golang SDK for [binance](https://www.binance.com) API.

[![Build Status](https://travis-ci.org/workfoxes/tripwire.svg?branch=main)](https://travis-ci.org/workfoxes/tripwire)
[![GoDoc](https://godoc.org/github.com/workfoxes/tripwire?status.svg)](https://godoc.org/github.com/workfoxes/tripwire)
[![Go Report Card](https://goreportcard.com/badge/github.com/workfoxes/tripwire)](https://goreportcard.com/report/github.com/workfoxes/tripwire)
[![codecov](https://codecov.io/gh/workfoxes/tripwire/branch/main/graph/badge.svg)](https://codecov.io/gh/workfoxes/tripwire)

All the REST APIs listed in [binance API document](https://github.com/binance-exchange/binance-official-api-docs) are implemented, as well as the websocket APIs.

For best compatibility, please use Go >= 1.17.3.

Make sure you have read binance API document before continuing.

### API List

Name | Description | Status
------------ | ------------ | ------------
[rest-api.md](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md) | Details on the Rest API (/api) | <input type="checkbox" checked> Implemented
[web-socket-streams.md](https://github.com/binance-exchange/binance-official-api-docs/blob/master/web-socket-streams.md) | Details on available streams and payloads | <input type="checkbox" checked>  Implemented
[user-data-stream.md](https://github.com/binance-exchange/binance-official-api-docs/blob/master/user-data-stream.md) | Details on the dedicated account stream | <input type="checkbox" checked>  Implemented
[margin-api.md](https://github.com/binance-exchange/binance-official-api-docs/blob/master/margin-api.md) | Details on the Margin API (/sapi) | <input type="checkbox" checked>  Implemented
[futures-api.md](https://binance-docs.github.io/apidocs/futures/en/#general-info) | Details on the Futures API (/fapi) | <input type="checkbox" checked>  Partially Implemented
[delivery-api.md](https://binance-docs.github.io/apidocs/delivery/en/#general-info) | Details on the Coin-M Futures API (/dapi) | <input type="checkbox" checked>  Partially Implemented

### Installation

```shell
go get github.com/workfoxes/tripwire
```

### Importing

```golang
import (
    "github.com/workfoxes/tripwire"
)
```

### Documentation

[![GoDoc](https://godoc.org/github.com/workfoxes/tripwire?status.svg)](https://godoc.org/github.com/workfoxes/tripwire)