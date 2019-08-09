test:
	go test ./...

lint:
	golangci-lint run # -v --enable-all
.PHONY: lint
