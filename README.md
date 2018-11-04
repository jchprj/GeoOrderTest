# GeoOrderTest

This is an order test based on geographic coordinates. First place an order, then get order list, any unassigned order can be taken.


API doc link: [API DOC](docs/swagger.md)

## Start

```
GeoOrderTest.exe -h
Use MySQL store data, also cache data in memory.
Handle path: /orders, support 3 methods:
  GET: list orders
  POST: place new order
  PATCH: take an unassigned order

Usage:
  GeoOrderTest --config config.yml [flags]

Flags:
      --config string      config file (default "config.yml")
  -h, --help               help for GeoOrderTest
      --logDir string      log file directory (default "logs")
      --logLevel string    log level (default "info")
      --logPrefix string   log file prefix
      --logStd             log to console as well
```

## Compile

Written in Golang 1.10.3, just go build.

## Deploy

Ubuntu 18.04 LTS, docker, start.sh, stop.sh, destroy.sh

### Config

Put Google Maps API key in config.yml under ThirdParty block:

`  GoogleMapsAPIKey: YourKey`

Other is to config http server and database connection.

### Docker

Docker build

Docker run

## Used Go libraries

- Cobra: A commander
- logrus: log system
- gorilla: mux router
- xorm: MySQL engine
- go-swagger: API doc

## Test

benchmark, concurrent

### Run Go test

### Client test
