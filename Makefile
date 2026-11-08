.PHONY: test check

test:
	go test ./...

check: test
	go vet ./...
