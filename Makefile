run:
	go run main.go

build:
	go build -o cli main.go

test:
	go test ./...

lint:
	golangci-lint run

hello:
	go run main.go hello

assistant:
	go run main.go ask assistant "What can you do?"

dev:
	go run main.go ask dev "Write a Go program that reads a file"

tutor:
	go run main.go ask tutor "Explain Docker to a beginner"

