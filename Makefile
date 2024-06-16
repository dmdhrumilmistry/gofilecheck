APP_NAME=gofilecheck

all: build test run

docker:
	@docker build -t $(APP_NAME) .

build:
	@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

run:
	@bin/$(APP_NAME)

test:
	@go test -v ./...

clean:
	@go mod tidy
	@go clean
	@rm -rf bin

dev:
	@air


