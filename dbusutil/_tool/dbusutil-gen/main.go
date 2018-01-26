package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"

	"pkg.deepin.io/lib/strv"
)

func astNodeToStr(fset *token.FileSet, node interface{}) (string, error) {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, node)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

type Property struct {
	Name  string
	Type  string
	Equal string
}

type Generator struct {
	buf bytes.Buffer
	pkg Package
}

func (g *Generator) printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(&g.buf, format, args...)
}

const propsMuField = "PropsMu"

func (g *Generator) generate() {
	g.printf("package %s\n", g.pkg.name)

	g.printf("import (\n")

	for _, imp := range g.pkg.extraImports {
		g.printf(imp + "\n")
	}
	g.printf("\"pkg.deepin.io/lib/dbusutil\"\n")

	// end import
	g.printf(")\n\n")

	for typ, props := range g.pkg.typePropMap {
		for _, prop := range props {

			// set property method
			returnVarType := ""
			if prop.Equal != "nil" {
				returnVarType = "(changed bool)"
			}

			g.printf("func (v *%s) setProp%s(service *dbusutil.Service, value %s) %s {\n",
				typ, prop.Name, prop.Type, returnVarType)

			g.printf("v.%s.Lock()\n", propsMuField)

			switch prop.Equal {
			case "nil":
				g.printf("v.%s = value\n", prop.Name)
			case "":
				g.printf("if v.%s != value {\n", prop.Name)
				g.printf("v.%s = value\n", prop.Name)
				g.printf("changed = true\n")
				g.printf("}\n")
			default:
				expr := fmt.Sprintf("%s(v.%s, value)", prop.Equal, prop.Name)
				if strings.HasPrefix(prop.Equal, "method:") {
					method := prop.Equal[len("method:"):]
					expr = fmt.Sprintf("v.%s.%s(value)", prop.Name, method)
				}

				g.printf("if !%s {\n", expr)
				g.printf("v.%s = value\n", prop.Name)
				g.printf("changed = true\n")
				g.printf("}\n")
			}

			g.printf("v.%s.Unlock()\n", propsMuField)

			if prop.Equal != "nil" {
				g.printf("if service != nil && changed {\n")
				g.printf("    service.EmitPropertyChanged(v, \"%s\", value)\n", prop.Name)
				g.printf("}\n")
				g.printf("return\n")
			} else {
				g.printf("if service != nil {\n")
				g.printf("    service.NotifyChange(v, \"%s\", value)\n", prop.Name)
				g.printf("}\n")
			}
			g.printf("}\n\n")

			// get property method
			g.printf("func (v *%s) getProp%s() %s {\n", typ, prop.Name, prop.Type)
			g.printf("    v.%s.RLock()\n", propsMuField)
			g.printf("    value := v.%s\n", prop.Name)
			g.printf("    v.%s.RUnlock()\n", propsMuField)
			g.printf("    return value\n")
			g.printf("}\n\n")
		}
	}
}

type Package struct {
	name        string
	typePropMap map[string][]Property
	//              ^type name
	extraImports []string
}

func (g *Generator) parseFiles(names []string, types strv.Strv) {
	log.Printf("parseFiles names: %v, types: %v\n", names, types)
	fs := token.NewFileSet()
	var typePropMap = make(map[string][]Property)

	for _, name := range names {
		if !strings.HasSuffix(name, ".go") {
			//not go file
			continue
		}

		f, err := parser.ParseFile(fs, name, nil, parser.ParseComments)
		if err != nil {
			log.Fatalf("failed to parse file %q: %s", name, err)
		}
		ast.Inspect(f, func(node ast.Node) bool {
			decl, ok := node.(*ast.GenDecl)

			if !ok || decl.Tok != token.TYPE {
				return true
			}

			for _, spec := range decl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				if !types.Contains(typeSpec.Name.Name) {
					continue
				}

				props := getProps(fs, structType)
				typePropMap[typeSpec.Name.Name] = props
			}
			return true
		})

	}

	g.pkg.typePropMap = typePropMap

}

func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		log.Println("warning: internal error: invalid Go generated:", err)
		return g.buf.Bytes()
	}
	return src
}

var (
	typeNames    string
	extraImports string
	outputFile   string
)

func init() {
	flag.StringVar(&typeNames, "type", "", "comma-separated list of type names; must be set")
	flag.StringVar(&extraImports, "import", "", "")
	flag.StringVar(&outputFile, "output", "", "output file")
}

func main() {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("dbusutil-gen: ")

	flag.Parse()

	types := strv.Strv(strings.Split(typeNames, ","))
	//goFile := os.Getenv("GOFILE")
	goPackage := os.Getenv("GOPACKAGE")

	files := flag.Args()

	var parsedExtraImports []string
	for _, imp := range strings.Split(extraImports, ",") {
		if imp == "" {
			continue
		}
		if strings.Contains(imp, "=") {
			parts := strings.SplitN(imp, "=", 2)
			pkg := parts[0]
			alias := parts[1]
			// pkg.deepin.io/lib/dbus1=dbus,bytes
			parsedExtraImports = append(parsedExtraImports, fmt.Sprintf("%s \"%s\"",
				alias, pkg))
		} else {
			parsedExtraImports = append(parsedExtraImports, `"`+imp+`"`)
		}
	}

	g := &Generator{
		pkg: Package{
			name:         goPackage,
			extraImports: parsedExtraImports,
		},
	}

	g.parseFiles(files, types)
	g.generate()

	code := g.format()

	if outputFile == "" {
		outputFile = g.pkg.name + "_dbusutil.go"
	}
	log.Println("output file:", outputFile)
	err := ioutil.WriteFile(outputFile, code, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func isExportField(fieldName string) bool {
	return unicode.IsUpper(rune(fieldName[0]))
}

func getProps(fs *token.FileSet, structType *ast.StructType) []Property {
	//ast.Print(fs, structType)

	var prevField *ast.Field
	var props []Property
	for _, field := range structType.Fields.List {
		if len(field.Names) != 1 {
			prevField = field
			continue
		}

		fieldName := field.Names[0].Name
		if !isExportField(fieldName) {
			prevField = field
			continue
		}
		if fieldName == propsMuField {
			prevField = field
			continue
		}

		var equal string
		if field.Doc != nil {
			comments := field.Doc.List
			if len(comments) > 0 {
				option := getOptionFromComment(comments[0].Text)
				if option != "" {
					log.Printf("field %s option %s", fieldName, option)
				}

				if option == "ignore" {
					prevField = field
					continue

				} else if option == "ignore-below" {
					break
				}

				equal = getEqualFunc(option)
			}
		}

		fieldType, err := astNodeToStr(fs, field.Type)
		if err != nil {
			log.Fatal(err)
		}

		if fieldType == "sync.RWMutex" && len(prevField.Names) == 1 {
			prevFieldName := prevField.Names[0].Name
			if prevFieldName+"Mu" == fieldName {
				// ignore this field and prev field
				props = props[:len(props)-1]
				prevField = field
				continue
			}
		}

		props = append(props, Property{
			Name:  fieldName,
			Type:  fieldType,
			Equal: equal,
		})

		prevField = field
	}
	return props
}

func getEqualFunc(option string) string {
	idx := strings.Index(option, "equal=")
	if idx != -1 {
		equal := option[idx+len("equal="):]
		return strings.TrimSpace(equal)
	}

	return ""
}

func getOptionFromComment(comment string) string {
	comment = strings.TrimPrefix(comment, "//")
	comment = strings.TrimSpace(comment)

	const prefix = "dbusutil-gen:"
	if !strings.HasPrefix(comment, prefix) {
		return ""
	}
	option := comment[len(prefix):]
	return strings.TrimSpace(option)
}
