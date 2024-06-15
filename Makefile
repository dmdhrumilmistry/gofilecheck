APP_NAME=gocheckfile

build:
	@go build -o bin/gofilecheck cmd/gofilecheck/main.go

run:
	@go run cmd/$(APP_NAME)/main.go

test:
	@go test -v ./...

clean:
	@go mod tidy
	@go clean
	@rm -rf bin/$(APP_NAME)

.PHONY: build test run

