all: build test-cover run

build:
	go build ./calculate_tax.go

run:
	go run ./calculate_tax.go

clean:
	go mod tidy

test:
	go test ./...

test-alt:
	go test -v -count=1 ./... -covermode=count

test-cover:
	go test -count=1 ./... -covermode=count

test-out:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html