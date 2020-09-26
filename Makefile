BINARY_NAME = todo-api

export GO111MODULE=on

run:
	go run main.go
build:
	go build -o ./dist/$(BINARY_NAME)
