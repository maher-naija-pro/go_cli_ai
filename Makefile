build:
	go build -o bin/myapp ./cmd/myapp

run:
	go run ./cmd/myapp

test:
	go test ./...

fmt:
	go fmt ./...

