package main

import "path"
import "encoding/xml"
import "encoding/json"
import "log"
import "os"
import "os/exec"
import "io"
import "flag"
import "bytes"

import "dlib/dbus"

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
	OutFile, XMLFile, Dest, ObjectPath, Interface, ObjectName, TestPath string
}
type _Config struct {
	Language     string
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

var INFOS *Infos

func parseInfo() {
	if INFOS != nil {
		panic("Don't call multime the function of parseInfo")
	}
	INFOS = new(Infos)
	var outputPath, inputFile string
	flag.StringVar(&outputPath, "out", "out", "the file to save the generated code")
	flag.StringVar(&inputFile, "in", "dbus.in.json", "the config file path")

	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(INFOS)
	if err != nil {
		panic(err)
	}

	if outputPath != "out" {
		INFOS.Config.OutputDir = outputPath
	} else if len(INFOS.Config.OutputDir) == 0 {
		INFOS.Config.OutputDir = outputPath
	}
	if INFOS.Config.Language == "GoLang" {
		for i, ifc := range INFOS.Interfaces {
			INFOS.Interfaces[i].OutFile = ifc.OutFile + ".go"
		}
	} else if INFOS.Config.Language == "PyQt" {
		for i, ifc := range INFOS.Interfaces {
			INFOS.Interfaces[i].OutFile = ifc.OutFile + ".py"
		}
	} else {
		log.Fatal(`Didn't no generate target language, please set Language to "Golang" or "PyQt"`)
	}
}

func main() {
	parseInfo()
	os.MkdirAll(INFOS.Config.OutputDir, 0755)
	var writer io.Writer
	var err error
	if INFOS.Config.Language == "GoLang" {
		if writer, err = os.Create(path.Join(INFOS.Config.OutputDir, "init.go")); err != nil {
		panic(err)
		}
	} else if INFOS.Config.Language == "PyQt" {
		if writer, err = os.Create(path.Join(INFOS.Config.OutputDir, "__init__.py")); err != nil {
		panic(err)
		}
	}
	getMainTemplate().Execute(writer, nil)

	writer.(*os.File).Close()
	defer func() {
		exec.Command("gofmt", "-w", INFOS.Config.OutputDir).Start()
	}()
	for _, ifc := range INFOS.Interfaces {
		file := path.Join(INFOS.Config.InputDir, ifc.XMLFile)
		var reader io.Reader
		writer, err = os.Create(path.Join(INFOS.Config.OutputDir, ifc.OutFile))
		if _, err := os.Stat(file); err == nil {
			reader, err = os.Open(file)
			if err != nil {
				panic(err.Error() + "(File:" + file + ")")
			}
			info := GetInterfaceInfo(reader, ifc.Interface)
			GenInterfaceCode(INFOS.Config.Language, INFOS.Config.PkgName, info, writer, INFOS.Config.DestName, ifc.Interface, ifc.ObjectName)
			/*if ifc.TestPath != "" {*/
			/*var test_writer io.Writer*/
			/*test_writer, err = os.Create(path.Join(INFOS.Config.OutputDir, path.Base(ifc.OutFile)+"_test.go"))*/
			/*genTest(ifc.TestPath, INFOS.Config.PkgName, ifc.ObjectName, test_writer, info)*/
			/*}*/
			reader.(*os.File).Close()
		} else {
			conn, _ := dbus.SystemBus()
			var xml string
			if err := conn.Object(ifc.Dest, dbus.ObjectPath(ifc.ObjectPath)).Call("org.freedesktop.DBus.Introspectable.Introspect", 0).Store(&xml); err != nil {
				panic(err.Error() + "Interface " + ifc.Interface + " is can't dynamic introspect")
			}
			GenInterfaceCode(INFOS.Config.Language, INFOS.Config.PkgName, GetInterfaceInfo(bytes.NewBufferString(xml), ifc.Interface), writer, INFOS.Config.DestName, ifc.Interface, ifc.ObjectName)

		}
		writer.(*os.File).Close()
	}
}

/*func genTest(testPath, pkgName string, objName string, writer io.Writer, info dbus.InterfaceInfo) {*/
/*funcs := template.FuncMap{*/
/*"TestPath": func() string { return testPath },*/
/*"PkgName":  func() string { return pkgName },*/
/*"ObjName":  func() string { return objName },*/
/*[>"GetTestValue": func(args []dbus.ArgInfo) string {<]*/
/*[>},<]*/
/*}*/
/*template.Must(template.New("testing").Funcs(funcs).Parse(__TEST_TEMPLATE)).Execute(writer, info)*/
/*}*/
