FROM golang:1.15 as builder

RUN mkdir -p /myapp
ADD . /myapp
WORKDIR /myapp

RUN go test -v ./...

