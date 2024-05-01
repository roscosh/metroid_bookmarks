FROM golang:1.22.1

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod .
RUN  go mod download

COPY . .
