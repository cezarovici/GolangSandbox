package main

import (
	"html/template"
	"io"
	"os"
)

type User struct {
	Name string
}

func randerTo(w io.Writer, templateFilePath string, model any) error {
	t, errParse := template.ParseGlob(templateFilePath)
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, model)
}

func writeTo(filePath, from string, model any) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	return randerTo(f, from, model)
}

func main() {
	user := User{
		Name: "Cezar",
	}

	writeTo("newTemplate", "template.tmpl", user)
}
