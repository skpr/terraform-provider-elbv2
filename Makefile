#!/usr/bin/make -f

# Run all lint checking with exit codes for CI.
lint:
	revive -config revive.toml -set_exit_status ./cmd/... ./internal/...

# Run tests with coverage reporting.
test:
	go test -cover ./...

.PHONY: *
