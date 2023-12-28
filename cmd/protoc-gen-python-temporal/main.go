package main

import (
	"log"
	"os"

	"github.com/nlm/protoc-gen-temporal/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	log.SetOutput(os.Stderr)
	generate, err := gen.Generate(gen.SupportedLanguages[gen.LangPython])
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	protogen.Options{}.Run(generate)
}
