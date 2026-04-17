.PHONY: run build test clean deps

run:
	go run cmd/server/main.go

build:
	go build -o bin/url-shortener cmd/server/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/
	go clean

deps:
	go mod download
	go mod tidy