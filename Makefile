.PHONY: run

run:
	@go run cmd/main.go

build:
	@go build -ldflags="-s -w" -o bin/skald cmd/main.go