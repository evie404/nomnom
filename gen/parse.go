package gen

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func readFileAst(filePath string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, filePath, nil, parser.ParseComments)
}

func ListEnumsTypesValues(decls []ast.Decl) []Enum {
	enumTypes := listEnumTypes(decls)
	enumValues := listEnumValues(decls)

	for i := range enumTypes {
		if vals, found := enumValues[enumTypes[i].Name]; found {
			enumTypes[i].Values = vals
		}
	}

	return enumTypes
}

func listEnumTypes(decls []ast.Decl) []Enum {
	var enums []Enum

	for _, decl := range decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		if genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			ident, ok := typeSpec.Type.(*ast.Ident)
			if !ok {
				continue
			}

			// fmt.Printf("Enum type name: %+v\n", typeSpec.Name)
			// fmt.Printf("Enum type: %+v\n", ident.Name)

			// TODO: there's probably a better way to match these
			switch ident.Name {
			case "string", "int":
				enums = append(enums, Enum{
					Name:     typeSpec.Name.Name,
					BaseType: ident.Name,
				})
			}
		}
	}

	return enums
}

func listEnumValues(decls []ast.Decl) map[string][]EnumValue {
	enumValues := map[string][]EnumValue{}

	for _, decl := range decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		if genDecl.Tok != token.CONST {
			continue
		}

		for _, spec := range genDecl.Specs {
			valSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			ident, ok := valSpec.Type.(*ast.Ident)
			if !ok {
				continue
			}

			for i := range valSpec.Names {
				name := valSpec.Names[i]
				value := valSpec.Values[i]

				lit, ok := value.(*ast.BasicLit)
				if !ok {
					continue
				}

				typeName := ident.Name

				switch lit.Kind {
				case token.STRING, token.INT:
					if _, found := enumValues[typeName]; !found {
						enumValues[typeName] = nil
					}

					enumValues[typeName] = append(enumValues[typeName], EnumValue{
						Name:  name.Name,
						Value: lit.Value,
					})
				}
			}
		}
	}

	return enumValues
}
