build:
	@ go build -o bin/stand cmd/main.go

run: build
	@ ./bin/stand

test:
	go test ./... -v 
install: build
	sudo mv ./bin/stand /usr/bin/stand

.PHONY: build run test
