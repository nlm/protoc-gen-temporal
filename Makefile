PROG=protoc-gen-temporal

.PHONY: build

build: $(PROG)

$(PROG): ./cmd/$(PROG)/*.go ./cmd/$(PROG)/*.tpl
	go build ./cmd/$(PROG)/

.PHONY: proto

proto: demopb/demo.pb.go demopb/demo.temporal.go

demopb/demo.pb.go: demopb/demo.proto
	protoc -I demopb --go_out=demopb --go_opt=paths=source_relative demo.proto

demopb/demo.temporal.go: demopb/demo.proto $(PROG)
	protoc -I demopb --temporal_out=demopb --temporal_opt=paths=source_relative --plugin $(PROG)=$(PROG) demo.proto

demo: ./cmd/demo/*.go
	go build -o demo ./cmd/demo

.PHONY: clean

clean:
	rm -f $(PROG) demopb/demo.temporal.go demopb/demo.pb.go
