# rungo

Single file go scripts relying on `$GOPATH` break when `GO111MODULE=on` which is enabled by default.

Go 1.17 is set to completely remove `GOPATH` and remove `GO111MODULE=off`

This app will create a temporary directory, copy your script into that directory, download dependencies using `go mod tidy` and run the script.

Usage:

```sh
go install github.com/fkautz/rungo
rungo main.go
```
