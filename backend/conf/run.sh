#!/bin/bash

# wsl linter fix
wsl --fix ./...

# gofumpt linter fix
golangci-lint run --disable-all -E gofumpt --fix

# swag documentation generation
swag init -g cmd/main.go

# build app
go build -o ./tmp/main ./cmd
