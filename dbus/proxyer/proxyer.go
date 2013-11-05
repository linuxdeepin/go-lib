package main

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

var __GLOBAL_TEMPLATE = `
package main
import "dlib/dbus"
var __conn *dbus.Conn = nil
func GetSessionBus() *dbus.Conn {
	if __conn  == nil {
		var err error
		__conn, err = dbus.SessionBus()
		if err != nil {
			panic(err)
		}
	}
	return __conn
}
`

var __IFC_TEMPLATE = `
var __obj{{ExportName}} *dbus.Object = nil
type {{ExportName}} struct {
	core *dbus.Object
}
{{$obj_name := .Name}}
{{range .Methods }}
func (__obj__ {{ExportName }}) {{.Name}} ({{GetParamterInsProto .Args}}) ({{GetParamterOutsProto .Args}}) {
	__obj__.core.Call("{{$obj_name}}.{{.Name}}", 0{{GetParamterNames .Args}}).Store({{GetParamterOuts .Args}})
	return
}
{{end}}

func Get{{ExportName}}() {{ExportName}} {
	if __obj{{ExportName}} == nil {
		return {{ExportName}}{GetSessionBus().Object("{{DestName}}", "{{PathName}}")}
	} else {
		return {{ExportName}}{__obj{{ExportName}}}
	}
}

`

var __ARG_PREFIX = "__arg_"

func GenInterfaceCode(reader io.Reader, writer io.Writer, dest, path, ifc_name, exportName string) {
	funcs := template.FuncMap{
		"PathName":   func() string { return path },
		"DestName":   func() string { return dest },
		"ExportName": func() string { return exportName },
		"GetParamterNames": func(args []dbus.ArgInfo) (ret string) {
			for _, arg := range args {
				if arg.Direction != "out" {
					ret += ", "
					ret += __ARG_PREFIX + arg.Name
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
					ret += "&" + __ARG_PREFIX + arg.Name
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
					ret += __ARG_PREFIX + arg.Name + " " + dbus.TypeFor(arg.Type)
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
					ret += __ARG_PREFIX + arg.Name + " " + dbus.TypeFor(arg.Type)
				}
			}
			return
		},
	}
	templ := template.Must(template.New(exportName).Funcs(funcs).Parse(__IFC_TEMPLATE))
	info := GetInterfaceInfo(reader, ifc_name)
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
	Type, File, Dest, Path, Name, ObjName string
}
type _Config struct {
	NotExportBus bool
	OutputPath   string
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
	fmt.Println(info)
	return info
}

func parse_info() Infos {
	var outputPath, inputFile string
	flag.StringVar(&outputPath, "out", "dbusproxy.go", "the file to save the generated code")
	flag.StringVar(&inputFile, "in", "dbus.in.json", "the config file path")
	infos := loadInfo(inputFile)
	if outputPath != "dbusproxy.go" {
		infos.Config.OutputPath = outputPath
	} else if len(infos.Config.OutputPath) == 0 {
		infos.Config.OutputPath = outputPath
	}
	return infos
}

func main() {
	infos := parse_info()
	writer, err := os.Create(infos.Config.OutputPath)
	if err != nil {
		panic(err)
	}
	writer.WriteString(__GLOBAL_TEMPLATE)
	defer func() {
		writer.Close()
		exec.Command("gofmt", "-w", infos.Config.OutputPath).Start()
	}()
	for _, ifc := range infos.Interfaces {
		var reader io.Reader
		if _, err := os.Stat(ifc.File); ifc.Type == "xml" && err == nil {
			reader, err = os.Open(ifc.File)
			if err != nil {
				panic(err.Error() + "(File:" + ifc.File + ")")
			}
			GenInterfaceCode(reader, writer, ifc.Dest, ifc.Path, ifc.Name, ifc.ObjName)
			reader.(*os.File).Close()
		} else if ifc.Type == "introspect" {
			conn, _ := dbus.SessionBus()
			var xml string
			if err := conn.Object(ifc.Dest, dbus.ObjectPath(ifc.Path)).Call("org.freedesktop.DBus.Introspectable.Introspect", 0).Store(&xml); err != nil {
				panic(err.Error() + "Interface " + ifc.Name + " is can't dynamic introspect")
			}
			GenInterfaceCode(bytes.NewBufferString(xml), writer, ifc.Dest, ifc.Path, ifc.Name, ifc.ObjName)

		}
	}
}
