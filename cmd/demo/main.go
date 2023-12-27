package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/nlm/protoc-gen-temporal/demopb"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func Client() {
	// Setup client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal("unable to create temporal client:", err)
	}
	defer c.Close()
	dc := demopb.NewDemoClient(c)

	// Start Workflow
	swo := client.StartWorkflowOptions{TaskQueue: *flagTaskQueue}
	wr, err := dc.HelloWorld(context.Background(), swo, &demopb.HelloWorldRequest{
		Name: "Alice",
	})
	if err != nil {
		log.Fatal("Unable to start the Workflow:", err)
	}

	// Retrieve result
	var result demopb.HelloWorldResponse
	err = wr.Get(context.Background(), &result)
	if err != nil {
		log.Fatal("unable to get result:", err)
	}
	log.Print("result:", result.Message)
}

// ------------------------------------------ //

type DemoWorker struct {
	demopb.UnimplementedDemoWorker
}

func (dw *DemoWorker) HelloWorld(ctx workflow.Context, req *demopb.HelloWorldRequest) (*demopb.HelloWorldResponse, error) {
	return &demopb.HelloWorldResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}, nil
}

func Worker() {
	// Setup client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal("unable to create temporal client:", err)
	}
	defer c.Close()

	// Create and register worker
	w := worker.New(c, *flagTaskQueue, worker.Options{})
	dw := &DemoWorker{}
	demopb.RegisterDemoWorker(w, dw)

	// Run worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal("error running worker:", err)
	}
}

// ------------------------------------------ //

var (
	flagMode      = flag.String("mode", "", "client or worker")
	flagTaskQueue = flag.String("taskqueue", "default", "task queue")
)

func main() {
	flag.Parse()
	switch *flagMode {
	case "worker":
		Worker()
	case "client":
		Client()
	default:
		log.Fatal("unknown mode: ", *flagMode)
	}
}
