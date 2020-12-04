package flutter

import (
	"deyforyou/colorizer"
	"fmt"
)

var (
	flutter = []assets{
		{
			dir: ".",
			files: []file{
				{
					name:    "Makefile",
					content: makeFile,
				},
			},
		},
		{
			dir: "lib",
			files: []file{
				{
					name:    "main.dart",
					content: mainCode,
				},
			},
		},
		{
			dir: "lib/screens",
			files: []file{
				{
					name:    "home.dart",
					content: homeCode,
				},
			},
		},
		{
			dir: "lib/screens/core",
			files: []file{
				{
					name:    "provider.dart",
					content: providerCode,
				},
				{
					name:    "app.dart",
					content: appCode,
				},
				{
					name:    "storage.dart",
					content: storageCode,
				},
				{
					name: "config.dart",
					content: configCode,
				},
			},
		},
		{
			dir: "lib/screens/core/data",
			files: []file{
				{
					name:    "example.dart",
					content: exampleCode,
				},
			},
		},
	}
)

// Flutter is
func Flutter() {
	for _, assets := range flutter {
		err := createDir(assets.dir)
		if err != nil {
			panic(err)
		}
		fmt.Println(colorizer.Yellow(fmt.Sprintf("Creating directory '%s' (success)", assets.dir)))
		for _, f := range assets.files {
			name := fmt.Sprintf("%s/%s", assets.dir, f.name)
			err := createFile(name, f.content, nil)
			if err != nil {
				panic(err)
			}
			fmt.Println(colorizer.Teal(fmt.Sprintf("Creating file '%s' (success)", name)))
		}
	}
	fmt.Println(colorizer.White(fmt.Sprintf("\nNow, you will find the models in '%s' directory.\n", colorizer.Purple("data"))))
	fmt.Println(colorizer.Magenta(fmt.Sprintf("Run 'make hive' to generate hive type for these models.\n")))
}
