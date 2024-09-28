#!/bin/bash

# wsl linter fix
go test -race -coverprofile=coverage.out ./...

# gofumpt linter fix
go tool cover -html=coverage.out -o test/coverage.html

# swag documentation generation
rm coverage.out