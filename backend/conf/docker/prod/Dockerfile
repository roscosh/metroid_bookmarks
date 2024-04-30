FROM golang:1.22.1-alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd

FROM scratch
WORKDIR /bin
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder main /bin/main
COPY --from=builder /build/conf /bin/conf
COPY --from=builder /build/schema /bin/schema
ENTRYPOINT ["/bin/main"]