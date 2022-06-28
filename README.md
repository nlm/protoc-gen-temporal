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
----- BEGIN PLUGIN -----
- demo/demo.proto
----- BEGIN FILE demo/demo.proto -----
Message: Demo
- Field: pipo string
Enum MyEnum
- Value: zero 0
- Value: one 1
- Value: two 2
----- END FILE demo/demo.proto -----
----- END PLUGIN -----
```