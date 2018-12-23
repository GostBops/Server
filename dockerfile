FROM golang:latest

MAINTAINER Razil "GostBop"

RUN go get github.com/go-sql-driver/mysql && go get github.com/dgrijalva/jwt-go && go get github.com/gorilla/mux

WORKDIR $GOPATH/src/github.com/GostBops/Server
ADD . $GOPATH/src/github.com/GostBops/Server
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./Server"]]