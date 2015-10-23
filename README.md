# The Go-Microservice generator 

## Installation

Install [Git](http://git-scm.com), [node.js](http://nodejs.org), and [Go 1.1](http://golang.org/).  The development mode also requires [SQLite](http://www.sqlite.org).

Install Yeoman:

    npm install -g yo

Install the Angular-Go-Martini generator:

    npm install -g generator-go-microservice


## Creating a go service

In a new directory, generate the service:

    yo go-microservice

Get the dependencies:

    goop install 

Run the service:

    goop exec go run server.go

Your service will run at [http://localhost:9001](http://localhost:9001).

