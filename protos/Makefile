# Makefile for protobuf generation

# Directories
PROTO_ROOT := proto
PROTO_DIR := $(PROTO_ROOT)/sso
GEN_DIR := ./gen/go/

# Protobuf compiler
PROTOC := protoc
PROTOC_FLAGS := -I=$(PROTO_ROOT)

# Go options
GO_OUT := --go_out=$(GEN_DIR)
GO_OPT := --go_opt=paths=source_relative
GO_GRPC_OUT := --go-grpc_out=$(GEN_DIR)
GO_GRPC_OPT := --go-grpc_opt=paths=source_relative

# Protobuf files
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)
GO_GEN_FILES := $(patsubst $(PROTO_DIR)/%.proto,$(GEN_DIR)/%.pb.go,$(PROTO_FILES))

.PHONY: proto

proto: $(GO_GEN_FILES)

$(GEN_DIR)/%.pb.go: $(PROTO_DIR)/%.proto
	@mkdir -p $(GEN_DIR)
	$(PROTOC) $(PROTOC_FLAGS) $< $(GO_OUT) $(GO_OPT) $(GO_GRPC_OUT) $(GO_GRPC_OPT)