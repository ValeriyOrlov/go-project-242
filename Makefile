.PHONY:
	build test lint clean install fmt
install:
	go mod tidy
build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
clean:
	rm -rf bin/
fmt:
	gofmt -w .
run:
	./bin/hexlet-path-size .
lint:
	golangci-lint run
test:
	go test -v ./tests