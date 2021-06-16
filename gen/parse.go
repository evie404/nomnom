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

func ListEnumsTypesValues(decls []ast.Decl) ([]*StringEnum, []*IntEnum) {
	stringTypes, intTypes := listEnumTypes(decls)
	stringVals, intVals := listEnumValues(decls)

	for i := range stringTypes {
		if vals, found := stringVals[stringTypes[i].Name]; found {
			stringTypes[i].Values = vals
		}
	}

	for i := range intTypes {
		if vals, found := intVals[intTypes[i].Name]; found {
			intTypes[i].Values = vals
		}
	}

	return stringTypes, intTypes
}

func listEnumTypes(decls []ast.Decl) ([]*StringEnum, []*IntEnum) {
	var stringEnums []*StringEnum
	var intEnums []*IntEnum

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
			case "string":
				stringEnums = append(stringEnums, &StringEnum{
					Name: typeSpec.Name.Name,
				})
			case "int":
				intEnums = append(intEnums, &IntEnum{
					Name: typeSpec.Name.Name,
				})
			}
		}
	}

	return stringEnums, intEnums
}

func listEnumValues(decls []ast.Decl) (map[string][]StringEnumValue, map[string][]IntEnumValue) {
	stringEnumValues := map[string][]StringEnumValue{}
	intEnumValues := map[string][]IntEnumValue{}

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

				if lit.Kind == token.STRING {
					if _, found := stringEnumValues[typeName]; !found {
						stringEnumValues[typeName] = nil
					}

					stringEnumValues[typeName] = append(stringEnumValues[typeName], StringEnumValue{
						Name:  name.Name,
						Value: lit.Value,
					})
				}

				// if lit.Kind == token.INT {
				// 	if _, found := intEnumValues[typeName]; !found {
				// 		intEnumValues[typeName] = nil
				// 	}

				// 	intEnumValues[typeName] = append(intEnumValues[typeName], IntEnumValue{
				// 		Name:  name.Name,
				// 		Value: lit.Value,
				// 	})
				// }

				// fmt.Printf("\tType name: %+v\n", ident.Name)
				// fmt.Printf("\tConst name: %+v\n", name)
				// // TODO: check all values of right type
				// fmt.Printf("\tValue: %+v (%v)\n", lit.Value, lit.Kind)
			}
		}
	}

	return stringEnumValues, intEnumValues
}
