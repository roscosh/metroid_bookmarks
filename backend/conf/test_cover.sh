#!/bin/bash

# Run all tests with the following flags:
# -race: detect race conditions
# -coverprofile=coverage.out: save test coverage data to coverage.out
go test -race -coverprofile=coverage.out ./cmd/... ./internal/... ./pkg/...

# Use the go tool cover to generate an HTML coverage report:
# -html=coverage.out: input coverage data
# -o ./coverage.html: output HTML report to coverage.html
go tool cover -html=coverage.out -o ./coverage.html

# Remove the temporary coverage data file
rm coverage.out
