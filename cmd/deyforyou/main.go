package main

import (
	"deyforyou/cmd/flutter"
	"flag"
)

var (
	project string
)

func init() {
	flag.StringVar(&project, "project", "golang", "")
}

func init() {
	flag.Parse()
}

func main() {
	switch project {
	case "flutter":
		flutter.Flutter()
	}
}
