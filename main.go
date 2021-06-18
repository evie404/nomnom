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
	write   = flag.Bool("w", false, "write result to (source) file instead of stdout")
	genTest = flag.Bool("t", false, "generate test files")
)

func main() {
	flag.Parse()
	filePaths := flag.Args()

	if len(filePaths) == 0 {
		printUsage()
		os.Exit(2)
	}

	err := generateFiles(filePaths)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "usage: usage [flags] [path ...]\n")
	flag.PrintDefaults()
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
			err = writeFile(filePath, ".enum", out)
			if err != nil {
				return err
			}
		} else {
			allOutputs = append(allOutputs, out...)
		}

		if *genTest {
			testOut, err := gen.GenerateEnumHelpersTests(pkgName, enums)
			if err != nil {
				return fmt.Errorf("generating enum helpers tests: %w", err)
			}

			err = writeFile(filePath, ".enum_test", testOut)
			if err != nil {
				return err
			}
		}
	}

	if !*write {
		println(string(allOutputs))
	}

	return nil
}

func writeFile(originalFilename string, extPrefix string, content []byte) error {
	ext := filepath.Ext(originalFilename)
	newFilePath := strings.Replace(originalFilename, ext, extPrefix+ext, -1)

	// On Windows, we need to re-set the permissions from the file. See golang/go#38225.
	var perms os.FileMode
	if fi, err := os.Stat(originalFilename); err == nil {
		perms = fi.Mode() & os.ModePerm
	}

	err := ioutil.WriteFile(newFilePath, content, perms)
	if err != nil {
		return fmt.Errorf("writing file `%s`: %w", newFilePath, err)
	}

	return nil
}
