.DEFAULT_GOAL := run

build:
	go build -o ./bin/app

run: build
	./bin/app

dev: build
	./bin/app -dev

test:
	go test ./... -v

cover:
	go test ./... -coverprofile=coverage.out
