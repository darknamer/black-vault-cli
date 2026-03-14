.PHONY: build proto test test-cover test-cover-profile test-cover-html
build:
	go build -o blackvault .

# Generate gRPC Go code (requires protoc, protoc-gen-go, protoc-gen-go-grpc)
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/proto/blackvault.proto

# Run all tests
test:
	go test ./...

# Run tests with coverage summary (percent per package)
test-cover:
	go test -cover ./...

# Run tests and write coverage profile to coverage.out (for cmd + test packages)
test-cover-profile:
	go test -coverprofile=coverage.out -covermode=atomic ./cmd ./test

# Generate coverage profile and HTML report (open with: make test-cover-html && open coverage.html)
test-cover-html: test-cover-profile
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"
