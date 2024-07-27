build:
	@ go build -o bin/stand cmd/main.go

run: build
	@ ./bin/stand

test:
	go test ./... -v 

.PHONY: build run test
