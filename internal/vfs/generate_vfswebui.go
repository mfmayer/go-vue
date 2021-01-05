// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = http.Dir("assets/webui")
	err := vfsgen.Generate(fs, vfsgen.Options{
		VariableName: "FileSystem",
		PackageName:  "vfswebui",
		Filename:     "internal/vfs/vfswebui/filesystem.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
