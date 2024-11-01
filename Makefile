PROTO ?= buildsafe/v1/search.proto

lint: buf-produce-bin
	api-linter --set-exit-status --output-format yaml  --descriptor-set-in descriptorset.bin $(PROTO)
 

lint-summary: buf-produce-bin
	api-linter --set-exit-status --output-format summary --descriptor-set-in descriptorset.bin $(PROTO)

buf-produce-bin:
	buf build  --output descriptorset.bin --as-file-descriptor-set

.PHONY: buf-setup
buf-setup:
	@go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@go get google.golang.org/protobuf/cmd/protoc-gen-go
	@go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	@go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	@go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc
	
.PHONY: buf-generate
buf-generate:
	buf generate
