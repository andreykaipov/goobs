package main

import (
	"fmt"
	"go/ast"
	"os"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/fatih/structtag"
	"golang.org/x/tools/go/packages"
)

// StructFieldMap is a map of a struct's field names to data about the field.
type StructFieldMap map[string]*FieldData

type FieldData struct {
	// Tags is a map representing a field's tags, e.g. `default:"hello"`
	Tags *structtag.Tags

	// Type is the Jen representation of a type. To get the string
	// representation of it, e.g. something like "string" or
	// "[]*myQualified.StructType", we can fmt.Sprintf("%#v", blah) it
	Type *Statement
}

func structFieldsToMap(s *ast.StructType, pkg *packages.Package) StructFieldMap {
	out := StructFieldMap{}

	for _, f := range s.Fields.List {
		jenType := findJenTypeOfField(f, pkg)

		data := &FieldData{
			Type: jenType,
			Tags: findFieldTags(f),
		}

		// anonymous field, i.e. an embedded field
		// and if it's a qualified type, we drop the qualifier
		if f.Names == nil {
			typeName := fmt.Sprintf("%#v", jenType)
			split := strings.Split(typeName, ".")
			if len(split) == 1 {
				out[typeName] = data
			} else {
				out[split[1]] = data
			}
			continue
		}

		for _, n := range f.Names {
			out[n.Name] = data
		}
	}

	return out
}

// TODO maybe get rid of structtag and just use the ast, parsing from the
// field's Tag field directly. not sure why i went with structtag originally.
// maybe it was too messy? i can't remember lol
func findFieldTags(f *ast.Field) *structtag.Tags {
	if f.Tag == nil {
		return &structtag.Tags{}
	}

	// get rid of the backticks
	trimmed := f.Tag.Value[1 : len(f.Tag.Value)-1]

	tag, err := structtag.Parse(trimmed)
	if err != nil {
		panic(err)
	}

	return tag
}

func findJenTypeOfField(field *ast.Field, pkg *packages.Package) *Statement {
	var f func(e interface{}) *Statement

	f = func(e interface{}) (typeName *Statement) {
		switch typ := e.(type) {
		case *ast.InterfaceType:
			typeName = Interface()
		case *ast.Ident:
			typeName = Id(typ.Name)
		case *ast.StarExpr:
			typeName = Op("*").Add(f(typ.X))
		case *ast.MapType:
			typeName = Map(f(typ.Key)).Add(f(typ.Value))
		case *ast.ArrayType:
			typeName = Index().Add(f(typ.Elt))
		case *ast.ChanType:
			typeName = Chan().Add(f(typ.Value))
		case *ast.StructType:
			fmt.Fprintf(os.Stderr, "TODO: anon structs\n")
			os.Exit(2)
		case *ast.SelectorExpr:
			// Find the fully qualified path from the package imports
			qualifier := fmt.Sprintf("%#v", f(typ.X))
			for _, p := range pkg.Imports {
				if qualifier == p.Name {
					qualifier = p.PkgPath
					break
				}
			}
			typeName = Qual(qualifier, typ.Sel.Name)
		default:
			panic(fmt.Errorf("unhandled case for expression %#v", typ))
		}

		return typeName
	}

	return f(field.Type)
}

func findStructs(pkgPath string) (map[string]StructFieldMap, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedTypes,
		Dir:  pkgPath,
	})
	if err != nil {
		return nil, err
	}

	if len(pkgs) > 1 {
		return nil, fmt.Errorf("really should only be one package in %s, no?", pkgPath)
	}

	pkg := pkgs[0]

	fmt.Printf("%s\n", pkg.Types)

	structs := map[string]StructFieldMap{}

	for _, file := range pkg.Syntax {
		for objName, obj := range file.Scope.Objects {
			switch obj.Kind {
			case ast.Typ:
				switch typ := obj.Decl.(*ast.TypeSpec).Type.(type) {
				case *ast.StructType:
					structs[objName] = structFieldsToMap(typ, pkg)
				case *ast.Ident:
					continue
				default:
					continue
				}
			case ast.Con:
				continue
			default:
				continue
			}

		}
	}

	return structs, nil
}
