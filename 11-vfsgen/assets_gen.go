//+build ignore

package main

import (
	"log"

	sourcepkg "11-vfsgen/assets"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(sourcepkg.ImageAssets, vfsgen.Options{
		PackageName:     "example",
		BuildTags:       "!dev",
		VariableName:    "ImageAssets",
		VariableComment: "Assets contains project assets.",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
