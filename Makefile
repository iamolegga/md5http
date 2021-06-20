.PHONY: build
build:
	go build -o=./md5http ./cmd/md5http/main.go

.PHONY: lint
lint:
	go vet ./...

.PHONY: test
test:
	go test ./...
