.PHONY: build
build:
	go build ./cmd/md5http/main.go

.PHONY: lint
lint:
	go vet ./...

.PHONY: test
test:
	go test ./...
