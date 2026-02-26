.PHONY: build proto
build:
	go build -o blackvault .

# Generate gRPC Go code (requires protoc, protoc-gen-go, protoc-gen-go-grpc)
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/proto/blackvault.proto
