PROG=protoc-gen-demo

.PHONY: build

build: $(PROG)

$(PROG): ./cmd/$(PROG)/*.go
	go build ./cmd/$(PROG)/

.PHONY: proto

proto: demo/demo.pb.go demo/demo.demo.go

demo/demo.pb.go: demo/demo.proto
	protoc -I demo --go_out=demo --go_opt=paths=source_relative demo.proto

demo/demo.demo.go: demo/demo.proto $(PROG)
	protoc -I demo --demo_out=demo --demo_opt=paths=source_relative --plugin $(PROG)=$(PROG) demo.proto

.PHONY: demo

demo: proto
	go run ./cmd/demo

.PHONY: clean

clean:
	rm -f $(PROG) demo/demo.demo.go demo/demo.pb.go
