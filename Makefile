binary = server
current_dir = $(notdir $(shell pwd))

.PHONY: clean build fmt gen-certs docker-build

all: test build

build: 
	go build -v -o $(binary) ./cmd/server
test: 
	go test -v ./cmd/server

server: test build

fmt:
	go fmt ./cmd/server

gen-certs:
	openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 \
		-subj "/CN=127.0.0.1" \
		-keyout ./certs/private.key -out ./certs/public.crt

run: server
	./$(binary)

docker-build:
	docker build --tag echo:latest --no-cache=true build/
