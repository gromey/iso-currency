#!/scripts/bash

go run ./internal/cmd
go vet ./...
go fmt ./...