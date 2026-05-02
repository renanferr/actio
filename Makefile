.PHONY: build test clean lint

BIN=bin/actio

build: $(BIN)

$(BIN):
	go build -o $(BIN) ./cmd/actio

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -rf bin coverage.out
