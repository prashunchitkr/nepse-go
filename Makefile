all: build-example

build-example:
	go build -o build/example ./cmd/example/main.go

run-example:
	go run ./cmd/example
