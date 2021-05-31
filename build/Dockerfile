#
FROM golang:1.15
MAINTAINER Costas Drogos <costas.drogos@gmail.com>

ENV GOPATH=/go
WORKDIR /workdir

RUN go get github.com/costasd/echo/cmd/server 

ENTRYPOINT /go/bin/server
