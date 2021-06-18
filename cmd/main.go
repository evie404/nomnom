package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rickypai/nomnom/gen"
)

var (
	write = flag.Bool("w", false, "write result to (source) file instead of stdout")
)

func main() {
	flag.Parse()
	filePaths := flag.Args()

	if len(filePaths) == 0 {
		println("usage: nomnom [path to go file]")
		return
	}

	err := generateFiles(filePaths)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func generateFiles(filePaths []string) error {
	var allOutputs []byte

	for _, filePath := range filePaths {
		fset := token.NewFileSet()
		astFile, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("parsing file: %w", err)
		}
		pkgName := astFile.Name.Name

		enums := gen.ListEnumsTypesValues(astFile.Decls)

		out, err := gen.GenerateEnumHelpers(pkgName, enums)
		if err != nil {
			return fmt.Errorf("generating enum helpers: %w", err)
		}

		if *write {
			ext := filepath.Ext(filePath)
			newFilePath := strings.Replace(filePath, ext, ".enum"+ext, -1)

			// On Windows, we need to re-set the permissions from the file. See golang/go#38225.
			var perms os.FileMode
			if fi, err := os.Stat(filePath); err == nil {
				perms = fi.Mode() & os.ModePerm
			}

			err = ioutil.WriteFile(newFilePath, out, perms)
			if err != nil {
				return fmt.Errorf("writing file `%s`: %w", newFilePath, err)
			}
		} else {
			allOutputs = append(allOutputs, out...)
		}
	}

	if !*write {
		println(string(allOutputs))
	}

	return nil
}
