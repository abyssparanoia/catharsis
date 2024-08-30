FROM golang:1.23-alpine

WORKDIR /go/src/github.com/abyssparanoia/catharsis/
COPY . .

ENV GO111MODULE=off

RUN apk --no-cache --update upgrade \
    && apk add --no-cache git alpine-sdk \
    && go get github.com/tockins/realize

ENV GO111MODULE=on