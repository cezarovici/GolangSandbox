package main

import (
	"log"
	"os"
	"text/template"
)

type File struct {
	FileName   string
	IsTestFile bool
	Package    string
}

func RenderToPath(templateFilePath, renderToPath string, model File) error {
	t, errParse := template.ParseFiles(templateFilePath)
	if errParse != nil {
		return errParse
	}

	file, errCreate := os.Create(renderToPath)
	if errCreate != nil {
		return errCreate
	}

	if errExec := t.Execute(file, model); errExec != nil {
		return errExec
	}

	defer file.Close()
	return nil
}

func main() {
	f := File{
		FileName:   "Foo",
		IsTestFile: true,
		Package:    "test",
	}

	log.Print(RenderToPath("templates/template", "renderFile", f))
}
