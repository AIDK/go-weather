build:
	@go build -o bin/go-weather cmd/main.go

run: build
	@go run cmd/main.go