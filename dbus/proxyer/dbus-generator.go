package main

import "path"
import "encoding/xml"
import "encoding/json"
import "strings"
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
	Target       string
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
	outputs    map[string]io.Writer
}

var INFOS *Infos

func parseInfo() {
	if INFOS != nil {
		panic("Don't call multime the function of parseInfo")
	}
	INFOS = new(Infos)
	var outputPath, inputFile string
	flag.StringVar(&outputPath, "out", "out", "the directory to save the generated code")
	flag.StringVar(&inputFile, "in", "dbus.in.json", "the config file path")
	flag.Parse()

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
	INFOS.outputs = make(map[string]io.Writer)
	os.MkdirAll(INFOS.Config.OutputDir, 0755)
	busType := strings.ToLower(INFOS.Config.BusType)
	INFOS.Config.BusType = busType
	if busType != "sesssion" && busType != "system" {
		log.Fatal("Didn't support bus type", busType)
	}
	if INFOS.Config.Target == "GoLang" {
		for i, ifc := range INFOS.Interfaces {
			name := ifc.OutFile + ".go"
			INFOS.Interfaces[i].OutFile = name
			if INFOS.outputs[name], err = os.Create(path.Join(INFOS.Config.OutputDir, name)); err != nil {
				panic(err)
			}
			renderInterfaceInit(INFOS.outputs[name])
		}
	} else if INFOS.Config.Target == "PyQt" {
		for i, ifc := range INFOS.Interfaces {
			name := ifc.OutFile + ".py"
			INFOS.Interfaces[i].OutFile = name
			if INFOS.outputs[name], err = os.Create(path.Join(INFOS.Config.OutputDir, name)); err != nil {
				panic(err)
			}
			renderInterfaceInit(INFOS.outputs[name])
		}
	} else {
		log.Fatal(`Didn't supported target , please set Target to "Golang" or "PyQt"`)
	}
}

func main() {
	parseInfo()
	var writer io.Writer
	var err error
	if INFOS.Config.Target == "GoLang" {
		if writer, err = os.Create(path.Join(INFOS.Config.OutputDir, "init.go")); err != nil {
			panic(err)
		}
	} else if INFOS.Config.Target == "PyQt" {
		if writer, err = os.Create(path.Join(INFOS.Config.OutputDir, "__init__.py")); err != nil {
			panic(err)
		}
	}
	renderMain(writer)
	writer.(*os.File).Close()

	defer func() {
		exec.Command("gofmt", "-w", INFOS.Config.OutputDir).Start()
		for _, w := range INFOS.outputs {
			w.(*os.File).Close()
		}
	}()
	for _, ifc := range INFOS.Interfaces {
		writer = INFOS.outputs[ifc.OutFile]

		inFile := path.Join(INFOS.Config.InputDir, ifc.XMLFile)
		var reader io.Reader
		if _, err := os.Stat(inFile); err == nil {
			reader, err = os.Open(inFile)
			if err != nil {
				panic(err.Error() + "(File:" + inFile + ")")
			}
			info := GetInterfaceInfo(reader, ifc.Interface)
			renderInterface(INFOS.Config.Target, INFOS.Config.PkgName, info, writer, INFOS.Config.DestName, ifc.Interface, ifc.ObjectName)
			/*if ifc.TestPath != "" {*/
			/*var test_writer io.Writer*/
			/*test_writer, err = os.Create(path.Join(INFOS.Config.OutputDir, path.Base(ifc.OutFile)+"_test.go"))*/
			/*render(ifc.TestPath, INFOS.Config.PkgName, ifc.ObjectName, test_writer, info)*/
			/*}*/
			reader.(*os.File).Close()
		} else {
			conn, _ := dbus.SystemBus()
			var xml string
			if err := conn.Object(ifc.Dest, dbus.ObjectPath(ifc.ObjectPath)).Call("org.freedesktop.DBus.Introspectable.Introspect", 0).Store(&xml); err != nil {
				panic(err.Error() + "Interface " + ifc.Interface + " is can't dynamic introspect")
			}
			renderInterface(INFOS.Config.Target, INFOS.Config.PkgName, GetInterfaceInfo(bytes.NewBufferString(xml), ifc.Interface), writer, INFOS.Config.DestName, ifc.Interface, ifc.ObjectName)

		}
	}
}
