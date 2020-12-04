package flutter

import (
	"deyforyou/colorizer"
	"deyforyou/internal"
	"deyforyou/templates"
	"fmt"
)

var (
	flutter = []internal.Assets{
		{
			Dir: ".",
			Files: []internal.File{
				{
					Name:    "Makefile",
					Content: templates.MakeFile,
				},
			},
		},
		{
			Dir: "lib",
			Files: []internal.File{
				{
					Name:    "main.dart",
					Content: templates.MainCode,
				},
			},
		},
		{
			Dir: "lib/screens",
			Files: []internal.File{
				{
					Name:    "home.dart",
					Content: templates.HomeCode,
				},
			},
		},
		{
			Dir: "lib/screens/core",
			Files: []internal.File{
				{
					Name:    "provider.dart",
					Content: templates.ProviderCode,
				},
				{
					Name:    "app.dart",
					Content: templates.AppCode,
				},
				{
					Name:    "storage.dart",
					Content: templates.StorageCode,
				},
				{
					Name:    "config.dart",
					Content: templates.ConfigCode,
				},
			},
		},
		{
			Dir: "lib/screens/core/data",
			Files: []internal.File{
				{
					Name:    "example.dart",
					Content: templates.ExampleCode,
				},
			},
		},
	}
)

// Flutter is
func Flutter() {
	for _, assets := range flutter {
		err := internal.CreateDir(assets.Dir)
		if err != nil {
			panic(err)
		}
		fmt.Println(colorizer.Yellow(fmt.Sprintf("Creating directory '%s' (success)", assets.Dir)))
		for _, f := range assets.Files {
			name := fmt.Sprintf("%s/%s", assets.Dir, f.Name)
			err := internal.CreateFile(name, f.Content, nil)
			if err != nil {
				panic(err)
			}
			fmt.Println(colorizer.Teal(fmt.Sprintf("Creating file '%s' (success)", name)))
		}
	}
	fmt.Println(colorizer.White(fmt.Sprintf("\nNow, you will find the models in '%s' directory.\n", colorizer.Purple("data"))))
	fmt.Println(colorizer.Magenta(fmt.Sprintf("Run 'make hive' to generate hive type for these models.\n")))
}
