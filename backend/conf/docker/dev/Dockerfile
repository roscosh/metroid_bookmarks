FROM golang:1.22.1

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@v1.52.3
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1
RUN go install go.uber.org/mock/mockgen@v0.4
RUN go install github.com/bombsimon/wsl/v4/cmd/wsl@v4.4.1

COPY go.mod .
RUN  go mod download

COPY . .
