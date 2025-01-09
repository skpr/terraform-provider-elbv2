#!/usr/bin/make -f

export CGO_ENABLED=0

define build_step
	GOOS=$(1) GOARCH=$(2) go build -o bin/terraform-provider-elbv2_$(1)-$(2) -ldflags='-extldflags "-static"' github.com/skpr/terraform-provider-elbv2/cmd/terraform-provider-elbv2
endef

# Builds the project.
build:
	$(call build_step,linux,amd64)
	$(call build_step,linux,arm64)

# Run all lint checking with exit codes for CI.
lint:
	revive -config revive.toml -set_exit_status ./cmd/... ./internal/...

# Run tests with coverage reporting.
test:
	go test -cover ./...

.PHONY: *
