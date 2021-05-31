binary = server
current_dir = $(notdir $(shell pwd))

.PHONY: clean build fmt

all: test build

build: 
	go build -v -o $(binary) ./cmd/server
test: 
	go test -v ./cmd/server

server: test build

fmt:
	go fmt ./cmd/server

run: server
	./$(binary)

docker-build:
	docker-build --tag echo:latest build/
