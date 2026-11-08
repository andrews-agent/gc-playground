.PHONY: test check

test:
	go test ./...

check: test
	./scripts/lint.sh
