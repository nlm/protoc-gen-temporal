package gen

import (
	_ "embed"
	"html/template"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

//go:embed templates/go.tpl
var goTemplate []byte

//go:embed templates/python.tpl
var pyTemplate []byte

type TemplateInput struct {
	File    *protogen.File
	GenFile *protogen.GeneratedFile
}

const (
	LangGo = iota
	LangPython
)

var funcs = template.FuncMap{
	"PyPackageName": func(s string) string {
		return strings.Replace(s, ".proto", "_pb2", -1)
	},
	"PyImportName": func(s string) string {
		return strings.Replace(s, ".proto", "__pb2", -1)
	},
}

var SupportedLanguages = map[int]Params{
	LangGo: {
		Template:       template.Must(template.New("go.tpl").Funcs(funcs).Parse(string(goTemplate))),
		FilenameSuffix: "_temporal.pb.go",
	},
	LangPython: {
		Template:       template.Must(template.New("python.tpl").Funcs(funcs).Parse(string(pyTemplate))),
		FilenameSuffix: "_pb2_temporal.py",
	},
}

type Params struct {
	Template       *template.Template
	FilenameSuffix string
}

func Generate(params Params) (func(*protogen.Plugin) error, error) {
	return func(plu *protogen.Plugin) error {
		for _, file := range plu.Files {
			if !file.Generate {
				continue
			}
			fileName := file.GeneratedFilenamePrefix + params.FilenameSuffix
			genFile := plu.NewGeneratedFile(fileName, file.GoImportPath)
			if err := params.Template.Execute(genFile, TemplateInput{
				File:    file,
				GenFile: genFile,
			}); err != nil {
				return err
			}
		}
		return nil
	}, nil
}
