# protoc-gen-demo

a demo on how to make a minimal protobuf code generator

## how to use

Build the generator

```bash
$ make build
go build -o protoc-gen-demo main.go
```

Use it with protoc

```bash
$ make demo
protoc --demo_out=demo --plugin protoc-gen-demo=protoc-gen-demo demo/demo.proto
go run ./cmd/demo
Hello, Demo!
```
