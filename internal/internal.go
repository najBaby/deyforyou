package internal

import (
	"os"
	"text/template"
)

type File struct {
	Name    string
	Content string
}

type Assets struct {
	Dir   string
	Files []File
}

func CreateDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func CreateFile(name, text string, data interface{}) error {
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
	return template.New(name).Funcs(Funcs)
}
