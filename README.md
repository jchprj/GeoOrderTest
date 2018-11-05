# GeoOrderTest

This is an order test based on geographic coordinates. First place an order, then get order list, any unassigned order can be taken.


## Deploy

`start.sh`

Tested under Ubuntu 18.04 LTS, start.sh will install Docker, and build a Docker image, and run the image.  

When building the Docker image, MySQL will be installed, database and table will be created.   

When running, port 8080 and 8081 will be mounted to the local system. The progresses are: go get the code, run go test, generate API doc, start the API server(8080) and API doc server(8081). 

### Config

*MUST* Put Google Maps API key in config.yml under ThirdParty block:

`  GoogleMapsAPIKey: YourKey`

Other parts are http server and database connection configures, no modification needed if use the start.sh.

# API doc

Markdown version documentation link: [API DOC](docs/swagger.md)

API documentation is generated from source code by swagger.  
To generate, run `gen.sh` under `/docs`. Should install [go-swagger](https://goswagger.io/) and [swagger-markdown](https://www.npmjs.com/package/swagger-markdown) first.  
As a demonstration, `httpserver.sh` under `/docs` will serve a web page at port 8081 for API doc.

## Command help

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

`go get -u github.com/jchprj/GeoOrderTest`  
`go install github.com/jchprj/GeoOrderTest`  
`./bin/GeoOrderTest --config config.yml`

## Used Go libraries

- [Cobra](https://github.com/spf13/cobra): a commander
- [Logrus](https://github.com/sirupsen/logrus): log system
- [gorilla/mux](https://github.com/gorilla/mux): mux router
- [xorm](https://github.com/go-xorm/xorm): MySQL engine
- [Swagger](https://github.com/go-swagger/go-swagger): API documentation

## Test


### Go test

`go test ./...`

Include some benchmark and concurrent tests.

Unit test files are:  

* api/handlers/list_test.go
* api/handlers/place_test.go
* api/handlers/take_test.go
* mgr/db_test.go
* mgr/mysql_test.go
* mgr/utils_test.go

