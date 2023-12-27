# protoc-gen-temporal

Using protobuf and gRPC services to generate Temporal workflows

## how to use

Build the generator

```bash
$ make build
go build ./cmd/protoc-gen-temporal/
```

## demo

Build proto

```bash
$ make proto
protoc -I demopb --temporal_out=demopb --temporal_opt=paths=source_relative --plugin protoc-gen-temporal=protoc-gen-temporal demo.proto
```

Build the demo app

```bash
$ make demo
go build -o demo ./cmd/demo
```

Start a temporal dev server

```bash
$ temporal server start-dev
```

Run the worker

```bash
$ ./demo -mode worker
```

Run the client

```bash
$ ./demo -mode client
result: Hello, Alice!
```
