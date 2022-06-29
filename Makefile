PROG=protoc-gen-demo

.PHONY: build

build: $(PROG)

$(PROG): ./cmd/$(PROG)/*.go
	go build ./cmd/$(PROG)/

.PHONY: proto demoproto demo clean

proto: demo/demo.proto
	protoc --go_out=demo demo/demo.proto

demoproto: demo/demo.proto $(PROG)
	protoc --demo_out=demo --plugin $(PROG)=$(PROG) demo/demo.proto

demo: demoproto
	go run ./cmd/demo

clean:
	rm -f $(PROG) demo/demo.demo.go
