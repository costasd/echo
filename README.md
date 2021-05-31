## About
A simple service to echo posted json back to the client

## Install

`go get github.com/costasd/echo/cmd/server`

## Features

A simple server listening on `127.0.0.1:3000` for incoming connections.
* If json is POSTed or PUT, json is returned, with an extra key, `echoed:true`
* In all other cases, an error is returned.

Optionally, TLS can be enabled on `127.0.0.1:3001`, by providing a certificate keypair.

## Build/Develop
### Locally with make
To build the project, just execute `make` from project's root dir:

```
$ make
go test -v ./cmd/server
=== RUN   TestEchoServer
=== RUN   TestEchoServer/Test_Post_Valid_Json
2021/05/31 16:42:23 200 [] /api/echo OK: echoed back
=== RUN   TestEchoServer/Test_Put_Valid_Json
2021/05/31 16:42:23 200 [] /api/echo OK: echoed back
=== RUN   TestEchoServer/Test_Get_Json
2021/05/31 16:42:23 405 [] /api/echo Error: only POST/PUT are allowed
=== RUN   TestEchoServer/Test_Post_Valid_Json_but_with_echoed_false
2021/05/31 16:42:23 200 [] /api/echo OK: echoed back
=== RUN   TestEchoServer/Test_Post_Valid_Json_but_with_echoed_true
2021/05/31 16:42:23 400 [] /api/echo Error: echoed cannot be set to true
=== RUN   TestEchoServer/Test_Post_Valid_Json_but_with_erroneous_content-type
2021/05/31 16:42:23 400 [] /api/echo Error: Content-Type must be set to application/json
--- PASS: TestEchoServer (0.00s)
    --- PASS: TestEchoServer/Test_Post_Valid_Json (0.00s)
    --- PASS: TestEchoServer/Test_Put_Valid_Json (0.00s)
    --- PASS: TestEchoServer/Test_Get_Json (0.00s)
    --- PASS: TestEchoServer/Test_Post_Valid_Json_but_with_echoed_false (0.00s)
    --- PASS: TestEchoServer/Test_Post_Valid_Json_but_with_echoed_true (0.00s)
    --- PASS: TestEchoServer/Test_Post_Valid_Json_but_with_erroneous_content-type (0.00s)
PASS
ok  	_/home/costasd/dev/echo/cmd/server	0.002s
go build -v -o server ./cmd/server
_/home/costasd/dev/echo/cmd/server

```

### Locally with Docker
To build a container, execute the relevant `make` target from project's root dir:
```
$ make docker-build
```

## Usage
After the server is built, it can be run either directly
```
./server
```

or via docker:
```
docker run -p 3000:3000 -it echo:latest
```

### TLS
Optionally, TLS can be configured. In that case, a certificate pair must be supplied.\\
A helper exists in a make target to generate a self-signed pair under `certs/`:
```
$ make gen-certs
openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 \
	-subj "/CN=127.0.0.1" \
	-keyout ./certs/private.key -out ./certs/public.crt
Generating a RSA private key
...............................................................................................................................++++
...................++++
writing new private key to './certs/private.key'
-----
```

Then, the server can be configured to listen for TLS connections on 3001, either directly,
```
./server -cert certs/public.crt -key certs/private.key
```

or via docker:
```
docker run -p 3000:3000 -p 3001:3001 -v ./certs:/certs -it echo:latest /go/bin/server -cert certs/public.crt -key certs/private.key
```

