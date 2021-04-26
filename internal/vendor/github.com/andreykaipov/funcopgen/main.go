package main

import (
	"flag"
	"fmt"
	"go/ast"
	"os"
	"sort"
	"strings"
	"unicode"

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

func firstRune(str string) (r rune) {
	for _, r = range str {
		return
	}
	return
}

func structFieldsToMap(s *ast.StructType) StructFieldMap {
	out := StructFieldMap{}

	for _, f := range s.Fields.List {
		jenType := findJenTypeOfField(f)

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

func findJenTypeOfField(field *ast.Field) *Statement {
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

var (
	fs           = flag.NewFlagSet("funcopgen", flag.ExitOnError)
	typeNames    = fs.String("type", "", "Comma-delimited list of type names")
	prefix       = fs.String("prefix", "", "Prefix to attach to functional options, e.g. WithColor, WithName, etc.")
	factory      = fs.Bool("factory", false, "If present, add a factory function for your type, e.g. NewAnimal(opt ...Option)")
	unexported   = fs.Bool("unexported", false, "If present, functional options are also generated for unexported fields.")
	uniqueOption = fs.Bool("unique-option", false,
		"If present, prepends the type to the Option type, e.g. AnimalOption.\n"+
			"Handy if generating for several structs within the same package.",
	)

	pkg *packages.Package
)

func init() {
	fs.Parse(os.Args[1:])

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage of %s:\n", fs.Name())
		fs.PrintDefaults()
	}

	if len(*typeNames) == 0 {
		fs.Usage()
		os.Exit(1)
	}

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedImports,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}

	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	// Since this tool is meant to be used from a go:generate comment, there
	// should only ever be one package.
	if len(pkgs) != 1 {
		fmt.Fprintf(os.Stderr, "expected only one package!")
		os.Exit(1)
	}

	pkg = pkgs[0]
}

func main() {
	types := map[string]interface{}{}
	for _, t := range strings.Split(*typeNames, ",") {
		types[t] = nil
	}

	// Find structs
	structs := map[string]StructFieldMap{}

	for _, file := range pkg.Syntax {
		for objName, obj := range file.Scope.Objects {
			if _, ok := types[objName]; !ok {
				continue
			}
			switch obj.Kind {
			case ast.Typ:
				switch typ := obj.Decl.(*ast.TypeSpec).Type.(type) {
				case *ast.StructType:
					structs[objName] = structFieldsToMap(typ)
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

	for t := range types {
		fields, ok := structs[t]
		if !ok {
			fmt.Fprintf(os.Stderr, "Unknown type %q in %q in package %q\n", t, pkg.Name, pkg.PkgPath)
			os.Exit(1)
		}

		// Sort the fields so we can traverse the map in a deterministic
		// order as we want the generated code to be the same between
		// subsequent runs.
		keys := make([]string, 0, len(fields))
		for k := range fields {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		f := NewFile(pkg.Name)

		f.HeaderComment("This file has been automatically generated. Don't edit it.")

		optionName := ""
		if *uniqueOption {
			optionName = t
		}
		optionName += "Option"

		f.Add(Type().Id(optionName).Func().Params(Op("*").Id(t)), Line())

		setDefaults := func(d Dict) {
			for _, field := range keys {
				tags := fields[field].Tags

				if tag, _ := tags.Get("default"); tag != nil {
					switch t := fields[field].Type; fmt.Sprintf("%#v", t) {
					case "string":
						d[Id(field)] = Lit(tag.Name)
					default:
						d[Id(field)] = Id(tag.Name)
					}
				}
			}
		}

		if *factory {
			f.Add(
				Func().Id("New"+t).Params(Id("opts").Op("...").Id(optionName)).Op("*").Id(t).Block(
					Id("o").Op(":=").Op("&").Id(t).Values(DictFunc(setDefaults)),
					Line(),
					For(Id("_, opt").Op(":=").Range().Id("opts")).Block(
						Id("opt").Call(Id("o")),
					),
					Line(),
					Return(Id("o")),
				),
				Line(),
			)
		}

		for _, field := range keys {
			typeName := fields[field].Type

			titledField := field
			if unicode.IsLower(firstRune(field)) {
				if !*unexported {
					continue
				}
				titledField = strings.Title(field)
			}

			f.Add(
				Func().Id(*prefix+titledField).Params(Id("x").Add(typeName)).Id(optionName).Block(
					Return(
						Func().Params(Id("o").Op("*").Id(t)).Block(
							Id("o").Dot(field).Op("=").Id("x"),
						),
					),
				),
				Line(),
			)
		}

		outFile := fmt.Sprintf("zz_generated.%s_funcop.go", strings.ToLower(t))

		if err := f.Save(outFile); err != nil {
			panic(err)
		}

		fmt.Printf("Generated functional options for `%s.%s`\n", pkg.Name, t)
	}
}
