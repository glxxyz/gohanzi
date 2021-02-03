# gohanzi
Rewrite of my Python hanzi analysis tool in Go

# Getting Started

Run a test server from project root then navigate to http://localhost:8080/

```
go run main.go
```

Working endpoints:

* http://localhost:8080/
* http://localhost:8080/homophones
* ...

# Tests


Run unit tests and display coverage, from project root:

```
go test github.com/glxxyz/gohanzi/containers -coverprofile=c.out && go tool cover -html=c.out
```
