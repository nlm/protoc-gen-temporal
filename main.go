package main

import (
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
)

func Write(writer io.Writer, data ...any) (int, error) {
	return writer.Write([]byte(fmt.Sprintln(data...)))
}

func Generate(gen *protogen.Plugin) error {
	fmt.Fprintln(os.Stderr, "----- BEGIN PLUGIN -----")

	for _, fileName := range gen.Request.FileToGenerate {
		fmt.Fprintln(os.Stderr, "-", fileName)
	}

	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		fileName := file.GeneratedFilenamePrefix + ".gen.txt"
		fileWriter := gen.NewGeneratedFile(fileName, "")
		writer := io.MultiWriter(os.Stderr, fileWriter)
		Write(writer, "----- BEGIN FILE", file.Desc.Path(), "-----")
		for _, message := range file.Messages {
			Write(writer, "Message:", message.Desc.Name())
			for _, field := range message.Fields {
				Write(writer, "- Field:", field.Desc.Name(), field.Desc.Kind())
			}
		}
		for _, enum := range file.Enums {
			Write(writer, "Enum", enum.Desc.Name())
			for _, value := range enum.Values {
				Write(writer, "- Value:", value.Desc.Name(), value.Desc.Number())
			}
		}
		Write(writer, "----- END FILE", file.Desc.Path(), "-----")
	}
	fmt.Fprintln(os.Stderr, "----- END PLUGIN -----")
	return nil
}

func main() {
	protogen.Options{}.Run(Generate)
}
