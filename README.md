# The Go-Microservice generator

## Installation

Install [Git](http://git-scm.com), [node.js](http://nodejs.org), and [Go 1.1](http://golang.org/).  The development mode also requires [SQLite](http://www.sqlite.org).

Install Yeoman:

    npm install -g yo

Install the Go-Microservice generator:

    npm install -g generator-go-microservice


## Creating a go service

In a new directory, generate the service:

    yo go-microservice

Get the dependencies:

    goop install

Run the service:

    goop exec go run server.go

Your service will run at [http://localhost:9001](http://localhost:9001).

## Simple apache Bench results for the 3 Libraries

`ab -n 5000 http://localhost:9001/`

### Echo

```
Complete requests:      5000
Failed requests:        0
Total transferred:      580000 bytes
HTML transferred:       70000 bytes
Requests per second:    4435.47 [#/sec] (mean)
Time per request:       0.225 [ms] (mean)
Time per request:       0.225 [ms] (mean, across all concurrent requests)
Transfer rate:          502.46 [Kbytes/sec] received
```

### http

```
Complete requests:      5000
Failed requests:        0
Total transferred:      655000 bytes
HTML transferred:       70000 bytes
Requests per second:    4660.73 [#/sec] (mean)
Time per request:       0.215 [ms] (mean)
Time per request:       0.215 [ms] (mean, across all concurrent requests)
Transfer rate:          596.25 [Kbytes/sec] received
```


### httprouter

```
Complete requests:      5000
Failed requests:        0
Total transferred:      655000 bytes
HTML transferred:       70000 bytes
Requests per second:    4935.01 [#/sec] (mean)
Time per request:       0.203 [ms] (mean)
Time per request:       0.203 [ms] (mean, across all concurrent requests)
Transfer rate:          631.33 [Kbytes/sec] received
```