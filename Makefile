.PHONY: build
build:
	go build -o bin/dartsly-match-service cmd/main.go

.PHONY: install
install:
	go mod download

.PHONY: compile_arm
compile_arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o bin/dartsly-match-service cmd/main.go

.PHONY: compile_amd64
compile_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/dartsly-match-service cmd/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run:
	go run cmd/main.go

