package main

import "path"
import "encoding/xml"
import "encoding/json"
import "fmt"
import "text/template"
import "os"
import "os/exec"
import "io"
import "flag"
import "bytes"

import "dlib/dbus"

func GenInterfaceCode(pkgName string, info dbus.InterfaceInfo, writer io.Writer, dest, ifc_name, exportName string) {
	fmt.Println("d:", dest, "i:", ifc_name, "e:", exportName)
	funcs := template.FuncMap{
		"PkgName":    func() string { return pkgName },
		"OBJ_NAME":   func() string { return "obj" },
		"TypeFor":    func(s string) string { return dbus.TypeFor(s) },
		"DestName":   func() string { return dest },
		"IfcName":    func() string { return ifc_name },
		"ExportName": func() string { return exportName },
		"GetParamterNames": func(args []dbus.ArgInfo) (ret string) {
			for _, arg := range args {
				if arg.Direction != "out" {
					ret += ", "
					ret += arg.Name
				}
			}
			return
		},
		"GetParamterOuts": func(args []dbus.ArgInfo) (ret string) {
			var notFirst = false
			for _, arg := range args {
				if arg.Direction != "in" {
					if notFirst {
						ret += ","
					}
					notFirst = true
					ret += "&" + arg.Name
				}
			}
			return
		},
		"GetParamterOutsProto": func(args []dbus.ArgInfo) (ret string) {
			var notFirst = false
			for _, arg := range args {
				if arg.Direction != "in" {
					if notFirst {
						ret += ","
					}
					notFirst = true
					ret += arg.Name + " " + dbus.TypeFor(arg.Type)
				}
			}
			return
		},
		"GetParamterInsProto": func(args []dbus.ArgInfo) (ret string) {
			var notFirst = false
			for _, arg := range args {
				if arg.Direction != "out" {
					if notFirst {
						ret += ","
					}
					notFirst = true
					ret += arg.Name + " " + dbus.TypeFor(arg.Type)
				}
			}
			return
		},
	}
	templ := template.Must(template.New(exportName).Funcs(funcs).Parse(__IFC_TEMPLATE))
	templ.Execute(writer, info)
}

func GetInterfaceInfo(reader io.Reader, ifc_name string) dbus.InterfaceInfo {
	decoder := xml.NewDecoder(reader)
	obj := dbus.NodeInfo{}
	decoder.Decode(&obj)
	for _, ifc := range obj.Interfaces {
		if ifc.Name == ifc_name {
			return ifc
		}
	}
	panic("No " + ifc_name + " interface")
}

type _Interface struct {
	GoFile, XMLFile, Dest, ObjectPath, Interface, ObjectName, TestPath string
}
type _Config struct {
	NotExportBus bool
	OutputDir    string
	InputDir     string
	PkgName      string
	DestName     string
	BusType      string
}

type Infos struct {
	Interfaces []_Interface
	Config     _Config
}

func loadInfo(path string) Infos {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(f)
	info := Infos{}
	err = dec.Decode(&info)
	if err != nil {
		panic(err)
	}
	return info
}

func parse_info() Infos {
	var outputPath, inputFile string
	flag.StringVar(&outputPath, "out", "out", "the file to save the generated code")
	flag.StringVar(&inputFile, "in", "dbus.in.json", "the config file path")
	infos := loadInfo(inputFile)
	if outputPath != "out" {
		infos.Config.OutputDir = outputPath
	} else if len(infos.Config.OutputDir) == 0 {
		infos.Config.OutputDir = outputPath
	}
	return infos
}

func main() {
	infos := parse_info()
	os.MkdirAll(infos.Config.OutputDir, 0755)
	writer, err := os.Create(path.Join(infos.Config.OutputDir, "init.go"))
	if err != nil {
		panic(err)
	}
	template.Must(template.New("main").Funcs(template.FuncMap{
		"GetBusType": func() string { return infos.Config.BusType },
		"PkgName":    func() string { return infos.Config.PkgName },
	}).Parse(__GLOBAL_TEMPLATE)).Execute(writer, nil)

	/*writer.WriteString(__GLOBAL_TEMPLATE)*/
	writer.Close()
	defer func() {
		exec.Command("gofmt", "-w", infos.Config.OutputDir).Start()
	}()
	for _, ifc := range infos.Interfaces {
		file := path.Join(infos.Config.InputDir, ifc.XMLFile)
		var reader io.Reader
		writer, err = os.Create(path.Join(infos.Config.OutputDir, ifc.GoFile))
		if _, err := os.Stat(file); err == nil {
			reader, err = os.Open(file)
			if err != nil {
				panic(err.Error() + "(File:" + file + ")")
			}
			info := GetInterfaceInfo(reader, ifc.Interface)
			GenInterfaceCode(infos.Config.PkgName, info, writer, infos.Config.DestName, ifc.Interface, ifc.ObjectName)
			if ifc.TestPath != "" {
				var test_writer io.Writer
				test_writer, err = os.Create(path.Join(infos.Config.OutputDir, path.Base(ifc.GoFile)+"_test.go"))
				genTest(ifc.TestPath, infos.Config.PkgName, ifc.ObjectName, test_writer, info)
			}
			reader.(*os.File).Close()
		} else {
			conn, _ := dbus.SystemBus()
			var xml string
			if err := conn.Object(ifc.Dest, dbus.ObjectPath(ifc.ObjectPath)).Call("org.freedesktop.DBus.Introspectable.Introspect", 0).Store(&xml); err != nil {
				panic(err.Error() + "Interface " + ifc.Interface + " is can't dynamic introspect")
			}
			GenInterfaceCode(infos.Config.PkgName, GetInterfaceInfo(bytes.NewBufferString(xml), ifc.Interface), writer, infos.Config.DestName, ifc.Interface, ifc.ObjectName)

		}
		writer.Close()
	}
}

var __TEST_TEMPLATE = `/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
package {{PkgName}}
import "testing"
{{range .Methods}}
func Test{{ObjName}}Method{{.Name}} (t *testing.T) {
	{{/*
	rnd := rand.New(rand.NewSource(99))
	r := Get{{ObjName}}("{{TestPath}}").{{.Name}}({{.Args}})
--*/}}

}
{{end}}

{{range .Properties}}
func Test{{ObjName}}Property{{.Name}} (t *testing.T) {
	t.Log("Get the property {{.Name}} of object {{ObjName}} ===> ",
		Get{{ObjName}}("{{TestPath}}").Get{{.Name}}())
}
{{end}}

{{range .Signals}}
func Test{{ObjName}}Signal{{.Name}} (t *testing.T) {
}
{{end}}
`

func genTest(testPath, pkgName string, objName string, writer io.Writer, info dbus.InterfaceInfo) {
	funcs := template.FuncMap{
		"TestPath": func() string { return testPath },
		"PkgName":  func() string { return pkgName },
		"ObjName":  func() string { return objName },
		/*"GetTestValue": func(args []dbus.ArgInfo) string {*/
		/*},*/
	}
	template.Must(template.New("testing").Funcs(funcs).Parse(__TEST_TEMPLATE)).Execute(writer, info)
}
