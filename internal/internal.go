package internal

import (
	"deyforyou/templates"

	"os"
	"text/template"
)

type file struct {
	name    string
	content string
}

type assets struct {
	dir   string
	files []file
}

func createDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func createFile(name, text string, data interface{}) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	t := template.Must(newTemplate(name).Parse(text))
	err = t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}

func newTemplate(name string) *template.Template {
	return template.New(name).Funcs(templates.Funcs)
}
