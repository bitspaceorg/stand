ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

build:
	@ go build -o bin/stand cmd/main.go

run: build
	@ ./bin/stand $(RUN_ARGS)

test:
	go test ./... -v 

.PHONY: build run test
