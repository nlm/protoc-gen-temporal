GO_PROG=protoc-gen-go-temporal
PY_PROG=protoc-gen-python-temporal
DEMOPB_DIR=demopb
DEMOPB_PROTO=$(DEMOPB_DIR)/demo.proto
INTERNAL_SOURCES=$(shell find internal)

.PHONY: build

build: $(GO_PROG) $(PY_PROG)

$(GO_PROG): $(wildcard ./cmd/$(GO_PROG)/*.go) $(INTERNAL_SOURCES)
	go build -o $@ ./cmd/$(GO_PROG)/

$(PY_PROG): $(wildcard ./cmd/$(PY_PROG)/*.go) $(INTERNAL_SOURCES)
	go build -o $@ ./cmd/$(PY_PROG)/

.PHONY: proto

proto: $(DEMOPB_DIR)/demo.pb.go $(DEMOPB_DIR)/demo_temporal.pb.go $(DEMOPB_DIR)/demo_pb2.py $(DEMOPB_DIR)/demo_pb2_temporal.py

$(DEMOPB_DIR)/demo.pb.go: $(DEMOPB_PROTO)
	protoc -I $(DEMOPB_DIR) --go_out=$(DEMOPB_DIR) --go_opt=paths=source_relative demo.proto

# $(DEMOPB_DIR)/demo_grpc.pb.go: $(DEMOPB_PROTO)
# 	protoc -I $(DEMOPB_DIR) --go-grpc_out=$(DEMOPB_DIR) --go-grpc_opt=paths=source_relative demo.proto
$(DEMOPB_DIR)/demo_pb2.py: $(DEMOPB_PROTO)
	protoc -I $(DEMOPB_DIR) --python_out=$(DEMOPB_DIR) --plugin protoc-gen-grpc-python=$(shell which grpc_python_plugin) demo.proto

$(DEMOPB_DIR)/demo_pb2_temporal.py: $(DEMOPB_PROTO) $(PY_PROG)
	protoc -I $(DEMOPB_DIR) --python-temporal_out=$(DEMOPB_DIR) --python-temporal_opt=paths=source_relative --plugin $(PY_PROG)=$(PY_PROG) demo.proto

$(DEMOPB_DIR)/demo_temporal.pb.go: $(DEMOPB_PROTO) $(GO_PROG)
	protoc -I $(DEMOPB_DIR) --go-temporal_out=$(DEMOPB_DIR) --go-temporal_opt=paths=source_relative --plugin $(GO_PROG)=$(GO_PROG) demo.proto

demo: ./cmd/demo/*.go
	go build -o demo ./cmd/demo

.PHONY: clean

clean:
	rm -f $(GO_PROG) $(DEMOPB_DIR)/demo_temporal.pb.go $(DEMOPB_DIR)/demo.pb.go
