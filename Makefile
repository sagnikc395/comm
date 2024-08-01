build:
    @go build -o bin/comm cmd/main.go

test:
    @go test -v ./...

run: build
    @./bin/comm
