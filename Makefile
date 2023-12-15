build:
	@go build -o bin/app.exe


run:build
	@./bin/app

test:
	@go test -v ./...