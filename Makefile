# Building/Running

.PHONY: run
run:
	go run ./cmd/tm_api -verbose -demo

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o bin/tm_api ./cmd/tm_api
	docker build --tag tm_api:latest .
	docker image save tm_api -o bin/tm_api.image.tar
	docker image rm tm_api -f

# Testing

.PHONY: test
test:
	go test -run=^Test -timeout 5m -cover ./...

.PHONY: bench
bench:
	go test -run=^$$ -cover -bench=. ./...

.PHONY: fuzz
fuzz:
	go test -run=^$$ -race -cover -fuzztime 5m -fuzz "FuzzGameString" ./src/turingmachine/game

.PHONY: security
security:
	@echo -e "\n> Checking for common security mistakes..."
	@gosec -track-suppressions -exclude=G302,G304,G404,G115 ./...
	@echo -e "\n> Checking imports for known vulnerabilities..."
	@govulncheck ./...

# Helpers

.PHONY: setup
setup:
	go mod tidy

.PHONY: clean
clean:
	go mod tidy
	rm -f bin/*
	rm -f 
	go clean
	go fmt ./...
	go vet ./...
