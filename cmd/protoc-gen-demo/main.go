package main

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/compiler/protogen"
)

func Generate(gen *protogen.Plugin) error {
	log.Print("----- START PLUGIN -----")
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		fileName := file.GeneratedFilenamePrefix + ".demo.go"
		genFile := gen.NewGeneratedFile(fileName, "")
		log.Print("----- BEGIN FILE ", file.Desc.Path(), " -----")
		genFile.P("package ", file.GoPackageName)
		genFile.P("import \"fmt\"")
		for _, message := range file.Messages {
			genFile.P("func (", message.GoIdent.GoName, ") HelloWorld() {")
			genFile.P("fmt.Println(\"Hello, ", message.GoIdent.GoName, "!\")")
			genFile.P("}")
		}
		log.Print("----- END FILE ", file.Desc.Path(), " -----")
	}
	log.Print("----- END PLUGIN -----")
	return nil
}

func main() {
	//log.SetOutput(os.Stderr)
	log.SetOutput(ioutil.Discard)
	protogen.Options{}.Run(Generate)
}
