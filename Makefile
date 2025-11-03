
.PHONY: all build clean fmt lint test

all: generate build

build: fmt
	-$(MAKE) lint || true
	go build -v ./...

generate:
	go generate ./pkg/image/nrgba
	go generate ./pkg/image/rgba

clean:
	go clean

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test -v ./...
