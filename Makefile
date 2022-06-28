PROG=protoc-gen-demo

.PHONY: build

build: $(PROG)

$(PROG): main.go
	go build -o protoc-gen-demo main.go

.PHONY: proto

proto: demo/demo.proto
	protoc --go_out=demo demo/demo.proto

demo: demo/demo.proto $(PROG)
	protoc --demo_out=demo --plugin protoc-gen-demo=protoc-gen-demo demo/demo.proto
