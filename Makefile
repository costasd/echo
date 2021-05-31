binary = server
current_dir = $(notdir $(shell pwd))

.PHONY: clean build release

all: test build

build: 
	go build -v -o $(binary) ./cmd/server
test: 
	go test -v ./cmd/server

server: test build

run: server
	./$(binary)

docker-build:
	docker-build --tag echo:latest build/
