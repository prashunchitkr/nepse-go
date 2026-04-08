all: build

build:
	go build -o build/example ./...

run-example:
	go run ./cmd/example
