FROM golang:1.18.3

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy