.PHONY: build
build:
	go build -o bin/dartsly-match-service cmd/main.go

.PHONY: compile_arm
compile_arm:
	GOOS=linux GOARCH=arm go build -o bin/dartsly-match-service cmd/main.go

.PHONY: compile_amd64
compile_amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/dartsly-match-service cmd/main.go

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: run
run:
	go run cmd/main.go

