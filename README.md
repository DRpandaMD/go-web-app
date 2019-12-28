# go-web-app
A go web app

### update readme for repo rebase

## Commands to run locally

### Docker

```docker build -t go-web-app .```
```docker run -d -p 8080:8080 --name running-app go-web-app:latest ```


### locally
run it

```go run main.go```

build it

```go build```

executable 

```./go-web-app```

Install dep & adding a package
```bash
dep version
dep init
dep ensure -add github.com/gorilla/mux
```

testing and cleaning up the testcache
```
go test ./...
go clean -testcache
```