BUILDDIR := release

.PHONY: release
release:
	@echo "Building caddy"
	@go build -o $(BUILDDIR)/caddy src/cmd/caddy.go
	@echo "Building caddy-shim"
	@go build -o $(BUILDDIR)/caddy-shim src/shim/caddy-shim.go

.PHONY: release-arm
release-arm:
	@echo "Building caddy"
	@GOOS=darwin GOARCH=arm64 go build -o $(BUILDDIR)/caddy src/cmd/caddy.go
	@echo "Building caddy-shim"
	@GOOS=darwin GOARCH=arm64 go build -o $(BUILDDIR)/caddy-shim src/shim/caddy-shim.go