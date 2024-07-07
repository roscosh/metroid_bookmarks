FROM golang:1.22.1

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go install github.com/air-verse/air@v1.52.3
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

COPY go.mod .
RUN  go mod download

COPY . .
