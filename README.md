# gohanzi
Rewrite of my Python hanzi analysis tool in Go

# Getting Started

Run a test server with, from within gohanzi/ and navigate to http://localhost:8080/

```
go run main.go
```

# Tests


Run unit tests and display coverage, from within gohanzi/

```
go test github.com/glxxyz/gohanzi/containers -coverprofile=c.out && go tool cover -html=c.out
```