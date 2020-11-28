package ent

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
)

// API is
type API interface {
	Extension() string
}

// NewAPI is
func NewAPI(name string) API {
	switch name {
	case "grpc":
		return &Grpc{}
	default:
		return &Graphql{}
	}
}

var (
	file        = "schema.%s"
	schema, api string
)

func init() {
	flag.StringVar(&schema, "schema", "", "")
	flag.StringVar(&api, "api", "graphql", "")
}

func init() {
	flag.Parse()
}

func main() {

}

func generateAPI(schema *string) {
	if schema != nil {
		g, e := entc.LoadGraph(*schema, &gen.Config{
			Header: "// Don't touch this file",
			IDType: &field.TypeInfo{Type: field.TypeInt},
		})
		if e != nil {
			panic(e)
		}

		api := NewAPI(api)
		createFile(fmt.Sprintln(file, api.Extension()))
	}
}

func createFile(name string) (*os.File, string, error) {
	f, e := os.Create(name)
	if e != nil {
		panic(e)
	}
	return f, name, nil
}

func createTemplate(name, text string) {

	t := template.Must(template.New(name).Funcs(nil).Parse(text))
}
