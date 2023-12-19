build:
	@go build -o bin/app.exe


run:build
	@./bin/app

seed:
	@go run ./scripts/seed.go

test:
	@go test -v ./...