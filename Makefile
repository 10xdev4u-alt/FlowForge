BINARY_NAME=flowforge
MAIN_PATH=main.go

.PHONY: build run test clean

build:
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

test:
	go test -v ./...

clean:
	go clean
	rm -rf bin/
